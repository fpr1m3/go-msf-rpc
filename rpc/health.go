package rpc

type healthCheckReq struct {
	_msgpack struct{} `msgpack:",asArray"`
	Method   string
}

type healthCheckRes struct {
	Status string `msgpack:"status"`
}

func (msf *Metasploit) HealthCheck() (healthCheckRes, error) {
	ctx := &healthCheckReq{
		Method: "health.check",
	}
	var res healthCheckRes
	if err := msf.send(ctx, &res); err != nil {
		return healthCheckRes{}, err
	}
	return res, nil
}
