package main

import (
	"io/ioutil" //This package provides functions for reading and writing files.
	"log" // This package provides a simple logging infrastructure.
	"os"
	"modification/modifications"
)
//The entry point for execution.
func main() {
	arguements := os.Args[1:]
	if len(arguements) != 2 {
		return
	}
	file := arguements[0]
	if file[len(file)-4:] != ".txt" {
		return
	}
	content, _ := ioutil.ReadFile(file)
	fixedSpaceCharacter := modification.SpaceCharacter(string(content))
	fixedPunctuation := modification.FixPunctuation(fixedSpaceCharacter)
	changedVowel := modification.ChangeVowelForward(fixedPunctuation)
	forwardModify := modification.ForwardModify(changedVowel)

	file2 := arguements[1]
	if file[len(file2)-4:] != ".txt" {
		return
	}
	f, err := os.Create(file2)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	if _, err := f.WriteString(forwardModify); err != nil {
		log.Fatal(err)
	}
}
