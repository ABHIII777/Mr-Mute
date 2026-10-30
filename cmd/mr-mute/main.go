package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	var path string
	
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the path of media file: ")
	path, _ = reader.ReadString('\n')
	path = strings.TrimSpace(path)

	if path == "" {
		log.Fatal("Path cannot be left empty. Exiting...")
		return
	}

	cmd := exec.Command("mpv", path)
	fmt.Print("Starting the media file. Please wait :)")

	err := cmd.Start()

	if err != nil {
		log.Fatal(err)

	}
	cmd.Wait()
}
