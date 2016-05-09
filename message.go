package xmail

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
	Att                 []string
	boundaryAlternative string
	boundaryMixed       string
}

func (m *Message) initBoundary() {
	m.boundaryAlternative = "a" + makeBoundary()
	m.boundaryMixed = "m" + makeBoundary()
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
	if len(m.Att) != 0 {
		body += fmt.Sprintf("Content-Type: multipart/mixed; boundary=\"%s\"\r\n\r\n", m.boundaryMixed)
		body += fmt.Sprintf("--%s\r\n", m.boundaryMixed)
	}

	content, err := m.content()
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

	if len(m.Att) != 0 {
		attms, err := m.attachment()
		if err != nil {
			return "", err
		}
		for _, attm := range attms {
			body += fmt.Sprintf("--%s\r\n", m.boundaryMixed)
			body += fmt.Sprintf("Content-Type: application/octet-stream; name=\"%s\"\r\n", attm["name"])
			body += fmt.Sprintf("Content-Transfer-Encoding: base64\r\n")
			body += fmt.Sprintf("Content-Disposition: attachment; filename=\"%s\"\r\n\r\n", attm["name"])
			body += fmt.Sprintf("%s\r\n\r\n", attm["data"])
		}
		body += fmt.Sprintf("--%s--\r\n", m.boundaryMixed)
	}
	return body, nil
}

func (m *Message) content() (string, error) {
	body, err := chunkSplit(base64.StdEncoding.EncodeToString(bytes.NewBufferString(m.Body).Bytes()))
	return body, err
}

func (m *Message) attachment() ([]map[string]string, error) {
	attms := make([]map[string]string, len(m.Att))
	for i, s := range m.Att {
		attm := make(map[string]string)
		b, err := ioutil.ReadFile(s)
		if err != nil {
			return nil, err
		}
		attm["name"] = filepath.Base(s)
		attm["data"], err = chunkSplit(base64.StdEncoding.EncodeToString(b))
		if err != nil {
			return nil, err
		}
		attms[i] = attm
	}
	return attms, nil
}

// String return a string of mail message.
func (m *Message) String() (string, error) {
	m.initBoundary()
	header, err := m.header()
	if err != nil {
		return "", err
	}
	body, err := m.body()
	if err != nil {
		return "", err
	}
	return header + body, nil
}

// Reader return a io.Reader of mail message.
func (m *Message) Reader() (io.Reader, error) {
	data, err := m.String()
	if err != nil {
		return nil, err
	}
	return strings.NewReader(data), nil
}
