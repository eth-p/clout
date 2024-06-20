//go:build windows && !windows_no_enable_vti
// +build windows,!windows_no_enable_vti

package color

func init() {
	Enable()
}
