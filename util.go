package mailx

import (
	"bytes"
	"io"
	"math/rand"
	"strconv"
	"time"
)

// Create boundary for MIME data.
func MakeBoundary() string {
	b := strconv.FormatInt(time.Now().UnixNano(), 10)
	b += strconv.FormatInt(rand.New(rand.NewSource(time.Now().UnixNano())).Int63(), 10)
	return b
}

// Chunk data using RFC 2045.
func ChunkSplit(s string) (string, error) {
	const LENTH = 76
	var data = make([]byte, 0)
	var block = make([]byte, LENTH)
	bfr := bytes.NewBufferString(s)
	bfw := bytes.NewBuffer(data)
	for {
		l, err := bfr.Read(block)
		if err == io.EOF {
			err = nil
			break
		}
		if err != nil {
			return "", err
		}
		_, err = bfw.Write(block[:l])
		if err != nil {
			return "", err
		}
		_, err = bfw.WriteString("\r\n")
		if err != nil {
			return "", err
		}
	}
	r := bfw.String()
	return r, nil
}
