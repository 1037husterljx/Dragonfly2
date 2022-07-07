package sn

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os/exec"
	"strings"
	"time"

	logger "d7y.io/dragonfly/v2/internal/dflog"
)

var (
	serialUrl = "http://100.100.100.200/latest/meta-data/serial-number"
)

var SN string

func init() {
	SN, _ = GetSN()
	logger.Infof("use sn :%s", SN)
}

func GetSN() (string, error) {
	var err error
	var sn string

	if sn, err = getSNFromHostinfo(); err == nil {
		return sn, err
	}
	logger.Infof("get SN from hostinfo failed: %v", err)

	if sn, err = getSNFromFile(); err == nil {
		return sn, err
	}
	logger.Infof("get SN from file failed: %v", err)

	if sn, err = getSNFromRemote(); err == nil {
		return sn, err
	}
	logger.Infof("get SN from remote failed: %v", err)

	if sn, err = getSNFromHardware(); err == nil {
		return sn, err
	}
	logger.Infof("get SN from hardware failed: %v", err)

	logger.Warnf("can not get SN from anywhere")
	return "", fmt.Errorf("get SN error")
}

// getSNFromHostinfo  从hostinfo命令获取sn
func getSNFromHostinfo() (string, error) {
	cmd := exec.Command("hostinfo")
	if info, err := cmd.CombinedOutput(); err != nil {
		return "", err
	} else {
		for _, line := range strings.Split(string(info), "\n") {
			s := strings.TrimSpace(line)
			if strings.HasPrefix(s, "sn") {
				return strings.TrimSpace(s[2:]), nil
			}
		}
	}
	return "", fmt.Errorf("can not find sn in hostinfo program")
}

// getSNFromFile 从/usr/sbin/staragent_sn文件获取sn
func getSNFromFile() (string, error) {
	// making it very simple. No need to close the file.
	if content, err := ioutil.ReadFile("/usr/sbin/staragent_sn"); err != nil {
		return "", err
	} else {
		return strings.TrimSpace(string(content)), nil
	}
}

// getSNFromRemote 云上vpc机器通过访问 http://100.100.100.200/latest/meta-data/serial-number 获取到本机sn
func getSNFromRemote() (string, error) {
	client := http.Client{Timeout: time.Second * 1}
	resp, err := client.Get("http://100.100.100.200/latest/meta-data/serial-number")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if body, err := ioutil.ReadAll(resp.Body); err != nil {
		return "", err
	} else {
		if resp.StatusCode == http.StatusOK {
			return strings.TrimSpace(string(body)), nil
		} else {
			return "", fmt.Errorf(string(body))
		}
	}
}

// getSNFromHardware refer: https://yuque.antfin-inc.com/staragent/help/gr7mw7
func getSNFromHardware() (string, error) {
	cmd := exec.Command("/usr/sbin/dmidecode -s system-serial-number")
	if info, err := cmd.CombinedOutput(); err != nil {
		return "", err
	} else {
		return strings.TrimSpace(string(info)), nil
	}
}
