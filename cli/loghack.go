package main

import (
	"github.com/comail/colog"
	"log"
)

func loghack(verbose bool) {
	b2i := map[bool]int{false: 0, true: 1}

	colog.SetDefaultLevel(colog.LTrace)

	colog.SetMinLevel([]colog.Level{ colog.LInfo, colog.LTrace }[b2i[verbose]])

	colog.SetFormatter(&colog.StdFormatter{
		Colors: true,
		Flag:   []int{ 0, log.Ldate | log.Ltime | log.Lshortfile }[b2i[verbose]],
	})
	colog.Register()
}
