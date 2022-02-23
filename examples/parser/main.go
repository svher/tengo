package main

import (
	"fmt"
	"github.com/d5/tengo/v2/parser"
	"github.com/d5/tengo/v2/token"
	"io/ioutil"
	"strings"
)

func main() {
	fileSet := parser.NewFileSet()
	src, err := ioutil.ReadFile("examples/testcases/main.tengo")
	if err != nil {
		panic(err)
	}
	sourceFile := fileSet.AddFile("(main)", -1, len(src))
	var errors parser.ErrorList
	s := parser.NewScanner(sourceFile, src, func(pos parser.SourceFilePos, msg string) {
		errors.Add(pos, msg)
	}, 0)

	var tok token.Token
	var lit string
	var pos parser.Pos

	for tok != token.EOF {
		tok, lit, pos = s.Scan()
		fmt.Println(int(tok), tok, strings.Replace(lit, "\n", "\\n", -1), pos)
	}
}
