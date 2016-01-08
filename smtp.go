package xmail

import(
    "net/smtp"
    "io"
    "crypto/tls"
    "strconv"
)

type SMTP struct{
    Server string
    Port int
    UserName string
    Password string
    IsTLS bool
}

// RFC 821,822,1869,2821
func (m *SMTP)Send(msg *Message)(err error){
    var client *smtp.Client
    var auth smtp.Auth

    var i int

    client, err = smtp.Dial(m.Server + ":" + strconv.Itoa(m.Port))
    if err != nil {
        return
    }
    err = client.Hello(m.Server)
    if err != nil {
        return
    }
    if m.IsTLS{
        err = client.StartTLS(&tls.Config{ServerName: m.Server, InsecureSkipVerify: true})
        if err != nil {
            return
        }
    }
    auth = smtp.PlainAuth("", m.UserName, m.Password, m.Server)
    err = client.Auth(auth)
    if err != nil {
        return
    }
    err = client.Mail(msg.From.Address)
    if err != nil {
        return
    }
    for i, _ = range msg.To{
        err = client.Rcpt(msg.To[i].Address)
        if err != nil {
            return
        }
    }
    var in io.WriteCloser
    in, err = client.Data()
    if err != nil {
        return
    }
    var r io.Reader
    r, err = msg.Reader()
    if err != nil {
        return
    }
    _, err = io.Copy(in, r)
    if err != nil {
        return
    }
    err = in.Close()
    if err != nil {
        return
    }
    err = client.Quit()
    if err != nil {
        return
    }
    return
}

