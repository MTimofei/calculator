package main

import (
	"fmt"
	"strconv"
	"strings"
)

func ConvExample(example, ns string) (operation string, variable []int, err error) {
	if ns == "Arabic" {
		operation, variable, err = ConvArabic(example)
		return operation, variable, err
	}
	if ns == "Ruman" {
		operation, variable, err = ConvRuman(example)
		return operation, variable, err
	}
	if ns == " " {
		err = fmt.Errorf("erro number system")
		return operation, variable, err
	}

	return operation, variable, err
}

func ConvArabic(example string) (operation string, variable []int, err error) {
	operationChar := []string{"/", "*", "-", "+"}

	for i := 0; i < len(operationChar); i++ {
		y := strings.Contains(example, operationChar[i])
		if y {
			variableS := strings.Split(example, operationChar[i])
			for _, i := range variableS {
				j, err := strconv.Atoi(i)
				if err != nil {
					panic(err)
				}
				variable = append(variable, j)
			}
			operation = operationChar[i]
			for n := 0; n < len(variable); n++ {
				if variable[n] > 10 {
					err = fmt.Errorf("error invalid variable")
					return operation, variable, err
				}
			}
			return operation, variable, err
		}
	}

	return operation, variable, err
}

func ConvRuman(example string) (operation string, variable []int, err error) {
	operationChar := []string{"/", "*", "-", "+"}

	for i := 0; i < len(operationChar); i++ {
		y := strings.Contains(example, operationChar[i])
		if y {
			variableS := strings.Split(example, operationChar[i])
			for y := 0; y < len(variableS); y++ {
				x := 0
				vaR := strings.Split(variableS[y], "")
				nI, nV, nX := 0, 0, 0
				for n := 0; n < len(vaR); n++ {

					if vaR[n] == "I" {
						x++

						if nI > 2 {
							err = fmt.Errorf("error invalid variable")
							return operation, variable, err
						}
						nI++
					}
					if vaR[n] == "V" {

						if nV > 0 {
							err = fmt.Errorf("error invalid variable")
							return operation, variable, err
						}
						if n == 0 {
							x = x + 5
						}
						if n > 0 {
							x = x - 5
						}
						nV++
					}
					if vaR[n] == "X" {
						if nX > 2 {
							err = fmt.Errorf("error invalid variable")
							return operation, variable, err
						}
						if n == 0 {
							x = x + 10
						}
						if n > 0 {
							x = x - 10
						}
						nX++
					}
					if x < 0 {
						x = x * (-1)
					}

				}
				variable = append(variable, x)
			}

			operation = operationChar[i]
			for n := 0; n < len(variable); n++ {
				if variable[n] > 10 {
					err = fmt.Errorf("error invalid variable")
					return operation, variable, err
				}
			}
			return operation, variable, err
		}

	}
	return operation, variable, err
}

func ConvAnswer(answer int, ns string) (answerForUs string, err error) {
	if ns == "Arabic" {
		answerForUs := strconv.Itoa(answer)
		return answerForUs, nil
	}
	if ns == "Ruman" {
		if answer < 0 {
			err := fmt.Errorf("error: The Roman numeral system has no negative value")
			return "", err
		}
		if answer >= 50 {
			x := answer / 50

			for i := 0; i < x; i++ {
				answerForUs = answerForUs + "L"
			}
			answer = answer - 50*x
		}
		if answer >= 40 {
			x := answer / 40

			for i := 0; i < x; i++ {
				answerForUs = answerForUs + "XL"
			}
			answer = answer - 40*x
		}
		if answer >= 10 {
			x := answer / 10

			for i := 0; i < x; i++ {
				answerForUs = answerForUs + "X"
			}
			answer = answer - 10*x
		}
		if answer >= 9 {
			x := answer / 9

			for i := 0; i < x; i++ {
				answerForUs = answerForUs + "IV"
			}
			answer = answer - 9*x
		}
		if answer >= 5 {

			x := answer / 5

			for i := 0; i < x; i++ {
				answerForUs = answerForUs + "V"
			}
			answer = answer - 5*x

		}
		if answer >= 4 {
			x := answer / 4

			for i := 0; i < x; i++ {
				answerForUs = answerForUs + "IV"
			}
			answer = answer - 4*x
		}
		if answer >= 1 {

			x := answer / 1

			for i := 0; i < x; i++ {
				answerForUs = answerForUs + "I"
			}

		}
		return answerForUs, err

	}
	return answerForUs, err
}
