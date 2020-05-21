--- src/go/plugins/system/uname/uname_freebsd.go.orig	2020-05-20 10:56:50.872106000 +0200
+++ src/go/plugins/system/uname/uname_freebsd.go	2020-05-20 10:38:19.374078000 +0200
@@ -0,0 +1,83 @@
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
+package uname
+
+import (
+	"errors"
+	"fmt"
+
+	"golang.org/x/sys/unix"
+)
+
+func arrayToString(unameArray *[256]byte) string {
+	var byteString [256]byte
+	var indexLength int
+	for ; indexLength < len(unameArray); indexLength++ {
+		if 0 == unameArray[indexLength] {
+			break
+		}
+		byteString[indexLength] = uint8(unameArray[indexLength])
+	}
+	return string(byteString[:indexLength])
+}
+
+func getUname(params []string) (uname string, err error) {
+	if len(params) > 0 {
+		return "", errors.New("Too many parameters.")
+	}
+
+	var utsname unix.Utsname
+	if err = unix.Uname(&utsname); err != nil {
+		err = fmt.Errorf("Cannot obtain system information: %s", err.Error())
+		return
+	}
+	uname = fmt.Sprintf("%s %s %s %s %s", arrayToString(&utsname.Sysname), arrayToString(&utsname.Nodename),
+		arrayToString(&utsname.Release), arrayToString(&utsname.Version), arrayToString(&utsname.Machine))
+
+	return uname, nil
+}
+
+func getHostname(params []string) (hostname string, err error) {
+	if len(params) > 0 {
+		return "", errors.New("Too many parameters.")
+	}
+
+	var utsname unix.Utsname
+	if err = unix.Uname(&utsname); err != nil {
+		err = fmt.Errorf("Cannot obtain system information: %s", err.Error())
+		return
+	}
+
+	return arrayToString(&utsname.Nodename), nil
+}
+
+func getSwArch(params []string) (uname string, err error) {
+	if len(params) > 0 {
+		return "", errors.New("Too many parameters.")
+	}
+
+	var utsname unix.Utsname
+	if err = unix.Uname(&utsname); err != nil {
+		err = fmt.Errorf("Cannot obtain system information: %s", err.Error())
+		return
+	}
+
+	return arrayToString(&utsname.Machine), nil
+}
