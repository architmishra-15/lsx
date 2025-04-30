// file_printer.go

package main

import (
	"fmt"
	"math"
	"os"
	"path/filepath"
	"strings"
	"unicode/utf8"
)

// check if a file is a framework config file
func isConfigFile(name string) string {
	switch name {
	case "tailwind.config.js":
		return "tailwind"
	case "vue.config.js":
		return "vue"
	case "vite.config.js", "vite.config.ts":
		return "vite"
	case "next.config.js", "next.config.ts":
		return "nextjs"
	case ".eslintrc.js", ".eslintrc.json", ".eslintrc.yml", ".eslintrc.yaml":
		return "eslint"
	default:
		return ""
	}
}

// displays file information with icons and type-based colors in a multi-column layout
func PrintFilesInColumns(files []os.FileInfo, numColumns int, flags ...Flags) {
	if len(files) == 0 {
		return
	}

	// Check if we have flag info
	var flagsInfo Flags
	if len(flags) > 0 {
		flagsInfo = flags[0]
	}

	// If using long format, handle differently
	if flagsInfo.LongFormat {
		printLongFormat(files, flagsInfo.HumanReadable)
		return
	}

	// Calculate how many rows we need
	numRows := int(math.Ceil(float64(len(files)) / float64(numColumns)))

	// Find the maximum filename width needed for each column, respecting maxFilenameWidth
	columnWidths := make([]int, numColumns)
	for col := 0; col < numColumns; col++ {
		for row := 0; row < numRows; row++ {
			idx := row + col*numRows
			if idx < len(files) {
				// Account for the file name length
				nameWidth := utf8.RuneCountInString(files[idx].Name())
				if nameWidth > maxFilenameWidth {
					nameWidth = maxFilenameWidth
				}
				if nameWidth > columnWidths[col] {
					columnWidths[col] = nameWidth
				}
			}
		}
	}

	// Print files in rows with icons and colors
	for row := 0; row < numRows; row++ {
		// Track if any file in this row has a long name
		hasLongName := false
		mainContent := make([]string, numColumns)
		overflowContent := make([]string, numColumns)
		iconColors := make([]string, numColumns)
		icons := make([]string, numColumns)

		// Prepare content for this row
		for col := 0; col < numColumns; col++ {
			idx := row + col*numRows
			if idx < len(files) {
				file := files[idx]
				name := file.Name()
				ext := strings.ToLower(filepath.Ext(name))
				isDir := file.IsDir()

				// Check for framework config files
				frameworkIcon := isConfigFile(name)
				if frameworkIcon != "" {
					iconColors[col] = Color["yellow"] // Or any color you want for framework icons
					icons[col] = Icons[frameworkIcon]
				} else {
					// Get icon and color based on file type
					colorCode, icon := getFileTypeColorAndIcon(file, ext, isDir)
					iconColors[col] = colorCode
					icons[col] = icon
				}

				// Handle ".git" directory specifically
				if isDir && name == ".git" {
					icons[col] = Icons["git_folder"]
				}

				// Handle .gitignore and .gitattributes as dotfiles
				if name == ".gitignore" || name == ".gitattributes" {
					icons[col] = Icons[name] // Assuming Icons[".gitignore"] and Icons[".gitattributes"] exist
				}

				// Handle long filenames
				if utf8.RuneCountInString(name) > maxFilenameWidth {
					hasLongName = true
					displayPart := truncateString(name, maxFilenameWidth)
					overflowPart := name[len(displayPart):]

					mainContent[col] = displayPart
					overflowContent[col] = overflowPart
				} else {
					mainContent[col] = name
					overflowContent[col] = ""
				}
			}
		}

		// Print main content row
		for col := 0; col < numColumns; col++ {
			idx := row + col*numRows
			if col < len(mainContent) && mainContent[col] != "" {
				colorCode := iconColors[col]
				icon := icons[col]
				name := mainContent[col]

				// Special formatting for directories - make them bold and underlined and add "/"
				if idx < len(files) && files[idx].IsDir() {
					name += "/"
				}

				// File display with colored icon and white filename
				fileDisplay := fmt.Sprintf("%s%s %s%s",
					colorCode,
					icon,
					Color["white"]+name, // Filename in white
					Color["reset"],
				)

				// Add padding for all columns except the last one
				if col < numColumns-1 {
					padding := columnWidths[col] - utf8.RuneCountInString(name) + 4
					fmt.Print(fileDisplay + strings.Repeat(" ", padding))
				} else {
					fmt.Print(fileDisplay)
				}
			}
		}
		fmt.Print("\n")

		// If there are long names, print overflow content
		if hasLongName {
			for col := 0; col < numColumns; col++ {
				if col < len(overflowContent) && overflowContent[col] != "" {
					colorCode := iconColors[col]
					overflow := overflowContent[col]

					// Preserve bold and underline for directories in overflow
					idx := row + col*numRows
					if idx < len(files) && files[idx].IsDir() {
						overflow = "/" + overflow // Add "/" for directory overflow
					}

					// Truncate if still too long
					if utf8.RuneCountInString(overflow) > maxFilenameWidth {
						overflow = truncateString(overflow, maxFilenameWidth-3) + "..."
					}

					// Indented and colored continuation
					overflowDisplay := fmt.Sprintf("%s  %s%s",
						colorCode,
						overflow,
						Color["reset"],
					)

					// Add padding for all columns except the last one
					if col < numColumns-1 {
						padding := columnWidths[col] - utf8.RuneCountInString(overflow) + 4
						fmt.Print(overflowDisplay + strings.Repeat(" ", padding))
					} else {
						fmt.Print(overflowDisplay)
					}
				} else if col < numColumns-1 {
					// Empty continuation line, but maintain column width
					fmt.Print(strings.Repeat(" ", columnWidths[col]+4))
				}
			}
			fmt.Print("\n")
		}
	}
}
