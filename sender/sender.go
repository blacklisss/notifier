package sender

import "notificationService/models"

type ISender interface {
	Send(notice models.Notice) error
}

type Sender struct {
	Transports []ISender
	Err        []error
}

func (s *Sender) Send(notice models.Notice) []error {
	for _, t := range s.Transports {
		if err := t.Send(notice); err != nil {
			s.Err = append(s.Err, err)
		}
	}

	if s.Err != nil {
		return s.Err
	}

	return nil
}

func NewSender(transports []ISender) *Sender {
	sender := &Sender{
		Transports: transports,
	}

	return sender
}
