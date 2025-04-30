package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// printLongFormat displays files in the long listing format like ls -l
func printLongFormat(files []os.FileInfo, humanReadable bool) {
	// Calculate total size in 1K blocks
	var totalSize int64
	for _, file := range files {
		totalSize += file.Size()
	}
	totalBlocks := (totalSize + 1023) / 1024

	// Print total line (Unix ls compatibility)
	if humanReadable {
		fmt.Printf("total %s\n", FormatFileSize(totalBlocks*1024, true))
	} else {
		fmt.Printf("total %d\n", totalBlocks)
	}

	// Find the maximum length for various fields to ensure alignment
	maxSizeLen := 0
	maxOwnerLen := 0
	maxGroupLen := 0

	// First pass to determine max field lengths
	for _, file := range files {
		// Size length
		size := file.Size()
		sizeStr := ""
		if humanReadable {
			sizeStr = FormatFileSize(size, true)
		} else {
			sizeStr = strconv.FormatInt(size, 10)
		}
		if len(sizeStr) > maxSizeLen {
			maxSizeLen = len(sizeStr)
		}

		// Get owner and group and find max lengths
		filePath := file.Name()
		owner, group := getFileOwner(filePath)
		if len(owner) > maxOwnerLen {
			maxOwnerLen = len(owner)
		}
		if len(group) > maxGroupLen {
			maxGroupLen = len(group)
		}
	}

	// Now print each file in Unix-like format
	for _, file := range files {
		name := file.Name()
		mode := file.Mode()

		// Default Unix-style permission format
		perms := "-"

		// File type
		if mode.IsDir() {
			perms = "d"
		}

		// Add simulated rwx permissions
		if mode.IsDir() {
			perms += "rwxr-xr-x" // directories
		} else if strings.HasSuffix(strings.ToLower(name), ".exe") {
			perms += "rwxr-xr--" // executables
		} else {
			perms += "rw-r--r--" // regular files
		}

		// Get links, owner and group
		links := 1
		if mode.IsDir() {
			links = 2 // Dirs often have 2+ links in Unix
		}

		owner, group := getFileOwner(name)

		// Format size
		size := file.Size()
		sizeStr := ""
		if humanReadable {
			sizeStr = FormatFileSize(size, true)
		} else {
			sizeStr = strconv.FormatInt(size, 10)
		}

		// Format modification time
		modTime := FormatModTime(file.ModTime())

		// Get color and icon for file
		ext := strings.ToLower(filepath.Ext(name))
		isDir := file.IsDir()
		colorCode, icon := getFileTypeColorAndIcon(file, ext, isDir)

		// Format display name
		displayName := name
		if isDir {
			displayName += "/"
		}

		// Print in Unix-like format without extra newlines
		fmt.Printf("%s %2d %-*s %-*s %*s %s %s%s %s%s%s\n",
			perms,
			links,
			maxOwnerLen, owner,
			maxGroupLen, group,
			maxSizeLen, sizeStr,
			modTime,
			colorCode, icon,
			Color["white"], displayName, Color["reset"],
		)
	}
}
