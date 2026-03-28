package fileio

import (
	"bufio"
	"fmt"
	"os"
)

func FileIOInGo() {
	fmt.Println("-- Reading a file with os.ReadFile() --")
	data, err := os.ReadFile("go.mod")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("content:", string(data))

	fmt.Println("-- Writing to a file with os.WriteFile() --")
	sample_data := "Loving go programming so far"
	err = os.WriteFile("OSWFile.log", []byte(sample_data), 0644)
	fmt.Println("content write to file complete:", sample_data)

	fmt.Println("-- Reading file with os.Open() and bufio.NewScanner() --")
	file, _ := os.Open("OSWFile.log")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	fmt.Println("-- Writing a file with os.Create() and bufio.NewWriter() --")
	file, _ = os.Create("sample_file.log")

	writer := bufio.NewWriter(file)
	writer.WriteString("writing into file using bufio.NewWriter")
	writer.Flush()
	fmt.Println("completed write to file")

	fmt.Println("-- Appending to file using os.OpenFile() and os.O_APPEND flag --")
	file, _ = os.OpenFile("sample_file.log", os.O_APPEND|os.O_WRONLY, 0644)
	file.WriteString("\nAppending new line using os.OpenFile() and os.O_APPEND flag")
	fmt.Println("Append operation completed")
}
