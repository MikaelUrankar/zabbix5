--- src/go/plugins/kernel/kernel_freebsd.go.orig	2020-05-21 11:39:04 UTC
+++ src/go/plugins/kernel/kernel_freebsd.go
@@ -0,0 +1,54 @@
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
+package kernel
+
+import (
+	"fmt"
+	"os/exec"
+	"strconv"
+	"strings"
+)
+
+func getMax(proc bool) (max uint64, err error) {
+	var out []byte
+	var result string
+
+	if proc {
+		out, err = exec.Command("/sbin/sysctl", "-n", "kern.pid_max").Output()
+		if err != nil {
+			err = fmt.Errorf("sysctl -n kern.pid_max failed: %v\n", err)
+		}
+		result = strings.TrimRight(string(out), "\n")
+		max, err = strconv.ParseUint(string(result), 10, 64)
+	} else {
+		out, err = exec.Command("/sbin/sysctl", "-n", "kern.maxfiles").Output()
+		if err != nil {
+			err = fmt.Errorf("sysctl -n kern.maxfiles failed: %v\n", err)
+		}
+		result = strings.TrimRight(string(out), "\n")
+		max, err = strconv.ParseUint(string(result), 10, 64)
+	}
+
+	if err != nil {
+		err = fmt.Errorf("Cannot obtain data from kernel.")
+	}
+
+	return
+}
