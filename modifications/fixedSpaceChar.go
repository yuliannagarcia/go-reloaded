package modification

import "strings"

func SpaceCharacter(file string) []string {
	fixedSpaces := strings.Fields(file)
	stringOfText := strings.Join(fixedSpaces, " ")
	var fixedSpaceStr string
	var fixedChar []string
	for i, v := range stringOfText {
		if rune(v) == ',' || rune(v) == '.' || rune(v) == ';' || rune(v) == ':' || rune(v) == '?' || rune(v) == '!' || rune(v) == '\'' || rune(v) == '(' || rune(v) == ')' {
			fixedSpaceStr += " "
			fixedSpaceStr += string(stringOfText[i])
			fixedSpaceStr += " "
		} else if v == '‘' {
			fixedSpaceStr += " "
			fixedSpaceStr += "'"
			fixedSpaceStr += " "
		} else {
			fixedSpaceStr += string(stringOfText[i])
		}
	}
	fixedChar = strings.Fields(fixedSpaceStr)
	return fixedChar
}
