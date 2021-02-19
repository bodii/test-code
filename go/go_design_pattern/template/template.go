package template

import "fmt"

// ISMS sms interface
type ISMS interface {
	send(content string, phone int) error
}

// sms struct
type sms struct {
	ISMS
}

// Valid valid sms
func (s *sms) Valid(content string) error {
	if len(content) > 63 {
		return fmt.Errorf("content is too long")
	}
	return nil
}

// Send send sms
func (s *sms) Send(content string, phone int) error {
	if err := s.Valid(content); err != nil {
		return err
	}

	return s.send(content, phone)
}

// TelecomSms telecom sms
type TelecomSms struct {
	*sms
}

// NewTelecomSms new TelecomSms
func NewTelecomSms() *TelecomSms {
	tel := &TelecomSms{}
	tel.sms = &sms{ISMS: tel}
	return tel
}

func (tel *TelecomSms) send(content string, phone int) error {
	fmt.Println("send by telecom success")
	return nil
}
