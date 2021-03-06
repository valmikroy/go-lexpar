package lexpar

import (
	"io"
	"strings"
	"testing"
)

func TestLexingLines(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, input io.Reader, desired_out []TokenInstance) {

		t.Helper()
		s := NewScanner(input)
		for _, ti := range desired_out {
			token, literal := s.Scan()
			if token != ti.Type {
				t.Error("Token mismatch")
				t.Error("Expected token ", ti.Type, " but got ", token, " with literal value ", literal)
			} else if literal != ti.Literal {
				t.Error("Literal mismatch")
				t.Error("Expected literal ", ti.Literal, " but got ", literal, " with token type", token)
			}

		}
	}

	t.Run("First line starting with char", func(t *testing.T) {
		t.Logf("Parsing string  'FRU Device Description : Custom device (ID 00)'")
		var r io.Reader = strings.NewReader("FRU Device Description : Custom device (ID 00)\n")
		var correct_parsed []TokenInstance
		correct_parsed = []TokenInstance{
			TokenInstance{KEY, "FRU Device Description"},
			TokenInstance{VAL, "Custom device (ID 00)"},
			TokenInstance{NEWLINE, "\n"},
		}

		assertCorrectMessage(t, r, correct_parsed)

	})
	t.Run("Line starting with space", func(t *testing.T) {
		t.Logf("Parsing string ' Board Serial          : QZZZ1213142424'")
		var r io.Reader = strings.NewReader(" Board Serial          : QZZZ1213142424\n")
		var correct_parsed []TokenInstance
		correct_parsed = []TokenInstance{
			TokenInstance{KEY, "Board Serial"},
			TokenInstance{VAL, "QZZZ1213142424"},
			TokenInstance{NEWLINE, "\n"},
		}

		assertCorrectMessage(t, r, correct_parsed)

	})

	t.Run("Value has delimiter character", func(t *testing.T) {
		t.Logf("Parsing string ' Board Mfg Date        : Sun May 10 21:41:00 2010'")
		var r io.Reader = strings.NewReader(" Board Mfg Date        : Sun May 10 21:41:00 2010\n")
		var correct_parsed []TokenInstance
		correct_parsed = []TokenInstance{
			TokenInstance{KEY, "Board Mfg Date"},
			TokenInstance{VAL, "Sun May 10 21:41:00 2010"},
			TokenInstance{NEWLINE, "\n"},
		}

		assertCorrectMessage(t, r, correct_parsed)

	})
	t.Run("Key has no value", func(t *testing.T) {
		t.Logf("Parsing string ' Board FRU ID          :'")
		var r io.Reader = strings.NewReader(" Board FRU ID          :\n")
		var correct_parsed []TokenInstance
		correct_parsed = []TokenInstance{
			TokenInstance{KEY, "Board FRU ID"},
			TokenInstance{VAL, ""},
			TokenInstance{NEWLINE, "\n"},
		}

		assertCorrectMessage(t, r, correct_parsed)

	})
	t.Run("Key has invalid key length", func(t *testing.T) {
		var r io.Reader = strings.NewReader("A : BBB\n")
		var correct_parsed []TokenInstance
		correct_parsed = []TokenInstance{
			TokenInstance{KEY, "A"},
			TokenInstance{VAL, "BBB"},
			TokenInstance{NEWLINE, "\n"},
		}

		assertCorrectMessage(t, r, correct_parsed)

	})
	t.Run("With no value field", func(t *testing.T) {
		var r io.Reader = strings.NewReader("ABCD\n")
		var correct_parsed []TokenInstance
		correct_parsed = []TokenInstance{
			TokenInstance{KEY, "ABCD"},
			TokenInstance{VAL, ""},
			TokenInstance{NEWLINE, "\n"},
		}

		assertCorrectMessage(t, r, correct_parsed)

	})
	/*
		t.Run("With no value for both key and value", func(t *testing.T) {
			var r io.Reader = strings.NewReader("\n")
			var correct_parsed []TokenInstance
			correct_parsed = []TokenInstance{
				TokenInstance{KEY, ""},
				TokenInstance{VAL, ""},
				TokenInstance{NEWLINE, "\n"},
			}

			assertCorrectMessage(t, r, correct_parsed)

		}) */
}
