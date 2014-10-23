package player

import (
	"github.com/michd/magestro-zone/singleplayer/things"
	"os"
)

type Player struct {
	name string
	// Player stats - stolen shamelessly form Fallout's SPECIAL system
	hp   int // Hitpoints
	str  int // Strength
	per  int // Perception
	end  int // Endurance
	char int // Charisma
	sma  int // Smarts (intelligence)
	agi  int // Agility
	luc  int // Luck

	loc *things.Area // Room or whatever the player finds themselves in
}

func (player *Player) Name() string {
	return player.name
}

func (player *Player) SetName(newName string) {
	player.name = newName
}

// TODO add a desc field, implement:
//   - Desc()
//   - SetDesc()
//   - Populate()

// TODO getter/setter for stats:
//   - GetStat(statName string) int
//   - SetStat(statName string, newValue int)

func (player Player) quit() {
	// TODO: any state saving here
	os.Exit(0)
}

// Find a thing in the context of this player
// Thing can be the area the player is in or something in the area.
// TODO: once player has inventory, look in inventory too
func (player *Player) find(rawThingName string) interface{} {
	if player.loc.Name() == rawThingName {
		return player.loc
	}

	if player.loc.Has(rawThingName) {
		return player.loc.Item(rawThingName)
	}

	return nil
}

func (player *Player) Loc() *things.Area {
	return player.loc
}

func (player *Player) SetLoc(area *things.Area) {
	player.loc = area
}

func (player *Player) Execute(command Command) (output string) {
	switch command.verb {
	case "examine":
		// Check if we can examine it
		if command.subject == "" {
			return "Examine what?"
		}

		if command.subject == "me" {
			return "I am " + player.Name() + ". Just another tiny collection of atoms in this vast universe."
		}

		thing := player.find(command.subject)

		if thing == nil {
			return "I see no " + command.subject + "."
		}

		if examinableThing, found := thing.(things.Examinable); found {
			return examinableThing.Examine()
		}

		return "This is not something I can examine."

	case "help":
		return "Type \"examine <thing>\" to learn more about a thing.\nType \"exit\" or \"quit\" to exit."

	case "exit", "quit":
		player.quit()
	}

	return "I don't understand what you want me to do."
}

// Create a new player given a name
func Create(name string) (newPlayer *Player) {
	newPlayer = new(Player)

	if name == "" {
		newPlayer.SetName("Altin") // A nod to the Galactic Mage by Jon Daulton
	} else {
		newPlayer.SetName(name)
	}

	// TODO: come up with a bunch of sensible init values for the player's stats

	return
}
