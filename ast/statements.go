package ast

import "github.com/ElaraLang/elara/lexer"

type ImportStatement struct {
	Token  lexer.Token
	Module Module
}

type NamespaceStatement struct {
	Token  lexer.Token
	Module Module
}

type ExpressionStatement struct {
	Token      lexer.Token
	Expression Expression
}

type DeclarationStatement struct {
	Token      lexer.Token
	Mutable    bool
	Lazy       bool
	Open       bool
	Identifier IdentifierLiteral
	Type       Type
	Value      Expression
}

type StructDefStatement struct {
	Token  lexer.Token
	Id     IdentifierLiteral
	Fields []StructField
}

type WhileStatement struct {
	Token     lexer.Token
	Condition Expression
	Body      Statement
}

type ExtendStatement struct {
	Token      lexer.Token
	Identifier IdentifierLiteral
	Alias      IdentifierLiteral
	Body       *BlockStatement
}

type BlockStatement struct {
	Token lexer.Token
	Block []Statement
}

type TypeStatement struct {
	Token      lexer.Token
	Identifier IdentifierLiteral
	InternalId IdentifierLiteral
	Contract   Type
}

type GenerifiedStatement struct {
	Token     lexer.Token
	Contract  NamedContract
	Statement Statement
}

type ReturnStatement struct {
	Token lexer.Token
	Value Expression
}
