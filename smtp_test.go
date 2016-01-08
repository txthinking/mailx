package xmail

import(
    "testing"
    "net/mail"
)

func TestSendMail(t *testing.T){
    p := "tx111111"
    s := &SMTP{
        Server: "smtp.ym.163.com",
        Port: 25,
        UserName: "bot@ym.txthinking.com",
        Password: p,
        IsTLS: false,
    }
    f := &mail.Address{
        Name: "雷锋",
        Address: "bot@ym.txthinking.com",
    }
    ts := []*mail.Address{
        &mail.Address{
            Name: "雷锋",
            Address: "cloud@txthinking.com",
        },
    }
    m := &Message{
        From: f,
        To: ts,
        Subject: "hello",
        Body: "哈哈",
        Att: []string{
            "/etc/hosts",
        },
    }

    err := s.Send(m)
    if err != nil{
        t.Fatal(err)
    }
}
