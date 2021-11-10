// +build android
// +build arm64

// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
// This file is based on the diff available at
// https://go-review.googlesource.com/c/sys/+/197540/

package cpuproxy

/*
#include <sys/auxv.h>

// getauxval is not available on Android until API level 20. Link it as a weak
// symbol and check whether it is not NULL before using it.
unsigned long getauxval(unsigned long type) __attribute__((weak));
*/
import "C"

import "sync"

// These constants are Linux specific.
const (
	_AT_HWCAP = 16 // hardware capability bit vector
)

// get returns the value of the getauxval auxiliary vector, or
// zero where the functionality is unavailable.
func get(t uint) uint {
	if C.getauxval == C.NULL {
		return 0
	}
	return uint(C.getauxval(C.ulong(t)))
}

// dogethwcap returns the value of _AT_HWCAP by calling
// the getauxval(3) function in libc (if available).
func dogethwcap() uint {
	return get(_AT_HWCAP)
}

// These variables allow to cache getauxval(3) results.
var (
	once  sync.Once
	hwcap uint
)

// gethwcap is like dogethwcap except that this function
// ensures we call getauxval(3) just once. After the first
// invocation we memoize the result.
func gethwcap() uint {
	once.Do(func() {
		hwcap = dogethwcap()
	})
	return hwcap
}

// HWCAP bits. These are exposed by Linux.
const (
	hwcap_AES     = 1 << 3
	hwcap_PMULL   = 1 << 4
	hwcap_SHA1    = 1 << 5
	hwcap_SHA2    = 1 << 6
	hwcap_CRC32   = 1 << 7
	hwcap_ATOMICS = 1 << 8
	hwcap_CPUID   = 1 << 11
)

// HasAES returns whether the CPU supports AES.
func HasAES() bool {
	return (gethwcap() & hwcap_AES) != 0
}

// HasGFMUL returns whether the CPU supports GFMUL.
func HasGFMUL() bool {
	return (gethwcap() & hwcap_PMULL) != 0
}
