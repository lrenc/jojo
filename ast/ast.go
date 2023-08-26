package ast

import "jojo/token"

/**
 * 基础节点
 */
type Node interface {
	TokenLiteral() string
}

/**
 * 语句
 */
type Statement interface {
	Node
	statementNode()
}

/**
 * 表达式
 */
type Expression interface {
	Node
	expressionNode()
}

/**
 * var 语句
 */
type VarStatement struct {
	/**
	 * token.var词法单元
	 */
	Token token.Token
	/**
	 * 变量名
	 */
	Name *Identifier
	/**
	 * 类型
	 */
	Type *Type
	/**
	 * 值（表达式）
	 */
	Value Expression
}

func (vs *VarStatement) statementNode() {}

func (vs *VarStatement) TokenLiteral() string {
	return vs.Token.Literal // var 字符串
}

/**
 * const 语句，和var结构一致
 */
type ConstStatement struct {
	/**
	 * token.const词法单元
	 */
	Token token.Token
	/**
	 * 变量名
	 */
	Name *Identifier
	/**
	 * 类型
	 */
	Type *Type
	/**
	 * 表达式
	 */
	Value Expression
}

func (cs *ConstStatement) statementNode() {}

func (cs *ConstStatement) TokenLiteral() string {
	return cs.Token.Literal // const 字符串
}

/**
 * 标识符
 */
type Identifier struct {
	Token token.Token // token.IDENTIFIER 词法单元
	Value string
}

func (i *Identifier) expressionNode() {}

func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

/**
 * 类型
 */
type Type struct {
	Token token.Token
	Value string
}

func (t *Type) TokenLiteral() string {
	return t.Token.Literal
}

/**
 * 程序由一系列语句组成
 */
type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}
