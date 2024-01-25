package main

import (
	"bufio"
	"eval"
	"log"
	"os"
	"strings"
)

func main() {
	for {
		reader := bufio.NewReader(os.Stdin)
		log.Print("Enter expression: ")
		rawExp, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		exp, err := eval.ParseInfixExpression(strings.TrimSuffix(rawExp, "\n"))
		if err != nil {
			log.Fatal(err)
		}

		evaluation, err := exp.Evaluate()
		if err != nil {
			log.Fatal(err)
		}

		log.Println("Result: ", evaluation)

	}
}
