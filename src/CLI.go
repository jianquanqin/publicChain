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
	fmt.Println("\tcreateWallet -- create wallet")
	fmt.Println("\tgetAddressList -- print all wallet's address")
	fmt.Println("\tcreateBlockChain -address --genesis address")
	fmt.Println("\ttransfer -from FROM -to TO -amount AMOUNT --transaction details")
	fmt.Println("\tprintChain -- output block's information")
	fmt.Println("\tgetBalance -address -- output balance")
	fmt.Println("\ttest -- test tool")
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
	testCmd := flag.NewFlagSet("test", flag.ExitOnError)
	getAddressListCmd := flag.NewFlagSet("getAddressList", flag.ExitOnError)
	createWalletCmd := flag.NewFlagSet("createWallet", flag.ExitOnError)
	transferBlockCmd := flag.NewFlagSet("transfer", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("printChain", flag.ExitOnError)
	createBlockChainCmd := flag.NewFlagSet("createBlockChain", flag.ExitOnError)
	getBalanceCmd := flag.NewFlagSet("getBalance", flag.ExitOnError)

	flagFrom := transferBlockCmd.String("from", "", "origin address")
	flagTo := transferBlockCmd.String("to", "", "destination address")
	flagAmount := transferBlockCmd.String("amount", "", "transfer amount")

	flagcreateBlockChainWithAddress := createBlockChainCmd.String("address", "", "create the address of genesis block")
	getBalanceWithAddress := getBalanceCmd.String("address", "", "inquire one's account")

	switch os.Args[1] {
	case "test":
		err := testCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "getAddressList":
		err := getAddressListCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "createWallet":
		err := createWalletCmd.Parse(os.Args[2:])
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
	case "createBlockChain":
		err := createBlockChainCmd.Parse(os.Args[2:])
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
	if testCmd.Parsed() {

		cli.Test()
	}
	if getAddressListCmd.Parsed() {

		cli.GetAddressList()
	}
	if createWalletCmd.Parsed() {

		cli.createWallet()
	}
	if transferBlockCmd.Parsed() {
		if *flagFrom == "" || *flagTo == "" || *flagAmount == "" {
			printUsage()
			os.Exit(1)
		}

		from := JSONToArray(*flagFrom)
		to := JSONToArray(*flagTo)

		//verify the validity of address before transaction occurs
		for index, fromAddress := range from {
			if IsValidForAddress([]byte(fromAddress)) == false || IsValidForAddress([]byte(to[index])) == false {
				fmt.Println("Address is invalid")
				printUsage()
				os.Exit(1)
			}
		}

		amount := JSONToArray(*flagAmount)
		cli.send(from, to, amount)
	}
	if printChainCmd.Parsed() {
		//fmt.Println("output all blocks' information")
		cli.printChain()
	}
	if createBlockChainCmd.Parsed() {

		if IsValidForAddress([]byte(*flagcreateBlockChainWithAddress)) == false {
			fmt.Println("address is invalid")
			printUsage()
			os.Exit(1)
		}
		cli.creatGenesisBlockChain(*flagcreateBlockChainWithAddress)

	}
	if getBalanceCmd.Parsed() {
		if *getBalanceWithAddress == "" {
			fmt.Println("the address shouldn't be null")
			printUsage()
			os.Exit(1)
		}

		if IsValidForAddress([]byte(*getBalanceWithAddress)) == false {
			fmt.Println("address is invalid")
			printUsage()
			os.Exit(1)
		}

		cli.getBalance(*getBalanceWithAddress)
	}
}
