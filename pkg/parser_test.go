package lexpar

import (
	"io"
	"io/ioutil"
	"strings"
	"testing"
)

func TestParsing(t *testing.T) {

	content, _ := ioutil.ReadFile("text.txt")
	input := string(content)
	var r io.Reader = strings.NewReader(input)
	var p *Parser = NewParser(r)
	p.Parse()
}
