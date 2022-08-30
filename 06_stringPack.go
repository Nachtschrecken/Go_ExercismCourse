package main

import "strings"

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	var welcomeMessage = "Welcome to the Tech Palace, " + strings.ToUpper(customer)
	return welcomeMessage
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	var addBorder = strings.Repeat("*", numStarsPerLine) + "\n" + welcomeMsg + "\n" + strings.Repeat("*", numStarsPerLine)
	return addBorder
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	var cleanupMessage = strings.ReplaceAll(oldMsg, "*", "")
	cleanupMessage = strings.TrimSpace(cleanupMessage)
	return cleanupMessage
}
