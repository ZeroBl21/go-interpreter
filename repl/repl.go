package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/ZeroBl21/go-interpreter/compiler"
	"github.com/ZeroBl21/go-interpreter/lexer"
	"github.com/ZeroBl21/go-interpreter/parser"
	"github.com/ZeroBl21/go-interpreter/vm"
)

const (
	RESET  = "\033[0m"
	BLUE   = "\033[34m"
	PROMPT = BLUE + ">> " + RESET
)

const MONKEY_FACE = `            __,__
   .--.  .-"     "-.  .--.
  / .. \/  .-. .-.  \/ .. \
 | |  '|  /   Y   \  |'  | |
 | \   \  \ 0 | 0 /  /   / |
  \ '- ,\.-"""""""-./, -' /
   ''-' /_   ^ ^   _\ '-''
       |  \._   _./  |
       \   \ '~' /   /
        '._ '-=-' _.'
           '-----'
`

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

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

		comp := compiler.New()
		if err := comp.Compile(program); err != nil {
			fmt.Fprintf(out, "Woops! Compilation failed:\n %s\n",
				err)
			continue
		}

		machine := vm.New(comp.Bytecode())
		if err := machine.Run(); err != nil {
			fmt.Fprintf(out, "Woops! Executing bytecode failed:\n%s\n",
				err)
		}

		lastPopped := machine.LastPoppedStackElem()
		io.WriteString(out, lastPopped.Inspect())
		io.WriteString(out, "\n")
	}
}

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, MONKEY_FACE)
	io.WriteString(out, "Woops! We ran into some monkey business here\n")
	io.WriteString(out, " parser errors:\n")

	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
