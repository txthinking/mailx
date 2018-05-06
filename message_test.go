package mailx

import (
	"net/mail"
	"testing"
)

func TestMessage(t *testing.T) {

	m := &Message{
		From: &mail.Address{
			Name:    "mailx",
			Address: "739f35c64d-48cf45@inbox.mailtrap.io",
		},
		FakeFrom: &mail.Address{
			Name:    "Obama",
			Address: "fakename@fakeaddress.com",
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
	if _, err := m.Reader(); err != nil {
		t.Fatal(err)
	}
}
