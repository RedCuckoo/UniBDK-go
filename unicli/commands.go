package main

import (
	"bufio"
	"context"
	"fmt"
	"github.com/RedCuckoo/UniBDK-go/proto/generated/proto"
	"github.com/RedCuckoo/UniBDK-go/unikeys"
	"github.com/pkg/errors"
	"github.com/tyler-smith/go-bip39"
	"golang.org/x/term"
	"google.golang.org/grpc"
	"os"
	"strings"
	"syscall"
)

func passwordPrompt(newPassword bool) ([]byte, error) {
	if newPassword {
		fmt.Printf("set password (min 8 characters len):")
	} else {
		fmt.Printf("enter password:")
	}
	var err error
	bytePassword := make([]byte, 0)
	exitCounter := 0
	for len(string(bytePassword)) < 8 && exitCounter < 3 {
		exitCounter++
		bytePassword, err = term.ReadPassword(syscall.Stdin)
		if err != nil {
			fmt.Printf("\ninvalid password: %s", err.Error())
			continue
		}
		if len(string(bytePassword)) < 8 {
			fmt.Printf("\ninvalid password length\ntry again:")
		}
	}

	if exitCounter == 3 {
		fmt.Println("\nthree invalid tries to set password")
		return nil, errors.New("three invalid tries to set password")
	}

	return bytePassword, nil
}

func createCommand() {
	err := os.Mkdir(STORAGE_DIR, 0777)
	if err != nil {
		if !strings.Contains(err.Error(), "file exists") {
			fmt.Print(err)
			return
		}
	}

	var storageFile *os.File
	defer storageFile.Close()

	storageFile, err = os.Open(STORAGE_FILE)
	if err != nil {
		if !strings.Contains(err.Error(), "no such file") {
			fmt.Print(err)
			return
		}
		// create new
		storageFile, err = os.Create(STORAGE_FILE)
		if err != nil {
			fmt.Print(err)
			return
		}
		fmt.Println("initializing new storage...")
		//fmt.Print("set password (min 8 characters len):")

		bytePassword, err := passwordPrompt(true)
		if err != nil {
			return
		}

		fmt.Println()
		mnemonic, err := unikeys.GenerateMnemonic()
		if err != nil {
			fmt.Printf("failed to generate mnemonic: %s\n", err.Error())
			return
		}

		fmt.Printf("memorize or write down in secure place generated mnemonic (enter to continue):\n" +
			mnemonic)
		_, err = fmt.Scanln()
		if err != nil {
			fmt.Printf("failed to read input: %s\n", err.Error())
			return
		}

		_, err = unikeys.NewKey(mnemonic)
		if err != nil {
			fmt.Printf("failed to create key from generated mnemonic: %s\n", err.Error())
			return
		}

		encrypt, err := Encrypt(mnemonic, string(bytePassword))
		if err != nil {
			fmt.Printf("failed to encrypt wallet, needed better password: %s\n", err.Error())
			return
		}

		_, err = storageFile.Write([]byte(encrypt))
		if err != nil {
			fmt.Printf("failed to save encrypted wallet to file: %s\n", err.Error())
			return
		}

		fmt.Println("successfully created new wallet")
	} else {
		line, b, err := bufio.NewReader(storageFile).ReadLine()
		if err != nil {
			// empty file
			if strings.Contains(err.Error(), "EOF") {
				err := os.Remove(STORAGE_FILE)
				if err != nil {
					fmt.Printf("failed to remove empty storage file: %s\n", err.Error())
					return
				}
				createCommand()
			}
			return
		}
		if b {
			fmt.Println("removing corrupted storage")
			err := os.Remove(STORAGE_FILE)
			if err != nil {
				fmt.Printf("failed to remove empty storage file: %s\n", err.Error())
				return
			}
			createCommand()
		}

		fmt.Print("wallet exists, would you like to reset it? (y/n): ")
		var response string
		_, err = fmt.Scanln(&response)
		if err != nil {
			fmt.Printf("failed to read input: %s\n", err.Error())
			return
		}

		if response == "y" {
			bytePassword, err := passwordPrompt(false)
			if err != nil {
				return
			}

			decrypt, err := Decrypt(string(line), string(bytePassword))
			if err != nil {
				fmt.Printf("failed to decrypt wallet: %s\n", err.Error())
				return
			}
			_, err = bip39.NewSeedWithErrorChecking(decrypt, "")
			if err != nil {
				fmt.Println("\ninvalid password")
				return
			}

			err = os.Remove(STORAGE_FILE)
			if err != nil {
				fmt.Printf("failed to reset wallet: %s\n", err.Error())
				return
			}
			createCommand()
		}
	}
}

func importCommand(mnemonic string) {
	err := os.Mkdir(STORAGE_DIR, 0777)
	if err != nil {
		if !strings.Contains(err.Error(), "file exists") {
			fmt.Print(err)
			return
		}
	}

	var storageFile *os.File
	defer storageFile.Close()

	storageFile, err = os.Open(STORAGE_FILE)
	if err != nil {
		fmt.Println("importing new storage...")
		//fmt.Print("set password (min 8 characters len):")

		bytePassword, err := passwordPrompt(true)
		if err != nil {
			return
		}

		fmt.Println()

		_, err = unikeys.NewKey(mnemonic)
		if err != nil {
			fmt.Printf("failed to create key from mnemonic: %s\n", err.Error())
			return
		}

		encrypt, err := Encrypt(mnemonic, string(bytePassword))
		if err != nil {
			fmt.Printf("failed to encrypt wallet, needed better password: %s\n", err.Error())
			return
		}

		_, err = storageFile.Write([]byte(encrypt))
		if err != nil {
			fmt.Printf("failed to save encrypted wallet to file: %s\n", err.Error())
			return
		}

		fmt.Println("successfully created new wallet")
	}
}

func getWallet() (*unikeys.Key, error) {
	defer fmt.Println()
	var storageFile *os.File
	defer storageFile.Close()

	storageFile, err := os.Open(STORAGE_FILE)
	if err != nil {
		fmt.Println("failed to open storage file")
		return nil, err
	}

	line, _, err := bufio.NewReader(storageFile).ReadLine()
	if err != nil {
		// empty file
		if strings.Contains(err.Error(), "EOF") {
			err := os.Remove(STORAGE_FILE)
			if err != nil {
				fmt.Printf("failed to remove empty storage file: %s\n", err.Error())
				return nil, err
			}
		}
		return nil, err
	}

	bytePassword, err := passwordPrompt(false)
	if err != nil {
		return nil, err
	}

	mnemonic, err := Decrypt(string(line), string(bytePassword))
	if err != nil {
		fmt.Printf("failed to decrypt wallet: %s\n", err.Error())
		return nil, err
	}
	_, err = bip39.NewSeedWithErrorChecking(mnemonic, "")

	key, err := unikeys.NewKey(mnemonic)
	if err != nil {
		fmt.Printf("failed to open wallet: %s\n", err.Error())
		return nil, err
	}

	return key, nil

}

func balance(blockchain int) {
	client, err := grpc.Dial("localhost:1706", grpc.WithInsecure(), grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(104857600)))
	if err != nil {
		fmt.Println("failed to connect to indexing server")
	}
	indexerClient := proto.NewIndexerClient(client)

	key, err := getWallet()
	if err != nil {
		return
	}

	var address string

	switch proto.Blockchains(blockchain) {
	case proto.Blockchains_Bitcoin:
		address, err = key.BtcAddress()
		if err != nil {
			fmt.Println("failed to get btc address")
		}
	case proto.Blockchains_Ethereum:
		address, err = key.EthAddress()
		if err != nil {
			fmt.Println("failed to get eth address")
		}
	}

	balances, err := indexerClient.AddressBalance(context.Background(), &proto.AddressBalanceRequest{
		Address:    address,
		Blockchain: proto.Blockchains(blockchain),
	})
	if err != nil {
		fmt.Printf("request failed:%s", err)
		return
	}

	fmt.Println(balances)
}

func sign(message string) {
	key, err := getWallet()
	if err != nil {
		return
	}
	blockchain, err := key.WalletByBlockchain("btc")
	if err != nil {
		fmt.Printf("failed to get wallet: %s\n", err.Error())
		return
	}
	data, err := blockchain.SignData([]byte(message))
	if err != nil {
		fmt.Printf("failed to sign data: %s\n", err.Error())
		return
	}
	fmt.Println(data)
}

func getHelp() string {
	return fmt.Sprintf("unibdk available commands:\n")
}
