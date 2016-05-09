// xmail (https://github.com/txthinking/xmail). Under the MIT license.

package xmail

import (
	"net/mail"
	"testing"
)

func TestMessage(t *testing.T) {
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
		From:     f,
		FakeFrom: ff,
		To:       ts,
		Subject:  "Xmail test smtp",
		Body:     "哈哈",
		Att: []string{
			"/etc/hosts",
		},
	}
	if _, err := m.String(); err != nil {
		t.Fatal(err)
	}
	if _, err := m.Reader(); err != nil {
		t.Fatal(err)
	}
}
