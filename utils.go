package main

import (
	"fmt"
	"log"
	"os"
)

const (
	BANNER string = `
    ____                     _____   _
   / __ )____  ____  ____ __/__  /  (_)___
  / __  / __ \/ __ \/ __ '__ \/ /  / / __ \
 / /_/ / /_/ / /_/ / / / / / / /__/ / /_/ /
/_____/\____/\____/_/ /_/ /_/____/_/ .___/
                                  /_/      v%s
				(@wakaka6)
`
	VERSION string = "0.1.0"
)

var (
	DIGITAL []string = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	LOWER   []string = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	SYMBOL  []string = []string{"!", "\"", "#", "$", "%", "&", "'", "(", ")", "*", "+", ",", "-", ".", "/", ":", ";", "<", "=", ">", "?", "@", "[", "\\", "]", "^", "_", "`", "{", "|", "}", "~"}
	UPPER   []string = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
)

func ShowBanner() {
	fmt.Printf(BANNER, VERSION)
}

func SayInfo(msg string) {
	log.SetPrefix("[+] ")
	log.Println(msg)
}

func SayError(err error) {
	if err != nil {
		log.SetPrefix("[-] ")
		log.Println(err)
	}
}

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func genPayloads(payload []string, num int) [][]string {
	var payloads [][]string
	for i := 0; i < num; i++ {
		payloads = append(payloads, payload)
	}

	return payloads
}

func OutputResult(method, result string) {
	if result == "" {
		log.SetPrefix("[-] ")
		log.Printf("%s failed.\n", method)
	} else {
		log.SetPrefix("[*] ")
		log.Printf("%s sucessful, the password is %s", method, result)
	}
}

func WritePWD2File(outFile, filename, pwd string) {
    outcome := fmt.Sprintf("%v: %v\n", filename, pwd)
    WriteFile(outFile, outcome)
    if verbose{
        SayInfo(fmt.Sprintf("Secussful write result to %s", outFile))
    }
}

func WriteFile(outFile, str string) {
	cf, err := os.OpenFile(outFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	CheckError(err)
	defer cf.Close()

	_, err = cf.Write([]byte(str))
	CheckError(err)
}
