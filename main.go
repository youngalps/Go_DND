package main

import (
	"fmt"

	character "github.com/youngalps/Go_DND/Character"
)

func main() {
	// player creation
	player := character.NewCharacter("PlayerName", "PlayerClass") // TODO: make this a input from character

	// Set the Player Level
	player.SetLevel(1) // this should not be a input all players start at 1

	fmt.Println("Enter your characters name")

	// Read and set each stat
	player.SetStatsFromInput()

	// Display
	fmt.Println("Character Name:", player.Name)
	fmt.Println("Character Class:", player.Class)
	fmt.Println("Character Level:", player.Level)
	fmt.Println("Character Stats:")
	fmt.Println("Strength:", player.Strength)
	fmt.Println("Dexterity:", player.Dexterity)
	fmt.Println("Constitution:", player.Constitution)
	fmt.Println("Intelligence:", player.Intelligence)
	fmt.Println("Wisdom:", player.Wisdom)
	fmt.Println("Charisma:", player.Charisma)

}
