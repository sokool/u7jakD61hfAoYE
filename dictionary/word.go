package dictionary

import (
	"fmt"
	"regexp"
	"strings"
)

type Word string

var alpha = regexp.MustCompile(`^[a-zA-Z]+$`)

func NewWord(s string) (Word, error) {
	if !alpha.MatchString(s) {
		return "", fmt.Errorf("invalid string format, %s expected", alpha.String())
	}

	return Word(strings.ToLower(s)), nil
}

func (w Word) IsZero() bool {
	return len(w) == 0
}

func (w Word) MarshalJSON() ([]byte, error) {
	if w.IsZero() {
		return []byte(`null`), nil
	}

	return []byte(`"` + w + `"`), nil
}

func (w *Word) UnmarshalJSON(b []byte) (err error) {
	s := len(b)
	if s < 2 || (b[0] != '"' && b[0] != b[s-1]) {
		return fmt.Errorf("invalid word json format, string required")
	}

	if *w, err = NewWord(string(b[1 : s-1])); err != nil {
		return err
	}

	return nil
}

func (w Word) append(r rune) Word {
	return Word(string(w) + string(r))
}

type Words map[Word]int

func (w Words) Recent() (t Word) {
	var r int
	for s, n := range w {
		if n >= r {
			r, t = n, s
		}
	}
	return
}
