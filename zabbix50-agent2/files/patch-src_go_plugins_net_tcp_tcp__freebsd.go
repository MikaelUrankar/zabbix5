--- src/go/plugins/net/tcp/tcp_freebsd.go.orig	2020-05-22 10:24:43 UTC
+++ src/go/plugins/net/tcp/tcp_freebsd.go
@@ -0,0 +1,35 @@
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
+package tcpudp
+
+import (
+	"errors"
+
+	"zabbix.com/pkg/plugin"
+)
+
+func exportSystemTcpListen(port uint16) (result interface{}, err error) {
+	return nil, errors.New("Not supported.")
+}
+
+func init() {
+	plugin.RegisterMetrics(&impl, "TCP",
+		"net.tcp.port", "Checks if it is possible to make TCP connection to specified port.")
+}