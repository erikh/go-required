package required

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Testing struct {
	A string `required:""`
	B string
	I int `required:""`
}

func TestAll(t *testing.T) {
	test := Testing{A: "hello", I: 10, B: "hello"}
	assert.Nil(t, Required(test))
	test.A = ""
	assert.NotNil(t, Required(test))
	test.A = "hello"
	test.I = 0
	assert.NotNil(t, Required(test))
}
