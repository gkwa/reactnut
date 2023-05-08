package cmd

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"strings"

	"github.com/Pallinder/go-randomdata"
	log "github.com/taylormonacelli/reactnut/cmd/logging"
)

func dostuff(basePath string) {
	adjective := randomdata.Adjective()
	noun := randomdata.Noun()
	concat := fmt.Sprintf("%s%s", adjective, noun)
	fullPath := filepath.Join(basePath, concat)

	for i := 0; i < 10; i++ {
		if !pathExists(fullPath) {
			break
		}
		adjective := randomdata.Adjective()
		noun := randomdata.Noun()
		concat := fmt.Sprintf("%s%s", adjective, noun)
		fullPath = filepath.Join(basePath, concat)
	}
	os.MkdirAll(fullPath, 0o755)
	fmt.Print(fullPath)
}

func pathExists(path string) bool {
	path, err := expandTilde(path)
	if err != nil {
		panic(err)
	}
	log.Logger.Trace("")
	log.Logger.Traceln(path) // output: /Users/username/Documents/example.txt

	// Use os.Stat() to get information about the path
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false
	}

	// Check if the error value is nil, which indicates that the path exists
	if err == nil {
		// Check if the path is a directory
		if fileInfo.Mode().IsDir() {
			log.Logger.Tracef("%s is a directory\n", path)
		} else {
			log.Logger.Tracef("%s is a file\n", path)
		}
	} else {
		log.Logger.Tracef("Path %s does not exist\n", path)
	}
	return true
}

func expandTilde(path string) (string, error) {
	if strings.HasPrefix(path, "~/") || path == "~" {
		currentUser, err := user.Current()
		if err != nil {
			return "", err
		}
		return strings.Replace(path, "~", currentUser.HomeDir, 1), nil
	}
	return path, nil
}
