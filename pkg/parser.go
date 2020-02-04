package lexpar

import (
	"encoding/json"
	"fmt"
	"io"
)

type Parser struct {
	s   *Scanner
	buf struct {
		tok         Token
		lit         string
		isUnscanned bool
	}
}

func NewParser(r io.Reader) *Parser {
	return &Parser{s: NewScanner(r)}
}

func (p *Parser) scan() (tok Token, lit string) {

	if p.buf.isUnscanned {
		p.buf.isUnscanned = false
		return p.buf.tok, p.buf.lit
	}

	tok, lit = p.s.Scan()

	p.buf.tok, p.buf.lit = tok, lit

	return
}

func (p *Parser) unscan() { p.buf.isUnscanned = true }

func (p *Parser) parseNewline() bool {

	tok, _ := p.scan()
	if tok == NEWLINE {
		return true
	}
	p.unscan()
	return false
}

func (p *Parser) parseILLEGAL() error {

	tok, l := p.scan()
	if tok == ILLEGAL {
		return fmt.Errorf("Illegal token '%s'\n", l)
	}
	return nil
}

func (p *Parser) parseEOF() bool {

	tok, _ := p.scan()
	if tok == EOF {
		return true
	}
	p.unscan()
	return false
}

func (p *Parser) parseVal() *string {
	var val *string
	tok, lit := p.scan()
	if tok == VAL {
		v := lit
		val = &v
	} else {
		p.unscan()
		val = nil
	}
	return val
}

func (p *Parser) parseKey() *string {
	var key *string
	tok, lit := p.scan()
	if tok == KEY {
		k := lit
		key = &k
	} else {
		p.unscan()
		key = nil
	}
	return key
}

func (p *Parser) Parse() {
	//	var r R = make(R)

	unit := make(map[string]string)
	for {

		err := p.parseILLEGAL()
		if err != nil {
			fmt.Println(err)
			break
		}

		k := p.parseKey()

		v := p.parseVal()

		if k != nil && p.parseNewline() {
			//fmt.Printf("%s : %s\n", *k, *v)
			unit[*k] = *v
			continue
		}

		if k == nil && v == nil {
			if p.parseNewline() {
				fmt.Println("Starting new section")
				//fmt.Printf("%v", unit)
				j, _ := json.Marshal(unit)
				fmt.Println(string(j))
				unit = make(map[string]string)
				continue
			} else {
				break
			}
		}
	}
}
