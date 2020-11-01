package ident

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPanicIdentification(t *testing.T) {
	defer func() {
		_ = recover()
		v := Identify()
		splitted := strings.Split(v, "/")
		filename := splitted[len(splitted)-1]
		assert.Equal(t, "ident_test.go:21", filename)

	}()

	var arr []string
	_ = arr[1]

}
