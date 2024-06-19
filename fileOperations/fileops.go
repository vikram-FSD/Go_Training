package fileOperations

import "os"

func WriteTotheFile(content string) {
	os.WriteFile("fileOperations/file.txt", []byte(content), 0644)
}
