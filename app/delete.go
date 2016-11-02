package app

import (
	"log"
	"os"
)

func Delete(files []string, dry bool) (err error) {
	for _, file := range files {
		if !dry {
			err := os.Remove(file)
			if err != nil {
				log.Printf("error: can not delete: " + file)
				return err
			}
		}

		log.Printf("info: delete: " + file)
	}
	return nil
}
