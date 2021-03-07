package vasker

import (
	"github.com/golang/protobuf/proto"
	cctpb "snowfrost.garden/vasker/cc_grammar"
)

func VarDecl(name string, ident *cctpb.Identifier) *cctpb.Declaration {
	ds := &cctpb.DeclarationSpecifier{
		Value: &cctpb.DeclarationSpecifier_TypeSpecifier{
			&cctpb.TypeSpecifier{
				Value: &cctpb.TypeSpecifier_SimpleTypeSpecifier{
					&cctpb.SimpleTypeSpecifier{
						Value: &cctpb.SimpleTypeSpecifier_DeclaredName{ident},
					},
				},
			},
		},
	}

	d := &cctpb.Declarator{
		Value: &cctpb.Declarator_DeclaredName{
			&cctpb.Identifier{
				Id: proto.String(name),
			},
		},
	}

	sd := &cctpb.SimpleDeclaration{}
	sd.Specifiers = append(sd.Specifiers, ds)
	sd.Declarators = append(sd.Declarators, d)

	return &cctpb.Declaration{
		Value: &cctpb.Declaration_BlockDeclaration{
			&cctpb.BlockDeclaration{
				Value: &cctpb.BlockDeclaration_SimpleDeclaration{sd},
			},
		},
	}
}

func DeclareAuto(name string) *cctpb.Declaration {
	return VarDecl(name, Id("auto"))
}

func StdStringCtor(s string) *cctpb.Expression {
	fc := &cctpb.FunctionCallExpression{
		Name: IdExpr(NsId("std", "string")),
	}
	AddFuncArg(fc, StringLiteralExpr(s))
	return &cctpb.Expression{
		Value: &cctpb.Expression_FunctionCallExpression{fc},
	}
}
