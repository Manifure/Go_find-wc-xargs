package main

import (
	"archive/tar"
	"compress/gzip"
	"flag"
	"io"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"
)

func addFileToTar(filename string, dir *string, wg *sync.WaitGroup) {
	defer wg.Done()

	currentTime := time.Now()
	timestamp := currentTime.Unix()
	timestampString := strconv.FormatInt(timestamp, 10)

	directoryToLog := filepath.Dir(filename) + "-"

	nameLogFile := filepath.Base(filename)

	fileNameWithoutLog := strings.TrimSuffix(nameLogFile, ".log")

	archiveName := *dir + directoryToLog + fileNameWithoutLog + "_" + timestampString + ".tar.gz"

	tarFile, err := os.Create(archiveName)
	if err != nil {
		log.Fatal(err)
	}
	defer tarFile.Close()

	gzipWriter := gzip.NewWriter(tarFile)
	defer gzipWriter.Close()

	tarWriter := tar.NewWriter(gzipWriter)
	defer tarWriter.Close()

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}

	header := &tar.Header{
		Name: nameLogFile,
		Mode: int64(stat.Mode()),
		Size: stat.Size(),
	}

	err = tarWriter.WriteHeader(header)
	if err != nil {
		log.Fatal(err)
	}

	_, err = io.Copy(tarWriter, file)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	var wg = new(sync.WaitGroup)

	dir := flag.String("a", "", "Directory")
	flag.Parse()

	if *dir != "" && (*dir)[len(*dir)-1] != '/' {
		println("Directory must end with /")
		os.Exit(1)
	}

	for i := 0; i < flag.NArg(); i++ {
		wg.Add(1)
		addFileToTar(flag.Arg(i), dir, wg)
	}
}
