package cmd

import (
	"bufio"
	"bytes"
	"log"
	"os/exec"
)

/*
author：willem
description：执行命令并阻塞打印log
date：2019-12-12 11:23
*/
func Exec(command string, params []string) {
	cmd := exec.Command(command, params...)

	log.Println("exec => ", cmd.String())

	stderr, _ := cmd.StderrPipe()
	stdout, _ := cmd.StdoutPipe()

	defer stderr.Close()
	defer stdout.Close()

	if err := cmd.Start(); err != nil {
		panic("started error => " + cmd.String())
	}

	logScan := bufio.NewScanner(stdout)
	go func() {
		for logScan.Scan() {
			log.Println(logScan.Text())
		}
	}()

	errBuf := bytes.NewBufferString("")
	scan := bufio.NewScanner(stderr)
	go func() {
		for scan.Scan() {
			s := scan.Text()
			errBuf.WriteString(s + "\n")
		}
	}()

	cmd.Wait()
	if !cmd.ProcessState.Success() {
		log.Fatalln("exec error => " + cmd.ProcessState.String() + ",message:" + errBuf.String())
	}
}

/*
author：willem
description：执行命令并立即返回
date：2020-05-27 19:38
*/
func DirectExec(command string, params []string) *exec.Cmd {
	cmd := exec.Command(command, params...)

	log.Println("direct exec => ", cmd.String())

	if err := cmd.Start(); err != nil {
		log.Fatalln("direct exec error => " + err.Error())
	}
	return cmd
}
