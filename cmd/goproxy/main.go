
package main

import (
"flag"
"github.com/timchenxiaoyu/goproxy"
)

func main() {

	// flag.String("log_dir", "", "If non-empty, write log files in this directory")
	addr := flag.String("http", ":8080", "proxy listen addr")
	flag.Parse()
	server := goproxy.NewServer(*addr)
	server.Start()
}
