package main

import (
	"flag"
	"fmt"
	"log"
	"time"
)

var (
	timeOption = flag.Int64("t", 0, "timestamp")
	full       = flag.Bool("f", false, "full format")
)

func main() {
	flag.Parse()
	tO := *timeOption
	if tO == 0 {
		log.Fatal("t option is mandatory")
	}

	t := time.Unix(tO, 0)

	full := *full
	if full == true {
		year, month, day := t.Date()
		fmt.Printf("%04d/%02d/%02dT%02d:%02d:%02d+09:00\n", year, month, day, t.Hour(), t.Minute(), t.Second())

	} else {
		fmt.Printf("%02d:%02d:%02d\n", t.Hour(), t.Minute(), t.Second())
	}
}
