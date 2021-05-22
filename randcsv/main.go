package main

import (
	"encoding/csv"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {

	path := "sample.csv"
	if len(os.Args) >= 2 {
		path = os.Args[1]
	}

	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)

	for idx := 0; idx < 100; idx++ {
		rand.Seed(time.Now().UnixNano())
		val := []string{strconv.Itoa(idx), strconv.Itoa(rand.Intn(100))}
		err = writer.Write(val)
		if err != nil {
			break
		}
	}
	writer.Flush()
}
