package main

import (
	"io"
	"log"
	"os"
	"os/exec"

	"github.com/kr/pty"
)

func main() {
	var name string
	var args []string

	if len(os.Args) == 1 {
		log.Fatal("Must provide at least a command")
	}
	if len(os.Args) > 2 {
		args = os.Args[2:]
	}
	name = os.Args[1]

	cmd := exec.Command(name, args...)
	p, err := pty.Start(cmd)
	if err != nil {
		log.Fatalf("Could not run command: %s", err)
	}
	defer p.Close()

	go io.Copy(os.Stdout, p)
	cmd.Wait()
}
