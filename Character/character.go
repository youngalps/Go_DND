package character

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	classes "github.com/youngalps/Go_DND/Classes"
)

type Character struct {
	Name         string
	Class        string
	Level        int
	Strength     int
	Dexterity    int
	Constitution int
	Intelligence int
	Wisdom       int
	Charisma     int
}

func (c *Character) SetCharacterName(name string) {
	c.Name = name
}

func (c *Character) SetCharacterClass(class string) {
	c.Class = class
}

func (c *Character) SetLevel(level int) {
	c.Level = level
}

// TODO: we are gonna make this done by roll
func (c *Character) SetStats(strength, dexterity, constitution, intelligence, wisdom, charisma int) {
	c.Strength = strength
	c.Dexterity = dexterity
	c.Constitution = constitution
	c.Intelligence = intelligence
	c.Wisdom = wisdom
	c.Charisma = charisma
}

func NewCharacter() *Character {
	character := &Character{}                     // create a new character
	availableClasses, err := classes.GetClasses() // available classes
	if err != nil {
		fmt.Println("Error getting classes")
	}

	character.SetLevel(1)

	// read user input
	reader := bufio.NewReader(os.Stdin)

	// *Player Name
	fmt.Print("Enter Name: ")
	name, _ := reader.ReadString('\n')
	if name == "" {
		fmt.Println("You must enter a name")
	} else {
		name = name[:len(name)-1]
		character.Name = name
	}

	//*Region Start
	//* Class Selection
	fmt.Println("Available Classes:")
	for i, class := range availableClasses {
		fmt.Printf("%d: %s\n", i+1, class.Name)
	}
	// prompt for class until valid one is chosen
	for {
		fmt.Print("Choose Class(enter the number): ")
		classChoiceStr, _ := reader.ReadString('\n')
		classChoiceStr = strings.TrimSpace(classChoiceStr)
		classChoice, err := strconv.Atoi(classChoiceStr)
		if err != nil {
			fmt.Println("Invalid input. Please enter a number")
			continue
		}

		if classChoice < 1 || classChoice > len(availableClasses) {
			fmt.Println("Invalid number. Please choose a number within the range.")
			continue
		}
		character.Class = availableClasses[classChoice-1].Index
		break
	}
	//*End Region Class Selection
	return character

}
