package bridge

// IMsgSender message sender interface
type IMsgSender interface {
	Send(msg string) error
}

// EmailMsgSender eamil message sender struct
type EmailMsgSender struct {
	emails []string
}

// NewEamilMsgSender create an new EamilMsgSender
func NewEamilMsgSender(emails []string) *EmailMsgSender {
	return &EmailMsgSender{emails: emails}
}

// Send send function
func (s *EmailMsgSender) Send(msg string) error {
	// send message
	return nil
}

// INotification notification interface
type INotification interface {
	Notify(msg string) error
}

// ErrorNotification error notificaton
type ErrorNotification struct {
	sender IMsgSender
}

// NewErrorNotification create an error notification
func NewErrorNotification(sender IMsgSender) *ErrorNotification {
	return &ErrorNotification{sender: sender}
}

// Notify notify error function
func (n *ErrorNotification) Notify(msg string) error {
	return n.sender.Send(msg)
}
