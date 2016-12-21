package main

import (
	"github.com/atomita/techlunch-deletelogfile/app"
	"github.com/jteeuwen/go-pkg-optarg"
	"log"
	"os"
	"time"
)

func main() {
	if len(optarg.Remainder) == 1 {
		cmd := optarg.Remainder[0]
		switch cmd {
		case "run":
			run()
		case "version":
			log.Printf("v0.0.1")
		}
	} else {
		optarg.Usage()
	}
}

func init() {
	dir, _ := os.Getwd()

	optarg.Header("run")
	optarg.Add("p", "period", "storage period days", 30)
	optarg.Add("d", "dir", "log files directory.", dir)
	optarg.Add("", "pattern", "file name pattern", "*.log")
	optarg.Add("", "really", "Really run", false)
	optarg.Add("", "verbose", "trace log", false)
	optarg.Header("version")

	ch := optarg.Parse()
	<-ch
}

func run() {
	// set defaults
	period := 30
	really := false
	verbose := false
	dir, _ := os.Getwd()
	pattern := "*.log"

	// read args
	for opt := range optarg.Parse() {
		switch opt.Name {
		case "period":
			period = opt.Int()
		case "really":
			really = opt.Bool()
		case "verbose":
			verbose = opt.Bool()
		case "dir":
			dir = opt.String()
		case "pattern":
			pattern = opt.String()
		}
	}

	loghack(verbose)

	now := time.Now()
	time := now.AddDate(0, 0, -1*period)

	log.Printf("debug: really: %v", really)
	log.Printf("debug: verbose: %v", verbose)
	log.Printf("debug: dir: %v", dir)
	log.Printf("debug: pattern: %v", pattern)
	log.Printf("debug: period: %v, %v", period, time)

	files, err := app.Find(dir+"/"+pattern, time)
	if err != nil {
		log.Printf("error: %v", err)
		os.Exit(1)
		return
	}

	if 0 == len(files) {
		log.Printf("info: The corresponding file is not found.")
		return
	}

	err = app.Delete(files, !really)
	if err != nil {
		log.Printf("error: %v", err)
		os.Exit(1)
		return
	}
}
