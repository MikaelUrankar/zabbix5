--- src/go/plugins/system/swap/swap_freebsd.go.orig	2020-05-28 13:16:24.250742000 +0200
+++ src/go/plugins/system/swap/swap_freebsd.go	2020-05-28 12:59:58.674912000 +0200
@@ -0,0 +1,85 @@
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
+package swap
+
+import (
+	"errors"
+
+	"zabbix.com/pkg/plugin"
+
+	"github.com/shirou/gopsutil/mem"
+)
+
+// Plugin -
+type Plugin struct {
+	plugin.Base
+}
+
+var impl Plugin
+
+// Export -
+func (p *Plugin) Export(key string, params []string, ctx plugin.ContextProvider) (result interface{}, err error) {
+	if key != "system.swap.size" {
+		return nil, plugin.UnsupportedMetricError
+	}
+
+	if len(params) > 2 {
+		return nil, errors.New("Too many parameters.")
+	}
+
+	var mode string
+	if len(params) == 2 && params[1] != "" {
+		mode = params[1]
+	}
+
+	if len(params) > 0 && params[0] != "" && params[0] != "all" {
+		return nil, errors.New("Invalid first parameter.")
+	}
+
+        sw, err := mem.SwapMemory()
+        if err != nil {
+		return nil, errors.New("Cannot get swap info.")
+        }
+
+	var total, avail uint64
+	total = sw.Total
+	avail = sw.Free
+
+	switch mode {
+	case "", "total":
+		return total, nil
+	case "free":
+		return avail, nil
+	case "used":
+		return total - avail, nil
+	case "pfree":
+		return float64(avail) / float64(total) * 100, nil
+	case "pused":
+		return float64(total-avail) / float64(total) * 100, nil
+	default:
+		return nil, errors.New("Invalid second parameter.")
+	}
+}
+
+func init() {
+	plugin.RegisterMetrics(&impl, "Swap",
+		"system.swap.size", "Returns Swap space size in bytes or in percentage from total.",
+	)
+}
