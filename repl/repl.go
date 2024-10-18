/*
Package repl implements the Read-Eval-Print Loop (REPL for the Leopard programming language.

This package provides the functionality to read input, evaluate expressions,
and print results or errors to the specified output
*/
package repl

import (
	"bufio"
	"fmt"
	"io"
	"leopard/evaluator"
	"leopard/lexer"
	"leopard/object"
	"leopard/parser"
)

const PROMPT = ">> "

// Start initializes the REPL, reading from the provided input and writing
// results to the provided output. It continues until EOF is reached.
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()

	for {
		fmt.Fprintf(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		evaluated := evaluator.Eval(program, env)
		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}
	}
}

// printParserErrors outputs the parsing errors into the specified writer.
func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, "Parser errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
