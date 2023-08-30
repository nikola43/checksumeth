package main

import (
	"fmt"
	"os"
	"regexp"

	"github.com/ethereum/go-ethereum/common"
)

func main() {
	// get arguments
	args := os.Args

	// validate number of arguments
	if len(args) < 2 {
		fmt.Println("no arguments provided")
		return
	}

	if len(args) > 2 {
		fmt.Println("too many arguments")
		return
	}

	// check if argument is valid hex number
	if !IsValidAddress(args[1]) {
		fmt.Println("invalid address")
		return
	}

	// checksum address
	address := common.HexToAddress(args[1])
	checksumAddress := common.HexToAddress(address.Hex()).Hex()
	fmt.Println(checksumAddress)

}

func IsValidAddress(iaddress interface{}) bool {
	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
	switch v := iaddress.(type) {
	case string:
		return re.MatchString(v)
	case common.Address:
		return re.MatchString(v.Hex())
	default:
		return false
	}
}
