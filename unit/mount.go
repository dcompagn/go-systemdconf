// DO NOT EDIT. This file is generated from systemd 244 by generatesdconf

package unit

import "github.com/sergeymakinen/go-systemdconf"

// MountSection represents information about the file system mount points it supervises
type MountSection struct {
	systemdconf.Section
	ExecOptions
	KillOptions
	ResourceControlOptions

	// Takes an absolute path of a device node, file or other resource to mount. See mount for details. If this refers to a device
	// node, a dependency on the respective device unit is automatically created. (See systemd.device for more information.)
	// This option is mandatory. Note that the usual specifier expansion is applied to this setting, literal percent characters
	// should hence be written as "%%".
	What systemdconf.Value

	// Takes an absolute path of a directory for the mount point; in particular, the destination cannot be a symbolic link. If the
	// mount point does not exist at the time of mounting, it is created. This string must be reflected in the unit filename. (See
	// above.) This option is mandatory.
	Where systemdconf.Value

	// Takes a string for the file system type. See mount for details. This setting is optional.
	Type systemdconf.Value

	// Mount options to use when mounting. This takes a comma-separated list of options. This setting is optional. Note that the
	// usual specifier expansion is applied to this setting, literal percent characters should hence be written as "%%".
	Options systemdconf.Value

	// Takes a boolean argument. If true, parsing of the options specified in Options= is relaxed, and unknown mount options are
	// tolerated. This corresponds with mount's -s switch. Defaults to off.
	SloppyOptions systemdconf.Value

	// Takes a boolean argument. If true, detach the filesystem from the filesystem hierarchy at time of the unmount operation,
	// and clean up all references to the filesystem as soon as they are not busy anymore. This corresponds with umount's -l switch.
	// Defaults to off.
	LazyUnmount systemdconf.Value

	// Takes a boolean argument. If true, force an unmount (in case of an unreachable NFS system). This corresponds with umount's
	// -f switch. Defaults to off.
	ForceUnmount systemdconf.Value

	// Directories of mount points (and any parent directories) are automatically created if needed. This option specifies
	// the file system access mode used when creating these directories. Takes an access mode in octal notation. Defaults to 0755.
	DirectoryMode systemdconf.Value

	// Configures the time to wait for the mount command to finish. If a command does not exit within the configured time, the mount
	// will be considered failed and be shut down again. All commands still running will be terminated forcibly via SIGTERM, and
	// after another delay of this time with SIGKILL. (See KillMode= in systemd.kill.) Takes a unit-less value in seconds, or
	// a time span value such as "5min 20s". Pass 0 to disable the timeout logic. The default value is set from DefaultTimeoutStartSec=
	// option in systemd-system.conf.
	TimeoutSec systemdconf.Value
}

// MountFile represents information about a file system mount point controlled and supervised by systemd
type MountFile struct {
	systemdconf.File

	Unit    UnitSection    // Generic information about the unit that is not dependent on the type of unit
	Mount   MountSection   // Information about the file system mount points it supervises
	Install InstallSection // Installation information for the unit
}
