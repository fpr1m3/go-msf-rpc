package rpc

// Plugins

type pluginLoadReq struct {
	_msgpack   struct{} `msgpack:",asArray"`
	Method     string
	Token      string
	PluginName string
	Options    map[string]string
}

type pluginLoadRes struct {
	Result string `msgpack:"result"`
}

type pluginUnLoadReq struct {
	_msgpack   struct{} `msgpack:",asArray"`
	Method     string
	Token      string
	PluginName string
}

type pluginUnLoadRes struct {
	Result string `msgpack:"result"`
}

type pluginLoadedReq struct {
	_msgpack struct{} `msgpack:",asArray"`
	Method   string
	Token    string
}

type pluginLoadedRes struct {
	Plugins []string `msgpack:"plugins"`
}

func (msf *Metasploit) PluginLoad(pluginName string, pluginOptions map[string]string) (pluginLoadRes, error) {
	ctx := &pluginLoadReq{
		Method:     "plugin.load",
		Token:      msf.token,
		PluginName: pluginName,
		Options:    pluginOptions,
	}
	var res pluginLoadRes
	if err := msf.send(ctx, &res); err != nil {
		return pluginLoadRes{}, err
	}
	return res, nil
}

func (msf *Metasploit) PluginUnLoad(pluginName string) (pluginUnLoadRes, error) {
	ctx := &pluginUnLoadReq{
		Method:     "plugin.unload",
		Token:      msf.token,
		PluginName: pluginName,
	}
	var res pluginUnLoadRes
	if err := msf.send(ctx, &res); err != nil {
		return pluginUnLoadRes{}, err
	}
	return res, nil
}

func (msf *Metasploit) PluginLoaded() (pluginLoadedRes, error) {
	ctx := &pluginLoadedReq{
		Method: "plugin.loaded",
		Token:  msf.token,
	}
	var res pluginLoadedRes
	if err := msf.send(ctx, &res); err != nil {
		return pluginLoadedRes{}, err
	}
	return res, nil
}
