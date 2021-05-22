package main

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {

	file, err := os.OpenFile("csvpath", os.O_WRONLY|os.O_CREATE, 0600)
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
		fmt.Println(val)
	}
	writer.Flush()
}
