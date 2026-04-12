package concurrency

import (
	"encoding/json"
	"fmt"
	"sync"
)

type FanOutFanIn struct {
	wg        *sync.WaitGroup
	userChan  chan User
	orderChan chan Order
	pmntChan  chan Payment
}

type User struct {
	ID   string
	Name string
	DOB  string
}

func (fofi *FanOutFanIn) getUserDetails() {
	defer fofi.wg.Done()
	fmt.Println("collecting user details")
	usr := User{
		ID:   "USR-1245",
		Name: "John",
		DOB:  "08-05-2000",
	}
	usrBS, err := json.MarshalIndent(usr, "", " ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("user details: %s\n", string(usrBS))
	fofi.userChan <- usr
}

type Order struct {
	ID    string
	Descr string
	Total float64
	Addr  string
}

func (fofi *FanOutFanIn) getOrderDetails() {
	defer fofi.wg.Done()
	fmt.Println("collecting order details for user")
	ord := Order{
		ID:    "ORD-1245",
		Descr: "Order placed for HP Laptop",
		Total: 75000,
		Addr:  "Manhattan, Vice City",
	}
	ordBS, err := json.MarshalIndent(ord, "", " ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("order details: %s\n", string(ordBS))
	fofi.orderChan <- ord
}

type Payment struct {
	ID   string
	Type string
}

func (fofi *FanOutFanIn) getPaymentDetails() {
	defer fofi.wg.Done()
	fmt.Println("collecting payment details")
	pmnt := Payment{
		ID:   "pmnt-1245",
		Type: "Credit",
	}
	pmntBS, err := json.MarshalIndent(pmnt, "", " ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("payment details: %s\n", string(pmntBS))
	fofi.pmntChan <- pmnt
}

type UserProfile struct {
	User    User    `json:"user"`
	Order   Order   `json:"order"`
	Payment Payment `json:"payment"`
}

/*
 * Fan-Out: the process of distrubuting a single task to multiple go routines is called
 * fan-out pattern
 * Fan-In: the process of collectnig all the broken results back to single unit is called
 * fan-in pattern
 */
func FanOutFanInPatternGo() {
	fmt.Println("-- Fanout Fanin Pattern in Go --")

	userChan := make(chan User)
	orderChan := make(chan Order)
	pmntChan := make(chan Payment)

	fofi := FanOutFanIn{
		wg:        &wg,
		userChan:  userChan,
		orderChan: orderChan,
		pmntChan:  pmntChan,
	}
	fnSlice := []func(){
		fofi.getUserDetails,
		fofi.getOrderDetails,
		fofi.getPaymentDetails,
	}

	for _, fn := range fnSlice {
		fofi.wg.Add(1)
		go fn()
	}

	go func() {
		wg.Wait()
		close(userChan)
		close(orderChan)
		close(pmntChan)

	}()

	users := make([]User, 0)
	orders := make([]Order, 0)
	pmnts := make([]Payment, 0)

	for range fnSlice {
		select {
		case user := <-userChan:
			users = append(users, user)
		case order := <-orderChan:
			orders = append(orders, order)
		case pmnt := <-pmntChan:
			pmnts = append(pmnts, pmnt)
		}
	}

	profiles := make([]*UserProfile, 0)
	for _, usr := range users {
		for _, order := range orders {
			for _, pmnt := range pmnts {
				profile := &UserProfile{
					User:    usr,
					Order:   order,
					Payment: pmnt,
				}
				profiles = append(profiles, profile)
				usrProfBS, _ := json.MarshalIndent(profile, "", " ")
				fmt.Printf("User Profile:\n%s\n", string(usrProfBS))
			}
		}
	}
}
