package main

import (
	"fmt"
	"os/exec"
)

func main() {
	title := "Hello, World!"
	message := "This is a notification from Go!"
	script := fmt.Sprintf(`display notification "%s" with title "%s"`, message, title)

	cmd := exec.Command("osascript", "-e", script)
	err := cmd.Run()
	if err != nil {
		fmt.Println("Failed to send notification:", err)
	}
}
