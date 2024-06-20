//go:build !windows
// +build !windows

package color

// Enable tells the operating system to enable ANSI color processing.
// This is a no-op on every platform except Windows.
func Enable() {
}
