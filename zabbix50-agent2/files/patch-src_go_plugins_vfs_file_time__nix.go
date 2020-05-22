--- src/go/plugins/vfs/file/time_nix.go.orig	2020-05-22 10:28:33 UTC
+++ src/go/plugins/vfs/file/time_nix.go
@@ -24,7 +24,8 @@ package file
 import (
 	"errors"
 	"fmt"
-	"syscall"
+
+	"golang.org/x/sys/unix"
 )
 
 // Export -
@@ -41,9 +42,9 @@ func (p *Plugin) exportTime(params []string) (result i
 		if len(params) == 1 || params[1] == "" || params[1] == "modify" {
 			return f.ModTime().Unix(), nil
 		} else if params[1] == "access" {
-			return f.Sys().(*syscall.Stat_t).Atim.Sec, nil
+			return f.Sys().(*unix.Stat_t).Atim.Sec, nil
 		} else if params[1] == "change" {
-			return f.Sys().(*syscall.Stat_t).Ctim.Sec, nil
+			return f.Sys().(*unix.Stat_t).Ctim.Sec, nil
 		} else {
 			return nil, errors.New("Invalid second parameter.")
 		}
