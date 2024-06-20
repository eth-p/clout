//go:build windows
// +build windows

package color

import (
	"os"
	"syscall"
)

// Enable tells the operating system to enable ANSI color processing.
// This is a no-op on every platform except Windows.
//
// ANSI Color codes in conhost.exe (the default terminal emulator) requires Windows 10 and an explicit call to
// Kernel32.dll!SetConsoleMode with the ENABLE_VIRTUAL_TERMINAL_INPUT flag.
//
// https://docs.microsoft.com/en-us/windows/console/setconsolemode
// https://stackoverflow.com/a/69542231
func Enable() {
	stdout := syscall.Handle(os.Stdout.Fd())

	var originalMode uint32
	syscall.GetConsoleMode(stdout, &originalMode)
	originalMode |= 0x0004

	syscall.
		MustLoadDLL("kernel32").
		MustFindProc("SetConsoleMode").
		Call(uintptr(stdout), uintptr(originalMode))
}
