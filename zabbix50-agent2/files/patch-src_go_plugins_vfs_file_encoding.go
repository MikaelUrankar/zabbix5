--- src/go/plugins/vfs/file/encoding.go.orig	2020-05-28 11:51:53 UTC
+++ src/go/plugins/vfs/file/encoding.go
@@ -25,6 +25,7 @@ package file
 //   return iconv(cd, &inbuf, inbytesleft, &outbuf, outbytesleft);
 // }
 //
+// #cgo freebsd LDFLAGS: -liconv
 // #cgo windows LDFLAGS: -liconv
 import "C"
 
