package lexpar

//  lexical token
type Token int

const DELIMITER rune = ':'
const KEY_LENGTH int = 23

// kinds of token
const (
        ILLEGAL Token = iota
        EOF
        WS
        NEWLINE

        KEY
        VAL
)

type TokenInstance struct {
        Type    Token
        Literal string
}

func isWhitespace(ch rune) bool {
        return ch == ' ' || ch == '\t'
}
func isNewLine(ch rune) bool {
        return ch == '\n'
}

func isDelimiter(ch rune) bool {
        return ch == DELIMITER
}

var eof = rune(0)
