package xmail

import(
    "time"
    "bytes"
    "io"
    "fmt"
    "encoding/base64"
    "strconv"
    "path/filepath"
    "math/rand"
    "io/ioutil"
    "net/mail"
    "github.com/txthinking/ant"
)

// Create boundary for MIME data.
func MakeBoundary()(b string){
    b = strconv.FormatInt(time.Now().UnixNano(), 10)
    b += strconv.FormatInt(rand.New(rand.NewSource(time.Now().UnixNano())).Int63(), 10)
    b = ant.MD5(b)
    return
}

// Chunk data using RFC 2045.
func ChunkSplit(s string)(r string, err error){
    const LENTH = 76
    var bfr, bfw *bytes.Buffer
    var data, block []byte

    data = make([]byte, 0)
    block = make([]byte, LENTH)
    bfr = bytes.NewBufferString(s)
    bfw = bytes.NewBuffer(data)
    var l int
    for {
        l, err = bfr.Read(block)
        if err == io.EOF{
            err = nil
            break
        }
        if err != nil{
            return
        }
        _, err = bfw.Write(block[:l])
        if err != nil{
            return
        }
        _, err = bfw.WriteString("\r\n")
        if err != nil{
            return
        }
    }
    r = bfw.String()
    return
}

type Message struct{
    From *mail.Address
    To []*mail.Address
    Subject string
    Body string
    Att []string
}

// MIME RFC2045
// https://en.wikipedia.org/wiki/MIME
// from/to can be "some@domain.com"
// or RFC 5322 address, e.g. "Barry Gibbs <bg@example.com>"
// NOTICE: If the name is not ASCII, you should be do some encoding,
// eg:"=?utf-8?B?"+base64.StdEncoding.EncodeToString([]byte("名字"))+"?="
// att is the file path for attachments or nil if no attachment.
func (m *Message)Reader()(r io.Reader, err error){
    var s string
    var i int
    var bs []byte
    var bf *bytes.Buffer = bytes.NewBufferString("")

    // prepare body data
    m.Body, err = ChunkSplit(base64.StdEncoding.EncodeToString(bytes.NewBufferString(m.Body).Bytes()))
    if err != nil {
        return
    }

    // prepare attachment data
    var attms []map[string]string
    var attm map[string]string
    if len(m.Att) != 0{
        attms = make([]map[string]string, len(m.Att))
        for i, s = range m.Att{
            attm = make(map[string]string)
            bs, err = ioutil.ReadFile(s)
            if err != nil{
                return
            }
            attm["name"] = filepath.Base(s)
            attm["data"], err = ChunkSplit(base64.StdEncoding.EncodeToString(bs))
            if err != nil {
                return
            }
            attms[i] = attm
        }
    }

    // prepare mail data
    var boundaryAlternative, boundaryMixed string

    bf.WriteString(fmt.Sprintf("Subject: %s\r\n", m.Subject))
    bf.WriteString(fmt.Sprintf("From: %s\r\n", m.From.String()))
    bf.WriteString(fmt.Sprintf("MIME-Version: 1.0\r\n"))
    s = ""
    for i, _ = range m.To{
        s += m.To[i].String() + ", "
    }
    bf.WriteString(fmt.Sprintf("To: %s\r\n", s[:len(s)-2]))

    boundaryAlternative = MakeBoundary()
    if len(m.Att) != 0{
        boundaryMixed = MakeBoundary()
        bf.WriteString(fmt.Sprintf("Content-Type: multipart/mixed; boundary=\"%s\"\r\n\r\n", boundaryMixed))
        bf.WriteString(fmt.Sprintf("--%s\r\n", boundaryMixed))
    }

    bf.WriteString(fmt.Sprintf("Content-Type: multipart/alternative; boundary=\"%s\"\r\n\r\n", boundaryAlternative))
    bf.WriteString(fmt.Sprintf("--%s\r\n", boundaryAlternative))
    bf.WriteString(fmt.Sprintf("Content-Type: text/plain; charset=utf-8\r\n"))
    bf.WriteString(fmt.Sprintf("Content-Transfer-Encoding: base64\r\n\r\n"))
    bf.WriteString(fmt.Sprintf("%s\r\n\r\n", m.Body))
    bf.WriteString(fmt.Sprintf("--%s\r\n", boundaryAlternative))
    bf.WriteString(fmt.Sprintf("Content-Type: text/html; charset=utf-8\r\n"))
    bf.WriteString(fmt.Sprintf("Content-Transfer-Encoding: base64\r\n\r\n"))
    bf.WriteString(fmt.Sprintf("%s\r\n\r\n", m.Body))
    bf.WriteString(fmt.Sprintf("--%s--\r\n", boundaryAlternative))

    if len(m.Att) != 0{
        for _, attm = range attms{
            bf.WriteString(fmt.Sprintf("--%s\r\n", boundaryMixed))
            bf.WriteString(fmt.Sprintf("Content-Type: application/octet-stream; name=\"%s\"\r\n", attm["name"]))
            bf.WriteString(fmt.Sprintf("Content-Transfer-Encoding: base64\r\n"))
            bf.WriteString(fmt.Sprintf("Content-Disposition: attachment; filename=\"%s\"\r\n\r\n", attm["name"]))
            bf.WriteString(fmt.Sprintf("%s\r\n\r\n", attm["data"]))
        }
        bf.WriteString(fmt.Sprintf("--%s--\r\n", boundaryMixed))
    }
    r = bf
    return
}

