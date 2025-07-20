package strings

import(
	"testing"
	"strings"
)

const msgIndex = "%s (parte: %s) - índices: esperado (%d) <> encontrado (%d)."

func TestIndex(t *testing.T){
	tests := []struct{
		text string
		part string
		expected int
	}{
		{"Micilini is Best!", "Micilini", 0},
		{"", "", 0},
		{"Hey", "hey", -1},
		{"GoLang", "L", 2},
	}

	for _, test := range tests{
		t.Logf("Result: %v", test)
		current := strings.Index(test.text, test.part)

		if current != test.expected{
			t.Errorf(msgIndex, test.text, test.part, test.expected, current)
		}
	}
}