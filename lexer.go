package groff

import (
	"bytes"
	"sync"
)

const EOF rune = 0

type Lexer struct {
	buf   bytes.Buffer
	nodes []Node
	src   *Src
}

type Src struct {
	src []rune
	pos *Pos
}

func (s *Src) Current() rune {
	return s.src[s.pos.current]
}

func (s *Src) Next() rune {
	if s.IsEOF() {
		return EOF
	}

	s.pos.current++
	return s.Current()
}

func (s *Src) IsEOF() bool {
	return s.EOF() <= s.pos.current
}

func (s *Src) EOF() int {
	return s.pos.last
}

type Pos struct {
	current int
	forward int
	last    int
}

func NewLexer(src string) *Lexer {
	return &Lexer{
		nodes: make([]Node, 0),
		src: &Src{
			src: []rune(src),
			pos: &Pos{
				last: len(src) - 1,
			},
		},
	}
}

func (l *Lexer) Buffer() string {
	return l.buf.String()
}

func (l *Lexer) Reset() {
	l.buf.Reset()
}

func (l *Lexer) Append(r rune) error {
	_, err := l.buf.WriteRune(r)
	return err
}

func (l *Lexer) AppendNode(n Node) {
	l.nodes = append(l.nodes, n)
}

func (l *Lexer) Scan() error {
	for ch := l.src.Current(); !l.src.IsEOF(); ch = l.src.Next() {
		switch l.Buffer() {
		case Comment:
			comment := l.ScanToEOL()
			l.AppendNode(NewCommentNode(comment))
			l.Reset()
		default:
			if err := l.Append(ch); err != nil {
				return err
			}
		}
	}
	return nil
}

type tmpBuffer struct {
	pool sync.Pool
}

func (t *tmpBuffer) Get() *bytes.Buffer {
	return t.pool.Get().(*bytes.Buffer)
}

func (t *tmpBuffer) Put(buf *bytes.Buffer) {
	buf.Reset()
	t.pool.Put(buf)
}

var tmpbuf = &tmpBuffer{
	pool: sync.Pool{
		New: func() interface{} {
			return &bytes.Buffer{}
		},
	},
}

func (l *Lexer) ScanToEOL() string {
	buf := tmpbuf.Get()
	defer tmpbuf.Put(buf)

	for !l.src.IsEOF() {
		ch := l.src.Next()
		if ch == '\n' {
			return buf.String()
		}
		buf.WriteRune(ch)
	}
	return buf.String()
}
