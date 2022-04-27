package main

import (
	"fmt"
	"strings"
)

/*
glob glob Silver is 34 Credits
glob prok Gold is 57800 Credits
pish pish Iron is 3910 Credits
xx
how much is pish tegj glob glob ?
how many Credits is glob prok Silver ?
how many Credits is glob prok Gold ?
how many Credits is glob prok Iron ?
how much wood could a woodchuck chuck if a woodchuck could chuck wood ?

output:
pish tegj glob glob is 42
xlii = 42
glob glob glob
iii = 3
glob prok Silver is 68 Credits
iv silver = 17
glob prok Gold is 57800 Credits
iv gold = 14450
glob prok Iron is 782 Credits
iv iron = 195.5
I have no idea what you are talking about
*/

const (
	iron   float32 = 195.5
	silver float32 = 17
	gold   float32 = 14450
)

func main() {
	translator("pish tegj glob glob")
	translator("glob prok gold")
	translator("glob prok silver")
	translator("glob prok iron")
	translator("how much wood could a woodchuck chuck if a woodchuck could chuck wood ?")
}

func translator(input string) {
	str := input
	// fmt.Println("string :", str)
	roman := roman(str)
	// fmt.Println("roman :", roman)
	romint := romantoint(roman)
	// fmt.Println("int :", romint)
	currency := currency(str)
	// fmt.Println("currency :", currency)
	if currency == 0 && romint != 0 {
		fmt.Println(input, "is", romint)
	} else if currency == 0 && romint == 0 {
		fmt.Println("I have no idea what you are talking about")
	} else {
		sum := romint * currency
		fmt.Println(input, "is", sum, "credits")
	}

}

func roman(str string) string {
	var parsed []string
	var romnum string
	parsed = strings.Split(str, " ")
	for _, data := range parsed {
		if data == "glob" {
			romnum += "I"
		} else if data == "prok" {
			romnum += "V"
		} else if data == "pish" {
			romnum += "X"
		} else if data == "tegj" {
			romnum += "L"
		}
	}

	return romnum
}

func romantoint(str string) float32 {
	var parsed []string
	var temp float32
	parsed = strings.Split(str, "")
	for i := 0; i < len(parsed); i++ {
		if parsed[i] == "I" {
			if i < len(parsed) {
				if i < len(parsed)-1 && parsed[i+1] == "V" {
					temp += 4
					i += 1
				} else if i < len(parsed)-1 && parsed[i+1] == "X" {
					temp += 9
					i += 1
				} else {
					temp += 1
				}
			}
		} else if parsed[i] == "V" {
			if i < len(parsed) {
				if i < len(parsed)-1 && parsed[i+1] == "I" {
					temp += 6
					i += 1
				}
			}
		} else if parsed[i] == "X" {
			if i < len(parsed) {
				if i < len(parsed)-1 && parsed[i+1] == "L" {
					temp += 40
					i += 1
				} else if i < len(parsed)-1 && parsed[i+1] == "C" {
					temp += 90
					i += 1
				} else {
					temp += 10
				}
			}
		} else if parsed[i] == "L" {
			if i < len(parsed) {
				temp += 50
			}
		}
	}

	return temp
}

func currency(str string) float32 {
	var parsed []string
	var temp float32
	parsed = strings.Split(str, " ")
	for _, data := range parsed {
		if data == "iron" {
			temp = iron
		} else if data == "silver" {
			temp = silver
		} else if data == "gold" {
			temp = gold
		}
	}
	return temp
}
