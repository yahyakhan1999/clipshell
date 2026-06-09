package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: clipshell <command>")
		os.Exit(1)
	}

	commandText := strings.Join(os.Args[1:], " ")

	cmd := exec.Command("bash", "-c", commandText)

	output, _ := cmd.CombinedOutput()

	result := fmt.Sprintf("> %s\n\n%s", commandText, string(output))

	clip := exec.Command("clip.exe")
	clip.Stdin = strings.NewReader(result)
	clip.Run()

	fmt.Print(string(output))
}
