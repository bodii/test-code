package memento

// InputText input text struct
type InputText struct {
	content string
}

// Append append content
func (in *InputText) Append(content string) {
	in.content += content
}

// GetText get content function
func (in *InputText) GetText() string {
	return in.content
}

// Snapshot snapshot struct
type Snapshot struct {
	content string
}

// Snapshot snapshot function
func (in *InputText) Snapshot() *Snapshot {
	return &Snapshot{content: in.content}
}

// Resore Resore snapshot to InputText
func (in *InputText) Resore(s *Snapshot) {
	in.content = s.GetText()
}

// GetText get snapshot content
func (s *Snapshot) GetText() string {
	return s.content
}
