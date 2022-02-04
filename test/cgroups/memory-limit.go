package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"strconv"
	"syscall"
)

// 挂载了 memory subsystem hierarchy 的根目录位置
const cgroupMemoryHierarchyMount = "/sys/fs/cgroup/memory"

func main() {
	fmt.Println("os args:", os.Args)
	if os.Args[0] == "/proc/self/exe" {
		// 容器进程
		fmt.Printf("current pid:%v\n", syscall.Getpid())
		cmd := exec.Command("sh", "-c", "stress --vm-bytes 200m --vm-keep -m 1")
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		checkErr(err, "/proc/self/exe run")
	}
	cmd := exec.Command("/proc/self/exe")
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Start()
	checkErr(err, "/proc/self/exe start")
	// 得到 fork 出来进程映射在外部命名空间的 pid
	fmt.Printf("%v\n", cmd.Process.Pid)
	// 在系统默认创建挂载了 memory subsystem Hierarchy 上创建 cgroup
	err = os.Mkdir(path.Join(cgroupMemoryHierarchyMount, "testmemorylimit"), 0755)
	checkErr(err, "Mkdir")
	// 将容器进程加入到这个 cgroup
	err = ioutil.WriteFile(path.Join(cgroupMemoryHierarchyMount, "testmemorylimit", "tasks"), []byte(strconv.Itoa(cmd.Process.Pid)), 0644)
	checkErr(err, "WriteFile tasks")
	//g限制 cgroup 进程使用
	err = ioutil.WriteFile(path.Join(cgroupMemoryHierarchyMount, "testmemorylimit", "memory.limit_in_bytes"), []byte("100m"), 0644)
	checkErr(err, "WriteFile limit_in_bytes")

	_, err = cmd.Process.Wait()
	checkErr(err, "cmd.Process.Wait")
}

func checkErr(err error, reason string) {
	if err != nil {
		panic(fmt.Sprintf("err:%v, reason:%s", err, reason))
	}
}
