package main

import(
    "flag"
    "strings"
    "fmt"
    "net/mail"
    "log"
    "github.com/txthinking/xmail"
)

func usage(){
    usage := `Usage:
    -h         This help.

    -server    required smtp server
    -port      required smtp port
    -username  required smtp user name
    -password  required smtp password
    -tls       optional smtp use or not tls

    -from      required from
    -to        required to, like "a@a.com" or more "a@a.com:b@a.com"
    -subject   optional subject
    -body      optional body
    -att       optional attachment, like "/path/to/a" or more "/path/to/a:/path/to/b"

Creator: Cloud <cloud@txthinking.com>
`
        fmt.Print(usage)
}

var h bool
var server string
var port int
var username string
var password string
var tls bool
var from string
var to string
var subject string
var body string
var att string

func main(){
    flag.BoolVar(&h, "h", false, "usage")
    flag.StringVar(&server, "server", "", "SMTP server")
    flag.IntVar(&port, "port", 0, "SMTP PORT")
    flag.StringVar(&username, "username", "", "SMTP user name")
    flag.StringVar(&password, "password", "", "SMTP password")
    flag.BoolVar(&tls, "tls", false, "Is TLS")
    flag.StringVar(&from, "from", "", "from")
    flag.StringVar(&to, "to", "", "to")
    flag.StringVar(&subject, "subject", "", "subject")
    flag.StringVar(&body, "body", "", "body")
    flag.StringVar(&att, "att", "", "attachment")
    flag.Parse()
    if h {
        usage()
        return
    }

    f, err := mail.ParseAddress(from)
    if err != nil{
        log.Fatal(err)
        return
    }
    var tos []*mail.Address = make([]*mail.Address, 0)
    for _, s := range strings.Split(to,":"){
        s = strings.TrimSpace(s)
        if s != ""{
            a, err := mail.ParseAddress(s)
            if err != nil{
                log.Fatal(err)
                return
            }
            tos = append(tos, a)
        }
    }
    var atts []string = make([]string, 0)
    for _, s := range strings.Split(att,":"){
        s = strings.TrimSpace(s)
        if s != ""{
            atts = append(atts, s)
        }
    }

    m := &xmail.Message{
        From: f,
        To: tos,
        Subject: subject,
        Body: body,
        Att: atts,
    }
    s := &xmail.SMTP{
        Server: server,
        Port: port,
        UserName: username,
        Password: password,
        IsTLS: tls,
    }
    err = s.Send(m)
    if err != nil{
        log.Fatal(err)
    }
}

