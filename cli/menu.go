package cli

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"toy-blockchain/blockchain"
	"toy-blockchain/ledger"
)

type CLI struct {
	Chain *blockchain.Blockchain
}

func (cli *CLI) Run() {
	scanner := bufio.NewScanner(os.Stdin)

	for {

		fmt.Println() // creates empty line before menu

		fmt.Println("===== TOY BLOCKCHAIN MENU =====")
		fmt.Println("1. Add Transaction")
		fmt.Println("2. Mine Transactions")
		fmt.Println("3. Generate Reward")
		fmt.Println("4. View Blockchain")
		fmt.Println("5. View Balance")
		fmt.Println("6. Check Validity")
		fmt.Println("7. Exit")
		fmt.Print("Choose option: ")

		scanner.Scan()
		choice := scanner.Text()

		switch choice {

		case "1":
			cli.addTransaction(scanner)

		case "2":

			cli.Chain.MineTransactions()

		case "3":

			fmt.Print("Miner name: ")

			scanner.Scan()

			miner := scanner.Text()

			cli.Chain.GenerateReward(miner)

		case "4":
			cli.printChain()

		case "5":

			cli.viewBalances()

		case "6":
			if cli.Chain.IsValid() {
				fmt.Println("Blockchain is VALID")
			} else {
				fmt.Println("Blockchain is INVALID")
			}

		case "7":
			fmt.Println("Exiting...")
			return

		default:
			fmt.Println("Invalid choice")
		}
	}
}

func (cli *CLI) addTransaction(scanner *bufio.Scanner) {

	fmt.Print("Sender: ")
	scanner.Scan()
	sender := scanner.Text()

	fmt.Print("Recipient: ")
	scanner.Scan()
	recipient := scanner.Text()

	fmt.Print("Amount: ")
	scanner.Scan()
	amountStr := scanner.Text()

	amount, err := strconv.ParseFloat(strings.TrimSpace(amountStr), 64)
	if err != nil {
		fmt.Println("Invalid amount")
		return
	}

	tx := ledger.Transaction{
		Sender:    sender,
		Recipient: recipient,
		Amount:    amount,
	}

	if !ledger.ValidateTransaction(tx) {
		fmt.Println("Invalid transaction")
		return
	}

	success := cli.Chain.AddTransaction(tx)

	if !success {

		fmt.Println("Transaction rejected: insufficient balance")

		return
	}

	fmt.Println("Transaction added to pool")
}

func (cli *CLI) printChain() {
	for _, block := range cli.Chain.Blocks {
		fmt.Println("\n--------------------")
		fmt.Println("Index:", block.Index)
		fmt.Println("Hash:", block.Hash)
		fmt.Println("PrevHash:", block.PreviousHash)
		fmt.Println("Transactions:", block.Transactions)
	}
}

func (cli *CLI) viewBalances() {

	fmt.Println("\n===== BALANCES =====")

	for user, balance := range cli.Chain.Balances {

		fmt.Println(user, ":", balance)

	}

}
