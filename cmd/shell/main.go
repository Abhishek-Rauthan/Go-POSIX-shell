package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Fprint

func main() {
	// Uncomment this block to pass the first stage
	for {

		fmt.Fprint(os.Stdout, "$ ")

		message, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input:", err)
		}

		message = strings.TrimSpace(message)

		commands := strings.Split(message, " ")

		var tokens []string
		for {
			start := strings.IndexAny(message, "'\"")
			if start == -1 {
				tokens = append(tokens, strings.Fields(message)...)
				break
			}
			ch := message[start]
			tokens = append(tokens, strings.Fields(message[:start])...)
			message = message[start+1:]
			end := strings.IndexByte(message, ch)
			token := message[:end]
			tokens = append(tokens, token)
			message = message[end+1:]
		}

		switch commands[0] {
		case "exit":
			code, err := strconv.Atoi(tokens[1])
			if err != nil {
				os.Exit(1)
			}
			os.Exit(code)

		case "pwd":
			pwd, _ := os.Getwd()
			fmt.Println(pwd)

		case "cd":
			if strings.TrimSpace(tokens[1]) == "~" {
				tokens[1] = os.Getenv("HOME")
			}
			if err := os.Chdir(tokens[1]); err != nil {
				fmt.Fprintf(os.Stdout, "%s: No such file or directory\n", tokens[1])
			}

		case "echo":
			fmt.Fprintf(os.Stdout, "%s\n", strings.Join(tokens[1:], " "))

		case "type":
			if tokens[1] == "echo" || tokens[1] == "exit" || tokens[1] == "type" || tokens[1] == "pwd" || tokens[1] == "cd" {
				fmt.Println(tokens[1] + " is a shell builtin")
			} else {
				fp, err := func() (string, error) {

					paths := strings.Split(os.Getenv("PATH"), ":")
					for _, path := range paths {
						fp := filepath.Join(path, tokens[1])
						if _, err := os.Stat(fp); err == nil {
							return fp, nil
						}
					}

					return "", errors.New("error happened")
				}()

				if err != nil {

					fmt.Println(tokens[1] + ": not found")
				} else {
					fmt.Println(fp)
				}
			}

		default:
			command := exec.Command(tokens[0], tokens[1:]...)
			command.Stderr = os.Stderr
			command.Stdout = os.Stdout
			err := command.Run()
			if err != nil {
				fmt.Printf("%s: command not found\n", tokens[0])
			}
		}
	}
}
