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
func (m *SMTP) Send(msg *Message) error {
    client, err := smtp.Dial(m.Server + ":" + strconv.Itoa(m.Port))
    if err != nil {
        return err
    }
    err = client.Hello(m.Server)
    if err != nil {
        return err
    }
    if m.IsTLS{
        err = client.StartTLS(&tls.Config{ServerName: m.Server, InsecureSkipVerify: true})
        if err != nil {
            return err
        }
    }
    auth := smtp.PlainAuth("", m.UserName, m.Password, m.Server)
    err = client.Auth(auth)
    if err != nil {
        return err
    }
    err = client.Mail(msg.From.Address)
    if err != nil {
        return err
    }
    for i, _ := range msg.To{
        err = client.Rcpt(msg.To[i].Address)
        if err != nil {
            return err
        }
    }
    var in io.WriteCloser
    in, err = client.Data()
    if err != nil {
        return err
    }
    var r io.Reader
    r, err = msg.Reader()
    if err != nil {
        return err
    }
    _, err = io.Copy(in, r)
    if err != nil {
        return err
    }
    err = in.Close()
    if err != nil {
        return err
    }
    err = client.Quit()
    if err != nil {
        return err
    }
    return nil
}

