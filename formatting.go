// formatting.go

package main

import (
	"fmt"
	"math"
	"os"
	"strings"
	"unicode/utf8"
)

// Maximum width for a filename before wrapping
const maxFilenameWidth = 20

// print a string slice in the specified number of columns with improved spacing
func PrintInColumns(items []string, numColumns int) {
	if len(items) == 0 {
		return
	}

	// Calculate how many rows are needed
	numRows := int(math.Ceil(float64(len(items)) / float64(numColumns)))

	// Find the maximum width needed for each column
	columnWidths := make([]int, numColumns)
	for col := 0; col < numColumns; col++ {
		for row := 0; row < numRows; row++ {
			idx := row + col*numRows
			if idx < len(items) {
				// Get the display width of the string (handling multi-byte characters)
				displayWidth := utf8.RuneCountInString(items[idx])
				if displayWidth > maxFilenameWidth {
					displayWidth = maxFilenameWidth
				}
				if displayWidth > columnWidths[col] {
					columnWidths[col] = displayWidth
				}
			}
		}
	}

	for row := 0; row < numRows; row++ {
		// Track if any item in this row needs an extra line
		needsExtraLine := false
		rowContent := make([]string, numColumns)
		overflowContent := make([]string, numColumns)

		for col := 0; col < numColumns; col++ {
			idx := row + col*numRows
			if idx < len(items) {
				item := items[idx]

				// Handle long filenames
				if utf8.RuneCountInString(item) > maxFilenameWidth {
					needsExtraLine = true
					displayPart := truncateString(item, maxFilenameWidth)
					overflowPart := item[len(displayPart):]

					rowContent[col] = displayPart
					overflowContent[col] = overflowPart
				} else {
					rowContent[col] = item
					overflowContent[col] = ""
				}
			}
		}

		// Print main content for this row
		for col := 0; col < numColumns; col++ {
			if col < len(rowContent) && rowContent[col] != "" {
				item := rowContent[col]

				// Add padding for all columns except the last one
				if col < numColumns-1 {
					padding := columnWidths[col] - utf8.RuneCountInString(item) + 4
					fmt.Print(item + strings.Repeat(" ", padding))
				} else {
					fmt.Print(item)
				}
			}
		}
		fmt.Print("\n")

		// If there's overflow content, print it on the next line
		if needsExtraLine {
			for col := 0; col < numColumns; col++ {
				if col < len(overflowContent) && overflowContent[col] != "" {
					// Truncate again if it's still too long
					overflow := overflowContent[col]
					if utf8.RuneCountInString(overflow) > maxFilenameWidth {
						overflow = truncateString(overflow, maxFilenameWidth-3) + "..."
					}

					// Add padding similar to main content
					if col < numColumns-1 {
						padding := columnWidths[col] - utf8.RuneCountInString(overflow) + 4
						fmt.Print(overflow + strings.Repeat(" ", padding))
					} else {
						fmt.Print(overflow)
					}
				} else if col < numColumns-1 {
					// Empty cell but need padding
					fmt.Print(strings.Repeat(" ", columnWidths[col]+4))
				}
			}
			fmt.Print("\n")
		}

		// End of row - add another line break for spacing between rows
		fmt.Print("\n")
	}
}

// determines appropriate icon and color based on file type
func getFileTypeColorAndIcon(file os.FileInfo, ext string, isDir bool) (string, string) {

	if isDir {
		return Color["blue"], Icons["folder"]
	}

	// Executable handling
	if file.Mode()&0111 != 0 {
		return Color["bright_green"], Icons["executable"]
	}

	name := file.Name()
	switch name {
	case "go.mod", "go.sum":
		return Color["cyan"], Icons["go.mod"]
	case "package.json":
		return Color["bright_green"] + Color["bold"], Icons["package.json"]
	case "tailwind.config.js", "tailwind.config.ts":
		return Color["bright_blue"], Icons["tailwind"]
	case "vue.config.js":
		return Color["bright_blue"], Icons["vue.config.js"]
	case ".eslintrc.js", ".eslintrc.json", ".eslintrc.yml", ".eslintrc.yaml":
		return Color["bright_blue"], Icons["eslint"]
	}
	// Check for a direct icon match
	if icon, exists := Icons[ext]; exists {
		return getColorForFileType(ext), icon
	}

	return Color["dim"], Icons["default"]
}

func getColorForFileType(ext string) string {
	switch {
	// Programming languages
	case isInList(ext, []string{".asm", ".go", ".py", ".js", ".ts", ".jsx", ".tsx", ".java", ".c", ".o", ".cpp", ".cs",
		".rb", ".php", ".swift", ".rs", ".dart", ".kt", ".scala", ".ex", ".exs", ".hs", ".pl", ".r", ".coffee", ".d", ".m", ".mat", ".ps1", ".jil", ".lua", ".ml", ".f90", ".vim", ".vimrc", ".bash", ".bashrc", ".zsh", ".zshrc", ".ads", ".cbl", ".sql"}):
		// Different colors for different programming languages
		switch ext {
		case ".java", ".go", ".jsx":
			return Color["cyan"]
		case ".py", ".asm", ".dart":
			return Color["bright_blue"]
		case ".rs":
			return Color["yellow"]
		case ".js", ".lock":
			return Color["bright_yellow"]
		case ".ts", ".tsx":
			return Color["blue"]
		case ".kt":
			return Color["red"]
		case ".c", ".cpp", ".cs":
			return Color["bright_green"]
		case ".rb":
			return Color["bright_red"]
		case ".php":
			return Color["bright_magenta"]
		case ".ex", ".exs":
			return Color["cyan"] // Elixir
		case ".hs":
			return Color["yellow"] // Haskell
		case ".pl":
			return Color["bright_cyan"] // Perl
		case ".r":
			return Color["green"] // R
		case ".scala":
			return Color["bright_magenta"]
		case ".sh", ".bash", ".zsh":
			return Color["green"]
		case ".lua":
			return Color["blue"]
		case ".d":
			return Color["magenta"]
		case ".m", ".mat", ".cbl":
			return Color["bright_yellow"]
		case ".ps1":
			return Color["bright_blue"]
		case ".jil":
			return Color["red"]
		case ".ml":
			return Color["yellow"]
		case ".f90":
			return Color["cyan"]
		case ".vim", ".vimrc":
			return Color["green"]
		case ".ads":
			return Color["bright_magenta"]
		case ".sql":
			return Color["bright_cyan"]
		default:
			return Color["bright_white"]
		}

	// Web technologies
	case isInList(ext, []string{".html", ".css", ".scss", ".sass", ".less", ".vue"}):
		switch ext {
		case ".html":
			return Color["yellow"] + Color["bold"]
		case ".css", ".scss", ".sass", ".less":
			return Color["bright_magenta"]
		default:
			return Color["yellow"]
		}

	// Data formats
	case isInList(ext, []string{".json", ".xml", ".yaml", ".yml", ".toml", ".csv"}):
		return Color["bright_yellow"]

	// Documents and text
	case isInList(ext, []string{".txt", ".md", ".pdf", ".doc", ".docx", ".odt", ".xls", ".xlsx", ".ppt", ".pptx", ".log", ".ipynb"}):
		switch ext {
		case ".pdf", ".doc", ".docx", ".odt":
			return Color["bright_red"]
		case ".xls", ".xlsx", ".ods":
			return Color["bright_green"]
		case ".ppt", ".pptx", ".odp":
			return Color["bright_yellow"]
		case ".log":
			return Color["cyan"]
		case ".ipynb":
			return Color["yellow"] + Color["bold"]
		default:
			return Color["white"]
		}

	// Media files
	case isInList(ext, []string{".jpg", ".jpeg", ".png", ".gif", ".bmp", ".svg", ".webp"}):
		return Color["bright_cyan"]

	case isInList(ext, []string{".mp3", ".wav", ".ogg", ".flac", ".aac", ".m4a", ".epub", ".mobi"}):
		return Color["magenta"]

	case isInList(ext, []string{".mp4", ".avi", ".mkv", ".mov", ".wmv", ".flv", ".webm"}):
		return Color["bright_magenta"]

	// Archives
	case isInList(ext, []string{".zip", ".tar", ".gz", ".bz2", ".7z", ".rar", ".xz"}):
		return Color["red"]

	// Scripts and configs
	case isInList(ext, []string{".sh", ".bash", ".zsh", ".fish", ".conf", ".cfg", ".ini", ".dockerfile", ".makefile", ".cmake", ".gitignore", ".gitattributes"}):
		return Color["green"]

	// Default for unknown types
	default:
		return Color["dim"]
	}
}

// Helper functions

// isInList checks if a string is in a slice of strings
func isInList(item string, list []string) bool {
	for _, i := range list {
		if i == item {
			return true
		}
	}
	return false
}

// safely truncates a string to the given width, respecting UTF-8 characters
func truncateString(s string, maxWidth int) string {
	if utf8.RuneCountInString(s) <= maxWidth {
		return s
	}

	runes := []rune(s)
	return string(runes[:maxWidth])
}
