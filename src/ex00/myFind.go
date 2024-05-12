package main

import (
	"flag"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

func main() {
	slFlag := flag.Bool("sl", false, "Search symlinks")
	dFlag := flag.Bool("d", false, "Search dirs")
	fFlag := flag.Bool("f", false, "Search files")
	extFlag := flag.String("ext", "", "Search extensions")
	flag.Parse()

	if *extFlag != "" && !*fFlag {
		println("-ext flag only works if the -f flag is specified")
		os.Exit(1)
	}

	if !*slFlag && !*dFlag && !*fFlag {
		*slFlag, *dFlag, *fFlag = true, true, true
	}

	err0 := filepath.WalkDir(flag.Arg(0), func(path string, d fs.DirEntry, err error) error {

		if err != nil {
			if os.IsPermission(err) {
				fmt.Printf("permission denied: ")
			} else {
				log.Fatal(err)
			}
		}

		if *dFlag && d.IsDir() {
			fmt.Printf("%s\n", path)
		}

		if *fFlag && !d.IsDir() {
			if *extFlag != "" {
				if filepath.Ext(path) == "."+*extFlag {
					fmt.Printf("%s\n", path)
				}
			} else {
				fmt.Printf("%s\n", path)
			}
		}

		sLink, err1 := filepath.EvalSymlinks(path)
		if *slFlag && d.Type()&os.ModeSymlink != 0 {
			fmt.Printf("%s -> ", path)
			if err1 != nil {
				fmt.Println("[broken]")
			} else {
				fmt.Println(sLink)
			}
		}

		return nil
	})
	if err0 != nil {
		log.Fatal(err0)
	}
}
