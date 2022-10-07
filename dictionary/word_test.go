package dictionary

import "testing"

func TestNewWord(t *testing.T) {
	type scenario struct {
		description string
		text        string
		expectErr   bool
	}

	cases := []scenario{
		{"a-z allowed", "bananas", false},
		{"A-Z allowed", "Apples", false},
		{"empty not allowed", "", true},
		{"spaces not allowed", "some nice sentence", true},
		{"special characters not allowed", "łódź", true},
	}

	for _, c := range cases {
		t.Run(c.description, func(t *testing.T) {
			if _, err := NewWord(c.text); (err == nil && c.expectErr) || (err != nil && !c.expectErr) {
				t.Fatalf("err: %s %v", err, c.expectErr)
			}
		})
	}
}

func TestWord_IsZero(t *testing.T) {
	//....todo
}

func TestWord_MarshalJSON(t *testing.T) {
	//....todo
}

func TestWord_UnmarshalJSON(t *testing.T) {
	//....todo
}
