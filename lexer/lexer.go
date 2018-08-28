package lexer

type Lexer struct {
	input        string
	position     int  //current position in input
	readPosition int  //current position being read (after current character)
	ch           byte //current character under examination
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0 //ASCII for NUL
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}
