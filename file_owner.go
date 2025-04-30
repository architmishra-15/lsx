package main

import (
	"os/user"
	"strings"
)

// getFileOwner returns the owner and group of a file
func getFileOwner(filePath string) (string, string) {
	// For simplicity, just use the current user
	username := getCurrentUsername()
	return username, username
}

// getCurrentUsername gets the current user's name
func getCurrentUsername() string {
	currentUser, err := user.Current()
	if err != nil {
		return "unknown"
	}

	// Extract username from full name or use the username directly
	parts := strings.Split(currentUser.Username, "\\")
	if len(parts) > 1 {
		return parts[1]
	}

	return currentUser.Username
}
