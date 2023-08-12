package lexer

import (
	"jojo/token"
)

type Lexer struct {
	input     []rune
	index     int  // 当前位置
	nextIndex int  // 下一个位置
	char      rune // 当前正在查看的字符
}

func New(input string) *Lexer {
	// 将字符串转换为 rune 数组
	l := &Lexer{input: []rune(input)}
	l.readChar()
	return l
}

// 读取下一个字符
func (l *Lexer) readChar() {
	if l.nextIndex >= len(l.input) {
		l.char = 0 // NUL
	} else {
		l.char = l.input[l.nextIndex]
	}
	l.index = l.nextIndex
	l.nextIndex += 1
}

// 获取下一个字符
func (l *Lexer) peekChar() rune {
	if l.nextIndex >= len(l.input) {
		return 0
	} else {
		return l.input[l.nextIndex]
	}
}

func newToken(tokenType token.TokenType, char rune) token.Token {
	return token.Token{Type: tokenType, Literal: string(char)}
}

// 获取下一个token
func (l *Lexer) NextToken() token.Token {

	var t token.Token

	l.skipWhitespace()
	switch l.char {
	// 等号
	case '=':
		if l.peekChar() == '=' {
			ch := l.char
			l.readChar()
			literal := string(ch) + string(l.char)
			t = token.Token{Type: token.EQ, Literal: literal}
		} else {
			t = newToken(token.ASSIGN, l.char)
		}

	// 运算符
	case '+':
		t = newToken(token.PLUS, l.char)
	case '-':
		t = newToken(token.MINUS, l.char)
	case '*':
		t = newToken(token.ASTERISK, l.char)
	case '/':
		t = newToken(token.SLASH, l.char)

	// 比较运算
	case '!':
		if l.peekChar() == '=' {
			ch := l.char
			l.readChar()
			literal := string(ch) + string(l.char)
			t = token.Token{Type: token.NEQ, Literal: literal}
		} else {
			t = newToken(token.BANG, l.char)
		}
	case '<':
		if l.peekChar() == '=' {
			ch := l.char
			l.readChar()
			literal := string(ch) + string(l.char)
			t = token.Token{Type: token.LET, Literal: literal}
		} else {
			t = newToken(token.LT, l.char)
		}
	case '>':
		if l.peekChar() == '=' {
			ch := l.char
			l.readChar()
			literal := string(ch) + string(l.char)
			t = token.Token{Type: token.GET, Literal: literal}
		} else {
			t = newToken(token.GT, l.char)
		}

	case ',':
		t = newToken(token.COMMA, l.char)
	case ';':
		t = newToken(token.SEMICOLON, l.char)
	case ':':
		t = newToken(token.COLON, l.char)

	case '(':
		t = newToken(token.LPAREN, l.char)
	case ')':
		t = newToken(token.RPAREN, l.char)

	case '[':
		t = newToken(token.LSAREN, l.char)
	case ']':
		t = newToken(token.RSAREN, l.char)

	case '{':
		t = newToken(token.LBRACE, l.char)
	case '}':
		t = newToken(token.RBRACE, l.char)
	case 0:
		t.Literal = ""
		t.Type = token.EOF

	default:
		if isLetter(l.char) {
			t.Literal = l.readIdentifier()
			t.Type = token.LookupIdent(t.Literal)
			return t
		} else if isDigit(l.char) {
			t.Type = token.NUMBER
			t.Literal = l.readNumber()
			return t
		} else {
			t = newToken(token.ILLEGAL, l.char)
		}
	}

	l.readChar()
	return t
}

// 读取完整的标识符
func (l *Lexer) readIdentifier() string {
	index := l.index
	for isLetter(l.char) {
		l.readChar()
	}
	return string(l.input[index:l.index])
}

// 是否是字母
func isLetter(char rune) bool {
	return 'a' <= char && char <= 'z' || 'A' <= char && char <= 'Z' || char == '_'
}

// 空白直接跳过
func (l *Lexer) skipWhitespace() {
	for l.char == ' ' || l.char == '\t' || l.char == '\n' || l.char == '\r' {
		l.readChar()
	}
}

func (l *Lexer) readNumber() string {
	index := l.index
	for isDigit(l.char) {
		l.readChar()
	}
	if l.char == '.' {
		l.readChar()
		for isDigit(l.char) {
			l.readChar()
		}
	}
	return string(l.input[index:l.index])
}

func isDigit(char rune) bool {
	return '0' <= char && char <= '9'
}
