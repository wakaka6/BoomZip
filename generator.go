package main

import (
	"bufio"
	"io"
	"os"
	"strings"
	"time"
)

const WAITTIME time.Duration = 1 * time.Second

func bruteforceFactory(pwdCh, result chan string, sets ...[]string) string {
	lens := func(i int) int { return len(sets[i]) }
	for ix := make([]int, len(sets)); ix[0] < lens(0); nextIndex(ix, lens) {
		var r []string
		for j, k := range ix {
			r = append(r, sets[j][k])
		}

		wg.Add(1)

		select {
		case pwdCh <- strings.Join(r, ""):
			continue
		case pwd := <-result: // return the found password on sucessful
            wg.Done()
			return pwd
		}

	}

	// wait a moment for gorutine execute over
	select {
	case pwd := <-result:
		return pwd
	case <-time.After(WAITTIME):
		return ""
	}
}

func nextIndex(ix []int, lens func(int) int) {
	for j := len(ix) - 1; j >= 0; j-- {
		ix[j]++
		if j == 0 || ix[j] < lens(j) {
			return
		}
		ix[j] = 0
	}
}

func dictionaryFactory(pwdCh, result chan string, filename string) string {

	f, err := os.Open(filename)
	CheckError(err)
	defer f.Close()

	r := bufio.NewReader(f)
	for {
		line, _, err := r.ReadLine()
		if err == io.EOF {
			break
		}
		CheckError(err)

		wg.Add(1)

		select {
		case pwdCh <- string(line):
            continue
		case pwd := <-result:
            wg.Done()
			return pwd

		}
	}

	// wait a moment for gorutine execute over
	select {
	case pwd := <-result:
		return pwd
	case <-time.After(WAITTIME):
		return ""
	}
}
