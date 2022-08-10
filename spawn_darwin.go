package qjs

import (
	"github.com/rosbit/go-expect"
)

func spawn(quickJSExecPath string, args ...string) (e *expect.Expect, err error) {
	return expect.SpawnPTY(quickJSExecPath, args...)
}
