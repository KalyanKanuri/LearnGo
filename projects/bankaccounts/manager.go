package bankaccounts

import (
	"fmt"
	"math/rand/v2"
)

type Account struct {
	Name          string
	AccountNumber string
	AccountType   string
	Balance       float64
	InterestRate  float64
	IsActive      bool
}

var accounts map[string]*Account
var accTypePrefixes []string

func init() {
	accounts = make(map[string]*Account)
	accTypePrefixes = []string{
		"SAV",
		"CRE",
		"CUR",
	}
}

func (a Account) String() string {
	return fmt.Sprintf("\nAccount Name: %s\nAccount Number: %s\nAccount Type: %s\nBalance: %.2f\nAccount Status: %t\n",
		a.Name, a.AccountNumber, a.AccountType, a.Balance, a.IsActive)
}

func generateAccNumber(accTypePrefix string) string {
	for _, accTypePre := range accTypePrefixes {
		if accTypePrefix == accTypePre {
			fmt.Println("Generating account number")
			accountNum := rand.IntN(10000)
			accountNumber := fmt.Sprintf("%s-%d", accTypePrefix, accountNum)
			for accNum := range accounts {
				if accNum == accountNumber {
					fmt.Println("clash in account number generated, retrying account number generation")
					accountNumber = generateAccNumber(accTypePrefix)
				} else {
					fmt.Println("created account number", accountNumber)
					return accountNumber
				}
			}
		}
	}
	return ""
}

func (acc *Account) createAccount() error {
	// check if account already exists and store
	for accNum, Account := range accounts {
		accExists := (Account.AccountType == acc.AccountType)
		if accNum == acc.AccountNumber && accExists {
			err := fmt.Errorf("There's already an account exists with account number %s for this account type %s",
				Account.AccountNumber,
				acc.AccountType,
			)
			return err
		} else {
			accounts[acc.AccountNumber] = acc
			fmt.Printf("New Account is created for user %s with account number %s\n",
				acc.Name,
				acc.AccountNumber,
			)
			accDetails := acc.String()
			fmt.Printf("-- Account Details -- %+v", accDetails)
			break
		}
	}
	return nil
}

func (acc *Account) deposit(amount float64) {
	fmt.Printf("\nProcessing deposit of amount %.2f into account %s\n",
		amount,
		acc.AccountNumber,
	)
	acc.Balance += amount
	fmt.Printf("Processed amount %.2f, into account %s, - updated balance %.2f\n",
		amount,
		acc.AccountNumber,
		acc.Balance,
	)
}

func (acc *Account) withdraw(amount float64) error {
	fmt.Printf("\nProcessing Withdraw of amount %.2f from account %s\n",
		amount,
		acc.AccountNumber,
	)
	if amount > acc.Balance {
		err := fmt.Errorf("Error processing withdrawal request, requested amount %.2f is greater than account balance %.2f\n",
			amount,
			acc.Balance,
		)
		return err
	}
	acc.Balance -= amount
	fmt.Printf("Processed withdraw of amount %.2f, from account %s - updated balance %.2f\n",
		amount,
		acc.AccountNumber,
		acc.Balance,
	)
	return nil
}

func (acc *Account) addInterest() {
	fmt.Printf("\nProcessing interest calculation for account %s with interest rate %.2f\n",
		acc.AccountNumber,
		acc.InterestRate,
	)
	interest := acc.Balance * acc.InterestRate
	acc.Balance += interest
	fmt.Printf("Processed interest %.2f into account %s - updated balance %.2f\n",
		interest,
		acc.AccountNumber,
		acc.Balance,
	)
}

func AccManager() {
	// create a dummy account to let accounts loop work
	dummyAccount := Account{
		Name:          "dummy",
		AccountNumber: "dummy-1234",
		AccountType:   "Dummy Account",
		Balance:       00,
		InterestRate:  0.02,
		IsActive:      false,
	}
	accounts[dummyAccount.AccountNumber] = &dummyAccount

	fmt.Println("-- Processing Account Creation --")
	savAcc := Account{
		Name:          "Kalyan Kanuri",
		AccountNumber: generateAccNumber("SAV"),
		AccountType:   "Savings Account",
		Balance:       3000,
		InterestRate:  0.02,
		IsActive:      true,
	}
	savAcc.createAccount()
	savAcc.deposit(1000)
	savAcc.withdraw(100)
	savAcc.addInterest()
	fmt.Println(accounts)
}
