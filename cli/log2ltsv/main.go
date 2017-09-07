package main

import (
	"bufio"
	"fmt"
	"github.com/tkuchiki/log2ltsv"
	"gopkg.in/alecthomas/kingpin.v2"
	"io"
	"log"
	"os"
)

var (
	format   = kingpin.Flag("format", "access log format (apache or nginx)").Default("nginx").String()
	filePath = kingpin.Flag("file", "access log").String()
)

func readFD(fpath string) (io.Reader, error) {
	s, err := os.Stdin.Stat()
	if err != nil {
		return nil, err
	}

	if s.Size() > 0 {
		return os.Stdin, nil
	}

	return os.Open(fpath)
}

func main() {
	kingpin.CommandLine.Help = "apache and nginx access log to ltsv format"
	kingpin.Version("0.1.0")
	kingpin.Parse()

	var parser log2ltsv.Parser
	switch *format {
	case "nginx":
		parser = log2ltsv.NewNginxParser()
	case "apache":
		parser = log2ltsv.NewApacheParser()
	}

	fd, ferr := readFD(*filePath)
	if ferr != nil {
		log.Fatal(ferr)
	}

	scanner := bufio.NewScanner(fd)

	for scanner.Scan() {
		ltsv, lerr := parser.ParseAndOutput(scanner.Text())
		if lerr != nil {
			log.Fatal(lerr)
		}
		fmt.Println(ltsv)
	}
}
