package main

import (
	"fmt"
	"strings"
)

func CheckAllChar(example string) (ns string, err error) {
	allowedChar := []string{"1", "2", "3", "4",
		"5", "6", "7", "8",
		"9", "0", "/", "*",
		"-", "+", "I", "V", "X"}
	//example = strings.ToUpper(example)
	exampleSl := strings.Split(example, "")
	err = CheckFirstChar(exampleSl, allowedChar)
	if err != nil {
		return ns, err
	}
	ns, err = Rout(exampleSl, allowedChar)
	if err != nil {
		return ns, err
	}
	err = CheckRepeatOperations(exampleSl, allowedChar)
	if err != nil {
		return ns, err
	}
	return ns, err

}

func CheckFirstChar(exampleSl, allowedChar []string) (err error) {
	for i := 10; i <= 14; i++ {
		if i == 14 {
			return nil
		}
		if exampleSl[0] == allowedChar[i] {
			err := fmt.Errorf("eorng first character")
			return err
		}
	}
	return err
}

func Rout(exampleSl, allowedChar []string) (ns string, err error) {
	ns = " "
	
	for i := 0; i < 10; i++ {
		if exampleSl[0] == allowedChar[i] {
			ns = "Arabic"
			err = CheckAllCharArabic(exampleSl, allowedChar)
			return ns, err
		}
	}
	if ns == " " {
		for i := 14; i <= len(allowedChar); i++ {
			if i == len(allowedChar) {
				err := fmt.Errorf("eorng first character")
				return ns, err
			}
			if exampleSl[0] == allowedChar[i] {
				ns = "Ruman"
				err = CheckAllCharRuman(exampleSl, allowedChar)
				return ns, err
			}
		}
		return ns, err
	}
	return ns, err
}

func CheckAllCharArabic(exampleSl, allowedChar []string) (err error) {

	for i := 0; i < len(exampleSl); i++ {

		for n := 0; n <= 14; n++ {

			if n == 14 {
				err := fmt.Errorf("error different number systems")
				return err
			}
			if exampleSl[i] == allowedChar[n] {
				break
			}

		}
		if err != nil {
			return err
		}

	}
	return err
}

func CheckAllCharRuman(exampleSl, allowedChar []string) (err error) {
	
	for i := 0; i < len(exampleSl); i++ {

		for n := 10; n <= len(allowedChar); n++ {

			if n == len(allowedChar) {
				err := fmt.Errorf("error different number systems")
				return err
			}
			if exampleSl[i] == allowedChar[n] {
				break
			}

		}
		if err != nil {
			return err
		}

	}
	return err
}

// проверка на повторение оприраций
func CheckRepeatOperations(exampleSl, allowedChar []string) (err error) {
	x := 0
	for i := 0; i < len(exampleSl); i++ {

		for n := 10; n < 14; n++ {

			if exampleSl[i] == allowedChar[n] {
				x++
				break
			}

		}

	}
	if x == 0 || x > 1 {
		err = fmt.Errorf("error in recording the example:repeat operations")
	}
	if x == 1 {
		err = nil
	}
	return err
}
