// Command program for all apps in solution
// Author: Thiep Wong
// Email: thiep.wong@gmail.com
package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"

	"github.com/go-ini/ini"
)

func main() {
	var showVersion bool
	var runSmartId bool

	cfg, err := ini.Load("conf/app.conf")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}
	fmt.Println("Bat dau chay cmd")
	//args := []string{"what", "ever", "you", "like"}
	//_a := []string{"name", "thep", "age", "345"}
	cmd := exec.Command("app/smartid/smartid", "-s", "8", "-n", "2")
	cc := cmd.Start()
	if cc != nil {
		fmt.Printf("Fail to read file: %v", cc)
	}

	fmt.Println("Da  chay cmd", cmd.Process.Pid)

	// Classic read of values, default section can be represented as empty string
	fmt.Println("App Mode:", cfg.Section("").Key("app_mode").String())
	fmt.Println("Data Path:", cfg.Section("paths").Key("data").String())

	flag.BoolVar(&showVersion, "version", false, "Print version information.")
	flag.BoolVar(&showVersion, "v", false, "Print version information.")

	flag.BoolVar(&runSmartId, "s", false, "Start smart id module")

	flag.Parse()
	// Show version and exit
	if showVersion {
		fmt.Println("Version 1.0.0")
		os.Exit(0)
	}

}
