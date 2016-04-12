package xmail

import(
    "net/http"
    "io"
    "strings"
    "io/ioutil"
    "errors"
    "bytes"
    "github.com/txthinking/ant"
)

const MAILGUN_API_URL = "https://api.mailgun.net/v3"

type Mailgun struct {
    Domain string
    APIKey string
}

func (m *Mailgun) Send(msg *Message) error {
    msgr, err := msg.Reader()
    if err != nil {
        return err
    }
    to := make([]string, len(msg.To))
    for i, _ := range msg.To{
        to[i] = msg.To[i].String()
    }
    bdry := MakeBoundary()
    body, err := ant.MultipartFormDataFromReader(
        map[string][]string{
            "to": []string{strings.Join(to, ",")},
        },
        map[string][]io.Reader{
            "message": []io.Reader{msgr},
        },
        bdry,
    )

    tspt := &http.Transport{
        TLSClientConfig:    nil,
        DisableCompression: true,
    }
    client := &http.Client{Transport: tspt}
    r, err := http.NewRequest("POST", MAILGUN_API_URL+"/"+m.Domain+"/messages.mime", body)
    if err != nil{
        return err
    }
    r.Header.Add("Content-Type", "multipart/form-data; boundary="+bdry)
    r.SetBasicAuth("api", m.APIKey)

    res, err := client.Do(r)
    defer res.Body.Close()
    if res.StatusCode == http.StatusOK{
        return nil
    }
    b, err := ioutil.ReadAll(res.Body)
    if err != nil{
        return err
    }
    err = errors.New(bytes.NewBuffer(b).String())
    return err
}

