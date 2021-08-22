package rpc

import "time"

// Jobs

type jobListReq struct {
	_msgpack struct{} `msgpack:",asArray"`
	Method   string
	Token    string
}

type jobListRes struct {
	Id string `msgpack:",omitempty"`
}

type jobInfoReq struct {
	_msgpack struct{} `msgpack:",asArray"`
	Method   string
	Token    string
	JobId    string
}

type jobInfoRes struct {
	Jid       uint32    `msgpack:"jid"`
	Name      string    `msgpack:"name"`
	StartTime time.Time `msgpack:"start_time"`
	UriPath   string    `msgpack:"uripath"`
	Datastore struct {
		EnableContextEncoding bool   `msgpack:"EnableContextEncoding"`
		DisablePayloadHandler bool   `msgpack:"DisblePayloadHandler"`
		Ssl                   bool   `msgpack:"SSL"`
		SslVersion            string `msgpack:"SSLVersion"`
		SrvHost               string `msgpack:"SRVHOST"`
		SrvPort               string `msgpack:"SRVPORT"`
		Payload               string `msgpack:"PAYLOAD"`
		Lhost                 string `msgpack:"LHOST"`
		Lport                 string `msgpack:"LPORT"`
	} `msgpack:"datastore"`
}

type jobStopReq struct {
	_msgpack struct{} `msgpack:",asArray"`
	Method   string
	Token    string
	JobId    string
}

type jobStopRes struct {
	Result string `msgpack:"result"`
}

// Jobs

func (msf *Metasploit) JobList() (jobListRes, error) {
	ctx := &jobListReq{
		Method: "console.tabs",
		Token:  msf.token,
	}
	var res jobListRes
	if err := msf.send(ctx, &res); err != nil {
		return jobListRes{}, err
	}
	return res, nil
}

func (msf *Metasploit) JobInfo(jobId string) (jobInfoRes, error) {
	ctx := &jobInfoReq{
		Method: "console.tabs",
		Token:  msf.token,
		JobId:  jobId,
	}
	var res jobInfoRes
	if err := msf.send(ctx, &res); err != nil {
		return jobInfoRes{}, err
	}
	return res, nil
}

func (msf *Metasploit) JobStop(jobId string) (jobStopRes, error) {
	ctx := &jobStopReq{
		Method: "console.tabs",
		Token:  msf.token,
		JobId:  jobId,
	}
	var res jobStopRes
	if err := msf.send(ctx, &res); err != nil {
		return jobStopRes{}, err
	}
	return res, nil
}
