package ast

import "jojo/token"

type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

// 一系列语句组成程序
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

// var 语句
type VarStatement struct {
	Token token.Token // token.var词法单元
	Name  *Identifier // 变量名
	Value Expression  // 表达式
}

func (vs *VarStatement) statementNode() {}

func (vs *VarStatement) TokenLiteral() string {
	return vs.Token.Literal // var 字符串
}

// const 语句，和var结构一致
type ConstStatement struct {
	Token token.Token // token.const词法单元
	Name  *Identifier // 变量名
	Value Expression  // 表达式
}

func (cs *ConstStatement) statementNode() {}

func (cs *ConstStatement) TokenLiteral() string {
	return cs.Token.Literal // const 字符串
}

type Identifier struct {
	Token token.Token // token.IDENT词法单元
	Value string
}

func (i *Identifier) expressionNode() {}

func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}
