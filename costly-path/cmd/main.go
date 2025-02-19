package main

import (
	"costly-path/pathfinder"
	"costly-path/pathhttp"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	sourcepath := os.Getenv("PATH_ADDRESS")
	var data [][]int
	err = pathhttp.GetJson(sourcepath, &data)
	if err != nil {
		panic(err)
	}

	ret, err := pathfinder.FindCostlyPath(data)
	if err != nil {
		panic(err)
	}

	fmt.Println(ret)
}
