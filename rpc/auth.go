package rpc

import "fmt"

// Auth
type loginReq struct {
	_msgpack struct{} `msgpack:",asArray"`
	Method   string
	Username string
	Password string
}

type loginRes struct {
	Result       string `msgpack:"result"`
	Token        string `msgpack:"token"`
	Error        bool   `msgpack:"error"`
	ErrorClass   string `msgpack:"error_class"`
	ErrorMessage string `msgpack:"error_message"`
}

type logoutReq struct {
	_msgpack    struct{} `msgpack:",asArray"`
	Method      string
	Token       string
	LogoutToken string
}

type logoutRes struct {
	Result string `msgpack:"result"`
}

func (msf *Metasploit) Login() error {
	ctx := &loginReq{
		Method:   "auth.login",
		Username: msf.user,
		Password: msf.pass,
	}

	var res loginRes
	if err := msf.send(ctx, &res); err != nil {
		fmt.Println("Failed at login")
		return err
	}
	msf.token = res.Token
	return nil
}

func (msf *Metasploit) Logout() error {
	ctx := &logoutReq{
		Method:      "auth.logout",
		Token:       msf.token,
		LogoutToken: msf.token,
	}

	var res logoutRes
	if err := msf.send(ctx, &res); err != nil {
		return err
	}
	msf.token = ""
	return nil
}
