package utils

import (
	"errors"
	"os"

	log "github.com/sirupsen/logrus"
)

// PathExists returns true if the given path exists, and false if it does not.
// If it cannot determine whether or not a path exists (e.g., because of a
// permissions problem), it will log an error and exit.
func PathExists(path string) bool {
	_, err := os.Stat(path)

	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return false
		}
		log.Fatalf("unable to check for %s: %v", path, err)
	}

	return true
}

// PathIsDirectory returns true if the given path is a directory, and false if
// it is not. If it cannot determine whether or not a path is a directory
// (e.g., because of a permissions problem), it will log an error and exit.
func PathIsDirectory(path string) bool {
	info, err := os.Stat(path)

	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return false
		}
		log.Fatalf("unable to check for %s: %v", path, err)
	}

	return info.IsDir()
}
