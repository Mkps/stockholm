package parse

import (
	"fmt"
	"os"
	"errors"
	"crypto/sha256"

	"stockholm/internal"
)

func hasFlag(args []string, flags ...string) bool {
	for _, arg := range args {
		for _, f := range flags {
			if arg == f {
				return true
			}
		}
	}
	return false
}

func isFlag(value string) bool {
	switch value {
	case "-r":
		return true
	case "-s":
		return true
	case "-h":
		return true
	case "-v":
		return true
	default:
		return false
	}
}

func getInfoFlags(argList []string) () {
	if hasFlag(argList, "-h", "--help") {
		fmt.Println("Stockholm help:")
		fmt.Println("-h, --help: Show this help menu")
		fmt.Println("-v, --version: Get version number")
		fmt.Println("-r, --reverse: Reverse the encryption on .ft files")
		fmt.Println("-s, --silent: Silences output during encryption or decryption")
		fmt.Println("Infection: ")
		fmt.Println("./stockholm [key]")
		fmt.Println("Reverse: ")
		fmt.Println("./stockholm --reverse [key]")
		os.Exit(0)
	}
	if hasFlag(argList, "-v", "--version") {
		fmt.Println("stockholm version 0.1.1")
		os.Exit(0)
	}
}

func getKeyValue(argList []string) ([]byte, error) {
	argc := len(argList)
	key_value := ""
	if argc >= 2 {
		key_value = argList[1]
		for index, value := range(argList) {
			if value == "-r" && index + 1 < argc {
				key_value = argList[index + 1]
			}
		}
	}

	if key_value == "" || isFlag(key_value) {
		return nil, errors.New("Error: Could not find key value")
	}
	if len(key_value) < 16 {
		return nil, errors.New(fmt.Sprintf("Error: Invalid key value [%s]", key_value))
	}
	hash := sha256.Sum256([]byte(key_value))
	return hash[:], nil
}

// We try to get the key in 3 ways: first from argv then from env then finally the builtin fallback
func Parser(argList []string) (*internal.StockholmOptions, error) {
	getInfoFlags(argList)
	ret := internal.StockholmOptions{
		Key: []byte(""),
		Silent: false,
		Reverse: false,
	}

	if hasFlag(argList, "-r", "--reverse") {
		ret.Reverse = true
	}

	key, err := getKeyValue(argList)
	if err != nil {
		return nil, err
	}
	ret.Key = key

	if hasFlag(argList, "-s", "--silent") {
		ret.Silent = true
	}
	return &ret, nil
}
