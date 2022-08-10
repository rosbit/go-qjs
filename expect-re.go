package qjs

import (
	"regexp"
)

var (
	ignoreRE   = regexp.MustCompile(`^QuickJS \- Type "\\h" for help[\r\n]`)
	promptRE   = regexp.MustCompile(`^qjs > `)                          // "qjs > "      prompt
	errRE      = regexp.MustCompile(`^[A-Z][a-z]*Error: .+?[\r\n]`)     // "xxxError: "  error message
	errAtRE    = regexp.MustCompile(`^\s+at .+?[\r\n]`)                 // "  at xxxx"   error at xxxx
	blankRE    = regexp.MustCompile(`^[\r\n]`)
	goalRE     = regexp.MustCompile(`^[\x61\x67]\x1b.*?[\r\n]`)  // goal(xxx,xxx)
	resultRE   = regexp.MustCompile(`^.+?[\r\n]`)

	jsonNameRE = regexp.MustCompile(`([\{\,] )([^ ]*?): `) // JSON name without quote
	jsonNameRepl = []byte(`${1}"${2}":`)
	functionId = []byte("function ")
	nan = []byte("NaN")
)

