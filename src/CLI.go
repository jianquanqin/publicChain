package src

import (
	"flag"
	"fmt"
	"log"
	"os"
)

type CLI struct{}

func printUsage() {
	fmt.Println("Usage:")
	fmt.Println("\tcreatBlockChain -address --genesis address")
	fmt.Println("\ttransfer -from FROM -to TO -amount AMOUNT --transaction details")
	fmt.Println("\tprintChain -- output block's information")
	fmt.Println("\tgetBalance -address -- output address's balance")
	fmt.Println("\tgetAddressLists -- output address list")
	fmt.Println("\tcreateNewWallet -- output address list")
}
func isValidArgs() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}
}
func (cli CLI) Run() {
	isValidArgs()
	//custom command
	getAddressListsCmd := flag.NewFlagSet("getAddressLists", flag.ExitOnError)
	transferBlockCmd := flag.NewFlagSet("transfer", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("printChain", flag.ExitOnError)
	creatBlockChainCmd := flag.NewFlagSet("creatBlockChain", flag.ExitOnError)
	getBalanceCmd := flag.NewFlagSet("getBalance", flag.ExitOnError)
	createNewWalletCmd := flag.NewFlagSet("createNewWallet", flag.ExitOnError)

	flagFrom := transferBlockCmd.String("from", "", "origin address")
	flagTo := transferBlockCmd.String("to", "", "destination address")
	flagAmount := transferBlockCmd.String("amount", "", "transfer amount")

	flagcreatBlockChainWithAddress := creatBlockChainCmd.String("address", "", "create the address of genesis block")
	getBalanceWithAddress := getBalanceCmd.String("address", "", "inquire one's account")

	switch os.Args[1] {
	case "createNewWallet":
		err := createNewWalletCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "getAddressLists":
		err := getAddressListsCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "transfer":
		err := transferBlockCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "printChain":
		err := printChainCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "creatBlockChain":
		err := creatBlockChainCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "getBalance":
		err := getBalanceCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	default:
		printUsage()
		os.Exit(1)
	}
	if createNewWalletCmd.Parsed() {
		cli.createWallet()
	}
	if getAddressListsCmd.Parsed() {
		cli.AddressLists()
	}
	if transferBlockCmd.Parsed() {
		if *flagFrom == "" || *flagTo == "" || *flagAmount == "" {
			printUsage()
			os.Exit(1)
		}

		from := JSONToArray(*flagFrom)
		to := JSONToArray(*flagTo)
		amount := JSONToArray(*flagAmount)

		cli.send(from, to, amount)
	}
	if printChainCmd.Parsed() {
		//fmt.Println("output all blocks' information")
		cli.printChain()
	}
	if creatBlockChainCmd.Parsed() {
		if *flagcreatBlockChainWithAddress == "" {
			fmt.Println("the address shouldn't be null")
			printUsage()
			os.Exit(1)
		}
		cli.creatGenesisBlockChain(*flagcreatBlockChainWithAddress)
	}
	if getBalanceCmd.Parsed() {
		if *getBalanceWithAddress == "" {
			fmt.Println("the address shouldn't be null")
			printUsage()
			os.Exit(1)
		}
		cli.getBalance(*getBalanceWithAddress)
	}
}
