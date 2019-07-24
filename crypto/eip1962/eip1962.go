package eip1962

/*
#cgo LDFLAGS: -L. -leip1962
#include "./src/wrapper.h"
extern int run(const char *i, uint32_t i_len, char *o, uint32_t *o_len, char *err, uint32_t *char_len);
extern int meter_gas(const char *i, uint32_t i_len, uint64_t *gas);
*/
import "C"

import (
	"errors"
	"math/big"
	"unsafe"
)

const maxOutputLen = 256 * 3 * 2

var (
	ErrInvalidMsgLen = errors.New("invalid data length, need >= bytes")
	ErrCallFailed    = errors.New("library call returned an error")
)

// Call calls the C++ implementation of the EIP
func Call(data []byte) ([]byte, error) {
	if len(data) == 0 {
		return nil, ErrInvalidMsgLen
	}
	ilen := len(data)
	outputBytes := make([]byte, maxOutputLen)
	olen := uint32(0)
	errStringBytes := make([]byte, maxOutputLen)
	errStringLen := uint32(0)

	var (
		inputdata  = (*C.char)(unsafe.Pointer(&data[0]))
		inputlen   = (C.uint32_t)(ilen)
		outputdata = (*C.char)(unsafe.Pointer(&outputBytes[0]))
		outputlen  = (*C.uint32_t)(unsafe.Pointer(&olen))
		errdata    = (*C.char)(unsafe.Pointer(&errStringBytes[0]))
		errlen     = (*C.uint32_t)(unsafe.Pointer(&errStringLen))
	)

	result := C.run(inputdata, inputlen, outputdata, outputlen, errdata, errlen)
	if result == 0 {
		// parse error string
		return nil, ErrCallFailed
	}

	return outputBytes[:olen], nil
}

// EstimateGas calls C++ implementation for a gas estimte
func EstimateGas(data []byte) (*big.Int, error) {
	return big.NewInt(1000000), nil
}
