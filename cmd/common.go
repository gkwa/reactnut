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

func concatWords() string {
	adjective := randomdata.Adjective()
	noun := randomdata.Noun()
	concat := strings.ToLower(fmt.Sprintf("%s%s", adjective, noun))
	log.Logger.Tracef("concatinated: %s", concat)
	return concat
}

func genPathStr(basePath string, fullPath *string) error {
	concat := concatWords()
	*fullPath = filepath.Join(basePath, concat)

	if filepath.IsAbs(*fullPath) {
		return nil
	}
	c, err := os.Getwd()
	if err != nil {
		return err
	}
	n := filepath.Join(c, *fullPath)
	*fullPath = n
	return nil
}

func dostuff(basePath string) (string, error) {
	log.Logger.Trace("calling dostuff")
	i := 0

	var fullPath string
	for i < 5 {
		log.Logger.Tracef("i:%d", i)
		err := genPathStr(basePath, &fullPath)
		if err != nil {
			log.Logger.Fatalf("can't create path %s", fullPath)
		}
		log.Logger.Debugf("fullPath: %s", fullPath)

		if !pathExists(fullPath) {
			break
		}
		i++
	}

	log.Logger.Tracef("making directory %s", fullPath)
	err := os.MkdirAll(fullPath, 0o755)
	if err != nil {
		log.Logger.Fatalf("failed to create path '%s'", fullPath)
	}
	if !pathExists(fullPath) {
		return "", err
	}
	return fullPath, err
}

func pathExists(path1 string) bool {
	path, err := expandTilde(path1)
	if err != nil {
		log.Logger.Fatalf("expanding tilde creates error for path: %s, error: %s",
			path, err)
	}
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
			log.Logger.Warningf("checking current user results in error: %s", err)
			return "", err
		}
		return strings.Replace(path, "~", currentUser.HomeDir, 1), nil
	}
	log.Logger.Tracef("returning path: %s", path)
	return path, nil
}
