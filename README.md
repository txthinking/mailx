## mailx
[![Build Status](https://travis-ci.org/txthinking/mailx.svg?branch=master)](https://travis-ci.org/txthinking/mailx)
[![Go Report Card](https://goreportcard.com/badge/github.com/txthinking/mailx)](https://goreportcard.com/report/github.com/txthinking/mailx)
[![GoDoc](https://godoc.org/github.com/txthinking/mailx?status.svg)](https://godoc.org/github.com/txthinking/mailx)

[🗣 News](https://t.me/txthinking_news)
[💬 Chat](https://join.txthinking.com)
[🩸 Youtube](https://www.youtube.com/txthinking) 
[❤️ Sponsor](https://github.com/sponsors/txthinking)

A lightweight SMTP mail library

❤️ A project by [txthinking.com](https://www.txthinking.com)

### Install

```
$ go get github.com/txthinking/mailx
```

### Example

```
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
```

## License

Licensed under The MIT License
