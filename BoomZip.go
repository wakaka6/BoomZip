package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
)

var (
	ver          bool
	verbose      bool
	goCount      int
	zipfile      string
	output       string
	burst        bool
	dictionary   string
	burstMin     int
	burstMax     int
	category     string
	customLetter string
)

func init() {
	log.SetFlags(log.LstdFlags)
	flag.BoolVar(&ver, "V", false, "Show version")
	flag.BoolVar(&ver, "Version", false, "Show version")
	flag.BoolVar(&verbose, "v", false, "verbose output")
	flag.IntVar(&goCount, "t", 3, "Set goroutine count for bruteforce the zip file")
	flag.StringVar(&zipfile, "i", "", "Path of the containing binary zip (e.g. xxx.zip)")
	flag.StringVar(&output, "o", "", "Output about password of zipfile")

	flag.BoolVar(&burst, "b", false, "Using bruteforce algorithem attack")
	flag.IntVar(&burstMin, "burstMin", 1, "start brute-force min length")
	flag.IntVar(&burstMax, "burstMax", 8, "start brute-force max length")
	flag.StringVar(&category, "l", "", "type [?1|?a|?A|?!|?#] \n?1 means 1234...\n?a means abcd...\n?A means ABCD...\n?! means !@#$...\n?# denote ?1?a?A?!")
	flag.StringVar(&customLetter, "p", "", "Using Custom letters to set brute-force payload (e.g a186)")

	flag.StringVar(&dictionary, "d", "", "Using dictionary attack")
}

var wg = sync.WaitGroup{}

func main() {
	defer wg.Wait()
	ShowBanner()
	if len(os.Args) < 2 {
		os.Args = append(os.Args, "-h")
	}
	flag.Parse()
	parseOption()
}

func parseOption() {
	if ver {
		fmt.Println(VERSION)
	}

	if zipfile == "" {
		SayInfo("You must given option -i")
		return
	}

	if !burst && dictionary == "" {
		burst = true // deafult attack method
	}

	pwdCh := make(chan string)
	result := make(chan string, 1)
	defer close(pwdCh)

	for i := 0; i < goCount; i++ { // start go-routine
		go findPWD(pwdCh, result)
	}

	if dictionary != "" {
		// dictionary-attack first if option -d and -b are set as the same time
		SayInfo("Start Dictionary Attack....")
		flag := DictionaryAttack(pwdCh, result, dictionary)
		if flag {
			return
		}
	}

	if burst {
		// Process character set
		var payload []string
		if customLetter != "" {
			payload = append(payload, strings.Split(customLetter, "")...)
		}
		if category != "" {
			letters := strings.Split(category, "?")[1:]
			for i, letter := range letters {
				if i == 0 && letter == "#" {
					payload = append(payload, DIGITAL...)
					payload = append(payload, LOWER...)
					payload = append(payload, UPPER...)
					payload = append(payload, SYMBOL...)
					break
				}
				switch letter {
				case "1":
					payload = append(payload, DIGITAL...)
				case "a":
					payload = append(payload, LOWER...)
				case "A":
					payload = append(payload, UPPER...)
				case "!":
					payload = append(payload, SYMBOL...)
				}
			}
		}

		// brute-force
		SayInfo("Start Brute-Force Attack....")
		flag := BruteForce(pwdCh, result, payload, burstMin, burstMax)
		if flag {
			return
		}
	}

	SayInfo("Not Found the password.(T T)")
	os.Exit(1)
}
