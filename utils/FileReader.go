package utils

import (
	"bufio"
	"io"
	"log"
	"os"
)

type FileReader struct {
	Filepath string
}

func (x *FileReader) ReadLines() []string {

	f, err := os.Open(x.Filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	s.Split(bufio.ScanLines)

	if err := s.Err(); err != nil {
		log.Fatal(err)
	}

	lines := []string{}
	for s.Scan() {
		line := s.Text()
		lines = append(lines, line)
	}

	return lines
}

func (x *FileReader) ReadText() string {

	f, err := os.Open(x.Filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	r := bufio.NewReader(f)
	b, err := io.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}

	return string(b)
}
