package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Command struct {
	Verb    string
	Subject string
}

func processCommand(command Command) {
	switch strings.ToLower(command.Verb) {
	case "examine":
		switch strings.ToLower(command.Subject) {
		case "room":
			fmt.Println("The room is dark an empty. You are alone and there is no way out.")
			fmt.Println("Game over.")
			os.Exit(0)
		case "":
			fmt.Println("Examine what?")
		default:
			fmt.Println("I see no", command.Subject)
		}
	case "":
		fmt.Println("uwotm8")
	default:
		fmt.Println("I have no idea how to", command.Verb)
	}
}

func main() {
	input := bufio.NewReader(os.Stdin)

	fmt.Println("Welcome to Magestro zone!")
	fmt.Println("You find yourself in a room.")

	for {
		var words []string
		fmt.Print("> ")
		line, _ := input.ReadString('\n')
		words = strings.Split(strings.Trim(line, " \n\t\r"), " ")

		command := new(Command)

		if len(words) > 0 {
			command.Verb = words[0]
		}

		if len(words) > 1 {
			command.Subject = words[1]
		}

		if strings.ToLower(command.Verb) == "exit" {
			break
		}

		processCommand(*command)

	}
	fmt.Println("Later!")
}
