package player

import (
	"github.com/michd/magestro-zone/singleplayer/things"
	"os"
)

type Player struct {
	name string
	desc string
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

func (player *Player) Desc() string {
	return player.desc
}

func (player *Player) SetDesc(newDesc string) {
	player.desc = newDesc
}

func (player *Player) Examine() string {
	return "I am " + player.Name() + ". Just another tiny collection of atoms in this vast universe."
}

func (player *Player) Populate(rawMap map[string]interface{}) {
	if _, ok := rawMap["name"]; ok {
		player.name, _ = rawMap["name"].(string)
	}

	if _, ok := rawMap["desc"]; ok {
		player.desc, _ = rawMap["desc"].(string)
	}

	if _, ok := rawMap["hp"]; ok {
		player.hp, _ = rawMap["hp"].(int)
	}

	if _, ok := rawMap["str"]; ok {
		player.str, _ = rawMap["str"].(int)
	}

	if _, ok := rawMap["per"]; ok {
		player.per, _ = rawMap["per"].(int)
	}

	if _, ok := rawMap["end"]; ok {
		player.end, _ = rawMap["end"].(int)
	}

	if _, ok := rawMap["char"]; ok {
		player.char, _ = rawMap["char"].(int)
	}

	if _, ok := rawMap["sma"]; ok {
		player.sma, _ = rawMap["sma"].(int)
	}

	if _, ok := rawMap["agi"]; ok {
		player.agi, _ = rawMap["agi"].(int)
	}

	if _, ok := rawMap["luc"]; ok {
		player.luc, _ = rawMap["luc"].(int)
	}
}

func (p *Player) GetStat(statName string) int {
	switch statName {
	case "hp":
		return p.hp
	case "str":
		return p.str
	case "per":
		return p.per
	case "end":
		return p.end
	case "char":
		return p.char
	case "sma":
		return p.sma
	case "agi":
		return p.agi
	case "luc":
		return p.luc
	}
	return nil
}

func (p *Player) SetStat(statName string, v int) {
	switch statName {
	case "hp":
		p.hp = v
	case "str":
		p.str = v
	case "per":
		p.str = v
	case "end":
		p.end = v
	case "char":
		p.char = v
	case "sma":
		p.sma = v
	case "agi":
		p.agi = v
	case "luc":
		p.luc = v
	}
}

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
			return player.Examine()
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
