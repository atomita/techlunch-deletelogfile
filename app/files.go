package app

import (
	"log"
	"os"
	"path/filepath"
	"time"
)

func Files(pattern string, before time.Time) (files []string, err error) {
	files, err = filepath.Glob(pattern)
	if err != nil {
		return
	}

	files, err = filter(files, func(file string) (bool, error) {
		fileinfo, err := os.Stat(file)
		if err != nil {
			return false, err
		}

		if fileinfo.IsDir() {
			log.Printf("trace: is directory: " + file)
			return false, nil
		}

		if before.Before(fileinfo.ModTime()) {
			log.Printf("trace: Within storage period: " + file)
			return false, nil
		}

		return true, nil
	})
	return
}

func filter(values []string, f func(string) (bool, error)) (filterd []string, err error) {
	filtered := make([]string, 0)
	for _, v := range values {
		flag, err := f(v)
		if err != nil {
			return filterd, err
		}
		if flag {
			filtered = append(filtered, v)
		}
	}
	return filtered, nil
}
