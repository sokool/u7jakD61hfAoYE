package dictionary_test

import (
	. "rollee/dictionary"
	"testing"
)

func TestStorage(t *testing.T) {
	type scenario struct {
		description string
		words       []Word
		query       Word
		expect      Word
	}

	cases := []scenario{
		{
			description: "empty dataset query for none",
		},
		{
			description: "three elements query one that not exists",
			words:       []Word{"aqua", "orange", "atlas"},
			query:       "beer",
			expect:      "",
		},
		{
			description: "two elements query one",
			words:       []Word{"aqua", "orange"},
			query:       "aqua",
			expect:      "aqua",
		},
		{
			description: "put three different case same words",
			words:       []Word{"aqua", "Aqua", "AQUA"},
			query:       "aqua",
			expect:      "aqua",
		},
		{
			description: "search word by it's prefix",
			words:       []Word{"bingo", "aqua", "bananas", "ball"},
			query:       "ban",
			expect:      "bananas",
		},
		{
			description: "search most frequent word when only prefix given",
			words:       []Word{"aqua", "atlas", "Aquarium", "aspirin", "ATLAs"},
			query:       "a",
			expect:      "atlas",
		},
		{
			description: "search most frequent word when only prefix given",
			words:       []Word{"a", "ab", "ab"},
			query:       "a",
			expect:      "ab",
		},
	}

	for _, c := range cases {
		t.Run(c.description, func(t *testing.T) {
			if w := NewStorage(c.words...).Search(c.query).Recent(); w != c.expect {
				t.Fatalf("expected:%v got:%v", c.expect, w)
			}
		})
	}
}
