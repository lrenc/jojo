package lexer

import (
	"jojo/token"
)

/**
 * 词法分析器
 */
type Lexer struct {
	/**
	 * 输入
	 */
	input []rune
	/**
	 * 当前字符的位置
	 */
	index int
	/**
	 * 下一个字符的位置
	 */
	nextIndex int
	/**
	 * 当前正在查看的字符
	 */
	char rune
}

/**
 * 是否是字母或_
 */
func isLetter(char rune) bool {
	return ('a' <= char && char <= 'z') || ('A' <= char && char <= 'Z') || char == '_'
}

/**
 * 是否是数字
 */
func isDigit(char rune) bool {
	return '0' <= char && char <= '9'
}

/**
 * 创建token
 */
func newToken(tokenType token.TokenType, char rune) token.Token {
	return token.Token{Type: tokenType, Literal: string(char)}
}

/**
 * 读取下一个字符，同时移动位置
 */
func (l *Lexer) readChar() {
	if l.nextIndex >= len(l.input) {
		l.char = 0 // NUL
	} else {
		l.char = l.input[l.nextIndex]
	}
	l.index = l.nextIndex
	l.nextIndex += 1
}

/**
 * 获取下一个字符
 */
func (l *Lexer) peekChar() rune {
	if l.nextIndex >= len(l.input) {
		return 0
	} else {
		return l.input[l.nextIndex]
	}
}

/**
 * 核心函数
 * 获取下一个token
 */
func (l *Lexer) NextToken() token.Token {

	var t token.Token
	// 跳过空白部分
	l.skipWhitespace()

	switch l.char {
	// 等号
	case '=':
		if l.peekChar() == '=' { // ==
			ch := l.char
			l.readChar()
			literal := string(ch) + string(l.char)
			t = token.Token{Type: token.EQ, Literal: literal}
		} else {
			t = newToken(token.ASSIGN, l.char)
		}

	// 加号
	case '+':
		if l.peekChar() == '=' { // +=
			ch := l.char
			l.readChar()
			literal := string(ch) + string(l.char)
			t = token.Token{Type: token.PE, Literal: literal}
		} else {
			t = newToken(token.PLUS, l.char)
		}

	// 减号
	case '-':
		if l.peekChar() == '=' { // -=
			ch := l.char
			l.readChar()
			literal := string(ch) + string(l.char)
			t = token.Token{Type: token.ME, Literal: literal}
		} else {
			t = newToken(token.MINUS, l.char)
		}

	// 乘号
	case '*':
		if l.peekChar() == '=' { // *=
			ch := l.char
			l.readChar()
			literal := string(ch) + string(l.char)
			t = token.Token{Type: token.AE, Literal: literal}
		} else {
			t = newToken(token.ASTERISK, l.char)
		}

	// 除号
	case '/':
		if l.peekChar() == '=' { // /=
			ch := l.char
			l.readChar()
			literal := string(ch) + string(l.char)
			t = token.Token{Type: token.SE, Literal: literal}
		} else {
			t = newToken(token.SLASH, l.char)
		}

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

	// 单引号
	case '\'':
		t.Type = token.STRING
		t.Literal = l.readString()
		return t

	// 双引号
	case '"':
		t.Type = token.STRING
		t.Literal = l.readString()
		return t

	case 0:
		t.Literal = ""
		t.Type = token.EOF

	default:
		if isLetter(l.char) {
			t.Literal = l.readIdentifier()
			t.Type = token.LookupIdentifier(t.Literal)
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

/**
 * 读取标识符，不区分是否为关键字
 */
func (l *Lexer) readIdentifier() string {
	index := l.index
	for isLetter(l.char) {
		l.readChar()
	}
	return string(l.input[index:l.index])
}

/**
 * 读取number
 */
func (l *Lexer) readNumber() string {
	index := l.index
	for isDigit(l.char) {
		l.readChar()
	}

	// 小数
	if l.char == '.' {
		l.readChar()
		for isDigit(l.char) {
			l.readChar()
		}
	}
	return string(l.input[index:l.index])
}

/**
 * 读取字符串
 * 当前只识别单双引号
 */
func (l *Lexer) readString() string {
	index := l.index
	char := l.char
	if char == '"' || char == '\'' {
		l.readChar()
		for l.char != char {
			l.readChar()
		}
		// 跳过后面的引号
		l.readChar()
	}
	return string(l.input[index:l.index])
}

/**
 * 忽略空白
 */
func (l *Lexer) skipWhitespace() {
	for l.char == ' ' || l.char == '\t' || l.char == '\n' || l.char == '\r' {
		l.readChar()
	}
}

func New(input string) *Lexer {
	// 将字符串转换为 rune 数组
	l := &Lexer{input: []rune(input)}
	l.readChar()
	return l
}
