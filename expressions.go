package vasker

import (
	cctpb "snowfrost.garden/vasker/cc_grammar"
	"github.com/golang/protobuf/proto"
)

func StringLiteralExpr(s string) *cctpb.Expression {
	return &cctpb.Expression{
		Value: &cctpb.Expression_LiteralExpression{
			&cctpb.Literal{
				Value: &cctpb.Literal_StringLiteral{s},
			},
		},
	}
}

// IntLiteralExpr wraps the provided int in an integral literal expression.
func IntLiteralExpr(i int32) *cctpb.Expression {
	return &cctpb.Expression{
		Value: &cctpb.Expression_LiteralExpression{
			&cctpb.Literal{
				Value: &cctpb.Literal_IntegerLiteral{int64(i)},
			},
		},
	}
}

func IdExpr(id *cctpb.Identifier) *cctpb.Expression {
	return &cctpb.Expression{
		Value: &cctpb.Expression_IdentifierExpression{id},
	}
}

// StringIdExpr turns id into an Identifier, and wraps it in an identifier
// expression.
func StringIdExpr(id string) *cctpb.Expression {
	return IdExpr(Id(id))
}

func Id(id string) *cctpb.Identifier {
	return &cctpb.Identifier{
		Id: proto.String(id),
	}
}

func NsId(ns string, id string) *cctpb.Identifier{
	return &cctpb.Identifier{
		Namespace: proto.String(ns),
		Id: proto.String(id),
	}
}

// PtrIndirect wraps the provided expression in a unary pointer-indirection
// expression.
func PtrIndirect(expr *cctpb.Expression) *cctpb.Expression {
	return &cctpb.Expression{
						Value: &cctpb.Expression_UnaryExpression{
							&cctpb.UnaryExpression{
								Operator: cctpb.UnaryExpression_POINTER_INDIRECTION.Enum(),
								Operand: expr,
							},
						},
					}

}

// ObjMember returns an expression of the syntax obj.member.
func ObjMember(obj *cctpb.Expression, member *cctpb.Expression) *cctpb.Expression {
	return &cctpb.Expression{
		Value: &cctpb.Expression_MemberAccessExpression{
			&cctpb.MemberAccessExpression{
				Operator: cctpb.MemberAccessExpression_MEMBER_OF_OBJECT.Enum(),
				Lhs: obj,
				Rhs: member,
			},
		},
	}
}

// PtrMember returns an expression of the syntax obj->member.
func PtrMember(obj *cctpb.Expression, member *cctpb.Expression) *cctpb.Expression {
	return &cctpb.Expression{
		Value: &cctpb.Expression_MemberAccessExpression{
			&cctpb.MemberAccessExpression{
				Operator: cctpb.MemberAccessExpression_MEMBER_OF_POINTER.Enum(),
				Lhs: obj,
				Rhs: member,
			},
		},
	}
}

// FuncCall turns the passed identifier into a function call, e.g. id().
func FuncCall(id *cctpb.Identifier) *cctpb.Expression {
	return &cctpb.Expression{
		Value: &cctpb.Expression_FunctionCallExpression{
			&cctpb.FunctionCallExpression{
				Name: IdExpr(id),
			},
		},
	}
}

func wrapFuncCallExpr(e *cctpb.Expression) *cctpb.FunctionCallExpression_ExpressionArg {
	return &cctpb.FunctionCallExpression_ExpressionArg{
		Value: &cctpb.FunctionCallExpression_ExpressionArg_Expression{e},
	}
}

// AddFuncArg adds the expression e to the arguments of the function call fce.
func AddFuncArg(fce *cctpb.FunctionCallExpression, e *cctpb.Expression) {
	fce.Arguments = append(fce.Arguments, wrapFuncCallExpr(e))
}
