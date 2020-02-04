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

type Unit map[string]string

type Fru []Unit

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
func (p *Parser) carriageReturn() int {
	cnt := p.s.pos
	p.s.pos = 0
	return cnt
}

func (p *Parser) parseNewline() bool {

	tok, _ := p.scan()
	if tok == NEWLINE {
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
	var fru Fru = make(Fru, 0)
	unit := make(map[string]string)
	for {

		k := p.parseKey()

		v := p.parseVal()

		p.parseNewline()

		pos := p.carriageReturn()

		if k != nil && v != nil {
			unit[*k] = *v
		}

		if pos == 0 {
			break
		} else if pos == 3 {
			fru = append(fru, unit)
			unit = make(map[string]string)
		}

	}
	j, _ := json.Marshal(fru)
	fmt.Println(string(j))
}
