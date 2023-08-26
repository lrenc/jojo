package token

type TokenType string

type Token struct {
	/**
	 * token 类型
	 */
	Type TokenType
	/**
	 * token 字面量
	 */
	Literal string
}

/**
 * 关键字
 */
const (
	/**
	 * var 定义变量
	 */
	VAR = "VAR"

	/**
	 * const 定义常量
	 */
	CONST = "CONST"

	/**
	 * fn 定义函数
	 */
	FUNCTION = "FUNCTION"

	/**
	 * 条件语句
	 */
	IF   = "IF"
	ELSE = "ELSE"

	/**
	 * 循环语句
	 */
	FOR      = "FOR"
	IN       = "IN"
	BREAK    = "BREAK"
	CONTINUE = "CONTINUE"

	/**
	 * switch 语句
	 */
	SWITCH  = "SWITCH"
	CASE    = "CASE"
	DEFAULT = "DEFAULT"

	/**
	 * return
	 */
	RETURN = "RETURN"

	/**
	 * 模块化
	 */
	IMPORT = "IMPORT"
	EXPORT = "EXPORT"

	/**
	 * 并发
	 */
	ASYNC = "ASYNC"

	/**
	 * 面向对象
	 */
	INTERFACE = "INTERFACE"
	STRUCT    = "STRUCT"
)

/**
 * 类型
 */
const (
	/**
	 * number
	 */
	NUMBER = "NUMBER"

	/**
	 * string
	 */
	STRING = "STRING"

	/**
	 * boolean
	 */
	BOOLEAN = "BOOLEAN"

	/**
	 * null
	 */
	NULL = "NULL"
)

const (
	/**
	 * true
	 */
	TRUE = "TRUE"

	/**
	 * false
	 */
	FALSE = "FALSE"
)

/**
 * 运算符
 */
const (
	/**
	 * 四则运算
	 */
	PLUS     = "+"
	MINUS    = "-"
	ASTERISK = "*"
	SLASH    = "/"

	PE = "+="
	ME = "-="
	AE = "*="
	SE = "/="

	/**
	 * 比较运算符
	 */
	LT  = "<"
	GT  = ">"
	LET = "<="
	GET = ">="
	EQ  = "=="
	NEQ = "!="

	/**
	 *
	 */
	ASSIGN = "="
	BANG   = "!"
)

/**
 * 括号
 */
const (
	LPAREN = "("
	RPAREN = ")"

	LSAREN = "["
	RSAREN = "]"

	LBRACE = "{"
	RBRACE = "}"
)

/**
 * 标点符号
 */
const (
	DOT       = "."
	COLON     = ":"
	COMMA     = ","
	SEMICOLON = ";"
)

/**
 * 其他
 */
const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"
)

/**
 * 标识符
 */
const IDENTIFIER = "IDENTIFIER"

// 保留字
var keepwords = map[string]TokenType{
	"var":       VAR,
	"const":     CONST,
	"fn":        FUNCTION,
	"if":        IF,
	"else":      ELSE,
	"for":       FOR,
	"in":        IN,
	"break":     BREAK,
	"continue":  CONTINUE,
	"switch":    SWITCH,
	"case":      CASE,
	"default":   DEFAULT,
	"return":    RETURN,
	"import":    IMPORT,
	"export":    EXPORT,
	"async":     ASYNC,
	"struct":    STRUCT,
	"interface": INTERFACE,
	"number":    NUMBER,
	"string":    STRING,
	"boolean":   BOOLEAN,
	"null":      NULL,
	"true":      TRUE,
	"false":     FALSE,
}

func LookupIdentifier(identifier string) TokenType {
	if keepword, ok := keepwords[identifier]; ok {
		return keepword
	}
	return IDENTIFIER
}
