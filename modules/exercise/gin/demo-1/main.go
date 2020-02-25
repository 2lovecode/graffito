package main

import (
	"github.com/EDDYCJY/fake-useragent"
	"log"
)
func main() {
	random := browser.Random()
	log.Printf("Random: %s", random)

	chrome := browser.Chrome()
	log.Printf("Chrome: %s", chrome)


}
