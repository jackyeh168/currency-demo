package main

import (
	dataimport "currency-demo/dataImport"
	"currency-demo/dataexport"
	"currency-demo/model"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"sync"
)

func managerThread(fileNameChannel chan<- string) {

	dirName := "./data/"
	var files []string

	err := filepath.Walk(dirName, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		fileNameChannel <- file
	}
	close(fileNameChannel)
}

func main() {
	// datagenerator.Generate()

	cpuNum := runtime.NumCPU()
	fileNameChannel := make(chan string)
	accessLogChannel := make(chan model.Data)

	var importWaitGroup, exportWaitGroup sync.WaitGroup

	go managerThread(fileNameChannel)

	for i := 0; i < cpuNum; i++ {
		importWaitGroup.Add(1)
		go dataimport.Import(fileNameChannel, accessLogChannel, &importWaitGroup)
	}

	for i := 0; i < 2; i++ {
		exportWaitGroup.Add(1)
		go dataexport.Export(accessLogChannel, &exportWaitGroup)
	}

	importWaitGroup.Wait()
	close(accessLogChannel)

	exportWaitGroup.Wait()
}
