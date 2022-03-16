package main

import (
	"bytes"
	"fmt"
	"io"

	"github.com/yeka/zip"
)

func findPWD(pwdCh, result chan string) {

	for pwd := range pwdCh {
		if verbose {
			SayInfo(fmt.Sprintf("Check %v", pwd))
			// SayInfo(fmt.Sprintf("go routine count: %v", runtime.NumGoroutine()))
		}
		flag := unZip(zipfile, pwd)
		if flag {
			// notice controller that find the password
			result <- pwd
			close(result)
		}

		wg.Done()

	}
}

func BruteForce(pwdCh, result chan string, payload []string, min, max int) bool {
	for i := min; i <= max; i++ {
		var payloads [][]string
		payloads = genPayloads(payload, i)
		pwd := bruteforceFactory(pwdCh, result, payloads...)
		if pwd != "" {
			OutputResult("Brute-Force", pwd)
			if output != "" {
				WritePWD2File(output, zipfile, pwd)
			}
			return true
		}
	}
	OutputResult("Brute-Force", "")
	return false
}

func DictionaryAttack(pwdCh, result chan string, filename string) bool {

	pwd := dictionaryFactory(pwdCh, result, filename)

	if pwd != "" {
		if pwd != "" {
			OutputResult("Brute-Force", pwd)
			if output != "" {
				WritePWD2File(output, zipfile, pwd)
			}
			return true
		}
	}

	return false
}

func unZip(filename string, password string) bool {
	r, err := zip.OpenReader(filename)
	if err != nil {
		return false
	}
	defer r.Close()

	buffer := new(bytes.Buffer)

	for _, f := range r.File {
		if f.FileInfo().IsDir() || !f.IsEncrypted() {
			continue
		}
		f.SetPassword(password)
		r, err := f.Open()
		if err != nil {
			return false
		}
		defer r.Close()
		n, err := io.Copy(buffer, r)
		if n == 0 || err != nil {
			return false
		}
		break
	}
	return true

}
