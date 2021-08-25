package rpc

// Console

type consoleCreateReq struct {
	_msgpack struct{} `msgpack:",asArray"`
	Method   string
	Token    string
}

type consoleCreateRes struct {
	Id     string `msgpack:"id"`
	Prompt string `msgpack:"prompt"`
	Busy   bool   `msgpack:"busy"`
}

type consoleDestroyReq struct {
	_msgpack  struct{} `msgpack:",asArray"`
	Method    string
	Token     string
	ConsoleId string
}

type consoleDestroyRes struct {
	Result string `msgpack:"result"`
}

type consoleListReq struct {
	_msgpack struct{} `msgpack:",asArray"`
	Method   string
	Token    string
}

type consoleListRes map[string][]struct {
	Id     string `msgpack:"id"`
	Prompt string `msgpack:"prompt"`
	Busy   bool   `msgpack:"busy"`
}

type consoleWriteReq struct {
	_msgpack  struct{} `msgpack:",asArray"`
	Method    string
	Token     string
	ConsoleId string
	Command   string
}

type consoleWriteRes struct {
	Wrote uint32 `msgpack:"wrote"`
}

type consoleReadReq struct {
	_msgpack  struct{} `msgpack:",asArray"`
	Method    string
	Token     string
	ConsoleId string
}

type consoleReadRes struct {
	Data   string `msgpack:"data"`
	Prompt string `msgpack:"prompt"`
	Busy   bool   `msgpack:"busy"`
}

type consoleSessionDetachReq struct {
	_msgpack  struct{} `msgpack:",asArray"`
	Method    string
	Token     string
	ConsoleId string
}

type consoleSessionDetachRes struct {
	Result string `msgpack:"result"`
}

type consoleSessionKillReq struct {
	_msgpack  struct{} `msgpack:",asArray"`
	Method    string
	Token     string
	ConsoleId string
}

type consoleSessionKillRes struct {
	Result string `msgpack:"result"`
}

type consoleTabsReq struct {
	_msgpack  struct{} `msgpack:",asArray"`
	Method    string
	Token     string
	ConsoleId string
	InputLine string
}

type consoleTabsRes struct {
	Tabs []string `msgpack:"tabs"`
}

// Console

func (msf *Metasploit) ConsoleCreate() (consoleCreateRes, error) {
	ctx := &consoleCreateReq{
		Method: "console.create",
		Token:  msf.token,
	}
	var res consoleCreateRes
	if err := msf.send(ctx, &res); err != nil {
		return consoleCreateRes{}, err
	}
	return res, nil
}

func (msf *Metasploit) ConsoleDestroy(consoleid string) (consoleDestroyRes, error) {
	ctx := &consoleDestroyReq{
		Method:    "console.destroy",
		Token:     msf.token,
		ConsoleId: consoleid,
	}
	var res consoleDestroyRes
	if err := msf.send(ctx, &res); err != nil {
		return consoleDestroyRes{}, err
	}
	return res, nil
}

func (msf *Metasploit) ConsoleList() (consoleListRes, error) {
	ctx := &consoleListReq{
		Method: "console.list",
		Token:  msf.token,
	}
	var res consoleListRes
	if err := msf.send(ctx, &res); err != nil {
		return consoleListRes{}, err
	}
	return res, nil
}

func (msf *Metasploit) ConsoleWrite(consoleId, command string) (consoleWriteRes, error) {
	ctx := &consoleWriteReq{
		Method:    "console.write",
		Token:     msf.token,
		ConsoleId: consoleId,
		Command:   command,
	}
	var res consoleWriteRes
	if err := msf.send(ctx, &res); err != nil {
		return consoleWriteRes{}, err
	}
	return res, nil
}

func (msf *Metasploit) ConsoleRead(consoleId string) (consoleReadRes, error) {
	ctx := &consoleReadReq{
		Method:    "console.read",
		Token:     msf.token,
		ConsoleId: consoleId,
	}
	var res consoleReadRes
	if err := msf.send(ctx, &res); err != nil {
		return consoleReadRes{}, err
	}
	return res, nil
}

func (msf *Metasploit) ConsoleSessionDetch(consoleId string) (consoleSessionDetachRes, error) {
	ctx := &consoleSessionDetachReq{
		Method:    "console.session_detach",
		Token:     msf.token,
		ConsoleId: consoleId,
	}
	var res consoleSessionDetachRes
	if err := msf.send(ctx, &res); err != nil {
		return consoleSessionDetachRes{}, err
	}
	return res, nil
}

func (msf *Metasploit) ConsoleSessionKill(consoleId string) (consoleSessionKillRes, error) {
	ctx := &consoleSessionKillReq{
		Method:    "console.session_kill",
		Token:     msf.token,
		ConsoleId: consoleId,
	}
	var res consoleSessionKillRes
	if err := msf.send(ctx, &res); err != nil {
		return consoleSessionKillRes{}, err
	}
	return res, nil
}

func (msf *Metasploit) ConsoleTabs(consoleId, inputLine string) (consoleTabsRes, error) {
	ctx := &consoleTabsReq{
		Method:    "console.tabs",
		Token:     msf.token,
		ConsoleId: consoleId,
		InputLine: inputLine,
	}
	var res consoleTabsRes
	if err := msf.send(ctx, &res); err != nil {
		return consoleTabsRes{}, err
	}
	return res, nil
}
