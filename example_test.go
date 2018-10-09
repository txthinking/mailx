package mailx_test

import (
	"log"
	"net/mail"

	"github.com/txthinking/mailx"
)

func Example() {
	server := &mailx.SMTP{
		Server:   "smtp.mailtrap.io",
		Port:     465,
		UserName: "e3f534cfe656f4",
		Password: "b6e38ddc0f1e9d",
	}

	message := &mailx.Message{
		From: &mail.Address{
			Name:    "mailx",
			Address: "739f35c64d-48cf45@inbox.mailtrap.io",
		},
		To: []*mail.Address{
			{
				Name:    "Cloud",
				Address: "cloud@txthinking.com",
			},
		},
		Subject: "Hello",
		Body:    "I <b>love</b> U.",
		Attachment: []string{
			"/etc/hosts",
		},
	}

	if err := server.Send(message); err != nil {
		log.Fatal(err)
	}
}
