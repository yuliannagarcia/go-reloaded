package modification

import "strings"

func FixPunctuation(fixedSpace []string) []string {
	var fixedPunctuation string
	char := []string{",", ".", ":", ";", "?", "!", "'", "(", ")"}

	for i, v := range fixedSpace {
		if len(fixedSpace) < 2 {
			fixedPunctuation += fixedSpace[i]
		} else if i != len(fixedSpace)-1 {
			if v == char[0] || v == char[1] || v == char[2] || v == char[3] || v == char[4] || v == char[5] {
				if v == char[5] && fixedSpace[i+1] == char[4] {
					fixedPunctuation += fixedSpace[i]
				} else if (v == char[1] && fixedSpace[i+1] == char[1]) || (v == char[1] && fixedSpace[i+1] == char[6]) {
					fixedPunctuation += fixedSpace[i]
				} else {
					fixedPunctuation += fixedSpace[i]
					fixedPunctuation += " "
				}
			} else if v == char[6] {
				count := 0
				for _, l := range fixedPunctuation {
					if string(l) == char[6] {
						count++
					}
				}
				if count%2 == 0 {
					fixedPunctuation += fixedSpace[i]
				} else {
					fixedPunctuation += fixedSpace[i]
					fixedPunctuation += " "
				}
			} else if v == char[7] {
				fixedPunctuation += fixedSpace[i]
			} else if v == char[8] {
				if fixedSpace[i+1] == "," || fixedSpace[i+1] == "." || fixedSpace[i+1] == ";" || fixedSpace[i+1] == ":" || fixedSpace[i+1] == ")" || fixedSpace[i+1] == "!" || fixedSpace[i+1] == "?" {
					fixedPunctuation += fixedSpace[i]
				} else {
					fixedPunctuation += fixedSpace[i]
					fixedPunctuation += " "
				}
			} else {
				yesChar := false
				for _, k := range char[:6] {
					if fixedSpace[i+1] == k {
						fixedPunctuation += fixedSpace[i]
						yesChar = true
					}
				}
				if yesChar == false {
					yesDigit := false
					for _, p := range v {
						if p > 47 && p < 58 && fixedSpace[i+1] == ")" {
							fixedPunctuation += fixedSpace[i]
							yesDigit = true
							break
						}
					}
					if yesDigit == false {
						if fixedSpace[i+1] == "," || fixedSpace[i+1] == "." || fixedSpace[i+1] == ";" || fixedSpace[i+1] == ":" || fixedSpace[i+1] == ")" || fixedSpace[i+1] == "!" || fixedSpace[i+1] == "?" {
							fixedPunctuation += fixedSpace[i]
						} else {
							fixedPunctuation += fixedSpace[i]
							fixedPunctuation += " "
						}
					}
				}
			}
		} else if i == len(fixedSpace)-1 {
			fixedPunctuation += fixedSpace[i]
		}
	}
	var fixedDontStr string
	for i, v := range fixedPunctuation {
		if len(fixedPunctuation) > 1 {
			if v == ' ' && fixedPunctuation[i+1] == '\'' && ((fixedPunctuation[i+2] == 't' && fixedPunctuation[i-1] == 'n') || fixedPunctuation[i+2] == 's') {
				continue
			} else if v == ' ' && (v != rune(fixedPunctuation[len(fixedPunctuation)-2]) || v != rune(fixedPunctuation[len(fixedPunctuation)-1])) {
				if v == ' ' && ((fixedPunctuation[i+1] == '.' && fixedPunctuation[i+2] == '.') || (fixedPunctuation[i+1] == '!' && fixedPunctuation[i+2] == '?')) {
					continue
				} else {
					fixedDontStr += string(fixedPunctuation[i])
				}
			} else {
				fixedDontStr += string(fixedPunctuation[i])
			}
		} else {
			fixedDontStr += string(fixedPunctuation[i])
		}
	}
	fixedPunct2 := strings.Fields(fixedDontStr)
	return fixedPunct2
}
