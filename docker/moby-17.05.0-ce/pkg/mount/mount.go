package mount

import "fmt"

// GetMounts retrieves a list of mounts for the current running process.
func GetMounts() ([]*Info, error) {
	return parseMountTable()
}

// Mounted determines if a specified mountpoint has been mounted.
// On Linux it looks at /proc/self/mountinfo and on Solaris at mnttab.
func Mounted(mountpoint string) (bool, error) {
	entries, err := parseMountTable()
	if err != nil {
	
		return false, err
	}
    fmt.Printf("yang test 3xcddddddddddddd vmounted\n")
	// Search the table for the mountpoint
	for _, e := range entries {
		if e.Mountpoint == mountpoint {
		    fmt.Printf("yang test mounted\n")
			return true, nil
		}
	}
	return false, nil
}

// Mount will mount filesystem according to the specified configuration, on the
// condition that the target path is *not* already mounted. Options must be
// specified like the mount or fstab unix commands: "opt1=val1,opt2=val2". See
// flags.go for supported option flags.
func Mount(device, target, mType, options string) error {
    
	flag, _ := parseOptions(options)
	if flag&REMOUNT != REMOUNT {
		if mounted, err := Mounted(target); err != nil || mounted {
			return err
		}
	}
	return ForceMount(device, target, mType, options)
}

// ForceMount will mount a filesystem according to the specified configuration,
// *regardless* if the target path is not already mounted. Options must be
// specified like the mount or fstab unix commands: "opt1=val1,opt2=val2". See
// flags.go for supported option flags.
func ForceMount(device, target, mType, options string) error {
	flag, data := parseOptions(options)
	return mount(device, target, mType, uintptr(flag), data)
}

// Unmount lazily unmounts a filesystem on supported platforms, otherwise
// does a normal unmount.
func Unmount(target string) error {

	if mounted, err := Mounted(target); err != nil || !mounted {
		return err
	}

	return unmount(target, mntDetach)
}

func ForceUnmount(target string) error {
	return unmount(target, mntDetach)
}


