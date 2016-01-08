## xmail

A email library

### Install
```
$ go get github.com/txthinking/xmail
```

### Usage

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
    msg := &Message{
        From: from,
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
