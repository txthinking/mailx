package xmail

import(
    "testing"
    "net/mail"
)

func TestMailgun(t *testing.T){
    k := ""
    s := &Mailgun{
        "mail.txthinking.com",
        k,
    }
    f := &mail.Address{
        Name: "雷锋",
        Address: "bot@mail.txthinking.com",
    }
    ts := []*mail.Address{
        &mail.Address{
            Name: "雷锋他爹",
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

