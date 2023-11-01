package util

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/pkg/errors"
	"golang.org/x/sys/unix"

	iscsiutil "github.com/longhorn/go-iscsi-helper/util"
)

const (
	OsReleasePath = "/etc/os-release"
)

// GetHostKernelRelease retrieves the kernel release version of the host.
func GetHostKernelRelease() (string, error) {
	initiatorNSPath := iscsiutil.GetHostNamespacePath(HostProcPath)

	output, err := iscsiutil.ForkAndSwitchToNamespace(initiatorNSPath, time.Second*5, func() (*string, error) {
		var uname unix.Utsname
		err := unix.Uname(&uname)
		if err != nil {
			return nil, err
		}
		nulIndex := string(bytes.TrimRight(uname.Release[:], "\x00"))
		return &nulIndex, err
	})
	if err != nil {
		return "", err
	}
	return RemoveNewlines(*output), nil
}

// GetHostOSDistro retrieves the operating system distribution of the host.
func GetHostOSDistro() (string, error) {
	initiatorNSPath := iscsiutil.GetHostNamespacePath(HostProcPath)

	output, err := iscsiutil.ForkAndSwitchToNamespace(initiatorNSPath, time.Minute, func() (*string, error) {
		out, err := os.ReadFile(OsReleasePath)
		outStr := string(out)
		return &outStr, err
	})

	if err != nil {
		return "", errors.Wrapf(err, "failed to read %v on host", OsReleasePath)
	}

	scanner := bufio.NewScanner(strings.NewReader(*output))
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "ID=") {
			osDistro := RemoveNewlines(strings.TrimPrefix(line, "ID="))
			return strings.ReplaceAll(osDistro, `"`, ""), nil
		}
	}
	return "", fmt.Errorf("failed to find ID field in %v", OsReleasePath)
}
