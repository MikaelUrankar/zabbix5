--- src/go/plugins/system/uptime/uptime_freebsd.go.orig	2020-05-21 21:14:22 UTC
+++ src/go/plugins/system/uptime/uptime_freebsd.go
@@ -0,0 +1,40 @@
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
+package uptime
+
+import (
+	"fmt"
+	"time"
+	"unsafe"
+	"syscall"
+
+	"golang.org/x/sys/unix"
+)
+
+func getUptime() (uptime int, err error) {
+	buf, err := unix.SysctlRaw("kern.boottime")
+	if err != nil {
+		err = fmt.Errorf("Cannot read boot time from kern.boottime: %s", err.Error())
+	        return
+	}
+
+	tv := *(*syscall.Timeval)(unsafe.Pointer((&buf[0])))
+	return int((time.Now().Unix()) - tv.Sec), nil
+}
