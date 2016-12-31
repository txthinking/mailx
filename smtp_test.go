package xmail

import (
	"net/mail"
	"testing"
)

func TestSMTP(t *testing.T) {
	s := &SMTP{
		Server:   "smtp.mailtrap.io",
		Port:     2525,
		UserName: "e3f534cfe656f4",
		Password: "b6e38ddc0f1e9d",
		IsTLS:    false,
	}
	f := &mail.Address{
		Name:    "Xmail",
		Address: "739f35c64d-9422d2@inbox.mailtrap.io",
	}
	ts := []*mail.Address{
		{
			Name:    "Cloud",
			Address: "cloud@txthinking.com",
		},
	}

	m := &Message{
		From:    f,
		To:      ts,
		Subject: "Xmail test smtp",
		Body:    "哈哈",
		Att: []string{
			"/etc/hosts",
		},
	}
	err := s.Send(m)
	if err != nil {
		t.Fatal(err)
	}

	m = &Message{
		From:    f,
		To:      ts,
		Subject: "Xmail test smtp",
		Body:    "哈哈",
		Att: []string{
			"/etc/hosts",
		},
	}
	err = s.Send(m)
	if err != nil {
		t.Fatal(err)
	}
}
