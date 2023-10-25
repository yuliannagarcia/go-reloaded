package main

import (
	"io/ioutil"
	"log"
	"os"
	"modification/modifications"
)

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		return
	}
	file := args[0]
	if file[len(file)-4:] != ".txt" {
		return
	}
	content, _ := ioutil.ReadFile(file)
	fixedSpaceCharacter := modification.SpaceCharacter(string(content))
	fixedPunctuation := modification.FixPunctuation(fixedSpaceCharacter)
	changedVowel := modification.ChangeVowelForward(fixedPunctuation)
	forwardModify := modification.ForwardChange(changedVowel)

	file2 := args[1]
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
