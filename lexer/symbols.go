package lexer

var SymbolsMap = map[string]string{
	"(": "(",
	")": ")",
	"{": "{",
	"}": "}",
	",": ",",
	".": ".",
	"-": "-",
	"+": "+",
	";": ";",
	"/": "/",
	"!": "!",
	"=": "=",
	"<": "<",
	">": ">",
}

var SymbolsNested = map[string]string{
	"!": "!=",
	"=": "==",
	"<": "<",
	">": ">=",
}

type SymbolLexer struct{}

func NewSymbolLexer() *SymbolLexer {
	return &SymbolLexer{}
}

func (s *SymbolLexer) Match(p Peekable) bool {
	char, err := p.PeekSingleRune()
	if err != nil {
		return false
	}

	_, exists := SymbolsMap[string(char)]
	return exists
}

func (s *SymbolLexer) Lex(r Readable) (result []rune) {
	chars, err := r.PeekRunes(2)
	if err != nil {
		return nil
	}

	if _, ok := SymbolsNested[string(chars)]; ok {
		result, err = r.ReadRunes(2)
	}

	if _, ok := SymbolsMap[string(chars[0])]; ok {
		result, err = r.ReadRunes(1)
	}

	if err != nil {
		return nil
	}

	return result
}