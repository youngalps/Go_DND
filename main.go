package main

import (
	"fmt"

	character "github.com/youngalps/Go_DND/Character"
)

func main() {
	// player creation

	character := character.NewCharacter()

	fmt.Println("Character Created:")
	fmt.Printf("Name: %s\n", character.Name)
	fmt.Printf("Class: %s\n", character.Class)
	fmt.Printf("Level: %d\n", character.Level)
	fmt.Printf("Strength: %d\n", character.Strength)
	fmt.Printf("Dexterity: %d\n", character.Dexterity)
	fmt.Printf("Constitution: %d\n", character.Constitution)
	fmt.Printf("Intelligence: %d\n", character.Intelligence)
	fmt.Printf("Wisdom: %d\n", character.Wisdom)
	fmt.Printf("Charisma: %d\n", character.Charisma)

}
