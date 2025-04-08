// main.go

package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	args := os.Args

	// Handle flags
	if len(args) > 1 {
		if args[1] == "-h" || args[1] == "--help" {
			showHelp()
			return
		}
	}

	var path string
	if len(args) > 1 {
		path = args[1]
	} else {
		path = "."
	}

	// Checking if the provided path is a directory
	fileInfo, err := os.Stat(path)
	if err == nil && fileInfo.IsDir() {
		// It's a directory, print its contents
		printDirectoryContents(path)
	} else {
		// It's not a directory or there's an error, handle as pattern
		processPath(path)
	}
}

func processPath(pattern string) {
	// 1. Handle `*.extension`
	if strings.HasPrefix(pattern, "*.") {
		ext := strings.TrimPrefix(pattern, "*")
		printFilesWithExtension(".", ext)
		return
	}

	// 2. Handle `<some_path>/*.extension`
	parts := strings.Split(pattern, "/")
	if len(parts) > 1 && strings.HasSuffix(parts[len(parts)-1], ".*") {
		path := strings.Join(parts[:len(parts)-1], "/")
		ext := strings.TrimPrefix(parts[len(parts)-1], "*")
		printFilesWithExtension(path, ext)
		return
	}

	// 3. Handle `<some_path>/<some_filename>` or `<some_filename>`
	// Check if the file exists
	_, err := os.Stat(pattern)
	if err == nil {
		// If exists, print it
		printFile(pattern)
		return
	} else if os.IsNotExist(err) {
		// If doesn't exist
		fmt.Println("Error: No such file or directory:", pattern)
		return
	} else {
		// Some other error
		log.Fatal(err)
	}

	// If none of the above, and it's not an existing directory,
	// it might be an invalid path
	fmt.Println("Error: No such file or directory:", pattern)
}

func printFilesWithExtension(dirPath string, ext string) {
	dirEntries, err := os.ReadDir(dirPath)
	if err != nil {
		log.Fatal(err)
	}

	fileInfos := make([]os.FileInfo, 0)
	for _, entry := range dirEntries {
		if !strings.HasSuffix(entry.Name(), ext) {
			continue
		}
		info, err := entry.Info()
		if err != nil {
			continue
		}
		fileInfos = append(fileInfos, info)
	}
	PrintFilesInColumns(fileInfos, 5)
}

func printFile(filePath string) {
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		log.Fatal(err)
	}
	PrintFilesInColumns([]os.FileInfo{fileInfo}, 1)
}

func printDirectoryContents(dirPath string) {
	dirEntries, err := os.ReadDir(dirPath)
	if err != nil {
		log.Fatal(err)
	}

	fileInfos := make([]os.FileInfo, 0)
	for _, entry := range dirEntries {
		// Skip dotfiles - files that start with a dot
		if strings.HasPrefix(entry.Name(), ".") {
			continue
		}
		info, err := entry.Info()
		if err != nil {
			continue
		}
		fileInfos = append(fileInfos, info)
	}
	PrintFilesInColumns(fileInfos, 5)
}
