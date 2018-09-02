package groff

func Parse(src string) ([]Node, error) {
	lexer := NewLexer(src)
	if err := lexer.Scan(); err != nil {
		return nil, err
	}
	return lexer.nodes, nil
}
