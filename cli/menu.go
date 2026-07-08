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
		fmt.Println("\n===== TOY BLOCKCHAIN MENU =====")
		fmt.Println("1. Add Transaction")
		fmt.Println("2. Mine Block")
		fmt.Println("3. View Blockchain")
		fmt.Println("4. Check Validity")
		fmt.Println("5. Exit")
		fmt.Print("Choose option: ")

		scanner.Scan()
		choice := scanner.Text()

		switch choice {

		case "1":
			cli.addTransaction(scanner)

		case "2":
		if len(cli.Chain.PendingTransactions) == 0 {
			fmt.Println("No transactions to mine")
			break
		}

		cli.Chain.AddBlock(cli.Chain.PendingTransactions)
		cli.Chain.PendingTransactions = nil

		fmt.Println("Block mined successfully!")

		case "3":
			cli.printChain()

		case "4":
			if cli.Chain.IsValid() {
			fmt.Println("Blockchain is VALID")
		} else {
			fmt.Println("Blockchain is INVALID")
		}

		case "5":
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

	cli.Chain.AddTransaction(tx)
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