package main

import (
	"fmt"
	"learning/go/advanced/concurrency"
	"learning/go/advanced/data_security"
	"learning/go/advanced/fileio"
	"learning/go/basics"
	"learning/go/core"
	"learning/go/datastructures"
	"learning/go/extensions"
	"learning/go/projects/bankaccounts"
	"learning/go/projects/contactbook"
	"learning/go/projects/filedownloader"
	"learning/go/projects/glogger"
	"learning/go/projects/order"
	"learning/go/projects/payrollprocessor"
)

func main() {
	// Printing in Go
	print("Started Learning Go")
	print("default print func from go runtime doesn't add new line")
	println("default println func from go runtime adds new line")
	fmt.Print("print func from fmt, provides additional formatting")
	fmt.Println("println func from fmt, provides new line with formatting")
	fmt.Printf("printf func from fmt, provides control over vars while printing\n")

	// Variables in Go
	basics.VarsInGo()

	// Constants in Go
	basics.ConstInGo()
	basics.IotaInGO()

	// 1st project custom logger
	fmt.Println("\n******* Custom Logger *******")
	glog.Log(0, "implemented custom logger")
	glog.Log(1, "used iota for log levels")
	glog.Log(2, "defined a type Loglevel")
	glog.Log(3, "created a func Log")
	glog.Log(4, "accepts LogLevel and a log msg")

	// Loops in Go
	basics.LoopsInGo()

	// Conditionals in Go
	basics.ConditionalsInGo()

	// 2nd project order processing
	fmt.Println("\n******* Order Processing *******")
	cart := order.AddToCart([]string{"Shirts", "Pants", "Shoes"})
	orderDetails := order.ProcessOrder(cart, true)
	fmt.Printf("Order Details: Order-ID -- %s\nItems -- %v\nBase Price -- %.2f\nFinal Price -- %.2f\nDiscount -- %.2f\n",
		orderDetails["orderID"],
		orderDetails["items"],
		orderDetails["basePrice"],
		orderDetails["finalPrice"],
		orderDetails["discount"],
	)

	// data structures in Go
	fmt.Println("\n******* Data Structures *******")
	datastructures.ArraysInGo()
	datastructures.SlicesInGo()
	datastructures.MapsInGo()

	// Pointers in Go
	core.PointersInGo()

	// 3rd project contact book
	fmt.Println("\n******* Contact Book *******")
	fmt.Println("-- Adding Contacts --")
	contact1 := contactbook.Contact{
		Name:  "John Doe",
		Age:   30,
		Email: "2P5oq@example.com",
		Phone: "123-456-7890",
	}
	contactbook.AddContact(&contact1)

	contact2 := contactbook.Contact{
		Name:  "Jane Doe",
		Age:   25,
		Email: "4rB2H@example.com",
		Phone: "987-654-3210",
	}
	contactbook.AddContact(&contact2)

	contact3 := contactbook.Contact{
		Name:  "John Doe",
		Age:   35,
		Email: "2P5oq@example.com",
		Phone: "123-456-7890",
	}
	contactbook.AddContact(&contact3)

	fmt.Println("-- Listing Contacts --")
	contactbook.ListContacts()

	fmt.Println("-- Getting Contact --")
	user, exists := contactbook.GetContact("John Doe")
	if exists {
		fmt.Printf("Contact found: Name: %s, Age: %d, Email: %s, Phone: %s\n",
			user.Name,
			user.Age,
			user.Email,
			user.Phone,
		)
	} else {
		fmt.Println("Contact not found.")
	}

	// Funcs In Go
	core.FuncsInGo()

	// Panic & Recover In Go
	core.PanicRecoverInGo()

	// OOPS in GO
	fmt.Println("\n******* OOPS in Go *******")
	fmt.Println("-- receivers & methods --")
	kalyan := core.NewEmployee(
		123,
		"Kanuri",
		"Kalyan",
		"kalyankanuri497@gmail.com",
		"+91 9963059178",
		true,
		"Software Development",
	)
	fmt.Printf("%+v\n", kalyan)
	fullname := kalyan.GetEmployeeFullName()
	fmt.Println("Employee FullName:", fullname)
	kalyan.UpdateActiveState()
	fmt.Printf("After IsActive state update -> %+v\n", kalyan.IsActive)

	// Generics in Go
	core.GenericsInGo()

	// 4th Project Payroll Process
	fmt.Println("\n******* Payroll Processor *******")
	salEmp := payrollprocessor.SalariedEmployee{
		ID:           123,
		Name:         "John Cena",
		AnnualSalary: 900000.0,
	}
	hourEmp := payrollprocessor.HourlyEmployee{
		ID:         136,
		Name:       "Jane Smith",
		HourlyRate: 36.0,
		HourlyPay:  1600.0,
	}
	empList := []payrollprocessor.SalaryPayer{
		salEmp,
		hourEmp,
	}
	payrollprocessor.ProcessPayroll(empList)

	// Embedding in Go
	core.EmbeddingInGo()

	// 5th Project Bank Account Manager
	fmt.Println("\n******* Bank Account Manager *******")
	bankaccounts.AccManager()

	fmt.Println("\n******* Regular Expressions in Go *******")
	extensions.RegexInGo()

	fmt.Println("\n******* Templates in Go *******")
	extensions.TextTemplatesInGo()

	fmt.Println("\n******* Buffer Reader in Go *******")
	extensions.BufferReaderInGo()

	fmt.Println("\n******* Concurrency in Go *******")
	concurrency.GoRoutinesInGo()
	concurrency.WaitGroupsInGo()
	concurrency.ChannelsInGo()
	concurrency.MutexInGo()

	// 6th Project Concurrent File Downloader
	fmt.Println("\n******* Concurrent File Downloader *******")
	filedownloader.TriggerDownload()

	// File ops In Go
	fmt.Println("\n******* File Operations in Go *******")
	fileio.FileIOInGo()
	fileio.HandleFilePaths()
	fileio.HandlerDirsInGo()
	fileio.EmbedFilesInGo()

	// Data Security & Encode-Decode in Go
	fmt.Println("\n******* DataSecurity Encoding-Decoding in Go *******")
	datasecurity.MarshalUnMarshalInGo()
}
