package mailx

import (
	"net/mail"
	"testing"
)

func TestSMTP(t *testing.T) {
	s := &SMTP{
		Server:   "smtp.mailtrap.io",
		Port:     465,
		UserName: "e3f534cfe656f4",
		Password: "b6e38ddc0f1e9d",
	}

	m := &Message{
		From: &mail.Address{
			Name:    "mailx",
			Address: "739f35c64d-9422d2@inbox.mailtrap.io",
		},
		To: []*mail.Address{
			{
				Name:    "Cloud",
				Address: "cloud@txthinking.com",
			},
		},
		Subject: "Test",
		Body:    "哈哈",
		Attachment: []string{
			"/etc/hosts",
		},
	}
	err := s.Send(m)
	if err != nil {
		t.Fatal(err)
	}
}
