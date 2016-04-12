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
        Server: "smtp.ym.163.com",
        Port: 25,
        UserName: "bot@ym.txthinking.com",
        Password: "password",
        IsTLS: false,
    }

    from := &mail.Address{
        Name: "雷锋",
        Address: "bot@ym.txthinking.com",
    }
    to := []*mail.Address{
        &mail.Address{
            Name: "雷锋",
            Address: "cloud@txthinking.com",
        },
    }
    fakefrom := &mail.Address{
        Name: "雷锋",
        Address: "bot@ym.txthinking.com",
    }

    msg := &xmail.Message{
        From: from,
        FakeFrom: fakeFrom, // Optional, if u want, a fake from address
        To: to,
        Subject: "hello",
        Body: "哈哈",
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
