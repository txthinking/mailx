package xmail_test

import (
	"github.com/txthinking/xmail"
	"log"
	"net/mail"
)

func Example() {
	server := &xmail.SMTP{
		Server:   "mailtrap.io",
		Port:     2525,
		UserName: "e3f534cfe656f4",
		Password: "b6e38ddc0f1e9d",
		IsTLS:    false,
	}

	from := &mail.Address{
		Name:    "Xmail",
		Address: "739f35c64d-48cf45@inbox.mailtrap.io",
	}
	tos := []*mail.Address{
		&mail.Address{
			Name:    "Cloud",
			Address: "cloud@txthinking.com",
		},
	}
	msg := &xmail.Message{
		From:    from,
		To:      tos,
		Subject: "Xmail test smtp",
		Body:    "哈哈",
		Att: []string{
			"/etc/hosts",
		},
	}

	if err := server.Send(msg); err != nil {
		log.Fatal(err)
	}
}
