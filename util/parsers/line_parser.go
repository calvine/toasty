package parsers

import (
	"bufio"
	"errors"
	"io"
)

type ParserState int32

var ErrNoParseFuncSet = errors.New("no parseFunc set for line parser")

const Uninitialized ParserState = -999
const EndOfFile ParserState = -99

type LineParseFuncResult struct {
	Line string
	EOF  bool
}

type LineParseFunc func([]byte) LineParseFuncResult

type LineParser interface {
	SetLineParseFunc(LineParseFunc)
	Next() (LineParseFuncResult, error)
}

type lineParser struct {
	scanner   *bufio.Scanner
	State     ParserState
	parseFunc LineParseFunc
}

func NewLineParser(r io.Reader) lineParser {
	lp := lineParser{}
	lp.scanner = bufio.NewScanner(r)
	lp.State = Uninitialized
	return lp
}

func (p *lineParser) SetLineParserFunc(f LineParseFunc) {
	p.parseFunc = f
}

func (p *lineParser) Next() (LineParseFuncResult, error) {
	if p.parseFunc == nil {
		return LineParseFuncResult{}, ErrNoParseFuncSet
	}
	success := p.scanner.Scan()
	if !success {
		err := p.scanner.Err()
		if err == nil {
			// end of file
			return LineParseFuncResult{"", true}, nil
		} else {
			// an error occured
			return LineParseFuncResult{"", false}, err
		}
	}
	data := p.scanner.Bytes()
	return p.parseFunc(data), nil
}
