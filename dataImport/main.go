package dataimport

import (
	"bufio"
	"currency-demo/model"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

func Import(fileNameChannel <-chan string, accessLogChannel chan<- model.Data, wg *sync.WaitGroup) {

	for {
		fileName, ok := <-fileNameChannel
		if !ok {
			break
		}

		f, _ := os.Open(fileName)
		scanner := bufio.NewScanner(f)

		for scanner.Scan() {
			line := scanner.Text()
			if line == "" {
				continue
			}
			arr := strings.Split(line, "\t")

			num, _ := strconv.Atoi(arr[1])
			fmt.Println(arr[0])
			accessLogChannel <- model.Data{arr[0], num}
		}
	}

	wg.Done()
}
