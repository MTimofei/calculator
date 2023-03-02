package main

import (
	"fmt"
)

func main() {

	var variable []int
	var example string
	var operation string
	var ns string
	var err error
	var answer int
	var answerForUs string

	example, err = Input()

	if err != nil {
		fmt.Println(err)
		return
	}
	ns, err = CheckAllChar(example)

	if err != nil {
		fmt.Println(err)
		return
	}
	operation, variable, err = ConvExample(example, ns)

	if err != nil {
		fmt.Println(err)
		return
	}
	answer, err = Calculating(operation, variable)

	if err != nil {
		fmt.Println(err)
		return
	}
	answerForUs, err = ConvAnswer(answer, ns)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(example, "=", answerForUs)

}
