package main

import (
	"bufio"
	"fmt"
	"main/utils"
	"main/valcalc"
	"os"
)

func main() {
	done := make(chan interface{})
	defer close(done)

	filePath, err := utils.ReadFileNameFromEnv("FILE_PATH")
	if err != nil {
		panic(err)
	}
	readFile, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	res := []<-chan valcalc.Result{}

	for fileScanner.Scan() {
		res = append(res, valcalc.EvaluateExpression(done, valcalc.ParseExpression(done, fileScanner.Text())))
	}

	for _, c := range res {
		res := <-c
		if res.Err != nil {
			fmt.Printf("Error: %v\n", res.Err)
		} else {
			fmt.Printf("equation=%d\n", res.Result)
		}
	}

}
