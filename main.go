package main

import (
	"flag"
	"github.com/rs/zerolog/log"
	"io"
	"os"
)

//  gocopy ­from /path/to/source ­to /path/to/dest ­offset 1024 ­limit 2048
func main() {
	var from string
	var to string
	var offset int64
	var limit int64
	flag.StringVar(&from, "from", "", "Source path")
	flag.StringVar(&to, "to", "", "Destination path")
	flag.Int64Var(&offset, "offset", 0, "Offset")
	flag.Int64Var(&limit, "limit", 0, "Limit")
	flag.Parse()
	gocopy(from, to, offset, limit)
}

func gocopy(from string, to string, offset int64, limit int64) {
	fileFrom, _ := os.Open(from)
	fileTo, _ := os.Create(to)

	_, err := fileFrom.Seek(offset, io.SeekStart)

	var reader io.Reader

	if limit > 0 {
		reader = io.LimitReader(fileFrom, limit)
	} else {
		reader = fileFrom
	}
	if err != nil {
		log.Print("Error : something terrible happen -> ", err)
	}

	_, err = io.Copy(fileTo, reader)

	if err != nil {
		log.Print("Error : something terrible happen -> ", err)
	}
}
