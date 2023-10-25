package modification

import (
	"strconv"
	"strings"
)

func ForwardModify(changedVowel []string) string {
	var removedArr []string
	for i, v := range changedVowel {
		// Capitalize first letter
		if v == "(cap)" || v == "(cap)," || v == "(cap)." || v == "(cap):" || v == "(cap);" || v == "(cap)!" || v == "(cap)?" || v == "(cap)..." || v == "(cap)!?" {
			if v[len(v)-1] != ')' {
				removedArr[len(removedArr)-1] = strings.Title(removedArr[len(removedArr)-1])
				removedArr[len(removedArr)-1] = removedArr[len(removedArr)-1] + string(v[len(v)-1])
			} else {
				removedArr[len(removedArr)-1] = strings.Title(removedArr[len(removedArr)-1])
			}
			// to Upper
		} else if v == "(up)" || v == "(up)," || v == "(up)." || v == "(up):" || v == "(up);" || v == "(up)!" || v == "(up)?" || v == "(up)..." || v == "(up)!?" {
			if v[len(v)-1] != ')' {
				removedArr[len(removedArr)-1] = strings.ToUpper(removedArr[len(removedArr)-1])
				removedArr[len(removedArr)-1] = removedArr[len(removedArr)-1] + string(v[len(v)-1])
			} else {
				removedArr[len(removedArr)-1] = strings.ToUpper(removedArr[len(removedArr)-1])
			}
			// to Lower
		} else if v == "(low)" || v == "(low)," || v == "(low)." || v == "(low):" || v == "(low);" || v == "(low)!" || v == "(low)?" || v == "(low)..." || v == "(low)!?" {
			if v[len(v)-1] != ')' {
				removedArr[len(removedArr)-1] = strings.ToLower(removedArr[len(removedArr)-1])
				removedArr[len(removedArr)-1] = removedArr[len(removedArr)-1] + string(v[len(v)-1])
			} else {
				removedArr[len(removedArr)-1] = strings.ToLower(removedArr[len(removedArr)-1])
			}
			// Hex to Decimal
		} else if v == "(hex)" || v == "(hex)," || v == "(hex)." || v == "(hex);" || v == "(hex):" || v == "(hex)!" || v == "(hex)?" || v == "(hex)..." || v == "(hex)!?" {
			hexInt, _ := strconv.ParseInt(removedArr[len(removedArr)-1], 16, 64)
			if v[len(v)-1] != ')' {
				removedArr[len(removedArr)-1] = strconv.Itoa(int(hexInt))
				removedArr[len(removedArr)-1] = removedArr[len(removedArr)-1] + string(v[len(v)-1])
			} else {
				removedArr[len(removedArr)-1] = strconv.Itoa(int(hexInt))
			}
			// Binar to Decimal
		} else if v == "(bin)" || v == "(bin)," || v == "(bin)." || v == "(bin);" || v == "(bin):" || v == "(bin)!" || v == "(bin)?" || v == "(bin)..." || v == "(bin)!?" {
			binInt, _ := strconv.ParseInt(changedVowel[i-1], 2, 64)
			if v[len(v)-1] != ')' {
				removedArr[len(removedArr)-1] = strconv.Itoa(int(binInt))
				removedArr[len(removedArr)-1] = removedArr[len(removedArr)-1] + string(v[len(v)-1])
			} else {
				removedArr[len(removedArr)-1] = strconv.Itoa(int(binInt))
			}
		} else if v == "(cap," || v == "(up," || v == "(low," {
			var indexStr string
			inter := changedVowel[i+1]
			for j, l := range inter {
				if l >= 48 && l <= 57 {
					indexStr += string(inter[j])
				} else if l == '-' {
					indexStr = ""
					break
				}
			}
			indexInt, _ := strconv.Atoi(indexStr)
			if len(removedArr)-indexInt < 0 {
				indexInt = len(removedArr)
			}
			for h := range removedArr[len(removedArr)-indexInt:] {
				if v == "(cap," {
					removedArr[len(removedArr)-indexInt:][h] = strings.Title(removedArr[len(removedArr)-indexInt:][h])
				} else if v == "(up," {
					removedArr[len(removedArr)-indexInt:][h] = strings.ToUpper(removedArr[len(removedArr)-indexInt:][h])
				} else if v == "(low," {
					removedArr[len(removedArr)-indexInt:][h] = strings.ToLower(removedArr[len(removedArr)-indexInt:][h])
				}
			}
		} else if i != 0 {
			if changedVowel[i-1] == "(cap," || changedVowel[i-1] == "(up," || changedVowel[i-1] == "(low," {
				if v[len(v)-1] != ')' {
					removedArr[len(removedArr)-1] = removedArr[len(removedArr)-1] + string(v[len(v)-1])
				}
			} else {
				removedArr = append(removedArr, changedVowel[i])
			}
		} else if i == 0 {
			removedArr = append(removedArr, changedVowel[i])
		}
	}
	removedStr := ""
	for i := range removedArr {
		if len(removedArr) > 1 {
			if i != len(removedArr)-1 {
				removedStr += removedArr[i]
				removedStr += " "
			} else {
				removedStr += removedArr[i]
				removedStr += "\n"
				break
			}
		} else {
			removedStr += removedArr[i]
		}
	}
	return removedStr
}
