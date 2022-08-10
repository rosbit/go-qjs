package qjs

import (
	"github.com/rosbit/go-expect"
)

func spawn(quickJSExePath string, args...string) (e *expect.Expect, err error) {
	return expect.Spawn(quickJSExePath, args...)
}
