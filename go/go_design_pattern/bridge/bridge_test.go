package bridge

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrorNotifyication_Notify(t *testing.T) {
	sender := NewEamilMsgSender([]string{"test@test.com"})

	n := NewErrorNotification(sender)
	err := n.Notify("test msg")

	assert.Nil(t, err)
}
