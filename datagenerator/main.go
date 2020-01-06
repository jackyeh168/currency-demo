package datagenerator

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/Pallinder/go-randomdata"
)

const totalLine = 10
const fileCount = 5
const bucketCount = 10000

func Generate() {
	for i := 0; i < fileCount; i++ {
		createFile(i)
	}
}

func createFile(i int) {
	fileName := fmt.Sprintf("./data/data%d.txt", i)

	f, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	for i := 0; i < totalLine; i++ {
		fmt.Fprintf(f, "%s\t%d\n", randomdata.IpV4Address(), r1.Intn(100))
		if err != nil {
			log.Fatal(err)
		}
	}
	f.Close()
}
