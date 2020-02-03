package lexpar

import (
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
	return false
}

func (p *Parser) parseEOF() bool {

	tok, _ := p.scan()
	if tok == EOF {
		return true
	}
	return false
}

func (p *Parser) parseVal() (*string, error) {
	var val *string
	tok, lit := p.scan()
	if tok == VAL {
		v := lit
		val = &v
	} else {
		p.unscan()
		return nil, nil
	}

	return val, nil
}

func (p *Parser) parseKey() (*string, error) {
	var key *string
	tok, lit := p.scan()
	if tok == KEY {
		k := lit
		key = &k
	} else {
		p.unscan()
		return nil, nil
	}

	return key, nil
}

func (p *Parser) Parse() error {
	//	var r R = make(R)

	for {
		k, err := p.parseKey()
		if err != nil {
			return err
		}

		v, err := p.parseVal()
		if err != nil {
			return err
		}

		//		fmt.Printf("%s : %s\n", *k, *v)
		if p.parseNewline() && k != nil {
			fmt.Printf("%s : %s", *k, *v)
		} else {
			fmt.Println("Starting new section")
		}

		if p.parseEOF() {
			fmt.Println("END")
			return nil
		}

		if k == nil && v == nil {
			return fmt.Errorf("stopping runaway loop")
		}

	}
}
