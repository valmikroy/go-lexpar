package lexpar

import (
	"bufio"
	"bytes"
	"io"
	"strings"
)

type Scanner struct {
	r   *bufio.Reader
	pos int
}

func NewScanner(r io.Reader) *Scanner {
	return &Scanner{
		r:   bufio.NewReader(r),
		pos: 0,
	}
}

func (s *Scanner) read() rune {
	ch, _, err := s.r.ReadRune()
	if err != nil {
		return eof
	}
	return ch
}

func (s *Scanner) unread() { _ = s.r.UnreadRune() }

func (s *Scanner) scanNewline() (tok Token, l string) {

	var buf bytes.Buffer

	for {

		if ch := s.read(); isNewLine(ch) {
			s.pos++
			buf.WriteRune(ch)
		} else {
			s.unread()
			break
		}

	}

	return NEWLINE, "\n"
}

func (s *Scanner) scanKey() (tok Token, l string) {

	var buf bytes.Buffer
	buf.WriteRune(s.read())

	for {
		if ch := s.read(); ch == eof {
			return EOF, ""
		} else if isDelimiter(ch) {
			s.unread()
			break
		} else if isNewLine(ch) {
			s.unread()
			break
		} else {
			buf.WriteRune(ch)
		}
	}

	s.pos++
	strTrim := strings.TrimSpace(buf.String())
	return KEY, strTrim
}

func (s *Scanner) scanVal() (tok Token, l string) {

	var buf bytes.Buffer

	for {
		if ch := s.read(); ch == eof {
			break
		} else if isNewLine(ch) {
			s.unread()
			break
		} else {
			buf.WriteRune(ch)
		}
	}

	s.pos++
	strTrim := strings.TrimSpace(buf.String())
	return VAL, strTrim
}

func (s *Scanner) Scan() (tok Token, l string) {
	ch := s.read()

	if ch == eof && s.pos >= 3 {
		return EOF, ""
	}

	if s.pos == 0 {
		s.unread()
		return s.scanKey()
	} else if s.pos == 1 {
		return s.scanVal()
	} else if s.pos >= 2 {
		return s.scanNewline()
	}
	return ILLEGAL, ""

}
