package state

import "fmt"

// IState state interface
type IState interface {
	Approval(m *Machine)
	Reject(m *Machine)
	GetName() string
}

// Machine state machine
type Machine struct {
	state IState
}

// Approval  state machine approval function
func (m *Machine) Approval() {
	m.state.Approval(m)
}

// Reject state machine reject function
func (m *Machine) Reject() {
	m.state.Reject(m)
}

// GetName state machine GetName function
func (m *Machine) GetName() string {
	return m.state.GetName()
}

// SetState update state
func (m *Machine) SetState(state IState) {
	m.state = state
}

// leaderApproveState struct
type leaderApproveState struct{}

func (leaderApproveState) Approval(m *Machine) {
	fmt.Println("leader approval success")
	m.SetState(GetFinanceApproveState())
}

// GetName leaderApproveState GetName function
func (leaderApproveState) GetName() string {
	return "leaderApproveState"
}

// Reject leaderApproveState reject function
func (leaderApproveState) Reject(m *Machine) {}

// GetLeaderApproveState get leaderApproveState
func GetLeaderApproveState() IState {
	return &leaderApproveState{}
}

// financeApproveState struct
type financeApproveState struct{}

func (financeApproveState) Approval(m *Machine) {
	fmt.Println("finance approve success")
	fmt.Println("payment operation")
}

// GetName financeApproveState GetName function
func (financeApproveState) GetName() string {
	return "financeApproveState"
}

// Reject financeApproveState reject function
func (financeApproveState) Reject(m *Machine) {
	m.SetState(GetLeaderApproveState())
}

// GetFinanceApproveState get financeApproveState
func GetFinanceApproveState() IState {
	return &financeApproveState{}
}
