// +build linux

package main

import (
	"log"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	cmd := exec.Command("/bin/bash")

	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | // hostname -b zxytest
			syscall.CLONE_NEWIPC | // 1. ipcs 查看所有的消息队列信息  2. ipcmk -Q 新增消息队列e
			syscall.CLONE_NEWPID | // echo $$
			syscall.CLONE_NEWNS | // mount -t tmpfs tmpfs /tmp/testtemp (宿主机挂着)
			syscall.CLONE_NEWUSER | // readlink /proc/$$/ns/user
			syscall.CLONE_NEWNET, // ifconfig 网络为空
		// CLONE_NEWUSER 需要指定uid和gid
		UidMappings: []syscall.SysProcIDMap{
			{
				ContainerID: 0,
				HostID:      0,
				Size:        1,
			},
		},
		GidMappings: []syscall.SysProcIDMap{
			{
				ContainerID: 0,
				HostID:      0,
				Size:        1,
			},
		},
	}

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}
