--- lib/utils/disk.go.orig	2025-02-28 20:13:31 UTC
+++ lib/utils/disk.go
@@ -22,12 +22,12 @@
 package utils
 
 import (
-	"errors"
-	"io/fs"
-	"os"
-	"os/user"
-	"path/filepath"
-	"strconv"
+
+
+
+
+
+
 	"syscall"
 
 	"github.com/gravitational/trace"
@@ -68,75 +68,12 @@ func FreeDiskWithReserve(dir string, reservedFreeDisk 
 // as it's not bullet proof and can report wrong results.
 func CanUserWriteTo(path string) (bool, error) {
 	// prevent infinite loops with a max dir depth
-	var fileInfo os.FileInfo
-	var err error
 
-	for i := 0; i < 20; i++ {
-		fileInfo, err = os.Stat(path)
-		if err == nil {
-			break
-		}
-		if errors.Is(err, fs.ErrNotExist) {
-			path = filepath.Dir(path)
-			continue
-		}
 
-		return false, trace.BadParameter("failed to find path: %+v", err)
 
-	}
 
-	var uid int
-	var gid int
-	if stat, ok := fileInfo.Sys().(*syscall.Stat_t); ok {
-		uid = int(stat.Uid)
-		gid = int(stat.Gid)
-	}
 
-	var usr *user.User
-	if ogUser := os.Getenv("SUDO_USER"); ogUser != "" {
-		usr, err = user.Lookup(ogUser)
-		if err != nil {
-			return false, trace.NotFound("could not determine original user: %+v", err)
-		}
-	} else {
-		usr, err = user.Current()
-		if err != nil {
-			return false, trace.NotFound("could not determine current user: %+v", err)
-		}
-	}
 
-	perm := fileInfo.Mode().Perm()
 
-	// file is owned by the user
-	if strconv.Itoa(uid) == usr.Uid {
-		// file has u+wx permissions
-		if perm&syscall.S_IWUSR != 0 &&
-			perm&syscall.S_IXUSR != 0 {
-			return true, nil
-		}
-	}
-
-	// file and user have a group in common
-	groupIDs, err := usr.GroupIds()
-	if err != nil {
-		return false, trace.NotFound("could not determine current user group ids: %+v", err)
-	}
-	for _, groupID := range groupIDs {
-		if strconv.Itoa(gid) == groupID {
-			// file has g+wx permissions
-			if perm&syscall.S_IWGRP != 0 &&
-				perm&syscall.S_IXGRP != 0 {
-				return true, nil
-			}
-			break
-		}
-	}
-
-	// file has o+wx permissions
-	if perm&syscall.S_IWOTH != 0 &&
-		perm&syscall.S_IXOTH != 0 {
-		return true, nil
-	}
-
-	return false, nil
+	return true, nil
 }
