package player

import (
	"strings"
)

type Command struct {
	verb    string
	subject string
}

// Parse a string to a Command
// Taking suggestions for better name as this will be
// player.ParseCommand from the outside, and it doesn't take a
// Command as argument.
func ParseCommand(rawCommand string) Command {
	words := strings.Split(
		strings.ToLower(strings.Trim(rawCommand, " \n\t\r")), " ")
	command := Command{}

	// TODO: much more flexible command parsing
	if len(words) > 0 {
		command.verb = strings.ToLower(words[0])
	}

	if len(words) > 1 {
		command.subject = words[1]
	}

	return command
}
