package request

import (
	"errors"
	"io"
	"strings"
)

type Request struct {
	RequestLine RequestLine
}

type RequestLine struct {
	HttpVersion   string
	RequestTarget string
	Method        string
}

func catch(e error) {
	if e != nil {
		panic(e)
	}
}

func RequestFromReader(reader io.Reader) (*Request, error) {
	b, err := io.ReadAll(reader)
	catch(err)
	line, err := parseRequestLine(b)

	request := Request{
		RequestLine: line,
	}

	return &request, err
}

func parseRequestLine(b []byte) (RequestLine, error) {
	str := string(b)
	parts := strings.Fields(str)
	ver := strings.TrimPrefix(parts[2], "HTTP/")

	err := errors.New(nil)

	if parts[1] != strings.ToUpper(parts[1]) {
		err = errors.New("method type bad")
	}

	requestLine := RequestLine{
		HttpVersion:   ver,
		RequestTarget: parts[1],
		Method:        parts[0],
	}

	return requestLine, err
}
