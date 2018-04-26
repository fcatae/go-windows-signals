package main

import (
	"fmt"
	"syscall"
	"unsafe"
)

var (
	kernel32        = syscall.NewLazyDLL("kernel32.dll")
	procCreateMutex = kernel32.NewProc("CreateMutexW")
	procOpenMutex   = kernel32.NewProc("OpenMutexW")
)

func createMutex(name string) (uintptr, error) {
	paramMutexName := uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(name)))

	ret, _, err := procCreateMutex.Call(0, 0, paramMutexName)

	switch int(err.(syscall.Errno)) {
	case 0:
		return ret, nil
	default:
		return ret, err
	}
}

func createMutex(name string) (uintptr, error) {
	paramMutexName := uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(name)))

	ret, _, err := procCreateMutex.Call(0, 0, paramMutexName)

	switch int(err.(syscall.Errno)) {
	case 0:
		return ret, nil
	default:
		return ret, err
	}
}

func openMutex(name string) (uintptr, error) {
	paramMutexName := uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(name)))

	ret, _, err := procOpenMutex.Call(
		0x00100000, // SYNCRONIZE
		0,          // Not inheritable
		paramMutexName)

	switch int(err.(syscall.Errno)) {
	case 0:
		return ret, nil
	default:
		return ret, err
	}
}

func main() {
	fmt.Println("CreateMutex")

	m, err := createMutex("SomeMutexNameGUID")

	fmt.Println("- return:", m, err)

	if err == nil {
		fmt.Println("- return:", m)

		m2, err2 := openMutex("SomeMutexNameGUID")
		fmt.Println("- return:", m2, err2)
	}

	fmt.Scanln()
}
