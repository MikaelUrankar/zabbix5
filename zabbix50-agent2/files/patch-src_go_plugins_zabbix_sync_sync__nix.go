--- src/go/plugins/zabbix/sync/sync_nix.go.orig	2020-05-28 13:14:12.129504000 +0200
+++ src/go/plugins/zabbix/sync/sync_nix.go	2020-05-28 13:14:53.508801000 +0200
@@ -33,7 +33,6 @@ func getMetrics() []string {
 		"net.tcp.service", "Checks if service is running and accepting TCP connections.",
 		"net.tcp.service.perf", "Checks performance of TCP service.",
 		"system.users.num", "Number of users logged in.",
-		"system.swap.size", "Swap space size in bytes or in percentage from total.",
 		"vfs.dir.count", "Directory entry count.",
 		"vfs.dir.size", "Directory size (in bytes).",
 		"vfs.fs.inode", "Number or percentage of inodes.",
