package main

import (
	"strconv"
	"strings"
)

func isHex(s string) bool {
	_, err := strconv.ParseInt(s, 16, 64)
	if err == nil {
		return true
	}
	return false
}

func isBin(s string) bool {
	for i := range s {
		if s[i] != '1' && s[i] != '0' {
			return false
		}
	}
	return true
}

func article(words []string) []string {
	for i := range words {
		if i != 0 && strings.ContainsAny(string(words[i][0]), "aeiouh") && words[i-1] == "a" {
			words[i-1] = "an"
		}
	}
	return words
}

func puncts1(words []string) []string {
	for i := range words {
		if words[i] == "..." || words[i] == "!?" {
			words[i-1] = words[i-1] + words[i]
		} else if len(words[i]) == 1 && strings.ContainsAny(words[i], ".,!?:;") {
			words[i-1] = words[i-1] + words[i]
		} else if len(words[i]) != 1 && checkPuncts(words[i]) {
			words[i-1] = words[i-1] + words[i]
		} else if len(words[i]) != 1 && strings.ContainsAny(string(words[i][0]), ".,!?:;") {
			words[i-1] = words[i-1] + string(words[i][0])
			words[i] = words[i][1:]
		}
	}

	newArr := []string{}

	for i := range words {
		if words[i] == "..." || words[i] == "!?" || (len(words[i]) == 1 && strings.ContainsAny(words[i], ".,!?:;")) || (len(words[i]) != 1 && checkPuncts(words[i])) {
			continue
		} else {
			newArr = append(newArr, words[i])
		}
	}
	return newArr
}

func checkPuncts(s string) bool {
	for i := range s {
		if strings.ContainsAny(string(s[i]), ".,!?:;") == false {
			return false
		}
	}
	return true
}

func puncts2(words []string) []string {
	mark := false
	n := 0
	for i := range words {
		if words[i] == "'" || words[i] == "‘" {
			n += 1
		}
	}
	for i := range words {
		if (words[i] == "'" || words[i] == "‘") && i+1 < len(words) && mark == false {
			words[i+1] = "'" + words[i+1]
			mark = true
		} else if (words[i] == "'" || words[i] == "‘") && mark == true {
			words[i-1] = words[i-1] + "'"
			mark = false
		}
	}
	newArr := []string{}
	if n%2 != 0 {
		for i := range words {
			if (words[i] == "'" || words[i] == "‘") && n != 1 {
				n -= 1
				continue
			} else if words[i] == "'" || words[i] == "‘" && n == 1 {
				newArr = append(newArr, words[i])
			} else {
				newArr = append(newArr, words[i])
			}
		}
	} else {
		for i := range words {
			if words[i] == "'" || words[i] == "‘" {
				n -= 1
				continue
			} else {
				newArr = append(newArr, words[i])
			}
		}
	}
	return newArr
}

func removeIndex(words []string, index int) []string {
	return append(words[:index], words[index+1:]...)
}

func BigFunc(words []string) []string {
	for i := 0; i < len(words); i++ {
		switch {
		case words[i] == "(hex)" && i >= 1 && isHex(words[i-1]):
			num, _ := strconv.ParseInt(words[i-1], 16, 64)
			words[i-1] = strconv.Itoa(int(num))
			words = removeIndex(words, i)
			i = 0
		case words[i] == "(bin)" && i >= 1 && isBin(words[i-1]):
			num, _ := strconv.ParseInt(words[i-1], 2, 64)
			words[i-1] = strconv.Itoa(int(num))
			words = removeIndex(words, i)
			i = 0
		case words[i] == "(up)" && i >= 1:
			words[i-1] = strings.ToUpper(words[i-1])
			words = removeIndex(words, i)
			i = 0
		case words[i] == "(up," && haveNum(words[i+1]) && i+1 < len(words):
			num := TrimAtoi(words[i+1], i)

			for j := i - num; j < i; j++ {
				words[j] = strings.ToUpper(words[j])
			}
			words = removeIndex(words, i)
			words = removeIndex(words, i)
			i = 0

		case words[i] == "(low)" && i >= 1:
			words[i-1] = strings.ToLower(words[i-1])
			words = removeIndex(words, i)
			i = 0
		case words[i] == "(low," && haveNum(words[i+1]) && i+1 < len(words):
			num := TrimAtoi(words[i+1], i)

			for j := i - num; j < i; j++ {
				words[j] = strings.ToLower(words[j])
			}

			words = removeIndex(words, i)
			words = removeIndex(words, i)
			i = 0

		case words[i] == "(cap)" && i >= 1:
			words[i-1] = strings.Title(words[i-1])
			words = removeIndex(words, i)
			i = 0
		case words[i] == "(cap," && haveNum(words[i+1]) && i+1 < len(words):
			num := TrimAtoi(words[i+1], i)
			for j := i - num; j < i; j++ {
				words[j] = strings.Title(strings.ToLower(words[j]))
			}
			words = removeIndex(words, i)
			words = removeIndex(words, i)
			i = 0

		}
	}

	newArr := []string{}

	for i := range words {
		if words[i] == "(bin)" || words[i] == "(hex)" || words[i] == "(cap)" || words[i] == "(up)" || words[i] == "(low)" {
			continue
		} else {
			newArr = append(newArr, words[i])
		}
	}

	return newArr
}

func TrimAtoi(s string, i int) int {
	n := 0
	num, err := strconv.Atoi(string(s[:len(s)-1]))
	if err == nil && num > 0 {
		for _, j := range []rune(s) {
			if j >= '0' && j <= '9' {
				n = n*10 + int(j-'0')
			}
		}
	}
	if i > n {
		return n
	}
	return i
}

// func up(words []string) []string {
// }
func haveNum(s string) bool {
	for _, v := range s[:len(s)-1] {
		if !(v >= '0' && v <= '9') {
			return false
		}
	}
	return true
}
