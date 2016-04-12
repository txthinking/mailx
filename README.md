## xmail
[![GoDoc](https://godoc.org/github.com/txthinking/xmail?status.svg)](https://godoc.org/github.com/txthinking/xmail)
[![Build Status](https://drone.io/github.com/txthinking/xmail/status.png)](https://drone.io/github.com/txthinking/xmail/latest)
[Binary](https://drone.io/github.com/txthinking/xmail/files/cmd/sendmail/sendmail)

A SMTP/Mailgun/etc mail library, allow to set a fake from address.

### Install
```
$ go get github.com/txthinking/xmail
```

### Example

```
package main

import(
    "log"
    "net/mail"
    "github.com/txthinking/xmail"
)

func main(){
    smtp := &xmail.SMTP{
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

    msg := &xmail.Message{
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
