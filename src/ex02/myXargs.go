package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	command := os.Args[1]
	args := os.Args[2:]

	var dir []string

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		text := scanner.Text()
		dir = append(dir, text)
	}

	for _, text := range dir {
		directory := strings.ReplaceAll(text, " ", "")
		cmd := exec.Command(command, append(args, directory)...)
		//fmt.Println("cmd", cmd)
		out, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf(string(out))
	}

}
