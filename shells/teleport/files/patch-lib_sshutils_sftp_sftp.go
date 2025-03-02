--- lib/sshutils/sftp/sftp.go.orig	2025-02-28 20:13:31 UTC
+++ lib/sshutils/sftp/sftp.go
@@ -42,7 +42,7 @@ import (
 
 	"github.com/gravitational/teleport"
 	"github.com/gravitational/teleport/lib/defaults"
-	"github.com/gravitational/teleport/lib/sshutils/scp"
+//	"github.com/gravitational/teleport/lib/sshutils/scp"
 )
 
 // Options control aspects of a file transfer
@@ -664,7 +664,7 @@ func getAtime(fi os.FileInfo) time.Time {
 		return time.Unix(int64(sftpfi.Atime), 0)
 	}
 
-	return scp.GetAtime(fi)
+	return fi.ModTime()
 }
 
 // NewProgressBar returns a new progress bar that writes to writer.
