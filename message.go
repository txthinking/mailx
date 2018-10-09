package mailx

import (
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/mail"
	"path/filepath"
	"strings"
	"time"
)

// Message is the email body.
type Message struct {
	From                *mail.Address
	FakeFrom            *mail.Address
	To                  []*mail.Address
	Subject             string
	Body                string
	Attachment          []string
	boundaryAlternative string
	boundaryMixed       string
}

func (m *Message) initBoundary() {
	m.boundaryAlternative = "a" + MakeBoundary()
	m.boundaryMixed = "m" + MakeBoundary()
}

func (m *Message) header() (string, error) {
	header := ""
	header += fmt.Sprintf("Date: %s\r\n", time.Now().Format(time.RFC1123Z))
	header += fmt.Sprintf("Subject: =?utf-8?B?%s?=\r\n", base64.StdEncoding.EncodeToString(bytes.NewBufferString(m.Subject).Bytes()))
	if m.FakeFrom != nil {
		header += fmt.Sprintf("From: %s\r\n", m.FakeFrom.String())
	} else {
		header += fmt.Sprintf("From: %s\r\n", m.From.String())
	}

	if len(m.To) == 0 {
		return "", errors.New("Miss to address")
	}
	to := ""
	for i := range m.To {
		to += m.To[i].String() + ", "
	}
	header += fmt.Sprintf("To: %s\r\n", to[:len(to)-2])
	header += fmt.Sprintf("MIME-Version: %s\r\n", "1.0")
	return header, nil
}

func (m *Message) body() (string, error) {
	body := ""
	if len(m.Attachment) != 0 {
		body += fmt.Sprintf("Content-Type: multipart/mixed; boundary=\"%s\"\r\n\r\n", m.boundaryMixed)
		body += fmt.Sprintf("--%s\r\n", m.boundaryMixed)
	}

	content, err := ChunkSplit(base64.StdEncoding.EncodeToString(bytes.NewBufferString(m.Body).Bytes()))
	if err != nil {
		return "", nil
	}

	body += fmt.Sprintf("Content-Type: multipart/alternative; boundary=\"%s\"\r\n\r\n", m.boundaryAlternative)
	body += fmt.Sprintf("--%s\r\n", m.boundaryAlternative)
	body += fmt.Sprintf("Content-Type: text/plain; charset=utf-8\r\n")
	body += fmt.Sprintf("Content-Transfer-Encoding: base64\r\n\r\n")
	body += fmt.Sprintf("%s\r\n\r\n", content)
	body += fmt.Sprintf("--%s\r\n", m.boundaryAlternative)
	body += fmt.Sprintf("Content-Type: text/html; charset=utf-8\r\n")
	body += fmt.Sprintf("Content-Transfer-Encoding: base64\r\n\r\n")
	body += fmt.Sprintf("%s\r\n\r\n", content)
	body += fmt.Sprintf("--%s--\r\n", m.boundaryAlternative)

	if len(m.Attachment) != 0 {
		for _, s := range m.Attachment {
			b, err := ioutil.ReadFile(s)
			if err != nil {
				return "", err
			}
			name := filepath.Base(s)
			data, err := ChunkSplit(base64.StdEncoding.EncodeToString(b))
			if err != nil {
				return "", err
			}
			body += fmt.Sprintf("--%s\r\n", m.boundaryMixed)
			body += fmt.Sprintf("Content-Type: application/octet-stream; name=\"%s\"\r\n", name)
			body += fmt.Sprintf("Content-Transfer-Encoding: base64\r\n")
			body += fmt.Sprintf("Content-Disposition: attachment; filename=\"%s\"\r\n\r\n", name)
			body += fmt.Sprintf("%s\r\n\r\n", data)
		}
		body += fmt.Sprintf("--%s--\r\n", m.boundaryMixed)
	}
	return body, nil
}

// Reader return a io.Reader of mail message.
func (m *Message) Reader() (io.Reader, error) {
	m.initBoundary()
	header, err := m.header()
	if err != nil {
		return nil, err
	}
	body, err := m.body()
	if err != nil {
		return nil, err
	}
	return strings.NewReader(header + body), nil
}
