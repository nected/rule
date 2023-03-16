package ast

import (
	"fmt"
	"go/token"
	"strconv"
	"time"
)

type Node interface {
	node()
	String() string
}

type Expr interface {
	Node
	expr()
	Args() []string
}

// GroupExpr represents a groupExpr
type GroupExpr struct {
	Id       string
	Children []Expr
}

func (*GroupExpr) expr() {}
func (*GroupExpr) node() {}
func (*GroupExpr) Args() []string {
	args := []string{}
	return args
}
func (r *GroupExpr) String() string {
	str := ""
	for _, child := range r.Children {
		if child != nil {
			str += child.String()
		}
	}
	return str
}

// BinaryExpr represents an operation between two expressions.
type BinaryExpr struct {
	Id  string
	Op  token.Token
	LHS Expr
	RHS Expr
}

func (*BinaryExpr) expr() {}
func (*BinaryExpr) node() {}
func (*BinaryExpr) Args() []string {
	args := []string{}
	return args
}
func (e *BinaryExpr) String() string {
	if e.LHS != nil && e.RHS != nil {
		return fmt.Sprintf("%s %s %s", e.LHS.String(), e.Op, e.RHS.String())
	}
	return ""
}

// UniaryExpr represents an operation one expressions.
type UniaryExpr struct {
	Id  string
	Op  token.Token
	LHS Expr
}

func (*UniaryExpr) expr() {}
func (*UniaryExpr) node() {}
func (*UniaryExpr) Args() []string {
	args := []string{}
	return args
}
func (e *UniaryExpr) String() string {
	if e.LHS != nil {
		return fmt.Sprintf("%s %s", e.Op, e.LHS.String())
	}
	return ""
}

// TerniaryExpr represents an operation between three expressions.
type TerniaryExpr struct {
	Id   string
	Op   token.Token
	LHS  Expr
	RHS  Expr
	RHS2 Expr
}

func (*TerniaryExpr) expr() {}
func (*TerniaryExpr) node() {}
func (*TerniaryExpr) Args() []string {
	args := []string{}
	return args
}
func (e *TerniaryExpr) String() string {
	if e.LHS != nil && e.RHS != nil && e.RHS2 != nil {
		return fmt.Sprintf("%s %s %s %s", e.LHS.String(), e.Op, e.RHS.String(), e.RHS2.String())
	}
	return ""
}

// VarRef represents a reference to a variable.
type VarRef struct {
	Id  string
	Val string
}

func (*VarRef) expr() {}
func (*VarRef) node() {}
func (*VarRef) Args() []string {
	args := []string{}
	return args
}
func (r *VarRef) String() string { return QuoteIdent(r.Val) }

// NumberLiteral represents a numeric literal.
type NumberLiteral struct {
	Id  string
	Val float64
}

func (*NumberLiteral) expr() {}
func (*NumberLiteral) node() {}
func (*NumberLiteral) Args() []string {
	args := []string{}
	return args
}
func (l *NumberLiteral) String() string { return strconv.FormatFloat(l.Val, 'f', 3, 64) }

// BooleanLiteral represents a boolean literal.
type BooleanLiteral struct {
	Id  string
	Val bool
}

func (*BooleanLiteral) expr() {}
func (*BooleanLiteral) node() {}
func (*BooleanLiteral) Args() []string {
	args := []string{}
	return args
}
func (l *BooleanLiteral) String() string {
	if l.Val {
		return "true"
	}
	return "false"
}

// StringLiteral represents a string literal.
type StringLiteral struct {
	Id  string
	Val string
}

func (*StringLiteral) expr() {}
func (*StringLiteral) node() {}
func (*StringLiteral) Args() []string {
	args := []string{}
	return args
}
func (l *StringLiteral) String() string { return Quote(l.Val) }

// TimeLiteral represents a point-in-time literal.
type TimeLiteral struct {
	Id  string
	Val time.Time
}

func (*TimeLiteral) expr() {}
func (*TimeLiteral) node() {}
func (*TimeLiteral) Args() []string {
	args := []string{}
	return args
}
func (l *TimeLiteral) String() string { return l.Val.UTC().Format("2006-01-02 15:04:05.999") }
