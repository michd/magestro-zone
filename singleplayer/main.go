package main

import (
	"bufio"
	"fmt"
	"github.com/michd/magestro-zone/singleplayer/player"
	"github.com/michd/magestro-zone/singleplayer/things"
	"os"
	"strings"
)

func prepareArea() *things.Area {
	var jsonBlob = []byte(`{
		"name": "room",
		"desc": "A generic looking bathroom.",
		"items": {
			"toilet": {
				"name": "toilet",
				"desc": "I can pee in this."
			},
			"sink": {
				"name": "sink",
				"desc": "I can wash my hands in there after peeing."
			}
		}
	}`)

	return things.ParseJson(jsonBlob)
}

func main() {
	in := bufio.NewReader(os.Stdin)
	var input, output string

	// Set up player
	fmt.Println("Hi there! What's your name?")
	fmt.Print("> ")
	input, _ = in.ReadString('\n')
	playerName := strings.Trim(input, "\n\r\t ")

	me := player.Create(playerName)
	me.SetLoc(prepareArea())

	fmt.Printf("Hello %s, welcome to Magestro Zone!\n", me.Name())
	fmt.Println("You find yourself in a room.")

	// Command loop
	for {
		fmt.Print(me.Name() + "> ")
		input, _ = in.ReadString('\n')
		output = me.Execute(player.ParseCommand(input))
		fmt.Println(output)
	}
	fmt.Println("Later!")
}
