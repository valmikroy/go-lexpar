package lexpar

import (
	"io"
	"strings"
	"testing"
)

func TestParsing(t *testing.T) {
	var r io.Reader = strings.NewReader("A:AA\n B:BB\n C:CC\n\n D:DD\n")

	var p *Parser = NewParser(r)
	p.Parse()
}
