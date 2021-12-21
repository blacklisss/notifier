package sender

import (
	"fmt"
	"notificationService/models"
	"reflect"
	"testing"
)

type mockTransport struct {
}

func (st *mockTransport) Send(notice models.Notice) error {
	fmt.Println(notice.Text)
	return nil
}

type mockTransportWithError struct {
}

func (st *mockTransportWithError) Send(notice models.Notice) error {
	return fmt.Errorf("error")
}

func TestNewSender(t *testing.T) {
	type args struct {
		transports []ISender
	}

	stT := new(mockTransport)
	stTArr := args{[]ISender{stT}}
	sender := &Sender{
		Transports: []ISender{stT},
	}

	tests := []struct {
		name string
		args args
		want *Sender
	}{
		{name: "NewSenderTest", args: stTArr, want: sender},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSender(tt.args.transports); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSender() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSender_Send(t *testing.T) {

	stT := new(mockTransport)
	notice := models.NewNotice("Test", nil)

	type fields struct {
		Transports []ISender
		Err        []error
	}
	type args struct {
		notice models.Notice
	}

	stTArr := fields{Transports: []ISender{stT}}
	argArr := args{notice: *notice}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   []error
	}{
		{name: "TestingSend", fields: stTArr, args: argArr, want: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Sender{
				Transports: tt.fields.Transports,
				Err:        tt.fields.Err,
			}
			if got := s.Send(tt.args.notice); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Send() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSender_SendWithError(t *testing.T) {

	stT := new(mockTransportWithError)
	notice := models.NewNotice("Test", nil)

	type fields struct {
		Transports []ISender
		Err        []error
	}
	type args struct {
		notice models.Notice
	}

	stTArr := fields{Transports: []ISender{stT}}
	argArr := args{notice: *notice}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   []error
	}{
		{name: "TestingSendWithError", fields: stTArr, args: argArr, want: []error{fmt.Errorf("error")}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Sender{
				Transports: tt.fields.Transports,
				Err:        tt.fields.Err,
			}
			if got := s.Send(tt.args.notice); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Send() = %v, want %v", got, tt.want)
			}
		})
	}
}
