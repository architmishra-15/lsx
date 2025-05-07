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
	flags, remainingArgs := handle_flag(args)

	// Determine the path to process - use the first non-flag argument if it exists
	// otherwise use the current directory
	var path string
	if len(remainingArgs) > 1 {
		path = remainingArgs[1]
	} else {
		path = "."
	}

	// Checking if the provided path is a directory
	fileInfo, err := os.Stat(path)
	if err == nil && fileInfo.IsDir() {
		// If -d flag is set, print the directory entry itself, not its contents
		if flags.DirectoryOnly {
			fileInfos := []os.FileInfo{fileInfo}
			PrintFilesInColumns(fileInfos, 1)
		} else {
			// It's a directory, print its contents
			printDirectoryContents(path, flags)
		}
	} else {
		// It's not a directory or there's an error, handle as pattern
		processPath(path, flags)
	}
}

func processPath(pattern string, flags Flags) {
	// 1. Handle `*.extension`
	if strings.HasPrefix(pattern, "*.") {
		ext := strings.TrimPrefix(pattern, "*")
		printFilesWithExtension(".", ext, flags)
		return
	}

	// 2. Handle `<some_path>/*.extension`
	parts := strings.Split(pattern, "/")
	if len(parts) > 1 && strings.HasSuffix(parts[len(parts)-1], ".*") {
		path := strings.Join(parts[:len(parts)-1], "/")
		ext := strings.TrimPrefix(parts[len(parts)-1], "*")
		printFilesWithExtension(path, ext, flags)
		return
	}

	// 3. Handle `<some_path>/<some_filename>` or `<some_filename>`
	// Check if the file exists
	_, err := os.Stat(pattern)
	if err == nil {
		printFile(pattern, flags)
		return
	} else if os.IsNotExist(err) {
		fmt.Println("Error: No such file or directory:", pattern)
		return
	} else {
		log.Fatal(err)
	}

	fmt.Println("Error: No such file or directory:", pattern)
}

func printFilesWithExtension(dirPath string, ext string, flags Flags) {
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
	PrintFilesInColumns(fileInfos, 5, flags)
}

func printFile(filePath string, flags Flags) {
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		log.Fatal(err)
	}
	PrintFilesInColumns([]os.FileInfo{fileInfo}, 1, flags)
}

func printDirectoryContents(dirPath string, flags Flags) {
	dirEntries, err := os.ReadDir(dirPath)
	if err != nil {
		log.Fatal(err)
	}

	fileInfos := make([]os.FileInfo, 0)
	for _, entry := range dirEntries {
		// Skip dotfiles - unless -a flag is set
		if strings.HasPrefix(entry.Name(), ".") && !flags.AllFiles {
			continue
		}
		info, err := entry.Info()
		if err != nil {
			continue
		}
		fileInfos = append(fileInfos, info)
	}

	// For long format, use only 1 column
	columns := 5
	if flags.LongFormat {
		columns = 1
	}

	PrintFilesInColumns(fileInfos, columns, flags)
}
