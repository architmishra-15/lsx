package main

import (
	"fmt"
	"os"
	"strings"
)

// Flags structure to hold command line flags
type Flags struct {
	LongFormat    bool // -l flag
	AllFiles      bool // -a flag
	DirectoryOnly bool // -d or --directory flag
	HumanReadable bool // -h or --human-readable flag
	Help          bool // --help flag

	Version bool // -v or --version flag
}

// Parse command line arguments and return a Flags structure
func parseFlags(args []string) (Flags, []string) {
	flags := Flags{}
	remaining := []string{args[0]} // Keep program name

	// Skip program name (args[0])
	for i := 1; i < len(args); i++ {
		arg := args[i]

		// Handle combined flags (e.g., -la)
		if strings.HasPrefix(arg, "-") && !strings.HasPrefix(arg, "--") && len(arg) > 2 {
			// Process each character in the combined flag
			for _, c := range arg[1:] {
				switch c {
				case 'l':
					flags.LongFormat = true
				case 'a':
					flags.AllFiles = true
				case 'd':
					flags.DirectoryOnly = true
				case 'h':
					flags.HumanReadable = true
				case 'v':
					flags.Version = true
				}
			}
			continue
		}

		// Handle individual flags
		switch arg {
		case "-l":
			flags.LongFormat = true
		case "-a":
			flags.AllFiles = true
		case "-d", "--directory":
			flags.DirectoryOnly = true
		case "-h", "--human-readable":
			flags.HumanReadable = true
		case "--help":
			flags.Help = true
		case "-v", "--version":
			flags.Version = true
		default:
			// If it's not a flag, add it to remaining arguments
			remaining = append(remaining, arg)
		}
	}

	return flags, remaining
}

// Process flags and execute relevant commands
func handle_flag(args []string) (Flags, []string) {
	flags, remaining := parseFlags(args)

	// Handle help and version flags first
	if flags.Help {
		showHelp()
		os.Exit(0)
	}

	if flags.Version {
		show_version()
		os.Exit(0)
	}

	// We're no longer modifying os.Args directly, instead we're returning the remaining args
	// for the caller to use
	return flags, remaining
}

// Using the showHelp function from show_help.go

// Display version information
func show_version() {
	fmt.Println("Vesion 1.0.2")
	fmt.Println("󰗦 2025 Archit Mishra")
}
