package main

import (
	"github.com/comail/colog"
	"log"
)

func loghack(verbose bool) {
	colog.SetDefaultLevel(colog.LTrace)
	if verbose {
		colog.SetMinLevel(colog.LTrace)
	} else {
		colog.SetMinLevel(colog.LWarning)
	}
	colog.SetFormatter(&colog.StdFormatter{
		Colors: true,
		Flag:   log.Ldate | log.Ltime | log.Lshortfile,
	})
	colog.Register()
}
