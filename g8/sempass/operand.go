package sempass

import (
	"math"
	"strconv"

	"e8vm.io/e8vm/g8/ast"
	"e8vm.io/e8vm/g8/parse"
	"e8vm.io/e8vm/g8/tast"
	"e8vm.io/e8vm/g8/types"
	"e8vm.io/e8vm/lex8"
)

func buildInt(b *Builder, op *lex8.Token) tast.Expr {
	ret, e := strconv.ParseInt(op.Lit, 0, 64)
	if e != nil {
		b.Errorf(op.Pos, "invalid integer: %s", e)
		return nil
	}

	if ret < math.MinInt32 {
		b.Errorf(op.Pos, "integer too small to fit in 32-bit")
		return nil
	} else if ret > math.MaxUint32 {
		b.Errorf(op.Pos, "integer too large to fit in 32-bit")
		return nil
	}

	ref := tast.NewConstRef(types.NewNumber(ret), ret)
	return &tast.Const{ref}
}

func buildChar(b *Builder, op *lex8.Token) tast.Expr {
	v, e := strconv.Unquote(op.Lit)
	if e != nil {
		b.Errorf(op.Pos, "invalid char: %s", e)
		return nil
	} else if len(v) != 1 {
		b.Errorf(op.Pos, "invalid char in quote: %q", v)
		return nil
	}
	ref := tast.NewConstRef(types.Int8, int64(v[0]))
	return &tast.Const{ref}
}

func buildString(b *Builder, op *lex8.Token) tast.Expr {
	v, e := strconv.Unquote(op.Lit)
	if e != nil {
		b.Errorf(op.Pos, "invalid string: %s", e)
		return nil
	}
	ref := tast.NewConstRef(types.String, v)
	return &tast.Const{ref}
}

func buildIdent(b *Builder, ident *lex8.Token) tast.Expr {
	s := b.scope.Query(ident.Lit)
	if s == nil {
		b.Errorf(ident.Pos, "undefined identifier %s", ident.Lit)
		return nil
	}

	panic("todo")
}

func buildOperand(b *Builder, op *ast.Operand) tast.Expr {
	if op.Token.Type == parse.Keyword && op.Token.Lit == "this" {
		if b.this == nil {
			b.Errorf(op.Token.Pos, "using this out of a method function")
			return nil
		}
		return &tast.This{b.this}
	}

	switch op.Token.Type {
	case parse.Int:
		return buildInt(b, op.Token)
	case parse.Char:
		return buildChar(b, op.Token)
	case parse.String:
		return buildString(b, op.Token)
	case parse.Ident:
		return buildIdent(b, op.Token)
	default:
		b.Errorf(op.Token.Pos, "invalid or not implemented: %d",
			op.Token.Type,
		)
		return nil
	}
}