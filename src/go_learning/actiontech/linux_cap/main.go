package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	builtin_syscall "syscall"
	"unsafe"
)

var command string
var arg string
var keepCaps bool

const (
	LINUX_CAPABILITY_VERSION_3 = 0x20080522
	LINUX_CAPABILITY_U32S_3    = 2
)

type Capability uint

func (c Capability) String() string {
	switch c {
	case CAP_CHOWN:
		return "CAP_CHOWN"
	case CAP_DAC_OVERRIDE:
		return "CAP_DAC_OVERRIDE"
	case CAP_DAC_READ_SEARCH:
		return "CAP_DAC_READ_SEARCH"
	case CAP_FOWNER:
		return "CAP_FOWNER"
	case CAP_FSETID:
		return "CAP_FSETID"
	case CAP_KILL:
		return "CAP_KILL"
	case CAP_SETGID:
		return "CAP_SETGID"
	case CAP_SETUID:
		return "CAP_SETUID"
	case CAP_SETPCAP:
		return "CAP_SETPCAP"
	case CAP_LINUX_IMMUTABLE:
		return "CAP_LINUX_IMMUTABLE"
	case CAP_NET_BIND_SERVICE:
		return "CAP_NET_BIND_SERVICE"
	case CAP_NET_BROADCAST:
		return "CAP_NET_BROADCAST"
	case CAP_NET_ADMIN:
		return "CAP_NET_ADMIN"
	case CAP_NET_RAW:
		return "CAP_NET_RAW"
	case CAP_IPC_LOCK:
		return "CAP_IPC_LOCK"
	case CAP_IPC_OWNER:
		return "CAP_IPC_OWNER"
	case CAP_SYS_MODULE:
		return "CAP_SYS_MODULE"
	case CAP_SYS_RAWIO:
		return "CAP_SYS_RAWIO"
	case CAP_SYS_CHROOT:
		return "CAP_SYS_CHROOT"
	case CAP_SYS_PTRACE:
		return "CAP_SYS_PTRACE"
	case CAP_SYS_PACCT:
		return "CAP_SYS_PACCT"
	case CAP_SYS_ADMIN:
		return "CAP_SYS_ADMIN"
	case CAP_SYS_BOOT:
		return "CAP_SYS_BOOT"
	case CAP_SYS_NICE:
		return "CAP_SYS_NICE"
	case CAP_SYS_RESOURCE:
		return "CAP_SYS_RESOURCE"
	case CAP_SYS_TIME:
		return "CAP_SYS_TIME"
	case CAP_SYS_TTY_CONFIG:
		return "CAP_SYS_TTY_CONFIG"
	case CAP_MKNOD:
		return "CAP_MKNOD"
	case CAP_LEASE:
		return "CAP_LEASE"
	case CAP_AUDIT_WRITE:
		return "CAP_AUDIT_WRITE"
	case CAP_AUDIT_CONTROL:
		return "CAP_AUDIT_CONTROL"
	case CAP_SETFCAP:
		return "CAP_SETFCAP"
	case CAP_MAC_OVERRIDE:
		return "CAP_MAC_OVERRIDE"
	case CAP_MAC_ADMIN:
		return "CAP_MAC_ADMIN"
	case CAP_SYSLOG:
		return "CAP_SYSLOG"
	case CAP_WAKE_ALARM:
		return "CAP_WAKE_ALARM"
	case CAP_BLOCK_SUSPEND:
		return "CAP_BLOCK_SUSPEND"
	default:
		return ""
	}
}

const (
	CAP_CHOWN Capability = iota
	CAP_DAC_OVERRIDE
	CAP_DAC_READ_SEARCH
	CAP_FOWNER
	CAP_FSETID
	CAP_KILL
	CAP_SETGID
	CAP_SETUID
	CAP_SETPCAP
	CAP_LINUX_IMMUTABLE
	CAP_NET_BIND_SERVICE
	CAP_NET_BROADCAST
	CAP_NET_ADMIN
	CAP_NET_RAW
	CAP_IPC_LOCK
	CAP_IPC_OWNER
	CAP_SYS_MODULE
	CAP_SYS_RAWIO
	CAP_SYS_CHROOT
	CAP_SYS_PTRACE
	CAP_SYS_PACCT
	CAP_SYS_ADMIN
	CAP_SYS_BOOT
	CAP_SYS_NICE
	CAP_SYS_RESOURCE
	CAP_SYS_TIME
	CAP_SYS_TTY_CONFIG
	CAP_MKNOD
	CAP_LEASE
	CAP_AUDIT_WRITE
	CAP_AUDIT_CONTROL
	CAP_SETFCAP
	CAP_MAC_OVERRIDE
	CAP_MAC_ADMIN
	CAP_SYSLOG
	CAP_WAKE_ALARM
	CAP_BLOCK_SUSPEND
)

type CapHeader struct {
	Version uint32
	Pid     int32
}

type CapData struct {
	Effective   uint32
	Permitted   uint32
	Inheritable uint32
}

type Cap struct {
	Header CapHeader
	Data   [LINUX_CAPABILITY_U32S_3]CapData
}

func (c *Cap) EncodeEffectiveCaps() uint32 {
	return c.Data[0].Effective
}

func (c *Cap) EncodePermittedCaps() uint32 {
	return c.Data[0].Permitted
}

func (c *Cap) EncodeInheritableCaps() uint32 {
	return c.Data[0].Inheritable
}

func (c *Cap) EffectiveCaps() []Capability {
	return DecodeCaps(c.EncodeEffectiveCaps())
}

func (c *Cap) PermittedCaps() []Capability {
	return DecodeCaps(c.EncodePermittedCaps())
}

func (c *Cap) InheritableCaps() []Capability {
	return DecodeCaps(c.EncodeInheritableCaps())
}

// DecodeCaps decode mask to Capabilities
func DecodeCaps(mask uint32) []Capability {
	data := []Capability{}
	for i := CAP_CHOWN; i <= CAP_BLOCK_SUSPEND; i++ {
		if mask&(1<<i) > 0 {
			data = append(data, i)
		}
	}
	return data
}

// EncodeCaps encode Capabilities to mask
func EncodeCaps(caps []Capability) uint32 {
	var mask uint32
	for _, cap := range caps {
		mask |= (1 << cap)
	}
	return mask
}

// gets the capabilities from process pid
func GetCap(pid int) (*Cap, error) {
	caps := new(Cap)
	caps.Header.Version = LINUX_CAPABILITY_VERSION_3
	caps.Header.Pid = int32(pid)
	_, _, e1 := builtin_syscall.Syscall(builtin_syscall.SYS_CAPGET,
		uintptr(unsafe.Pointer(&caps.Header)),
		uintptr(unsafe.Pointer(&caps.Data[0])),
		0)
	if e1 != 0 {
		return nil, e1
	}
	return caps, nil
}
func main() {
	flag.StringVar(&command, "c", "ls", "command to exec")
	flag.BoolVar(&keepCaps, "k", true, "keep caps")
	flag.Parse()
	//cmd := exec.Command(command, arg)
	cmd := exec.Command("bash", "--noprofile", "--norc", "-c", command)

	if keepCaps {
		fmt.Println("setting caps...")
		caps, err := GetCap(os.Getpid())
		if err != nil {
			fmt.Printf("ERR: %v\n", err.Error())
		}
		uintCaps := make([]uintptr, 0)
		for _, cap := range caps.Data {
			for _, capInt := range DecodeCaps(cap.Effective) {
				uintCaps = append(uintCaps, uintptr(capInt))
			}
		}
		cmd.SysProcAttr = &builtin_syscall.SysProcAttr{
			AmbientCaps: uintCaps,
		}
	}

	bs, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("ERR: %v\n", err.Error())
	}
	fmt.Println(string(bs))
}
