package board

import "testing"

func TestTrieStartWith(t *testing.T) {
	trie := Trie{}

	words := []string{
		"ЯЗВЯЩИЙ",
		"ЯЗЕВ",
		"ЯЗОВ",
		"ЯЗЫК",
		"ЯЗЬ",
		"ПРОГРАММА",
	}
	for _, word := range words {
		trie.AddString([]rune(word))
	}

	cases := []struct {
		in   string
		want bool
	}{
		{"ЯЗВЯЩИЙ", true},
		{"ЯЗЕВ", true},
		{"ЯЗОВ", true},
		{"ЯЗЫК", true},
		{"ЯЗЬ", true},
		{"ПРОГРАММА", true},

		{"КНЯ", false},
		{"Я", true},
		{"ЯЗ", true},
		{"ПРОГ", true},
	}

	for _, c := range cases {
		got := trie.StartWith([]rune(c.in))
		if got != c.want {
			t.Errorf("StartWith(%q) == %t, want %t", c.in, got, c.want)
		}
	}
}

func TestTrieContains(t *testing.T) {
	trie := Trie{}

	words := []string{
		"КНЯЖЕНИЕ",
		"КНЯЖЕНИКА",
		"КНЯЖЕСКИЙ",
		"КНЯЖЕСТВО",
	}
	for _, word := range words {
		trie.AddString([]rune(word))
	}

	cases := []struct {
		in   string
		want bool
	}{
		{"АРБУЗ", false},
		{"КНЯЖЕНИЕ", true},
		{"КНЯЖЕНИКА", true},
		{"КНЯЖЕСКИЙ", true},
		{"КНЯЖЕСТВО", true},
		{"КНЯ", false},
		{"КНЯЖ", false},
		{"КНЯЖЕ", false},
		{"КНЯЖЕС", false},
		{"КНЯЖЕК", false},
	}

	for _, c := range cases {
		got := trie.Contains([]rune(c.in))
		if got != c.want {
			t.Errorf("Contains(%q) == %t, want %t", c.in, got, c.want)
		}
	}
}
