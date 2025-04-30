// file_utils.go

package main

import (
	"fmt"
	"os"
	"time"
)

// Constants for formatting output
const (
	KB = 1024
	MB = 1024 * KB
	GB = 1024 * MB
	TB = 1024 * GB
)

// FormatFileSize formats file size based on flags
func FormatFileSize(size int64, humanReadable bool) string {
	if !humanReadable {
		return fmt.Sprintf("%d", size)
	}

	// Human-readable format
	switch {
	case size >= TB:
		return fmt.Sprintf("%.1fT", float64(size)/float64(TB))
	case size >= GB:
		return fmt.Sprintf("%.1fG", float64(size)/float64(GB))
	case size >= MB:
		return fmt.Sprintf("%.1fM", float64(size)/float64(MB))
	case size >= KB:
		return fmt.Sprintf("%.1fK", float64(size)/float64(KB))
	default:
		return fmt.Sprintf("%dB", size)
	}
}

// FormatPermissions formats the file permissions as a string
func FormatPermissions(info os.FileInfo) string {
	mode := info.Mode()
	perms := ""
	
	// File type
	if mode.IsDir() {
		perms += "d"
	} else if mode&os.ModeSymlink != 0 {
		perms += "l"
	} else {
		perms += "-"
	}
	
	// User permissions
	perms += formatPermissionBits(mode, 0400, 0200, 0100)
	
	// Group permissions
	perms += formatPermissionBits(mode, 040, 020, 010)
	
	// Other permissions
	perms += formatPermissionBits(mode, 04, 02, 01)
	
	return perms
}

// formatPermissionBits helps format permission bits for a specific user category
func formatPermissionBits(mode os.FileMode, r, w, x os.FileMode) string {
	result := ""
	if mode&r != 0 {
		result += "r"
	} else {
		result += "-"
	}
	
	if mode&w != 0 {
		result += "w"
	} else {
		result += "-"
	}
	
	if mode&x != 0 {
		result += "x"
	} else {
		result += "-"
	}
	
	return result
}

// FormatModTime formats the modification time of a file
func FormatModTime(modTime time.Time) string {
	// Recent files show time, older files show year
	if time.Since(modTime) > 6*30*24*time.Hour { // Older than ~6 months
		return modTime.Format("Jan _2  2006")
	}
	return modTime.Format("Jan _2 15:04")
}
