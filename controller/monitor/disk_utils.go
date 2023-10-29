package monitor

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
	"golang.org/x/sys/unix"

	grpccodes "google.golang.org/grpc/codes"
	grpcstatus "google.golang.org/grpc/status"

	iscsiutil "github.com/longhorn/go-iscsi-helper/util"

	"github.com/longhorn/longhorn-manager/engineapi"
	"github.com/longhorn/longhorn-manager/util"

	longhorn "github.com/longhorn/longhorn-manager/k8s/pkg/apis/longhorn/v1beta2"
)

const (
	defaultBlockSize = 4096
)

// GetDiskStat returns the disk stat of the given directory
func getDiskStat(diskType longhorn.DiskType, name, path string, client *engineapi.DiskService) (stat *util.DiskStat, err error) {
	switch diskType {
	case longhorn.DiskTypeFilesystem:
		return getFilesystemTypeDiskStat(path)
	case longhorn.DiskTypeBlock:
		return getBlockTypeDiskStat(client, name, path)
	default:
		return nil, fmt.Errorf("unknown disk type %v", diskType)
	}
}

func getFilesystemTypeDiskStat(path string) (stat *util.DiskStat, err error) {
	return util.GetDiskStat(path)
}

func getBlockTypeDiskStat(client *engineapi.DiskService, name, path string) (stat *util.DiskStat, err error) {
	if client == nil {
		return nil, errors.New("disk service client is nil")
	}

	info, err := client.DiskGet(string(longhorn.DiskTypeBlock), name, path)
	if err != nil {
		return nil, err
	}
	return &util.DiskStat{
		DiskID:           info.ID,
		Path:             info.Path,
		Type:             info.Type,
		TotalBlocks:      info.TotalBlocks,
		FreeBlocks:       info.FreeBlocks,
		BlockSize:        info.BlockSize,
		StorageMaximum:   info.TotalSize,
		StorageAvailable: info.FreeSize,
	}, nil
}

// GetDiskConfig returns the disk config of the given directory
func getDiskConfig(diskType longhorn.DiskType, name, path string, client *engineapi.DiskService) (*util.DiskConfig, error) {
	switch diskType {
	case longhorn.DiskTypeFilesystem:
		return getFilesystemTypeDiskConfig(path)
	case longhorn.DiskTypeBlock:
		return getBlockTypeDiskConfig(client, name, path)
	default:
		return nil, fmt.Errorf("unknown disk type %v", diskType)
	}
}

func getFilesystemTypeDiskConfig(path string) (*util.DiskConfig, error) {
	nsPath := iscsiutil.GetHostNamespacePath(util.HostProcPath)
	filePath := filepath.Join(path, util.DiskConfigFile)

	fmt.Println("getFilesystemTypeDiskConfig", path)

	output, err := iscsiutil.ForkAndSwitchToNamespace(nsPath, func() (*[]byte, error) {
		out, err := os.ReadFile(filePath)
		return &out, err
	})
	fmt.Println("getFilesystemTypeDiskConfig output", path, output, err)
	if err != nil {
		return nil, fmt.Errorf("cannot find config file %v on host: %v", filePath, err)
	}

	cfg := &util.DiskConfig{}
	if err := json.Unmarshal(*output, cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal %v content %v on host: %v", filePath, output, err)
	}
	return cfg, nil
}

func getBlockTypeDiskConfig(client *engineapi.DiskService, name, path string) (config *util.DiskConfig, err error) {
	if client == nil {
		return nil, errors.New("disk service client is nil")
	}

	info, err := client.DiskGet(string(longhorn.DiskTypeBlock), name, path)
	if err != nil {
		if grpcstatus.Code(err) == grpccodes.NotFound {
			return nil, errors.Wrapf(err, "cannot find disk info")
		}
		return nil, err
	}
	return &util.DiskConfig{
		DiskUUID: info.UUID,
	}, nil
}

// GenerateDiskConfig generates a disk config for the given directory
func generateDiskConfig(diskType longhorn.DiskType, name, uuid, path string, client *engineapi.DiskService) (*util.DiskConfig, error) {
	switch diskType {
	case longhorn.DiskTypeFilesystem:
		return generateFilesystemTypeDiskConfig(path)
	case longhorn.DiskTypeBlock:
		return generateBlockTypeDiskConfig(client, name, uuid, path, defaultBlockSize)
	default:
		return nil, fmt.Errorf("unknown disk type %v", diskType)
	}
}

func generateFilesystemTypeDiskConfig(path string) (*util.DiskConfig, error) {
	cfg := &util.DiskConfig{
		DiskUUID: util.UUID(),
	}
	encoded, err := json.Marshal(cfg)
	if err != nil {
		return nil, fmt.Errorf("BUG: Cannot marshal %+v: %v", cfg, err)
	}

	nsPath := iscsiutil.GetHostNamespacePath(util.HostProcPath)
	filePath := filepath.Join(path, util.DiskConfigFile)

	_, err = iscsiutil.ForkAndSwitchToNamespace(nsPath, func() (*interface{}, error) {
		_, err := os.Stat(filePath)
		return nil, err
	})
	if err == nil {
		return nil, fmt.Errorf("disk cfg on %v exists, cannot override", filePath)
	}

	defer func() {
		if err != nil {
			if derr := util.DeleteDiskPathReplicaSubdirectoryAndDiskCfgFile(nsPath, path); derr != nil {
				err = errors.Wrapf(err, "cleaning up disk config path %v failed with error: %v", path, derr)
			}

		}
	}()

	_, err = iscsiutil.ForkAndSwitchToNamespace(nsPath, func() (*interface{}, error) {
		return nil, os.WriteFile(filePath, encoded, 0644)
	})
	if err != nil {
		return nil, fmt.Errorf("cannot write to disk cfg on %v: %v", filePath, err)
	}
	if err := util.CreateDiskPathReplicaSubdirectory(path); err != nil {
		return nil, err
	}
	_, err = iscsiutil.ForkAndSwitchToNamespace(nsPath, func() (*interface{}, error) {
		unix.Sync()
		return nil, nil
	})

	return cfg, nil
}

func generateBlockTypeDiskConfig(client *engineapi.DiskService, name, uuid, path string, blockSize int64) (*util.DiskConfig, error) {
	if client == nil {
		return nil, errors.New("disk service client is nil")
	}

	info, err := client.DiskCreate(string(longhorn.DiskTypeBlock), name, uuid, path, blockSize)
	if err != nil {
		return nil, err
	}
	return &util.DiskConfig{
		DiskUUID: info.UUID,
	}, nil
}

// DeleteDisk deletes the disk with the given name and uuid
func DeleteDisk(diskType longhorn.DiskType, diskName, diskUUID string, client *engineapi.DiskService) error {
	if client == nil {
		return errors.New("disk service client is nil")
	}

	return client.DiskDelete(string(diskType), diskName, diskUUID)
}

// getSpdkReplicaInstanceNames returns the replica lvol names of the given disk
func getSpdkReplicaInstanceNames(client *engineapi.DiskService, diskType, diskName string) (map[string]string, error) {
	if client == nil {
		return nil, errors.New("disk service client is nil")
	}

	instances, err := client.DiskReplicaInstanceList(diskType, diskName)
	if err != nil {
		return nil, err
	}

	instanceNames := map[string]string{}
	for name := range instances {
		instanceNames[name] = name
	}

	return instanceNames, nil
}
