package main

import "os"

func checkIfFileOrDirectoryExists(filename string) error {
	_, FileOrFolderError := os.Stat(filename)
	return FileOrFolderError
}

func createDirectory(dirPath string) {
	os.Mkdir(dirPath, 0666)
}

func readFromFile(filename string) ([]byte, error) {
	return os.ReadFile(filename)
}

func writeToFile(filename string, byteSlice []byte) error {
	return os.WriteFile(filename, byteSlice, 0666)
}
