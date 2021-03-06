package main

import (
	"flag"
	"fmt"
	"os"

	module "github.com/baidu/openedge/module"
	"github.com/baidu/openedge/module/logger"
)

func main() {
	f, err := module.ParseFlags("")
	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to parse argument:", err.Error())
		return
	}
	if f.Help {
		flag.Usage()
		return
	}

	m, err := New(f.Config)
	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to create module:", err.Error())
		logger.Log.WithError(err).Errorf("failed to create module")
		return
	}
	defer m.Close()
	err = m.Start()
	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to start module:", err.Error())
		logger.Log.WithError(err).Errorf("failed to start module")
		return
	}
	module.Wait()
}
