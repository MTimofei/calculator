package main

func Calculating(operation string, variable []int) (answer int, err error) {

	if operation == "+" {
		answer = 0
		for x := 0; x < len(variable); x++ {
			answer = answer + variable[x]
		}
		return answer, nil
	}
	if operation == "-" {
		answer = variable[0]
		for x := 1; x < len(variable); x++ {
			answer = answer - variable[x]
		}
		return answer, nil
	}
	if operation == "*" {
		answer = variable[0]
		for x := 1; x < len(variable); x++ {
			answer = answer * variable[x]
		}
		return answer, nil
	}
	if operation == "/" {
		answer = variable[0]
		for x := 1; x < len(variable); x++ {
			answer = answer / variable[x]
		}
		return answer, nil
	}
	return answer, err
}
