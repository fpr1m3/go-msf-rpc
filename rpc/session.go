package rpc

import "fmt"

// Session
type sessionListReq struct {
	_msgpack struct{} `msgpack:",asArray"`
	Method   string
	Token    string
}

type SessionListRes map[uint32]struct {
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

type sessionMeterpreterWriteReq struct {
	_msgpack  struct{} `msgpack:",asArray"`
	Method    string
	Token     string
	SessionID uint32
	Command   string
}

type sessionMeterpreterWriteRes struct {
	Result string `msgpack:"result"`
}

type sessionMeterpreterReadReq struct {
	_msgpack  struct{} `msgpack:",asArray"`
	Method    string
	Token     string
	SessionID uint32
}

type sessionMeterpreterReadRes struct {
	Data string `msgpack:"data"`
}

type sessionMeterpreterRunSingleReq struct {
	_msgpack  struct{} `msgpack:",asArray"`
	Method    string
	Token     string
	SessionID uint32
	Command   string
}

type sessionMeterpreterRunSingleRes sessionMeterpreterWriteRes

type sessionMeterpreterDetachReq struct {
	_msgpack  struct{} `msgpack:",asArray"`
	Method    string
	Token     string
	SessionID uint32
}

type sessionMeterpreterDetachRes sessionMeterpreterWriteRes

type sessionMeterpreterKillReq struct {
	_msgpack  struct{} `msgpack:",asArray"`
	Method    string
	Token     string
	SessionID uint32
}

type sessionMeterpreterKillRes sessionMeterpreterWriteRes

type sessionMeterpreterTabsReq struct {
	_msgpack  struct{} `msgpack:",asArray"`
	Method    string
	Token     string
	SessionID uint32
	InputLine string
}

type sessionMeterpreterTabsRes struct {
	Tabs []string `msgpack:"tabs"`
}

type sessionCompatibleModulesReq struct {
	_msgpack  struct{} `msgpack:",asArray"`
	Method    string
	Token     string
	SessionID uint32
}

type sessionCompatibleModulesRes struct {
	Modules []string `msgpack:"modules"`
}

type sessionShellUpgradeReq struct {
	_msgpack   struct{} `msgpack:",asArray"`
	Method     string
	Token      string
	SessionID  uint32
	IpAddress  string
	PortNumber uint32
}

type sessionShellUpgradeRes sessionMeterpreterWriteRes

type sessionRingClearReq struct {
	_msgpack  struct{} `msgpack:",asArray"`
	Method    string
	Token     string
	SessionID uint32
}

type sessionRingClearRes sessionMeterpreterWriteRes

type sessionRingPutReq struct {
	_msgpack  struct{} `msgpack:",asArray"`
	Method    string
	Token     string
	SessionID uint32
	Command   string
}

type sessionRingPutRes struct {
	WriteCount uint32 `msgpack:"write_count"`
}

func (msf *Metasploit) SessionList() (SessionListRes, error) {
	req := &sessionListReq{
		Method: "session.list",
		Token:  msf.token,
	}

	var res SessionListRes
	if err := msf.send(req, &res); err != nil {
		return nil, err
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

func (msf *Metasploit) SessionMeterpreterWrite(session uint32, command string) (sessionMeterpreterWriteRes, error) {
	ctx := &sessionMeterpreterWriteReq{
		Method:    "session.meterpreter_write",
		Token:     msf.token,
		SessionID: session,
		Command:   command,
	}

	var res sessionMeterpreterWriteRes
	if err := msf.send(ctx, &res); err != nil {
		return sessionMeterpreterWriteRes{}, err
	}

	return res, nil
}

func (msf *Metasploit) SessionMeterpreterRead(session uint32) (sessionMeterpreterReadRes, error) {
	ctx := &sessionMeterpreterReadReq{
		Method:    "session.meterpreter_read",
		Token:     msf.token,
		SessionID: session,
	}

	var res sessionMeterpreterReadRes
	if err := msf.send(ctx, &res); err != nil {
		return sessionMeterpreterReadRes{}, err
	}
	return res, nil
}

func (msf *Metasploit) SessionMeterpreterRunSingle(session uint32, command string) (sessionMeterpreterRunSingleRes, error) {
	ctx := &sessionMeterpreterRunSingleReq{
		Method:    "session.meterpreter_run_single",
		Token:     msf.token,
		SessionID: session,
		Command:   command,
	}

	var res sessionMeterpreterRunSingleRes
	if err := msf.send(ctx, &res); err != nil {
		return sessionMeterpreterRunSingleRes{}, err
	}

	return res, nil
}

func (msf *Metasploit) SessionMeterpreterSessionDetach(session uint32) (sessionMeterpreterDetachRes, error) {
	ctx := &sessionMeterpreterDetachReq{
		Method:    "session.meterpreter_session_detach",
		Token:     msf.token,
		SessionID: session,
	}

	var res sessionMeterpreterDetachRes
	if err := msf.send(ctx, &res); err != nil {
		return sessionMeterpreterDetachRes{}, err
	}
	return res, nil
}

func (msf *Metasploit) SessionMeterpreterSessionKill(session uint32) (sessionMeterpreterKillRes, error) {
	ctx := &sessionMeterpreterKillReq{
		Method:    "session.meterpreter_session_kill",
		Token:     msf.token,
		SessionID: session,
	}

	var res sessionMeterpreterKillRes
	if err := msf.send(ctx, &res); err != nil {
		return sessionMeterpreterKillRes{}, err
	}
	return res, nil
}

func (msf *Metasploit) SessionMeterpreterTabs(session uint32, inputLine string) (sessionMeterpreterTabsRes, error) {
	ctx := &sessionMeterpreterTabsReq{
		Method:    "session.meterpreter_tabs",
		Token:     msf.token,
		SessionID: session,
		InputLine: inputLine,
	}

	var res sessionMeterpreterTabsRes
	if err := msf.send(ctx, &res); err != nil {
		return sessionMeterpreterTabsRes{}, err
	}
	return res, nil
}

func (msf *Metasploit) SessionCompatibleModules(session uint32) (sessionCompatibleModulesRes, error) {
	ctx := &sessionCompatibleModulesReq{
		Method:    "session.compatible_modules",
		Token:     msf.token,
		SessionID: session,
	}

	var res sessionCompatibleModulesRes
	if err := msf.send(ctx, &res); err != nil {
		return sessionCompatibleModulesRes{}, err
	}
	return res, nil
}

func (msf *Metasploit) SessionShellUpgrade(session uint32, lhostAddress string, lportNumber uint32) (sessionShellUpgradeRes, error) {
	ctx := &sessionShellUpgradeReq{
		Method:     "session.shell_upgrade",
		Token:      msf.token,
		SessionID:  session,
		IpAddress:  lhostAddress,
		PortNumber: lportNumber,
	}

	var res sessionShellUpgradeRes
	if err := msf.send(ctx, &res); err != nil {
		return sessionShellUpgradeRes{}, err
	}
	return res, nil
}

func (msf *Metasploit) SessionRingClear(session uint32) (sessionRingClearRes, error) {
	ctx := &sessionRingClearReq{
		Method:    "session.ring_clear",
		Token:     msf.token,
		SessionID: session,
	}

	var res sessionRingClearRes
	if err := msf.send(ctx, &res); err != nil {
		return sessionRingClearRes{}, err
	}
	return res, nil
}

func (msf *Metasploit) SessionRingLast(session uint32) (sessionRingLastRes, error) {
	ctx := &sessionRingLastReq{
		Method:    "session.ring_last",
		Token:     msf.token,
		SessionID: session,
	}

	var res sessionRingLastRes
	if err := msf.send(ctx, &res); err != nil {
		return sessionRingLastRes{}, err
	}
	return res, nil
}

func (msf *Metasploit) SessionRingPut(session uint32, command string) (sessionRingPutRes, error) {
	ctx := &sessionRingPutReq{
		Method:    "session.ring_put",
		Token:     msf.token,
		SessionID: session,
		Command:   command,
	}

	var res sessionRingPutRes
	if err := msf.send(ctx, &res); err != nil {
		return sessionRingPutRes{}, err
	}
	return res, nil
}
