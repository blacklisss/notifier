package models

type Noticer interface {
	SetText(t string)
	SetRecipientID(r interface{})
}

type Notice struct {
	Text        string
	RecipientID interface{}
}

func NewNotice(text string, rID interface{}) *Notice {
	return &Notice{
		text,
		rID,
	}
} // создание нового извещение

func (n *Notice) SetText(t string) {
	n.Text = t
} // метод для передачи текста сообщения для извещения

func (n *Notice) SetRecipientID(r interface{}) {
	n.RecipientID = r
} // метод для передачи идентификатора адресата для извещения
