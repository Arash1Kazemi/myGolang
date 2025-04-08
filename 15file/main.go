package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Ha")

	content := "this need to go in a file"

	file, err := os.Create("./mylcogofile.txt")
	checkNilError(err)
	defer file.Close()

	length, err := io.WriteString(file, content)
	checkNilError(err)

	fmt.Printf("Wrote %d bytes to the file\n", length)
	readFile("./mylcogofile.txt");
}

func readFile(fileName string) {
	dataByte, err := os.ReadFile(fileName)
	checkNilError(err)
	fmt.Println(string(dataByte))
}

func checkNilError(err error) {
	if err != nil {
		fmt.Println("Error:", err)
		panic(err)
	}
}
