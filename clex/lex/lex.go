package lex

import (
	"unicode/utf8"
)

type itemType int
type statefn func(*lexer) statefn

const (
	itemError = iota + 1
	itemEnd
)

const eof = -1

type lexer struct {
	src    string
	start  int
	pos    int
	len    int
	lineno int
	state  statefn
	tok    chan token
}

func (l *lexer) next() rune {
}

func (l *lexer) peek() rune {
}

func (l *lexer) pook() rune {
}

func (l *lexer) emit(typ itemType) {
}

func (l *lexer) skipUntil(c rune) rune {
}

type token struct {
	typ itemType
	sym string
}

func (l *lexer) lexText() statefn {
}

func (l *lexer) lexIdentify() statefn {
}

func (l *lexer) lexString() statefn {
}

func (l *lexer) lexDigit() statefn {
}

func (l *lexer) lexType() statefn {
}

func (l *lexer) init(src string) {
	l.src = src
}

func lex(src string) lexer {
	l := make(lexer)
	l.init(src)

	for l != nil {
		l = l.state()
	}

}
func Lex(src string) <-chan token {
	l := lex(str)
	return l.tok
}
