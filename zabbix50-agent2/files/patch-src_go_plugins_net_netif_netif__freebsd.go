--- src/go/plugins/net/netif/netif_freebsd.go.orig      2020-05-20 10:56:50 UTC
+++ src/go/plugins/net/netif/netif_freebsd.go
@@ -0,0 +1,193 @@
+/*
+** Zabbix
+** Copyright (C) 2001-2020 Zabbix SIA
+**
+** This program is free software; you can redistribute it and/or modify
+** it under the terms of the GNU General Public License as published by
+** the Free Software Foundation; either version 2 of the License, or
+** (at your option) any later version.
+**
+** This program is distributed in the hope that it will be useful,
+** but WITHOUT ANY WARRANTY; without even the implied warranty of
+** MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
+** GNU General Public License for more details.
+**
+** You should have received a copy of the GNU General Public License
+** along with this program; if not, write to the Free Software
+** Foundation, Inc., 51 Franklin Street, Fifth Floor, Boston, MA  02110-1301, USA.
+**/
+
+package netif
+
+import (
+	"bufio"
+	"bytes"
+	"encoding/json"
+	"errors"
+	"strconv"
+	"strings"
+        "log"
+        "os/exec"
+
+	"zabbix.com/pkg/plugin"
+	"zabbix.com/pkg/std"
+)
+
+var stdOs std.Os
+
+var mapNetStatIn = map[string]uint{
+	"name":       0,
+	"mtu":        1,
+	"network":    2,
+	"address":    3,
+	"packets":    4,
+	"errors":     5,
+	"dropped":    6,
+	"bytes":      7,
+}
+
+var mapNetStatOut = map[string]uint{
+	"packets":    8,
+	"errors":     9,
+	"bytes":      10,
+	"collisions": 11,
+}
+
+func (p *Plugin) addStatNum(statName string, mapNetStat map[string]uint, statNums *[]uint) error {
+	if statNum, ok := mapNetStat[statName]; ok {
+		*statNums = append(*statNums, statNum)
+	} else {
+		return errors.New(errorInvalidSecondParam)
+	}
+	return nil
+}
+
+func (p *Plugin) getNetStats(networkIf string, statName string, dir dirFlag) (result uint64, err error) {
+	var statNums []uint
+
+	if dir&dirIn != 0 {
+		if err = p.addStatNum(statName, mapNetStatIn, &statNums); err != nil {
+			return
+		}
+	}
+
+	if dir&dirOut != 0 {
+		if err = p.addStatNum(statName, mapNetStatOut, &statNums); err != nil {
+			return
+		}
+	}
+
+        out, err := exec.Command("/usr/bin/netstat", "-bnI", networkIf).Output()
+        if err != nil {
+                log.Fatal("netstat failed: %v\n", err)
+        }
+        file := bytes.NewBufferString(string(out))
+
+	var total uint64
+loop:
+	for sLines := bufio.NewScanner(file); sLines.Scan(); {
+		dev := strings.Split(sLines.Text(), " ")
+
+		if len(dev) > 1 && networkIf == strings.TrimSpace(dev[0]) {
+			stats := strings.Fields(sLines.Text())
+
+			if len(stats) >= 12 {
+				for _, statNum := range statNums {
+					var res uint64
+
+					if res, err = strconv.ParseUint(stats[statNum], 10, 64); err != nil {
+						break loop
+					}
+					total += res
+				}
+				return total, nil
+			}
+			break
+		}
+	}
+	err = errors.New("Cannot find information for this network interface")
+	return
+}
+
+func (p *Plugin) getDevDiscovery() (netInterfaces []msgIfDiscovery, err error) {
+        out, err := exec.Command("/sbin/ifconfig", "-l").Output()
+        if err != nil {
+                log.Fatal("/sbin/ifconfig -a failed: %v\n", err)
+        }
+
+	netInterfaces = make([]msgIfDiscovery, 0)
+	netif := strings.Fields(string(out))
+	for _, i := range netif {
+		netInterfaces = append(netInterfaces, msgIfDiscovery{i})
+	}
+
+	return netInterfaces, nil
+}
+
+// Export -
+func (p *Plugin) Export(key string, params []string, ctx plugin.ContextProvider) (result interface{}, err error) {
+	var direction dirFlag
+	var mode string
+
+	switch key {
+	case "net.if.discovery":
+		if len(params) > 0 {
+			return nil, errors.New(errorParametersNotAllowed)
+		}
+		var devices []msgIfDiscovery
+		if devices, err = p.getDevDiscovery(); err != nil {
+			return
+		}
+		var b []byte
+		if b, err = json.Marshal(devices); err != nil {
+			return
+		}
+		return string(b), nil
+	case "net.if.collisions":
+		if len(params) > 1 {
+			return nil, errors.New(errorTooManyParams)
+		}
+
+		if len(params) < 1 || params[0] == "" {
+			return nil, errors.New(errorEmptyIfName)
+		}
+		return p.getNetStats(params[0], "collisions", dirOut)
+	case "net.if.in":
+		direction = dirIn
+	case "net.if.out":
+		direction = dirOut
+	case "net.if.total":
+		direction = dirIn | dirOut
+	default:
+		/* SHOULD_NEVER_HAPPEN */
+		return nil, errors.New(errorUnsupportedMetric)
+	}
+
+	if len(params) < 1 || params[0] == "" {
+		return nil, errors.New(errorEmptyIfName)
+	}
+
+	if len(params) > 2 {
+		return nil, errors.New(errorTooManyParams)
+	}
+
+	if len(params) == 2 && params[1] != "" {
+		mode = params[1]
+	} else {
+		mode = "bytes"
+	}
+
+	return p.getNetStats(params[0], mode, direction)
+}
+
+func init() {
+	stdOs = std.NewOs()
+
+	plugin.RegisterMetrics(&impl, "NetIf",
+		"net.if.collisions", "Returns number of out-of-window collisions.",
+		"net.if.in", "Returns incoming traffic statistics on network interface.",
+		"net.if.out", "Returns outgoing traffic statistics on network interface.",
+		"net.if.total", "Returns sum of incoming and outgoing traffic statistics on network interface.",
+		"net.if.discovery", "Returns list of network interfaces. Used for low-level discovery.")
+
+}
