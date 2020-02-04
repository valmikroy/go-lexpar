package lexpar

import (
	"io"
	"io/ioutil"
	"strings"
	"testing"
)

func TestFruParsing(t *testing.T) {

	content, _ := ioutil.ReadFile("resource/fru.out")
	input := string(content)
	var r io.Reader = strings.NewReader(input)
	var p *Parser = NewParser(r)
	p.Parse()
}

func TestLanParsing(t *testing.T) {

	content, _ := ioutil.ReadFile("resource/lan_print.out")
	input := string(content)
	var r io.Reader = strings.NewReader(input)
	var p *Parser = NewParser(r)
	p.Parse()
}

func TestChassisStatusParsing(t *testing.T) {

	content, _ := ioutil.ReadFile("resource/chassis_status.out")
	input := string(content)
	var r io.Reader = strings.NewReader(input)
	var p *Parser = NewParser(r)
	p.Parse()
}
