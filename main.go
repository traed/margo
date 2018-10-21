package main

import (
	driveapi "margo/driveApi"
	"os"
)

func main() {
	var filepath string
	if len(os.Args) < 2 {
		filepath = "test.md"
	} else {
		filepath = os.Args[1]
	}

	driveapi.UploadFile(filepath)
}
