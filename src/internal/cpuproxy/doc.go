// Package cpuproxy is a getauxval aware proxy for internal/cpu.
//
// The problem that we want to solve is that on Android there are
// cases where reading /proc/self/auxv is not possible.
//
// This causes crypto/tls to not choose AES where it would otherwise
// be possible, in turn causing censorship. See also the
// https://github.com/ooni/probe/issues/1444 issue for more details.
//
// Ideally we would like to call getauxval(3) when initializing
// the runtime package. However, runtime cannot use CGO. Doing that
// leads to an import loop, so we cannot build.
//
// We could also try to parse /proc/cpuinfo (I didn't explore this route).
//
// The solution chosen here is to export predicates on the CPU
// functionality. We limit ourselves to what we need in order to
// choose AES in crypto/tls when the CPU supports it.
//
// The predicates use internal/cpu values for every GOOS/GOARCH except
// android/arm64, where we call getauxval(3).
package cpuproxy
