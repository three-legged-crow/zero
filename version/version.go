// three-legged-crow.zero
//
// @(#)version.go  April 8th, 2022
// Copyright(c) 2022, Leyton Goth. All rights reserved.

// Package version 用于打印服务版本号
package version

import (
	"os"
	"strings"

	"github.com/three-legged-crow/zero/strings0/table"
)

// Print 打印版本号, 当输入的参数为 --version, -V 时, 会直接退出程序.
// 版本信息如下:
//
//	+-----------------+-------------------------+
//	| Binary          | Version                 |
//	+-----------------+-------------------------+
//	| mainland-server | leyton@2022.04.09.00.07 |
//	+-----------------+-------------------------+
func Print(v string) {
	slashIndex := 0
	if i := strings.LastIndex(os.Args[0], "/"); i > 0 {
		slashIndex = i
	}
	s := os.Args[0][slashIndex+1:]

	tab := table.New(3)
	tab.AddHeader("Binary", "Node IP", "Version")
	tab.AppendRow(s, os.Getenv("POD_IP"), v)
	tab.Print()

	// 识别为打印版本意图, 直接退出程序
	if len(os.Args) > 1 && (os.Args[1] == "--version" || os.Args[1] == "-V") {
		os.Exit(0)
	}
}
