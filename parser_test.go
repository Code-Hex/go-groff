package groff

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/k0kubun/pp"
)

func TestParse(t *testing.T) {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	f, err := os.Open("file.man")
	if err != nil {
		return err
	}
	defer f.Close()
	src, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}
	nodes, err := Parse(string(src))
	if err != nil {
		return err
	}
	pp.Println(nodes)
	return nil
}
