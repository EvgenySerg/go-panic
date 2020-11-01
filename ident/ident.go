package ident

import (
	"fmt"
	"runtime"
	"strings"
)

// Identify returns name of package, file name and line number of place where panic occurs
func Identify() string {
	var name, file string
	var line int
	var pc [16]uintptr

	n := runtime.Callers(3, pc[:])
	for _, pc := range pc[:n] {
		fn := runtime.FuncForPC(pc)
		if fn == nil {
			continue
		}
		file, line = fn.FileLine(pc)
		name = fn.Name()
		if !strings.HasPrefix(name, "runtime.") {
			break
		}
	}

	return fmt.Sprintf("name: %v, file: %v:%v", name, file, line)
}
