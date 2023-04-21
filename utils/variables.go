package utils

import (
	"fmt"
	"path/filepath"
)

func getCSVFilepath() string {
	absPath, err := filepath.Abs(CSV_FILE_PATH)
	if err != nil {
		fmt.Println("Error getting absolute path:", err)
	}
	return absPath
}

var CsvFilePath = getCSVFilepath()
