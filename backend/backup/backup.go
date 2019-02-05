package backup

import (
		"log"
		"io/ioutil"
)

func List(backupDir string) ([]string, error) {
	files, err := ioutil.ReadDir(backupDir)
	if err != nil {
		log.Println("Cannot get list of files in ", backupDir, err)
		return nil, err
	}
	var names []string
	for _, x := range files {
		names = append(names, x.Name())
	}

	return names, nil
}