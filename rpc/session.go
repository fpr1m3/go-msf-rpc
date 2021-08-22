package rpc

import "fmt"

// Session
type sessionListReq struct {
	_msgpack struct{} `msgpack:",asArray"`
	Method   string
	Token    string
}

type SessionListRes struct {
	ID          uint32 `msgpack:",omitempty"`
	Type        string `msgpack:"type"`
	TunnelLocal string `msgpack:"tunnel_local"`
	TunnelPeer  string `msgpack:"tunnel_peer"`
	ViaExploit  string `msgpack:"via_exploit"`
	ViaPayload  string `msgpack:"via_payload"`
	Description string `msgpack:"desc"`
	Info        string `msgpack:"info"`
	Workspace   string `msgpack:"workspace"`
	SessionHost string `msgpack:"session_host"`
	SessionPort int    `msgpack:"session_port"`
	Username    string `msgpack:"username"`
	UUID        string `msgpack:"uuid"`
	ExploitUUID string `msgpack:"exploit_uuid"`
}

type sessionWriteReq struct {
	_msgpack  struct{} `msgpack:",asArray"`
	Method    string
	Token     string
	SessionID uint32
	Command   string
}

type sessionWriteRes struct {
	WriteCount string `msgpack:"write_count"`
}

type sessionReadReq struct {
	_msgpack    struct{} `msgpack:",asArray"`
	Method      string
	Token       string
	SessionID   uint32
	ReadPointer string
}

type sessionReadRes struct {
	Seq  uint32 `msgpack:"seq"`
	Data string `msgpack:"data"`
}

type sessionRingLastReq struct {
	_msgpack  struct{} `msgpack:",asArray"`
	Method    string
	Token     string
	SessionID uint32
}

type sessionRingLastRes struct {
	Seq uint32 `msgpack:"seq"`
}

func (msf *Metasploit) SessionList() (map[uint32]SessionListRes, error) {
	req := &sessionListReq{
		Method: "session.list",
		Token:  msf.token,
	}

	res := make(map[uint32]SessionListRes)
	if err := msf.send(req, &res); err != nil {
		return nil, err
	}

	for id, session := range res {
		session.ID = id
		res[id] = session
	}
	return res, nil

}

func (msf *Metasploit) SessionReadPointer(session uint32) (uint32, error) {
	ctx := &sessionRingLastReq{
		Method:    "session.ring_last",
		Token:     msf.token,
		SessionID: session,
	}

	var sesRingLast sessionRingLastRes
	if err := msf.send(ctx, &sesRingLast); err != nil {
		return 0, err
	}

	return sesRingLast.Seq, nil
}

func (msf *Metasploit) SessionWrite(session uint32, command string) error {
	ctx := &sessionWriteReq{
		Method:    "session.shell_write",
		Token:     msf.token,
		SessionID: session,
		Command:   command,
	}

	var res sessionWriteRes
	if err := msf.send(ctx, &res); err != nil {
		return err
	}

	return nil
}

func (msf *Metasploit) SessionRead(session uint32, readPointer uint32) (string, error) {
	ctx := &sessionReadReq{
		Method:      "session.shell_read",
		Token:       msf.token,
		SessionID:   session,
		ReadPointer: string(readPointer),
	}

	var res sessionReadRes
	if err := msf.send(ctx, &res); err != nil {
		return "", err
	}

	return res.Data, nil
}
func (msf *Metasploit) SessionExecute(session uint32, command string) (string, error) {
	readPointer, err := msf.SessionReadPointer(session)
	if err != nil {
		return "", err
	}
	msf.SessionWrite(session, command)
	data, err := msf.SessionRead(session, readPointer)
	if err != nil {
		return "", err
	}
	return data, nil
}

func (msf *Metasploit) SessionExecuteList(session uint32, commands []string) (string, error) {
	var results string
	for _, command := range commands {
		tCommand := fmt.Sprintf("%s\n", command)
		result, err := msf.SessionExecute(session, tCommand)
		if err != nil {
			return results, err
		}
		results += result
	}

	return results, nil
}
