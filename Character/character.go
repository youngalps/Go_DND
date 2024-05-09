package character

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

func NewCharacter(name, class string) *Character {
	return &Character{
		Name:  name,
		Class: class,
	}
}

func (c *Character) SetLevel(level int) {
	c.Level = level
}

func (c *Character) SetStats(strength, dexterity, constitution, intelligence, wisdom, charisma int) {
	c.Strength = strength
	c.Dexterity = dexterity
	c.Constitution = constitution
	c.Intelligence = intelligence
	c.Wisdom = wisdom
	c.Charisma = charisma
}

// TODO move to a util folder
func (c *Character) SetStatsFromInput() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter Strength: ")
	strengthStr, _ := reader.ReadString('\n')
	strength, _ := strconv.Atoi(strengthStr)
	c.Strength = strength

	fmt.Print("Enter Dexterity: ")
	dexterityStr, _ := reader.ReadString('\n')
	dexterity, _ := strconv.Atoi(dexterityStr)
	c.Dexterity = dexterity

	fmt.Print("Enter Constitution: ")
	constitutionStr, _ := reader.ReadString('\n')
	constitution, _ := strconv.Atoi(constitutionStr)
	c.Constitution = constitution

	fmt.Print("Enter Intelligence: ")
	intelligenceStr, _ := reader.ReadString('\n')
	intelligence, _ := strconv.Atoi(intelligenceStr)
	c.Intelligence = intelligence

	fmt.Print("Enter Wisdom: ")
	wisdomStr, _ := reader.ReadString('\n')
	wisdom, _ := strconv.Atoi(wisdomStr)
	c.Wisdom = wisdom

	fmt.Print("Enter Charisma: ")
	charismaStr, _ := reader.ReadString('\n')
	charisma, _ := strconv.Atoi(charismaStr)
	c.Charisma = charisma
}
