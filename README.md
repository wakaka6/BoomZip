# BoomZip

BoomZip is a commandline tools to crack by password protecting the Zip files. It uses go-routine to achieve a fast crack. BoomZiper supports highly customizable character sets for brute force cracking and dictionary attacks.




## Installation
First of all, clone the repositories
```sh
git clone https://github.com/wakaka6/BoomZip.git
```

Then, enter the directory and build it
```sh
cd BoomZip && go build
```

or 
```sh
go get github.com/wakaka6/BoomZip
```

## Help
```
./BoomZip -h

    ____                     _____   _
   / __ )____  ____  ____ __/__  /  (_)___
  / __  / __ \/ __ \/ __ '__ \/ /  / / __ \
 / /_/ / /_/ / /_/ / / / / / / /__/ / /_/ /
/_____/\____/\____/_/ /_/ /_/____/_/ .___/
                                  /_/      v0.1.0
				(@wakaka6)
Usage of ./BoomZip:
  -V	Show version
  -Version
    	Show version
  -b	Using bruteforce algorithem attack
  -burstMax int
    	start brute-force max length (default 8)
  -burstMin int
    	start brute-force min length (default 1)
  -d string
    	Using dictionary attack
  -i string
    	Path of the containing binary zip (e.g. xxx.zip)
  -l string
    	type [?1|?a|?A|?!|?#] 
    	?1 means 1234...
    	?a means abcd...
    	?A means ABCD...
    	?! means !@#$...
    	?# denote ?1?a?A?!
  -o string
    	Output about password of zipfile
  -p string
    	Using Custom letters to set brute-force payload (e.g a186)
  -t int
    	Set goroutine count for bruteforce the zip file (default 3)
  -v	verbose output
```

## Example Usage
Using Brute-force attacks and limit letters extent at 4~8
```sh
./BoomZip -v -i example.zip -l "?#" -b --burstMin=4 -o result.txt 
```
Use both the dictionary attack and the brute force attack with the specified custom character set, in which case the dictionary attack is performed first.
```sh
./Boomzip -v -i example.zip -p 1234 -b --burstMin=4 --burstMax=4 -d password.txt -o result.txt 
```


