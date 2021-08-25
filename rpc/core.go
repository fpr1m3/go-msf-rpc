package rpc

// Core
type coreAddModulePathReq struct {
	_msgpack struct{} `msgpack:",asArray"`
	Method   string
	Token    string
	Path     string
}

type coreAddModulePathRes struct {
	Exploits  uint32 `msgpack:"exploits"`
	Auxiliary uint32 `msgpack:"auxiliary"`
	Post      uint32 `msgpack:"post"`
	Encoders  uint32 `msgpack:"encoders"`
	Nops      uint32 `msgpack:"nops"`
	Payloads  uint32 `msgpack:"payloads"`
}

type coreModuleStatsReq struct {
	_msgpack struct{} `msgpack:",asArray"`
	Method   string
	Token    string
}

type coreModuleStatsRes struct {
	Exploits  uint32 `msgpack:"exploits"`
	Auxiliary uint32 `msgpack:"auxiliary"`
	Post      uint32 `msgpack:"post"`
	Encoders  uint32 `msgpack:"encoders"`
	Nops      uint32 `msgpack:"nops"`
	Payloads  uint32 `msgpack:"payloads"`
}

type coreReloadModulesReq struct {
	_msgpack struct{} `msgpack:",asArray"`
	Method   string
	Token    string
}

type coreReloadModulesRes struct {
	Exploits  uint32 `msgpack:"exploits"`
	Auxiliary uint32 `msgpack:"auxiliary"`
	Post      uint32 `msgpack:"post"`
	Encoders  uint32 `msgpack:"encoders"`
	Nops      uint32 `msgpack:"nops"`
	Payloads  uint32 `msgpack:"payloads"`
}

type coreSaveReq struct {
	_msgpack struct{} `msgpack:",asArray"`
	Method   string
	Token    string
}

type coreSaveRes struct {
	Result string `msgpack:"result"`
}

type coreSetgReq struct {
	_msgpack    struct{} `msgpack:",asArray"`
	Method      string
	Token       string
	OptionName  string
	OptionValue string
}

type coreSetgRes struct {
	Result string `msgpack:"result"`
}

type coreUnSetgReq struct {
	_msgpack   struct{} `msgpack:",asArray"`
	Method     string
	Token      string
	OptionName string
}

type coreUnSetgRes struct {
	Result string `msgpack:"result"`
}

type coreThreadListReq struct {
	_msgpack struct{} `msgpack:",asArray"`
	Method   string
	Token    string
}

type coreThreadListRes map[int]struct {
	Status   string `msgpack:"status"`
	Critical bool   `msgpack:"critical"`
	Name     string `msgpack:"name"`
	Started  string `msgpack:"started"`
}

type coreThreadKillReq struct {
	_msgpack struct{} `msgpack:",asArray"`
	Method   string
	Token    string
	ThreadId string
}

type coreThreadKillRes struct {
	Result string `msgpack:"result"`
}

type coreVersionReq struct {
	_msgpack struct{} `msgpack:",asArray"`
	Method   string
	Token    string
}

type coreVersionRes struct {
	Version string `msgpack:"version"`
	Ruby    string `msgpack:"ruby"`
	Api     string `msgpack:"api"`
}

type coreStopReq struct {
	_msgpack struct{} `msgpack:",asArray"`
	Method   string
	Token    string
}

type coreStopRes struct {
	Result string `msgpack:"result"`
}

func (msf *Metasploit) CoreAddModulePath(path string) (coreAddModulePathRes, error) {
	ctx := &coreAddModulePathReq{
		Method: "core.add_module_path",
		Token:  msf.token,
		Path:   path,
	}

	var res coreAddModulePathRes
	if err := msf.send(ctx, &res); err != nil {
		return coreAddModulePathRes{}, err
	}
	return res, nil
}

func (msf *Metasploit) CoreModuleStats() (coreModuleStatsRes, error) {
	ctx := &coreModuleStatsReq{
		Method: "core.module_stats",
		Token:  msf.token,
	}

	var res coreModuleStatsRes
	if err := msf.send(ctx, &res); err != nil {
		return coreModuleStatsRes{}, err
	}
	return res, nil
}

func (msf *Metasploit) CoreReloadModules() (coreReloadModulesRes, error) {
	ctx := &coreReloadModulesReq{
		Method: "core.reload_modules",
		Token:  msf.token,
	}

	var res coreReloadModulesRes
	if err := msf.send(ctx, &res); err != nil {
		return coreReloadModulesRes{}, err
	}
	return res, nil
}

func (msf *Metasploit) CoreSave() (coreSaveRes, error) {
	ctx := &coreSaveReq{
		Method: "core.save",
		Token:  msf.token,
	}

	var res coreSaveRes
	if err := msf.send(ctx, &res); err != nil {
		return coreSaveRes{}, nil
	}
	return res, nil
}

func (msf *Metasploit) CoreSetg(optionName, optionValue string) (coreSetgRes, error) {
	ctx := &coreSetgReq{
		Method:      "core.setg",
		Token:       msf.token,
		OptionName:  optionName,
		OptionValue: optionValue,
	}

	var res coreSetgRes
	if err := msf.send(ctx, &res); err != nil {
		return coreSetgRes{}, nil
	}
	return res, nil
}

func (msf *Metasploit) CoreUnSetg(optionName string) (coreUnSetgRes, error) {
	ctx := &coreUnSetgReq{
		Method:     "core.unsetg",
		Token:      msf.token,
		OptionName: optionName,
	}

	var res coreUnSetgRes
	if err := msf.send(ctx, &res); err != nil {
		return coreUnSetgRes{}, nil
	}
	return res, nil
}

func (msf *Metasploit) CoreThreadList() (coreThreadListRes, error) {
	ctx := &coreThreadListReq{
		Method: "core.thread_list",
		Token:  msf.token,
	}

	var res coreThreadListRes
	if err := msf.send(ctx, &res); err != nil {
		return coreThreadListRes{}, nil
	}
	return res, nil
}

func (msf *Metasploit) CoreThreadKill(threadId string) (coreThreadKillRes, error) {
	ctx := &coreThreadKillReq{
		Method:   "core.thread_kill",
		Token:    msf.token,
		ThreadId: threadId,
	}

	var res coreThreadKillRes
	if err := msf.send(ctx, &res); err != nil {
		return coreThreadKillRes{}, nil
	}
	return res, nil
}

func (msf *Metasploit) CoreVersion() (coreVersionRes, error) {
	ctx := &coreVersionReq{
		Method: "core.version",
		Token:  msf.token,
	}
	var res coreVersionRes
	if err := msf.send(ctx, &res); err != nil {
		return coreVersionRes{}, err
	}
	return res, nil
}

func (msf *Metasploit) CoreStop() (coreStopRes, error) {
	ctx := &coreStopReq{
		Method: "core.stop",
		Token:  msf.token,
	}
	var res coreStopRes
	if err := msf.send(ctx, &res); err != nil {
		return coreStopRes{}, err
	}
	return res, nil
}
