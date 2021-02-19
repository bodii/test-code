package template

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_sms_Send(t *testing.T) {
	tel := NewTelecomSms()
	err := tel.Send("test", 1233434)
	assert.NoError(t, err)
}
