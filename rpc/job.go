package rpc

// Jobs

type jobListReq struct {
	_msgpack struct{} `msgpack:",asArray"`
	Method   string
	Token    string
}

type jobListRes map[string]string

type jobInfoReq struct {
	_msgpack struct{} `msgpack:",asArray"`
	Method   string
	Token    string
	JobId    string
}

type jobInfoRes struct {
	Jid       int                    `msgpack:"jid"`
	Name      string                 `msgpack:"name"`
	StartTime int                    `msgpack:"start_time"`
	UriPath   interface{}            `msgpack:"uripath,omitempty"`
	Datastore map[string]interface{} `msgpack:"datastore,omitempty"`
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
		Method: "job.list",
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
		Method: "job.info",
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
		Method: "job.stop",
		Token:  msf.token,
		JobId:  jobId,
	}
	var res jobStopRes
	if err := msf.send(ctx, &res); err != nil {
		return jobStopRes{}, err
	}
	return res, nil
}
