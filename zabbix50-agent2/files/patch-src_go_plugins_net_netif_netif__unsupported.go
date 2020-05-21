--- src/go/plugins/net/netif/netif_unsupported.go.orig	2020-05-21 17:57:39 UTC
+++ src/go/plugins/net/netif/netif_unsupported.go
@@ -1,4 +1,4 @@
-// +build !linux,!windows
+// +build !freebsd,!linux,!windows
 
 /*
 ** Zabbix
