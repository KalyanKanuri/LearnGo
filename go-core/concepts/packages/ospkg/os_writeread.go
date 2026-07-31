package ospkg

import (
	"fmt"
	"os"
)

func ExecOSSim() {
	fCreated, err := os.Create("notes.txt")
	if err != nil {
		fmt.Println("Error while creating file", err)
		return
	}
	defer fCreated.Close()

	fCreated.WriteString("Hello Backend Engineer!")
	fCreated.Sync()

	fOpened, err := os.Open("notes.txt")
	if err != nil {
		fmt.Println("Error while opening the file")
		return
	}
	defer fOpened.Close()

	buf := make([]byte, 2*1024)
	bytesRead, err := fOpened.Read(buf)
	if err != nil {
		fmt.Println("Error while reading the file", err)
		return
	}

	fmt.Println(string(buf[:bytesRead]))
}
