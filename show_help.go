// show_help.go

package main

import "fmt"

func showHelp() {
	fmt.Println("Usage: myls [options] [path]")
	fmt.Println("Options:")
	fmt.Println("  -h, --help  Show this help message")
	fmt.Println("  [path]      Path to list (default: current directory)")
	fmt.Println()
	fmt.Println("Pattern matching:")
	fmt.Println("  *.extension             List files with extension")
	fmt.Println("  path/*.extension        List files with extension in path")
	fmt.Println("  path/filename           Show details for filename")
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println("  myls")
	fmt.Println("  myls /home/user/documents")
	fmt.Println("  myls *.txt")
	fmt.Println("  myls /var/log/*.log")
	fmt.Println("  myls /etc/passwd")
}
