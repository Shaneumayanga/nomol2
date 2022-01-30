package sites

import (
	"bufio"
	"log"
	"os"
)

var sites []string

func GetSites() []string {
	return sites
}

func Prepare() {
	file, err := os.Open("./sites/sites.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		sites = append(sites, scanner.Text())
	}
}
