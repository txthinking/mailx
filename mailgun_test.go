package xmail

import (
	"net/mail"
	"testing"
)

func TestMailgun(t *testing.T) {
	const k := ""

	if k == ""{
		return
	}

	s := &Mailgun{
		"sandbox29318.mailgun.org",
		k,
	}

	f := &mail.Address{
		Name:    "Xmail",
		Address: "postmaster@sandbox29318.mailgun.org",
	}
	ff := &mail.Address{
		Name:    "Obama",
		Address: "fakename@fakeaddress.com",
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
		Subject: "Xmail test mailgun",
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
		From:     f,
		FakeFrom: ff,
		To:       ts,
		Subject:  "Xmail test mailgun",
		Body:     "哈哈",
		Att: []string{
			"/etc/hosts",
		},
	}
	err = s.Send(m)
	if err != nil {
		t.Fatal(err)
	}
}
