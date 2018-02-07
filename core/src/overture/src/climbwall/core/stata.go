//  +build DRCLO

package core

import (
	"bytes"
	"climbwall/utils"
	"encoding/binary"
	"net"
	"time"
)

var last time.Time

func StatPackage(s uint64, r uint64) {

	t := time.Now()

	elapsed := t.Sub(last)

	if elapsed.Nanoseconds() > 1000000000 {

		sendimp(10, 10)
		last = t
	}
}

func sendimp(s uint64, r uint64) {
	conn, err := net.Dial("unix", "stat_path")
	if err != nil {
		utils.Logger.Println("create statistics socket failed！", err)
		return
	}

	defer conn.Close()

	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, s)
	bufh := new(bytes.Buffer)
	binary.Write(bufh, binary.LittleEndian, r)
	out := append(buf.Bytes(), bufh.Bytes()...)

	conn.Write(out)
	utils.Logger.Println("statistics send ok!!")
}
