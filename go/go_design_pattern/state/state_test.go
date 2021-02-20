package state

import (
	"testing"

	assert "github.com/stretchr/testify/assert"
)

func TestMachine_GetStateName(t *testing.T) {
	m := &Machine{state: GetLeaderApproveState()}
	assert.Equal(t, "leaderApproveState", m.GetName())
	m.Approval()
	assert.Equal(t, "financeApproveState", m.GetName())
	m.Reject()
	assert.Equal(t, "leaderApproveState", m.GetName())
	m.Approval()
	assert.Equal(t, "financeApproveState", m.GetName())
	m.Approval()
}
