package term

import (
	"syscall"
	"unsafe"

	"golang.org/x/sys/unix"
)

func EnableRawMode(fd uintptr) (*syscall.Termios, error) {
	// Get current terminal settings
	origTermios := &syscall.Termios{}
	_, _, errno := syscall.Syscall6(syscall.SYS_IOCTL, fd, ioctlReadTermios(), uintptr(unsafe.Pointer(origTermios)), 0, 0, 0)
	if errno != 0 {
		return nil, errno
	}

	// Make a copy to modify
	raw := *origTermios

	raw.Lflag &^= syscall.ICANON | syscall.ECHO | syscall.ECHOE | syscall.ISIG

	// Apply raw settings
	_, _, errno = syscall.Syscall6(syscall.SYS_IOCTL, fd, ioctlWriteTermios(), uintptr(unsafe.Pointer(&raw)), 0, 0, 0)
	if errno != 0 {
		return nil, errno
	}

	return origTermios, nil
}

func DisableRawMode(fd uintptr, origTermios *syscall.Termios) error {
	_, _, errno := syscall.Syscall6(syscall.SYS_IOCTL, fd, ioctlWriteTermios(), uintptr(unsafe.Pointer(origTermios)), 0, 0, 0)
	if errno != 0 {
		return errno
	}
	return nil
}

func ioctlReadTermios() uintptr {
	// works only on darwin
	return syscall.TIOCGETA
}

func ioctlWriteTermios() uintptr {
	// works only on darwin
	return syscall.TIOCSETA
}

func GetSize(fd uintptr) (uint16, uint16) {
	// todo handle error
	uws, _ := unix.IoctlGetWinsize(int(fd), unix.TIOCGWINSZ)
	return uws.Row, uws.Col
}
