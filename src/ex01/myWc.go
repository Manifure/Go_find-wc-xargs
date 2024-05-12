package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sync"
	"unicode/utf8"
)

func countWords(filePath string, wg *sync.WaitGroup) {
	defer wg.Done()
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	wCount := 0
	for scanner.Scan() {
		wCount++
	}
	fmt.Printf("%d\t%s\n", wCount, filePath)
}

func countLines(filePath string, wg *sync.WaitGroup) {
	defer wg.Done()
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	lCount := 0
	for scanner.Scan() {
		lCount++
	}
	fmt.Printf("%d\t%s\n", lCount, filePath)
}

func countChars(filePath string, wg *sync.WaitGroup) {
	defer wg.Done()
	text, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	count := utf8.RuneCount(text)
	fmt.Printf("%d\t%s\n", count, filePath)
}

func main() {
	var wg = new(sync.WaitGroup)

	wFlag := flag.Bool("w", false, "count words")
	lFlag := flag.Bool("l", false, "count lines")
	mFlag := flag.Bool("m", false, "count characters")
	flag.Parse()

	if flag.NFlag() > 1 {
		println("Too many arguments!\nOnly one flag can be used")
		flag.PrintDefaults()
		os.Exit(1)
	}

	if *wFlag || flag.NFlag() < 1 {
		for i := 0; i < flag.NArg(); i++ {
			wg.Add(1)
			go countWords(flag.Arg(i), wg)
		}
	} else if *lFlag {
		for i := 0; i < flag.NArg(); i++ {
			wg.Add(1)
			go countLines(flag.Arg(i), wg)
		}
	} else if *mFlag {
		for i := 0; i < flag.NArg(); i++ {
			wg.Add(1)
			go countChars(flag.Arg(i), wg)
		}
	}
	wg.Wait()
}
