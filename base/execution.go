package base

import (
	"github.com/ElaraLang/elara/ast"
	"github.com/ElaraLang/elara/interpreter"
	"github.com/ElaraLang/elara/lexer"
	"github.com/ElaraLang/elara/parser"
	"time"
)

func Execute(fileName *string, code chan rune, scriptMode bool) (results ast.Statement, lexTime, parseTime, execTime time.Duration) {
	output := make(chan lexer.Token)
	go lexer.Lex(code, output)

	parserOutput := make(chan ast.Statement)
	parseErrors := make(chan parser.ParseError)
	psr := parser.NewParser(output, parserOutput, parseErrors)
	go psr.Parse()

	//if len(parseErrors) != 0 {
	//	file := "Unknown File"
	//	if fileName != nil {
	//		file = *fileName
	//	}
	//	_, _ = os.Stderr.WriteString(fmt.Sprintf("Syntax Errors found in %s: \n", file))
	//	for err := range parseErrors {
	//		_, _ = os.Stderr.WriteString(fmt.Sprintf("%s\n", err))
	//	}
	//	return
	//}

	interpreterOutput := make(chan *interpreter.Value)
	evaluator := interpreter.NewInterpreter(parserOutput, parseErrors, interpreterOutput)

	evaluator.Exec(scriptMode)
	return
}
