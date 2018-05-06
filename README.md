## mailx
[![Build Status](https://travis-ci.org/txthinking/mailx.svg?branch=master)](https://travis-ci.org/txthinking/mailx)
[![Go Report Card](https://goreportcard.com/badge/github.com/txthinking/mailx)](https://goreportcard.com/report/github.com/txthinking/mailx)
[![GoDoc](https://godoc.org/github.com/txthinking/mailx?status.svg)](https://godoc.org/github.com/txthinking/mailx)

A SMTP/Mailgun/etc mail library, allow to set a fake from address.

### Install
```
$ go get github.com/txthinking/mailx
```

### Example

```
package main

import(
    "log"
    "net/mail"
    "github.com/txthinking/mailx"
)

func main(){
    smtp := &mailx.SMTP{
        Server: "smtp.example.com",
        Port: 25,
        UserName: "tom@txthinking.com",
        Password: "password",
        IsTLS: false,
    }

    from := &mail.Address{
        Name: "Tom",
        Address: "tom@txthinking.com",
    }
    to := []*mail.Address{
        &mail.Address{
            Name: "Jerry",
            Address: "jerry@txthinking.com",
        },
    }
    fakefrom := &mail.Address{
        Name: "Obama",
        Address: "obama@whitehouse.gov",
    }

    msg := &mailx.Message{
        From: from,
        FakeFrom: fakeFrom, // Optional, if u want, a fake from address
        To: to,
        Subject: "hello",
        Body: "I <b>love</b> you.",
        Att: []string{
            "/etc/hosts",
        },
    }

    err := smtp.Send(msg)
    if err != nil{
        log.Fatal(err)
    }
}
```
