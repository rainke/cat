package repl

import (
	"bufio"
	"cat/evaluator"
	"cat/lexer"
	"cat/parser"
	"fmt"
	"io"
)

const PROMPT = ">> "

const CatLogo = `
  _______ _     _ _     _ _     _
 |__   __| |   (_) |   | | |   | |
    | |  | |__  _| | __| | | __| |
    | |  | '_ \| | |/ _' | |/ _' |
    | |  | | | | | | (_| | | (_| |
    |_|  |_| |_|_|_|\__,_|_|\__,_|
`

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprint(out, PROMPT)
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

		evaluated := evaluator.Eval(program)

		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}
	}
}

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, CatLogo)
	io.WriteString(out, "\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
