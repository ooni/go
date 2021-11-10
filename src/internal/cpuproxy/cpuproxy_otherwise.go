// +build !android !arm64

package cpuproxy

import "internal/cpu"

// HasAES returns whether the CPU supports AES.
func HasAES() bool {
	return cpu.X86.HasAES || cpu.ARM64.HasAES
}

// HasGFMUL returns whether the CPU supports GFMUL.
func HasGFMUL() bool {
	return cpu.X86.HasPCLMULQDQ || cpu.ARM64.HasPMULL
}
