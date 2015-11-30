package lib

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"

	"github.com/emirozer/cf-fastpush-plugin/utils"
)

type FileEntry struct {
	Checksum     string
	Modification int64
	Content      []byte
}

type Status struct {
	Health string
}

var cmd *exec.Cmd
var lock = sync.RWMutex{}
var cmdRaw = ""
var store = map[string]*FileEntry{}

const (
	ENV_RESTART_REGEX = "FASTPUSH_RESTART_REGEX"
	ENV_IGNORE_REGEX  = "FASTPUSH_IGNORE_REGEX"
	ENV_APP_DIRS      = "FASTPUSH_APP_DIRS"
)

func ListFiles() map[string]*FileEntry {
	for _, dir := range GetAppDirs() {
		err := filepath.Walk(dir, func(path string, f os.FileInfo, err error) error {
			if f.IsDir() {
				return nil
			}
			if store[path] != nil && store[path].Modification == f.ModTime().Unix() {
				// cache hit
				return nil
			}
			fileEntry := FileEntry{}
			checksum, _ := utils.ChecksumsForFile(path)
			fileEntry.Checksum = checksum.SHA256
			fileEntry.Modification = f.ModTime().Unix()
			lock.RLock()
			store[path] = &fileEntry
			lock.RUnlock()
			return nil
		})
		if err != nil {
			log.Println(err)
		}
	}
	return store
}

func GetAppDirs() []string {
	appDirsRaw := os.Getenv(ENV_APP_DIRS)
	if len(appDirsRaw) > 0 {
		return strings.Fields(appDirsRaw)
	}
	return []string{"./"}
}
