package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"time"
)

// ANSI escape codes for colors
const (
	ColorGreen  = "\033[32m"
	ColorYellow = "\033[33m"
	ColorRed    = "\033[31m"
	ColorReset  = "\033[0m"
)

func main() {
	// Validate arguments
	if len(os.Args) < 2 {
		log.Fatal("Usage: ./goping <hostname>")
	}
	host := os.Args[1]

	// Create the command
	cmd := exec.Command("ping", host)

	// Get a pipe to read the command's output
	out, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal("Error creating pipe:", err)
	}

	// Start the command
	if err := cmd.Start(); err != nil {
		log.Fatal("Error starting command:", err)
	}

  timeRegexp := regexp.MustCompile(`time=(\d+(\.\d+)?)`)
	// Read the command's output line by line
	scanner := bufio.NewScanner(out)
	for scanner.Scan() {
		fmt.Println(colorText(scanner.Text(), timeRegexp))
	}

	// Check for errors while scanning
	if err := scanner.Err(); err != nil {
		log.Fatal("Error reading output:", err)
	}

	// Wait for the command to finish
	if err := cmd.Wait(); err != nil {
		log.Fatal("Command finished with error:", err)
	}
}

func colorText(line string, timeRegexp *regexp.Regexp) string {
  match := timeRegexp.FindStringSubmatch(line)
  if len(match) > 1 {
    timeValue, err := strconv.ParseFloat(match[1], 64)
    currentTime := time.Now()
    timeStampLine := currentTime.Format(time.RFC3339) + " :: "
    if err == nil {
      switch {
      case timeValue < 50:
        return timeStampLine + ColorGreen + line + ColorReset
      case timeValue < 150:
        return timeStampLine + ColorYellow + line + ColorReset
      default:
        return timeStampLine + ColorRed + line + ColorReset
      }
    }
  }
  return line
}
