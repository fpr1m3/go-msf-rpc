package rpc

// Modules

type moduleExploitsReq struct {
	_msgpack struct{} `msgpack:",asArray"`
	Method   string
	Token    string
}

type moduleExploitsRes struct {
	Modules []string `msgpack:"modules"`
}

type moduleAuxiliaryReq struct {
	_msgpack struct{} `msgpack:",asArray"`
	Method   string
	Token    string
}

type moduleAuxiliaryRes struct {
	Modules []string `msgpack:"modules"`
}

type modulePostReq struct {
	_msgpack struct{} `msgpack:",asArray"`
	Method   string
	Token    string
}

type modulePostRes struct {
	Modules []string `msgpack:"modules"`
}

type modulePayloadsReq struct {
	_msgpack struct{} `msgpack:",asArray"`
	Method   string
	Token    string
}

type modulePayloadsRes struct {
	Modules []string `msgpack:"modules"`
}

type moduleEncodersReq struct {
	_msgpack struct{} `msgpack:",asArray"`
	Method   string
	Token    string
}

type moduleEncodersRes struct {
	Modules []string `msgpack:"modules"`
}

type moduleNopsReq struct {
	_msgpack struct{} `msgpack:",asArray"`
	Method   string
	Token    string
}

type moduleNopsRes struct {
	Modules []string `msgpack:"modules"`
}

type moduleInfoReq struct {
	_msgpack   struct{} `msgpack:",asArray"`
	Method     string
	Token      string
	ModuleType string
	ModuleName string
}

type moduleInfoRes struct {
	Name        string     `msgpack:"name"`
	Description string     `msgpack:"description"`
	License     string     `msgpack:"license"`
	FilePath    string     `msgpack:"filepath"`
	Version     string     `msgpack:"version"`
	Rank        string     `msgpack:"rank"`
	References  [][]string `msgpack:"references"`
	Authors     []string   `msgpack:"authors"`
}

type moduleOptionsReq struct {
	_msgpack   struct{} `msgpack:",asArray"`
	Method     string
	Token      string
	ModuleType string
	ModuleName string
}

type moduleOptionsRes map[string]struct {
	Type     string      `msgpack:"type"`
	Required bool        `msgpack:"required"`
	Advanced bool        `msgpack:"advanced"`
	Evasion  bool        `msgpack:"evasion"`
	Desc     string      `msgpack:"desc"`
	Default  interface{} `msgpack:"default"`
	Enums    []string    `msgpack:"enums,omitempty"`
}

type moduleCompatiblePayloadsReq struct {
	_msgpack   struct{} `msgpack:",asArray"`
	Method     string
	Token      string
	ModuleName string
}

type moduleCompatiblePayloadsRes struct {
	Payloads []string `msgpack:"payloads"`
}

type moduleTargetCompatiblePayloadsReq struct {
	_msgpack   struct{} `msgpack:",asArray"`
	Method     string
	Token      string
	ModuleName string
	ArchNumber uint32
}

type moduleTargetCompatiblePayloadsRes struct {
	Payloads []string `msgpack:"payloads"`
}

type moduleCompatibleSessionsReq struct {
	_msgpack   struct{} `msgpack:",asArray"`
	Method     string
	Token      string
	ModuleName string
}

type moduleCompatibleSessionsRes struct {
	Sessions []string `msgpack:"sessions"`
}

type moduleEncodeReq struct {
	_msgpack      struct{} `msgpack:",asArray"`
	Method        string
	Token         string
	Data          string
	EncoderModule string
	Options       map[string]string
}

type moduleEncodeRes struct {
	Encoded []byte `msgpack:"encoded"`
}

type moduleExecuteReq struct {
	_msgpack   struct{} `msgpack:",asArray"`
	Method     string
	Token      string
	ModuleType string
	ModuleName string
	Options    map[string]string
}

type moduleExecuteRes struct {
	JobId uint32 `msgpack:"job_id"`
}

func (msf *Metasploit) ModuleExploits() (moduleExploitsRes, error) {
	ctx := &moduleExploitsReq{
		Method: "module.exploits",
		Token:  msf.token,
	}
	var res moduleExploitsRes
	if err := msf.send(ctx, &res); err != nil {
		return moduleExploitsRes{}, err
	}
	return res, nil
}

func (msf *Metasploit) ModuleAuxiliary() (moduleAuxiliaryRes, error) {
	ctx := &moduleAuxiliaryReq{
		Method: "module.auxiliary",
		Token:  msf.token,
	}
	var res moduleAuxiliaryRes
	if err := msf.send(ctx, &res); err != nil {
		return moduleAuxiliaryRes{}, err
	}
	return res, nil
}

func (msf *Metasploit) ModulePost() (modulePostRes, error) {
	ctx := &modulePostReq{
		Method: "module.post",
		Token:  msf.token,
	}
	var res modulePostRes
	if err := msf.send(ctx, &res); err != nil {
		return modulePostRes{}, err
	}
	return res, nil
}

func (msf *Metasploit) ModulePayloads() (modulePayloadsRes, error) {
	ctx := &modulePayloadsReq{
		Method: "module.payloads",
		Token:  msf.token,
	}
	var res modulePayloadsRes
	if err := msf.send(ctx, &res); err != nil {
		return modulePayloadsRes{}, err
	}
	return res, nil
}

func (msf *Metasploit) ModuleEncoders() (moduleEncodersRes, error) {
	ctx := &moduleEncodersReq{
		Method: "module.encoders",
		Token:  msf.token,
	}
	var res moduleEncodersRes
	if err := msf.send(ctx, &res); err != nil {
		return moduleEncodersRes{}, err
	}
	return res, nil
}

func (msf *Metasploit) ModuleNops() (moduleNopsRes, error) {
	ctx := &moduleNopsReq{
		Method: "module.nops",
		Token:  msf.token,
	}
	var res moduleNopsRes
	if err := msf.send(ctx, &res); err != nil {
		return moduleNopsRes{}, err
	}
	return res, nil
}

func (msf *Metasploit) ModuleInfo(moduleType, moduleName string) (moduleInfoRes, error) {
	ctx := &moduleInfoReq{
		Method:     "module.info",
		Token:      msf.token,
		ModuleType: moduleType,
		ModuleName: moduleName,
	}
	var res moduleInfoRes
	if err := msf.send(ctx, &res); err != nil {
		return moduleInfoRes{}, err
	}
	return res, nil
}

func (msf *Metasploit) ModuleOptions(moduleType, moduleName string) (moduleOptionsRes, error) {
	ctx := &moduleOptionsReq{
		Method:     "module.options",
		Token:      msf.token,
		ModuleType: moduleType,
		ModuleName: moduleName,
	}
	var res moduleOptionsRes
	if err := msf.send(ctx, &res); err != nil {
		return moduleOptionsRes{}, err
	}
	return res, nil
}

func (msf *Metasploit) ModuleCompatiblePayloads(moduleName string) (moduleCompatiblePayloadsRes, error) {
	ctx := &moduleCompatiblePayloadsReq{
		Method:     "module.compatible_payloads",
		Token:      msf.token,
		ModuleName: moduleName,
	}
	var res moduleCompatiblePayloadsRes
	if err := msf.send(ctx, &res); err != nil {
		return moduleCompatiblePayloadsRes{}, err
	}
	return res, nil
}

func (msf *Metasploit) ModuleTargetCompatiblePayloads(moduleName string, targetNumber uint32) (moduleTargetCompatiblePayloadsRes, error) {
	ctx := &moduleTargetCompatiblePayloadsReq{
		Method:     "module.target_compatible_payloads",
		Token:      msf.token,
		ModuleName: moduleName,
		ArchNumber: targetNumber,
	}
	var res moduleTargetCompatiblePayloadsRes
	if err := msf.send(ctx, &res); err != nil {
		return moduleTargetCompatiblePayloadsRes{}, err
	}
	return res, nil
}

func (msf *Metasploit) ModuleCompatibleSessions(moduleName string) (moduleCompatibleSessionsRes, error) {
	ctx := &moduleCompatibleSessionsReq{
		Method:     "module.compatible_sessions",
		Token:      msf.token,
		ModuleName: moduleName,
	}
	var res moduleCompatibleSessionsRes
	if err := msf.send(ctx, &res); err != nil {
		return moduleCompatibleSessionsRes{}, err
	}
	return res, nil
}

func (msf *Metasploit) ModuleEncode(data, encoderModule string, moduleOptions map[string]string) (moduleEncodeRes, error) {
	ctx := &moduleEncodeReq{
		Method:        "module.encode",
		Token:         msf.token,
		Data:          data,
		EncoderModule: encoderModule,
		Options:       moduleOptions,
	}
	var res moduleEncodeRes
	if err := msf.send(ctx, &res); err != nil {
		return moduleEncodeRes{}, err
	}
	return res, nil
}

func (msf *Metasploit) ModuleExecute(moduleType, moduleName string, moduleOptions map[string]string) (moduleExecuteRes, error) {
	ctx := &moduleExecuteReq{
		Method:     "module.execute",
		Token:      msf.token,
		ModuleType: moduleType,
		ModuleName: moduleName,
		Options:    moduleOptions,
	}
	var res moduleExecuteRes
	if err := msf.send(ctx, &res); err != nil {
		return moduleExecuteRes{}, err
	}
	return res, nil
}

func (msf *Metasploit) GetModuleRequires(moduleType, moduleName string) ([]string, error) {
	var returnValues []string

	options, err := msf.ModuleOptions(moduleType, moduleName)

	if err != nil {
		return nil, err
	}

	for key, option := range options {
		if option.Required {
			returnValues = append(returnValues, key)
		}
	}
	return returnValues, nil
}
