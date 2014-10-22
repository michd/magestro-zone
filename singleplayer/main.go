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
	areaFile := "content/areas/testbathroom.json"
	if area, err := things.AreaFromFile(areaFile); err == nil {
		return area
	} else {
		fmt.Println("Error loading area: " + err.Error())
		return nil
	}
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
