// list_view.go
package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// PrintListView displays files in Unix-like ls -l format
func PrintListView(files []os.FileInfo, flags Flags) {
	// Calculate total blocks (typically in 1K blocks on Unix)
	totalBlocks := calculateTotalBlocks(files)

	// Print total line (Unix ls compatibility)
	if flags.HumanReadable {
		fmt.Printf("total %s\n", FormatFileSize(totalBlocks*1024, true))
	} else {
		fmt.Printf("total %d\n", totalBlocks)
	}

	// Find the maximum length for size field to ensure alignment
	maxSizeLen := findMaxSizeLength(files, flags.HumanReadable)

	// Now print each file in long format
	for _, file := range files {
		printFileListView(file, maxSizeLen, flags.HumanReadable)
	}
}

// calculateTotalBlocks calculates the total blocks used by files
// On Unix, this is typically in 1K blocks
func calculateTotalBlocks(files []os.FileInfo) int64 {
	var totalSize int64
	for _, file := range files {
		totalSize += file.Size()
	}

	// Convert to 1K blocks (rounded up)
	return (totalSize + 1023) / 1024
}

// findMaxSizeLength finds the maximum string length of file sizes
func findMaxSizeLength(files []os.FileInfo, humanReadable bool) int {
	maxLen := 0
	for _, file := range files {
		size := file.Size()
		sizeStr := ""

		if humanReadable {
			sizeStr = FormatFileSize(size, true)
		} else {
			sizeStr = strconv.FormatInt(size, 10)
		}

		if len(sizeStr) > maxLen {
			maxLen = len(sizeStr)
		}
	}
	return maxLen
}

// getOwnerAndGroup attempts to get the owner and group of a file
func getOwnerAndGroup(file os.FileInfo) (string, string) {
	// For simplicity on Windows, just return generic owner/group
	// On Unix-like systems, you would extract this from the stat structure
	return "user", "group"
}

// getLinkCount gets the number of hard links (simplified for Windows)
func getLinkCount(file os.FileInfo) int {
	// On Windows, this concept doesn't directly map the same as in Unix
	// For directories, we'll show 2+ to match Unix behavior
	if file.IsDir() {
		return 2 // Simplified - in reality would count subdirectories
	}
	return 1 // Regular files typically have 1 link in Windows
}

// printFileListView prints a single file in ls -l format
func printFileListView(file os.FileInfo, maxSizeLen int, humanReadable bool) {
	// Debug output
	// fmt.Println("Debug: PrintListView is being called!")
	// Format permissions - For Windows, we'll simulate Unix-style permissions
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
	} else if strings.HasSuffix(strings.ToLower(file.Name()), ".exe") {
		perms += "rwxr-xr--" // executables
	} else {
		perms += "rw-r--r--" // regular files
	}

	// Get link count
	linkCount := getLinkCount(file)

	// Get owner and group
	owner, group := getOwnerAndGroup(file)

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

	// Get file name with appropriate styling
	name := file.Name()
	ext := strings.ToLower(filepath.Ext(name))
	isDir := file.IsDir()
	colorCode, icon := getFileTypeColorAndIcon(file, ext, isDir)

	// Add trailing slash for directories
	displayName := name
	if isDir {
		displayName += "/"
	}

	// Format final output line with columns aligned
	fmt.Printf("%s %2d %8s %-8s %*s %s %s%s %s%s%s\n",
		perms,
		linkCount,
		owner,
		group,
		maxSizeLen, sizeStr,
		modTime,
		colorCode, icon,
		Color["white"], displayName, Color["reset"],
	)
}
