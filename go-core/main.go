package main

import (
	"coreconcepts/concepts"
	"coreconcepts/concepts/concurrency"
	concpatterns "coreconcepts/concepts/concurrency/patterns"
	contextGo "coreconcepts/concepts/context"
	"coreconcepts/concepts/funcs"
	"coreconcepts/concepts/packages/bufiopkg"
	"coreconcepts/concepts/packages/iopkg"
	"coreconcepts/concepts/packages/ospkg"
	"fmt"
	"os"
)

func main() {
	fmt.Println("***** Arrays In Go *****")
	concepts.ArraysInGo()
	fmt.Println()

	fmt.Println("***** Slices In Go *****")
	concepts.SlicesInGo()
	fmt.Println()

	fmt.Println("***** Maps In Go *****")
	concepts.MapsInGo()
	fmt.Println()

	fmt.Println("***** Structs In Go *****")
	concepts.StructsInGo()
	fmt.Println()

	fmt.Println("***** GoTo Statement in Go *****")
	concepts.GoToStatement()
	fmt.Println()

	fmt.Println("***** Variadic Functions In Go *****")
	funcs.VariadicFunc(1, 2, 3, 4, 5)
	funcs.VariadicFunc()
	fmt.Println()

	fmt.Println("***** Anonymous Functions In Go *****")
	funcs.AnonymousFunc()
	fmt.Println()

	fmt.Println("***** Closures In Go *****")
	funcs.Closures()
	fmt.Println("")

	fmt.Println("***** Interfaces In Go *****")
	concepts.InterfacesInGo()
	fmt.Println()

	fmt.Println("***** Generics In Go *****")
	concepts.GenericsInGo()
	fmt.Println()

	fmt.Println("***** Error handling In Go *****")
	concepts.ErrorHandlingInGo()
	fmt.Println()

	fmt.Println("***** Concurrency In Go *****")

	fmt.Println("-- Producer Consumer Pattern --")
	concurrency.GenerateNums()
	fmt.Println()

	fmt.Println("-- Squares Pipeline --")
	concurrency.GenerateSqrs()
	fmt.Println()

	fmt.Println("-- Select Executor --")
	concurrency.ExecuteSelect()
	fmt.Println()

	fmt.Println("-- Default Executor --")
	concurrency.ExecuteDefault()
	fmt.Println()

	fmt.Println("-- Time Case Executor --")
	concurrency.ExecuteTimeCase()
	fmt.Println()

	fmt.Println("-- Context in concurrency --")
	fmt.Println("--- Context With Cancel ---")
	contextGo.ContextWithCancel()
	fmt.Println()

	fmt.Println("--- Context With Timeout ---")
	contextGo.ContextWithTimeout()
	fmt.Println()

	fmt.Println("--- Context With Deadline ---")
	contextGo.ContextWithDeadline()
	fmt.Println()

	fmt.Println("***** Concurrency Patterns *****")

	fmt.Println("-- Worker Pool Pattern --")
	concpatterns.ExecWP()
	fmt.Println()

	fmt.Println("-- Fan Out Fan In Pattern --")
	concpatterns.ExecFOFI()
	fmt.Println()

	fmt.Println("-- Graceful Shutdown --")
	concpatterns.ExecGSDown()
	fmt.Println()

	fmt.Println("***** Packages In Go *****")
	fmt.Println("-- OS Package --")
	ospkg.ExecOSSim()

	fmt.Println("-- IO Package --")
	file, err := os.Open("notes.txt")
	if err != nil {
		fmt.Println(err)
	}
	iopkg.IOReader(file)
	file.Seek(0, 0) // reposition the internal cursor at cpu level in the file

	fmt.Println("-- Bufio Package --")
	bufiopkg.ReadFunc(file)
	file.Seek(0, 0)
	bufiopkg.ReadByteFunc(file)
	file.Seek(0, 0)
	bufiopkg.ReadRuneFunc(file)
	file.Seek(0, 0)
	bufiopkg.ReadStringFunc(file)
	file.Seek(0, 0)
	bufiopkg.ReadBytesFunc(file)
	file.Seek(0, 0)
	bufiopkg.PeekFunc(file)
	file.Seek(0, 0)
	bufiopkg.DiscardFunc(file)
	wFile, err := os.OpenFile("output.txt", os.O_CREATE, 0777)
	if err != nil {
		fmt.Println("Error opening file", err)
		return
	}
	defer wFile.Close()
	bufiopkg.WriteFunc(wFile)
	bufiopkg.WriteStringFunc(wFile)
	bufiopkg.WriteByteFunc(wFile)
	bufiopkg.WriteRuneFunc(wFile)
}
