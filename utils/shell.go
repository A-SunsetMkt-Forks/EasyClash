package utils

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"syscall"
	"time"

	"github.com/pkg/errors"
)

func Exec(bin string, arg ...string) (_stdout string, err error) {
	cmd := exec.Command(bin, arg...)
	var stdin, stdout, stderr bytes.Buffer
	cmd.Stdin = &stdin
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err = cmd.Run()
	outStr, errStr := string(stdout.Bytes()), string(stderr.Bytes())
	if err != nil {
		return outStr, errors.Wrap(err, errStr)
	}
	return outStr, nil
}

func SaveFile(path string, content string) error {
	f, _ := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	defer f.Close()
	_, err := f.Write([]byte(content))
	if err != nil {
		return err
	}
	return nil
}

func GetFileContent(path string) (string, error) {
	_, err := os.Stat(path)
	if err != nil {
		return "", err
	}
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	text := string(content)
	return text, nil
}

func Mkdir(dir string) error {
	_, err := os.Stat(dir)
	exist := true
	if err != nil {
		if os.IsNotExist(err) {
			exist = false
		} else {
			return errors.Wrap(err, "SaveDir.Stat")
		}
	}
	if exist {
		return nil
	}
	err = os.Mkdir(dir, os.ModePerm)
	if err != nil {
		return errors.Wrap(err, "Mkdir err")
	}
	return nil
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func downloadFile(url, path string) error {
	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// 创建一个文件用于保存
	fmt.Println("1111", path)
	out, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return err
	}
	defer out.Close()

	// 然后将响应流和文件流对接起来
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}
	return nil
}

func PidExists(pid int) (bool, error) {
	if pid <= 0 {
		return false, fmt.Errorf("invalid pid %v", pid)
	}
	proc, err := os.FindProcess(pid)
	if err != nil {
		return false, err
	}
	err = proc.Signal(syscall.Signal(0))
	if err == nil {
		return true, nil
	}
	if err.Error() == "os: process already finished" {
		return false, nil
	}
	errno, ok := err.(syscall.Errno)
	if !ok {
		return false, err
	}
	switch errno {
	case syscall.ESRCH:
		return false, nil
	case syscall.EPERM:
		return true, nil
	}
	return false, err
}

func KillProcess(pid int) bool {
	if pid == 0 {
		return false
	}
	killErr := syscall.Kill(pid, syscall.SIGKILL)
	return killErr == nil
}

func Fork(name string, arg ...string) (*os.Process, error) {
	cmd := exec.Command(name, arg...)
	var out bytes.Buffer
	cmd.Env = os.Environ()
	cmd.Stdin = nil
	cmd.Stdout = &out
	cmd.Stderr = nil
	cmd.ExtraFiles = nil
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Setsid: true,
	}
	if err := cmd.Start(); err != nil {
		return nil, err
	}
	count := 0
	for {
		if count > 200 {
			break
		}
		if cmd.Process != nil {
			return cmd.Process, nil
		}

		time.Sleep(time.Second)
		fmt.Println(count)
		count++
	}
	return nil, errors.New("fork process timeout")
}

func runInWindows(cmd string) (string, error) {
	result, err := exec.Command("cmd", "/c", cmd).Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(result)), err
}

func RunCommand(cmd string) (string, error) {
	if runtime.GOOS == "windows" {
		return runInWindows(cmd)
	} else {
		return runInLinux(cmd)
	}
}

func runInLinux(cmd string) (string, error) {
	fmt.Println("Running Linux cmd:" + cmd)
	result, err := exec.Command("/bin/sh", "-c", cmd).Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(result)), err
}

//根据进程名判断进程是否运行
func CheckProRunning(serverName string) (bool, error) {
	a := `ps ux | awk '/` + serverName + `/ && !/awk/ {print $2}'`
	pid, err := RunCommand(a)
	if err != nil {
		return false, err
	}
	return pid != "", nil
}

//根据进程名称获取进程ID
func GetPid(serverName string) (string, error) {
	a := `ps ux | awk '/` + serverName + `/ && !/awk/ {print $2}'`
	pid, err := RunCommand(a)
	return pid, err
}
