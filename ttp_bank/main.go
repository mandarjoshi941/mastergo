package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/appliedgocourses/bank"
)

func main() {
	if len(os.Args) < 2 {
		usage()
		return
	}

	// Restore the bank data.
	// Call bank.Load() here and handle any error.
	bank.Load()

	// Save the bank data.
	// Use a deferred function for calling bank.Save().
	// (See the lecture on function behavior.)
	// If Save() returns an error, print it.
	defer bank.Save()

	// Perform the action.
	// os.Args[0] is the path to the executable.
	// os.Args[1] is the first parameter - the action we want to perform:
	// create, list, update, transfer, or history.

	switch os.Args[1] {

	case "list":
		s := bank.ListAccounts()
		fmt.Println(s)
	// TODO: more cases
	case "create":
		newAccName := os.Args[2]
		acc := bank.NewAccount(newAccName)
		fmt.Println("Account", acc.Name, "created successfully")
	case "update":
		i, _ := strconv.ParseInt(os.Args[3], 0, 32)
		newbal, err := update(os.Args[2], int(i))
		if err == nil {
			fmt.Println("Update successful new balnce is", newbal)
		} else {
			fmt.Errorf(err.Error())
		}
	case "transfer":
		i, _ := strconv.ParseInt(os.Args[4], 0, 32)
		newBal1, newBal2, err := transfer(os.Args[2], os.Args[3], int(i))
		if err == nil {
			fmt.Println("Transfer successful, new balance of", os.Args[2], "is", newBal1, "and", os.Args[3], "is", newBal2)
		} else {
			fmt.Errorf(err.Error())
		}
	case "history":
		hist, err := history(os.Args[2])
		if err == nil {
			fmt.Println(hist)
		} else {
			fmt.Errorf(err.Error())
		}
	default:
		fmt.Println("Please choose proper option")
		usage()

	}
}

func usage() {
	fmt.Println(`Usage:

bank create <name>                     Create an account.
bank list                              List all accounts.
bank update <name> <amount>            Deposit or withdraw money.
bank transfer <name> <name> <amount>   Transfer money between two accounts.
bank history <name>                    Show an account's transaction history.
`)
}

// update takes a name and an amount, deposits the amount if it
// is greater than zero, or withdraws it if it is less than zero,
// and returns the new balance and any error that occurred.
func update(name string, amount int) (int, error) {
	acc, err := bank.GetAccount(name)
	if err == nil {
		if amount > 0 {
			return bank.Deposit(acc, amount)
		}
		return bank.Withdraw(acc, amount*-1)
	}
	return 0, err
}

// transfer takes two names and an amount, transfers the amount from
// the account belonging to name #1 to the account belonging to name #2,
// and returns the new balances of both accounts and any error that occurred.
func transfer(name, name2 string, amount int) (int, int, error) {
	acc, err := bank.GetAccount(name)
	acc2, err := bank.GetAccount(name2)
	if err == nil {
		return bank.Transfer(acc, acc2, amount)
	}
	return 0, 0, err
}

// history takes an account name, retrieves the account, and calls bank.History()
// to get the history closure function. Then it calls the closure in a loop,
// formatting the return values and appending the result to the output string, until the boolean return parameter of the closure is `false`.
func history(name string) (string, error) {
	var result string
	acc, err := bank.GetAccount(name)
	if err == nil {
		hisFunc := bank.History(acc)
		for {
			trasctionAmount, resultingBalance, more := hisFunc()
			if !more {
				break
			}
			result = result + fmt.Sprint(trasctionAmount) + " " + fmt.Sprint(resultingBalance) + "\n"
		}
	}
	return result, err
}
