package modification

func ChangeVowelForward(fixedPunctuation []string) []string {
	for i, v := range fixedPunctuation {
		if len(fixedPunctuation) > 4 {
			if v == "a" && (fixedPunctuation[i+1][0] == 'a' || fixedPunctuation[i+1][0] == 'e' || fixedPunctuation[i+1][0] == 'i' || fixedPunctuation[i+1][0] == 'o' || fixedPunctuation[i+1][0] == 'u' || fixedPunctuation[i+1][0] == 'h') {
				fixedPunctuation[i] = "an"
			} else if v == "a" && (fixedPunctuation[i+1][0] == 'A' || fixedPunctuation[i+1][0] == 'E' || fixedPunctuation[i+1][0] == 'I' || fixedPunctuation[i+1][0] == 'O' || fixedPunctuation[i+1][0] == 'U' || fixedPunctuation[i+1][0] == 'H') {
				fixedPunctuation[i] = "an"
			} else if v == "A" && (fixedPunctuation[i+1][0] == 'a' || fixedPunctuation[i+1][0] == 'e' || fixedPunctuation[i+1][0] == 'i' || fixedPunctuation[i+1][0] == 'o' || fixedPunctuation[i+1][0] == 'u' || fixedPunctuation[i+1][0] == 'h') {
				fixedPunctuation[i] = "An"
			} else if v == "A" && (fixedPunctuation[i+1][0] == 'A' || fixedPunctuation[i+1][0] == 'E' || fixedPunctuation[i+1][0] == 'I' || fixedPunctuation[i+1][0] == 'O' || fixedPunctuation[i+1][0] == 'U' || fixedPunctuation[i+1][0] == 'H') {
				fixedPunctuation[i] = "AN"
			} else if v == "an" && fixedPunctuation[i+1][0] != 'a' && fixedPunctuation[i+1][0] != 'e' && fixedPunctuation[i+1][0] != 'i' && fixedPunctuation[i+1][0] != 'o' && fixedPunctuation[i+1][0] != 'u' && fixedPunctuation[i+1][0] != 'h' && fixedPunctuation[i+1][0] >= 95 && fixedPunctuation[i+1][0] <= 122 {
				fixedPunctuation[i] = "a"
			} else if v == "an" && fixedPunctuation[i+1][0] != 'A' && fixedPunctuation[i+1][0] != 'E' && fixedPunctuation[i+1][0] != 'I' && fixedPunctuation[i+1][0] != 'O' && fixedPunctuation[i+1][0] != 'U' && fixedPunctuation[i+1][0] != 'H' && fixedPunctuation[i+1][0] >= 65 && fixedPunctuation[i+1][0] <= 90 {
			} else if v == "An" && fixedPunctuation[i+1][0] != 'a' && fixedPunctuation[i+1][0] != 'e' && fixedPunctuation[i+1][0] != 'i' && fixedPunctuation[i+1][0] != 'o' && fixedPunctuation[i+1][0] != 'u' && fixedPunctuation[i+1][0] != 'h' && fixedPunctuation[i+1][0] >= 95 && fixedPunctuation[i+1][0] <= 122 {
				fixedPunctuation[i] = "A"
			} else if v == "AN" && fixedPunctuation[i+1][0] != 'A' && fixedPunctuation[i+1][0] != 'E' && fixedPunctuation[i+1][0] != 'I' && fixedPunctuation[i+1][0] != 'O' && fixedPunctuation[i+1][0] != 'U' && fixedPunctuation[i+1][0] != 'H' && fixedPunctuation[i+1][0] >= 65 && fixedPunctuation[i+1][0] <= 90 {
				fixedPunctuation[i] = "A"
			}
		} else {
			break
		}
	}
	return fixedPunctuation
}
