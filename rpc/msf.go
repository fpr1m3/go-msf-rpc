package rpc

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"net/http"

	"github.com/vmihailenco/msgpack/v5"
)

type Metasploit struct {
	host  string
	user  string
	pass  string
	token string
}

func New(host, user, pass string) (*Metasploit, error) {
	msf := &Metasploit{
		host: host,
		user: user,
		pass: pass,
	}

	if err := msf.Login(); err != nil {
		return nil, err
	}

	return msf, nil
}

func (msf *Metasploit) send(req interface{}, res interface{}) error {
	buf := new(bytes.Buffer)
	msgpack.NewEncoder(buf).Encode(req)
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	dest := fmt.Sprintf("%s/api", msf.host)
	response, err := http.Post(dest, "binary/message-pack", buf)
	// responseBytes, _ := httputil.DumpResponse(response, true)
	// log.Printf("Response dump: %s\n", string(responseBytes))
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if err := msgpack.NewDecoder(response.Body).Decode(&res); err != nil {
		return err
	}
	return nil
}
