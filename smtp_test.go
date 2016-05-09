// xmail (https://github.com/txthinking/xmail). Under the MIT license.

package xmail

import (
	"net/mail"
	"testing"
)

func TestSMTP(t *testing.T) {
	s := &SMTP{
		Server:   "mailtrap.io",
		Port:     2525,
		UserName: "e3f534cfe656f4",
		Password: "b6e38ddc0f1e9d",
		IsTLS:    false,
	}
	f := &mail.Address{
		Name:    "Xmail",
		Address: "739f35c64d-48cf45@inbox.mailtrap.io",
	}
	ff := &mail.Address{
		Name:    "Obama",
		Address: "fakename@fakeaddress.com",
	}
	ts := []*mail.Address{
		&mail.Address{
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
		From:     f,
		FakeFrom: ff,
		To:       ts,
		Subject:  "Xmail test smtp",
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
