package schema

import "path/filepath"

// Config represent the AuroraBoot
// configuration
type Config struct {
	// CloudConfig to use for generating installation mediums
	CloudConfig string `yaml:"cloud_config"`

	// Disable Netboot
	DisableNetboot bool `yaml:"disable_netboot"`

	// Disable HTTP Server
	DisableHTTPServer bool `yaml:"disable_http_server"`

	// Disable manual ISO boot
	DisableISOboot bool `yaml:"disable_iso"`

	// PixieCore HTTPServer Port
	NetBootHTTPPort string `yaml:"netboot_http_port"`

	// PixieCore Listen addr
	NetBootListenAddr string `yaml:"netboot_listen_addr"`

	State string `yaml:"state_dir"`

	ListenAddr string `yaml:"listen_addr"`

	// ISO block configuration
	ISO ISO `yaml:"iso"`

	// Netboot block configuration
	NetBoot NetBoot `yaml:"netboot"`

	Disk Disk `yaml:"disk"`
}

type Disk struct {
	RAW bool `yaml:"raw"`
	GCE bool `yaml:"gce"`
	VHD bool `yaml:"vhd"`

	ARM *ARMDiskOptions `yaml:"arm"`
}

type NetBoot struct {
	Cmdline string `yaml:"cmdline"`
}

type ISO struct {
	DataPath string `yaml:"data"`
}

func (c Config) StateDir(s ...string) string {
	d := "/tmp"
	if c.State != "" {
		d = c.State
	}

	return filepath.Join(append([]string{d}, s...)...)
}

type ARMDiskOptions struct {
	Model       string     `yaml:"model"`
	LVM         bool       `yaml:"lvm"`
	DiskSize    SizeOption `yaml:"size"`
	EFIOverlay  string     `yaml:"efi_overlay_dir"`
	PrepareOnly bool       `yaml:"prepare_only"`
}

type SizeOption struct {
	Disk              string `yaml:"size"`
	StatePartition    string `yaml:"state_partition"`
	Images            string `yaml:"images"`
	RecoveryPartition string `yaml:"recovery_partition"`
}
