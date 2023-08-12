package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL" // 非法的
	EOF     = "EOF"

	// 运算符
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	SLASH    = "/"
	ASTERISK = "*"
	BANG     = "!"

	LT  = "<"
	GT  = ">"
	LET = "<="
	GET = ">="
	EQ  = "=="
	NEQ = "!="

	// 分隔符
	COMMA     = ","
	SEMICOLON = ";"
	COLON     = ":"

	// 括号
	LPAREN = "("
	RPAREN = ")"

	LSAREN = "["
	RSAREN = "]"

	LBRACE = "{"
	RBRACE = "}"

	// 关键字
	// 标识符
	IDENT = "IDENT"
	// 字面量
	// INT = "INT"
	VAR      = "VAR"
	CONST    = "CONST"
	FUNCTION = "FUNCTION"

	TRUE   = "TRUE"
	FALSE  = "FALSE"
	IF     = "IF"
	ELSE   = "ELSE"
	RETURN = "RETURN"

	// 类型
	NUMBER  = "NUMBER"
	STRING  = "STRING"
	BOOLEAN = "BOOLEAN"
	NULL    = "NULL"

	IMPORT = "IMPORT"
	EXPORT = "EXPORT"
)

// 关键字
var keywords = map[string]TokenType{

	"var":   VAR,   // 变量
	"const": CONST, // 常量
	"fn":    FUNCTION,

	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,

	"number":  NUMBER,
	"string":  STRING,
	"boolean": BOOLEAN,
	"null":    NULL,

	"import": IMPORT,
	"export": EXPORT,
}

func LookupIdent(ident string) TokenType {
	if t, ok := keywords[ident]; ok {
		return t
	}
	// 不是关键字就是标识符
	// TODO 需要考虑字符串等情况
	return IDENT
}
