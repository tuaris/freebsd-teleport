--- lib/sshutils/scp/local.go.orig	2025-02-28 20:13:31 UTC
+++ lib/sshutils/scp/local.go
@@ -104,7 +104,7 @@ func makeFileInfo(filePath string) (FileInfo, error) {
 	return &localFileInfo{
 		filePath:   filePath,
 		fileInfo:   f,
-		accessTime: GetAtime(f),
+		accessTime: f.ModTime(),
 	}, nil
 }
 
