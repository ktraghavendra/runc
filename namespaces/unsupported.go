// +build !linux

package namespaces

import (
	"github.com/docker/libcontainer/cgroups"
)

func Exec(container *libcontainer.Container, term Terminal, rootfs, dataPath string, args []string, createCommand CreateCommand, startCallback func()) (int, error) {
	return -1, ErrUnsupported
}

func Init(container *libcontainer.Container, uncleanRootfs, consolePath string, syncPipe *SyncPipe, args []string) error {
	return ErrUnsupported
}

func InitializeNetworking(container *libcontainer.Container, nspid int, pipe *SyncPipe) error {
	return ErrUnsupported
}

func SetupCgroups(container *libcontainer.Container, nspid int) (cgroups.ActiveCgroup, error) {
	return nil, ErrUnsupported
}

func GetNamespaceFlags(namespaces map[string]bool) (flag int) {
	return 0
}
