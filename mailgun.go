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

type Mailgun struct{
    Domain string
    APIKey string
}

func (m *Mailgun)Send(msg *Message)(err error){
    var rd io.Reader
    rd, err = msg.Reader()
    if err != nil {
        return
    }
    var to []string = make([]string, len(msg.To))
    var i int
    for i, _ = range msg.To{
        to[i] = msg.To[i].String()
    }
    var br io.Reader
    var bd string = MakeBoundary()
    br, err = ant.MultipartFormDataFromReader(
        map[string][]string{
            "to": []string{strings.Join(to, ",")},
        },
        map[string][]io.Reader{
            "message": []io.Reader{rd},
        },
        bd,
    )

    var tr *http.Transport = &http.Transport{
        TLSClientConfig:    nil,
        DisableCompression: true,
    }
    var client *http.Client = &http.Client{Transport: tr}
    var r *http.Request
    r, err = http.NewRequest("POST", MAILGUN_API_URL+"/"+m.Domain+"/messages.mime", br)
    if err != nil{
        return
    }
    r.Header.Add("Content-Type", "multipart/form-data; boundary="+bd)
    r.SetBasicAuth("api", m.APIKey)
    var res *http.Response
    res, err = client.Do(r)
    if res.StatusCode == http.StatusOK{
        return
    }
    var b []byte
    b, err = ioutil.ReadAll(res.Body)
    if err != nil{
        return
    }
    res.Body.Close()
    err = errors.New(bytes.NewBuffer(b).String())
    return
}
