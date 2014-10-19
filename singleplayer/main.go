package main

import (
	"bufio"
	"fmt"
	"github.com/michd/magestro-zone/singleplayer/player"
	"github.com/michd/magestro-zone/singleplayer/things"
	"os"
	"strings"
)

// Temporary function to generate an example area to put the player into
func prepareArea() *things.Area {
	toilet := new(things.Item)
	toilet.SetName("toilet")
	toilet.SetDesc("I can pee in this.")

	sink := new(things.Item)
	sink.SetName("sink")
	sink.SetDesc("I can wash my hands in there after peeing.")

	area := new(things.Area)
	area.SetName("room")
	area.SetDesc("A generic looking bathroom.")
	area.AddItem(toilet)
	area.AddItem(sink)

	return area
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
