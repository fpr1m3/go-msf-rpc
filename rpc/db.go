package rpc

//import (
//	"time"
//)

// DB objects

type dbAddWorkspaceReq struct {
	_msgpack      struct{} `msgpack:",asArray"`
	Method        string
	Token         string
	WorkspaceName string
}

type dbAddWorkspaceRes struct {
	Result string `msgpack:"result"`
}

type dbAnalyzeHostReq struct {
	_msgpack struct{} `msgpack:",asArray"`
	Method   string
	Token    string
	Options  map[string]string
}

type dbAnalyzeHostRes struct {
	Host []struct {
		Address string `msgpack:"address"`
		Modules []map[string]struct {
			MType string `msgpack:"mtype"`
			MName string `msgpack:"mname"`
		} `msgpack:"modules"`
	} `msgpack:"host"`
}

type dbClientsReq struct {
	_msgpack struct{} `msgpack:",asArray"`
	Method   string
	Token    string
	Options  map[string]interface{}
}

type dbClientsRes struct {
	Clients []struct {
		Host      string `msgpack:"host"`
		UaString  string `msgpack:"ua_string"`
		UaName    string `msgpack:"ua_name"`
		UaVersion string `msgpack:"ua_ver"`
		CreatedAt uint32 `msgpack:"created_at"`
		UpdatedAt uint32 `msgpack:"updated_at"`
	} `msgpack:"clients"`
}

type dbConnectReq struct {
	_msgpack struct{} `msgpack:",asArray"`
	Method   string
	Token    string
	Options  map[string]interface{}
}

type dbConnectRes struct {
	Result string `msgpack:"result"`
}

type dbCreateCrackedCredentialReq struct {
	_msgpack struct{} `msgpack:",asArray"`
	Method   string
	Token    string
	Options  map[string]interface{}
}

type dbCreateCrackedCredentialRes map[string]interface{}

type dbCreateCredentialReq struct {
	_msgpack struct{} `msgpack:",asArray"`
	Method   string
	Token    string
	Options  map[string]interface{}
}

type dbCreateCredentialRes struct {
	Username    string `msgpack:"username"`
	Private     string `msgpack:"private"`
	PrivateType string `msgpack:"private_type"`
	RealmValue  string `msgpack:"realm_value"`
	RealmKey    string `msgpack:"realm_key"`
	Host        string `msgpack:"host"`
	ServiceName string `msgpack:"sname"`
	Status      string `msgpack:"status"`
}

type dbCredsReq struct {
	_msgpack      struct{} `msgpack:",asArray"`
	Method        string
	Token         string
	WorkspaceName map[string]interface{}
}

type dbCredsRes struct {
	Creds []struct {
		User        string `msgpack:"user"`
		Pass        string `msgpack:"pass"`
		UpdatedAt   uint32 `msgpack:"updated_at"`
		Type        string `msgpack:"type"`
		Host        string `msgpack:"host"`
		Port        uint32 `msgpack:"port"`
		Proto       string `msgpack:"proto"`
		ServiceName string `msgpack:"sname"`
	} `msgpack:"creds"`
}

type dbCurrentWorkspaceReq struct {
	_msgpack struct{} `msgpack:",asArray"`
	Method   string
	Token    string
}

type dbCurrentWorkspaceRes struct {
	Workspace   string `msgpack:"workspace"`
	WorkspaceId uint32 `msgpack:"workspace_id"`
}

type dbDeleteClientReq struct {
	_msgpack struct{} `msgpack:",asArray"`
	Method   string
	Token    string
	Options  map[string]interface{}
}

type dbDeleteClientRes struct {
	Result  string `msgpack:"result"`
	Deleted []map[string]struct {
		Address  string `msgpack:"address"`
		UAString string `msgpack:"ua_string"`
	} `msgpack:"deleted"`
}

type dbDeleteCredsReq struct {
	_msgpack struct{} `msgpack:",asArray"`
	Method   string
	Token    string
	Options  map[string]interface{}
}

type dbDeleteCredsRes struct {
	Result  string `msgpack:"result"`
	Deleted []struct {
		Creds []struct {
			Username    string `msgpack:"user"`
			Password    string `msgpack:"pass"`
			UpdatedAt   uint32 `msgpack:"updated_at"`
			Type        string `msgpack:"type"`
			Host        string `msgpack:"host"`
			Port        uint32 `msgpack:"port"`
			Protocol    string `msgpack:"proto"`
			ServiceName string `msgpack:"sname"`
		} `msgpack:"creds"`
	} `msgpack:"deleted"`
}

type dbDeleteHostReq struct {
	_msgpack struct{} `msgpack:",asArray"`
	Method   string
	Token    string
	Options  map[string]interface{}
}

type dbDeleteHostRes struct {
	Result  string   `msgpack:"result"`
	Deleted []string `msgpack:"deleted"`
}

type dbDeleteNoteReq struct {
	_msgpack struct{} `msgpack:",asArray"`
	Method   string
	Token    string
	Options  map[string]interface{}
}

type dbDeleteNoteRes struct {
	Result  string `msgpack:"result"`
	Deleted []map[string]struct {
		Address  string `msgpack:"address"`
		Port     uint32 `msgpack:"port"`
		Protocol string `msgpack:"proto"`
		NoteType string `msgpack:"ntype"`
	} `msgpack:"deleted"`
}

//del_service

//del_vuln

type dbDeleteWorkspaceReq struct {
	_msgpack      struct{} `msgpack:",asArray"`
	Method        string
	Token         string
	WorkspaceName string
}

type dbDeleteWorkspaceRes struct {
	Result string `msgpack:"result"`
}

//disconnect

//driver

type dbReportClientReq struct {
	_msgpack struct{} `msgpack:",asArray"`
	Method   string
	Token    string
	Options  map[string]interface{}
}

type dbReportClientRes struct {
	Result string `msgpack:"result"`
}

type dbReportEventReq struct {
	_msgpack struct{} `msgpack:",asArray"`
	Method   string
	Token    string
	Options  map[string]interface{}
}

type dbReportEventRes struct {
	Result string `msgpack:"result"`
}

type dbReportHostReq struct {
	_msgpack struct{} `msgpack:",asArray"`
	Method   string
	Token    string
	Options  map[string]interface{}
}

type dbReportHostRes struct {
	Result string `msgpack:"result"`
}

type dbEventsReq struct {
	_msgpack struct{} `msgpack:",asArray"`
	Method   string
	Token    string
	Options  map[string]interface{}
}

type dbEventsRes struct {
	Events []struct {
		Host      interface{}       `msgpack:"host"`
		CreatedAt uint32            `msgpack:"created_at"`
		UpdatedAt uint32            `msgpack:"updated_at"`
		Name      string            `msgpack:"name"`
		Critical  interface{}       `msgpack:"critical"`
		Username  interface{}       `msgpack:"username"`
		Info      map[string]string `msgpack:"info"`
	} `msgpack:"events"`
}

type dbStatusReq struct {
	_msgpack struct{} `msgpack:",asArray"`
	Method   string
	Token    string
}

type dbStatusRes map[string]string

type dbWorkspacesReq struct {
	_msgpack struct{} `msgpack:",asArray"`
	Method   string
	Token    string
}

type dbWorkspacesRes struct {
	Workspaces []struct {
		Id        uint32 `msgpack:"id"`
		Name      string `msgpack:"name"`
		CreatedAt int64  `msgpack:"created_at"`
		UpdatedAt int64  `msgpack:"updated_at"`
	} `msgpack:"workspaces"`
}

// DB Funcs

func (msf *Metasploit) DBAddWorkspace(workspaceName string) (dbAddWorkspaceRes, error) {
	ctx := &dbAddWorkspaceReq{
		Method:        "db.add_workspace",
		Token:         msf.token,
		WorkspaceName: workspaceName,
	}
	var res dbAddWorkspaceRes
	if err := msf.send(ctx, &res); err != nil {
		return dbAddWorkspaceRes{}, err
	}
	return res, nil
}

// Need data to test.
func (msf *Metasploit) DBAnalyzeHost(workspace, host string) (dbAnalyzeHostRes, error) {
	ctx := &dbAnalyzeHostReq{
		Method:  "db.analyze_host",
		Token:   msf.token,
		Options: map[string]string{"workspace": workspace, "host": host},
	}
	var res dbAnalyzeHostRes
	if err := msf.send(ctx, &res); err != nil {
		return dbAnalyzeHostRes{}, err
	}
	return res, nil
}

// Need data to test.
func (msf *Metasploit) DBClients(options map[string]interface{}) (dbClientsRes, error) {
	ctx := &dbClientsReq{
		Method:  "db.clients",
		Token:   msf.token,
		Options: options,
	}
	var res dbClientsRes
	if err := msf.send(ctx, &res); err != nil {
		return dbClientsRes{}, err
	}
	return res, nil
}

func (msf *Metasploit) DBConnect(options map[string]interface{}) (dbConnectRes, error) {
	ctx := &dbConnectReq{
		Method:  "db.connect",
		Token:   msf.token,
		Options: options,
	}
	var res dbConnectRes
	if err := msf.send(ctx, &res); err != nil {
		return dbConnectRes{}, err
	}
	return res, nil
}

func (msf *Metasploit) DBCreateCrackedCredential(options map[string]interface{}) (dbCreateCrackedCredentialRes, error) {
	ctx := &dbCreateCrackedCredentialReq{
		Method:  "db.create_cracked_credential",
		Token:   msf.token,
		Options: options,
	}
	var res dbCreateCrackedCredentialRes
	if err := msf.send(ctx, &res); err != nil {
		return dbCreateCrackedCredentialRes{}, err
	}
	return res, nil
}

func (msf *Metasploit) DBCreateCredential(options map[string]interface{}) (dbCreateCredentialRes, error) {
	ctx := &dbCreateCredentialReq{
		Method:  "db.create_credential",
		Token:   msf.token,
		Options: options,
	}
	var res dbCreateCredentialRes
	if err := msf.send(ctx, &res); err != nil {
		return dbCreateCredentialRes{}, err
	}
	return res, nil
}

func (msf *Metasploit) DBCreds(workspace map[string]interface{}) (dbCredsRes, error) {
	if len(workspace) == 0 {
		workspace = map[string]interface{}{"workspace": "default"}
	}
	ctx := &dbCredsReq{
		Method:        "db.creds",
		Token:         msf.token,
		WorkspaceName: workspace,
	}
	var res dbCredsRes
	if err := msf.send(ctx, &res); err != nil {
		return dbCredsRes{}, err
	}
	return res, nil
}

func (msf *Metasploit) DBCurrentWorkspace() (dbCurrentWorkspaceRes, error) {
	ctx := &dbStatusReq{
		Method: "db.current_workspace",
		Token:  msf.token,
	}
	var res dbCurrentWorkspaceRes
	if err := msf.send(ctx, &res); err != nil {
		return dbCurrentWorkspaceRes{}, err
	}
	return res, nil
}

func (msf *Metasploit) DBDeleteClient(options map[string]interface{}) (dbDeleteClientRes, error) {
	ctx := &dbDeleteClientReq{
		Method:  "db.del_client",
		Token:   msf.token,
		Options: options,
	}
	var res dbDeleteClientRes
	if err := msf.send(ctx, &res); err != nil {
		return dbDeleteClientRes{}, err
	}
	return res, nil
}

func (msf *Metasploit) DBDeleteCreds(options map[string]interface{}) (dbDeleteCredsRes, error) {
	ctx := &dbDeleteCredsReq{
		Method:  "db.del_creds",
		Token:   msf.token,
		Options: options,
	}
	var res dbDeleteCredsRes
	if err := msf.send(ctx, &res); err != nil {
		return dbDeleteCredsRes{}, err
	}
	return res, nil
}

func (msf *Metasploit) DBDeleteHost(options map[string]interface{}) (dbDeleteHostRes, error) {
	ctx := &dbDeleteHostReq{
		Method:  "db.del_host",
		Token:   msf.token,
		Options: options,
	}
	var res dbDeleteHostRes
	if err := msf.send(ctx, &res); err != nil {
		return dbDeleteHostRes{}, err
	}
	return res, nil
}

func (msf *Metasploit) DBDeleteNote(options map[string]interface{}) (dbDeleteNoteRes, error) {
	ctx := &dbDeleteNoteReq{
		Method:  "db.del_note",
		Token:   msf.token,
		Options: options,
	}
	var res dbDeleteNoteRes
	if err := msf.send(ctx, &res); err != nil {
		return dbDeleteNoteRes{}, err
	}
	return res, nil
}

func (msf *Metasploit) DBDeleteWorkspace(workspaceName string) (dbDeleteWorkspaceRes, error) {
	ctx := &dbDeleteWorkspaceReq{
		Method:        "db.del_workspace",
		Token:         msf.token,
		WorkspaceName: workspaceName,
	}
	var res dbDeleteWorkspaceRes
	if err := msf.send(ctx, &res); err != nil {
		return dbDeleteWorkspaceRes{}, err
	}
	return res, nil
}

func (msf *Metasploit) DBReportClient(options map[string]interface{}) (dbReportClientRes, error) {
	ctx := &dbReportClientReq{
		Method:  "db.report_client",
		Token:   msf.token,
		Options: options,
	}
	var res dbReportClientRes
	if err := msf.send(ctx, &res); err != nil {
		return dbReportClientRes{}, err
	}
	return res, nil
}

func (msf *Metasploit) DBReportEvent(options map[string]interface{}) (dbReportEventRes, error) {
	ctx := &dbReportEventReq{
		Method:  "db.report_event",
		Token:   msf.token,
		Options: options,
	}
	var res dbReportEventRes
	if err := msf.send(ctx, &res); err != nil {
		return dbReportEventRes{}, err
	}
	return res, nil
}

func (msf *Metasploit) DBReportHost(options map[string]interface{}) (dbReportHostRes, error) {
	ctx := &dbReportHostReq{
		Method:  "db.report_host",
		Token:   msf.token,
		Options: options,
	}
	var res dbReportHostRes
	if err := msf.send(ctx, &res); err != nil {
		return dbReportHostRes{}, err
	}
	return res, nil
}

func (msf *Metasploit) DBEvents(options map[string]interface{}) (dbEventsRes, error) {
	ctx := &dbEventsReq{
		Method:  "db.events",
		Token:   msf.token,
		Options: options,
	}
	var res dbEventsRes
	if err := msf.send(ctx, &res); err != nil {
		return dbEventsRes{}, err
	}
	return res, nil
}

func (msf *Metasploit) DBStatus() (dbStatusRes, error) {
	ctx := &dbStatusReq{
		Method: "db.status",
		Token:  msf.token,
	}
	var res dbStatusRes
	if err := msf.send(ctx, &res); err != nil {
		return dbStatusRes{}, err
	}
	return res, nil
}

func (msf *Metasploit) DBWorkspaces() (dbWorkspacesRes, error) {
	ctx := &dbWorkspacesReq{
		Method: "db.workspaces",
		Token:  msf.token,
	}
	var res dbWorkspacesRes
	if err := msf.send(ctx, &res); err != nil {
		return dbWorkspacesRes{}, err
	}
	return res, nil
}
