package magic

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDetect(t *testing.T) {
	inst, err := New(MAGIC_MIME, "./magic.mgc")
	assert.NoError(t, err)
	ret, err := inst.File("./2401.01663.pdf")
	assert.NoError(t, err)
	assert.Equal(t, "application/pdf", ret)

	data, err := os.ReadFile("./2401.01663.pdf")
	assert.NoError(t, err)
	ret, err = inst.Buffer(data)
	assert.NoError(t, err)
	assert.Equal(t, "application/pdf", ret)
}
