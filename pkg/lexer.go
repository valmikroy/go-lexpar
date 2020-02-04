package lexpar

import (
	"bufio"
	"bytes"
	"io"
	"strings"
)

type Scanner struct {
	r *bufio.Reader
}

func NewScanner(r io.Reader) *Scanner {
	return &Scanner{
		r: bufio.NewReader(r),
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

func (s *Scanner) scanWhiteSpace() (tok Token, l string) {

	var buf bytes.Buffer
	buf.WriteRune(s.read())

	for {
		if ch := s.read(); ch == eof {
			break
		} else if !isWhitespace(ch) {
			s.unread()
			break
		} else {
			buf.WriteRune(ch)
		}
	}

	return WS, buf.String()
}

func (s *Scanner) scanKey() (tok Token, l string) {

	var buf bytes.Buffer
	buf.WriteRune(s.read())
	var cnt int = 1

	for {
		if ch := s.read(); ch == eof {
			break
		} else if isDelimiter(ch) && cnt == KEY_LENGTH {
			s.unread()
			break
		} else if isNewLine(ch) {
			s.unread()
			break
		} else {
			buf.WriteRune(ch)
		}
		cnt++
	}

	strTrim := strings.TrimSpace(buf.String())
	if cnt < KEY_LENGTH {
		return ILLEGAL, strTrim
	}
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

	strTrim := strings.TrimSpace(buf.String())
	return VAL, strTrim
}

func (s *Scanner) Scan() (tok Token, l string) {
	ch := s.read()

	if isDelimiter(ch) {
		return s.scanVal()
	} else if !isNewLine(ch) {
		s.unread()
		return s.scanKey()
	}

	switch ch {
	case eof:
		return EOF, ""
	case '\n':
		return NEWLINE, string(ch)
	}
	return ILLEGAL, strings.TrimSpace(string(ch))
}
