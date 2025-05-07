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
	remaining := []string{args[0]}

	for i := 1; i < len(args); i++ {
		arg := args[i]

		// Handle combined flags (like -la or -lh)
		if strings.HasPrefix(arg, "-") && !strings.HasPrefix(arg, "--") && len(arg) > 2 {
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
			remaining = append(remaining, arg)
		}
	}

	return flags, remaining
}

// Process flags and execute relevant commands
func handle_flag(args []string) (Flags, []string) {
	flags, remaining := parseFlags(args)

	if flags.Help {
		showHelp()
		os.Exit(0)
	}

	if flags.Version {
		show_version()
		os.Exit(0)
	}

	return flags, remaining
}

// Display version information
func show_version() {
	fmt.Println("Vesion 1.1.0")
	fmt.Println("󰗦 2025 Archit Mishra")
}
