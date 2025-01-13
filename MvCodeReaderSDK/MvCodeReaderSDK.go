package MvCodeReaderSDK

/*
#cgo CFLAGS: -I../include
#cgo windows amd64 LDFLAGS: -L${SRCDIR}/../lib/win/64 -lMvCodeReaderCtrl -static
#include "MvCodeReaderCtrl.h"
#include <stdlib.h>
#include <stdio.h>
// normally you will have to define function or variables
// in another separate C file to avoid the multiple definition
// errors, however, using "static inline" is a nice workaround
// for simple functions like this one.

void __stdcall CallBackGetOneFrameTimeoutEx2(unsigned char * pData, MV_CODEREADER_IMAGE_OUT_INFO_EX2* pstFrameInfo, void* pUser) {
	go_callback_output(pData, pstFrameInfo, pUser);
}
*/
import "C"
import (
	"log"
	"unsafe"
)

// C 语言类型	CGO 类型	Go 语言类型
// char	C.char	byte
// singed char	C.schar	int8
// unsigned char	C.uchar	uint8
// short	C.short	int16
// unsigned short	C.ushort	uint16
// int	C.int	int32
// unsigned int	C.uint	uint32
// long	C.long	int32
// unsigned long	C.ulong	uint32
// long long int	C.longlong	int64
// unsigned long long int	C.ulonglong	uint64
// float	C.float	float32
// double	C.double	float64
// size_t	C.size_t	uint

const MaxFrameSize = 10 * 1024 * 1024 * 6

func GetSDKVersion() int32 {

	return int32(C.MV_CODEREADER_GetSDKVersion())

}

// EnumDevices
func EnumDevices(nTLayerType uint32) ([]MvCodeReaderDeviceInfo, error) {
	pstDevList := C.struct__MV_CODEREADER_DEVICE_INFO_LIST_{}

	err := Err(int32(C.MV_CODEREADER_EnumDevices((*C.struct__MV_CODEREADER_DEVICE_INFO_LIST_)(unsafe.Pointer(&pstDevList)), C.uint(nTLayerType))))

	if err != nil {
		return nil, err
	}

	var res []MvCodeReaderDeviceInfo
	num := int(uint32(pstDevList.nDeviceNum))
	for i := 0; i < num; i++ {
		res = append(res, *(*MvCodeReaderDeviceInfo)(unsafe.Pointer(pstDevList.pDeviceInfo[0])))
	}

	return res, nil
}

// EnumCodeReader
func EnumCodeReader() ([]MvCodeReaderDeviceInfo, error) {
	pstDevList := C.struct__MV_CODEREADER_DEVICE_INFO_LIST_{}

	err := Err(int32(C.MV_CODEREADER_EnumCodeReader((*C.struct__MV_CODEREADER_DEVICE_INFO_LIST_)(unsafe.Pointer(&pstDevList)))))

	if err != nil {
		return nil, err
	}

	var res []MvCodeReaderDeviceInfo
	num := int(uint32(pstDevList.nDeviceNum))
	for i := 0; i < num; i++ {
		res = append(res, *(*MvCodeReaderDeviceInfo)(unsafe.Pointer(pstDevList.pDeviceInfo[0])))
	}

	return res, nil
}

// EnumIdDevices
func EnumIdDevices() ([]MvCodeReaderDeviceInfo, error) {
	pstDevList := C.struct__MV_CODEREADER_DEVICE_INFO_LIST_{}

	err := Err(int32(C.MV_CODEREADER_EnumIDDevices((*C.struct__MV_CODEREADER_DEVICE_INFO_LIST_)(unsafe.Pointer(&pstDevList)))))

	if err != nil {
		return nil, err
	}

	var res []MvCodeReaderDeviceInfo
	num := int(uint32(pstDevList.nDeviceNum))
	for i := 0; i < num; i++ {
		res = append(res, *(*MvCodeReaderDeviceInfo)(unsafe.Pointer(pstDevList.pDeviceInfo[0])))
	}

	return res, nil
}

func IsDeviceAccessible(pstDevInfo MvCodeReaderDeviceInfo, nAccessMode uint32) bool {

	code := byte(C.MV_CODEREADER_IsDeviceAccessible((*C.struct__MV_CODEREADER_DEVICE_INFO_)(unsafe.Pointer(&pstDevInfo)), C.uint(nAccessMode)))

	return code == 1
}

type Device struct {
	handle unsafe.Pointer
}

func CreateHandle(pstDevInfo MvCodeReaderDeviceInfo) (*Device, error) {
	d := &Device{}
	errCode := Err(int32(C.MV_CODEREADER_CreateHandle(&d.handle, (*C.struct__MV_CODEREADER_DEVICE_INFO_)(unsafe.Pointer(&pstDevInfo)))))

	return d, errCode
}

func CreateHandleBySerialNumber(serial string) (*Device, error) {
	d := &Device{}
	errCode := Err(int32(C.MV_CODEREADER_CreateHandleBySerialNumber(&d.handle, ((*C.char)(unsafe.Pointer(&serial))))))

	return d, errCode
}

func (d *Device) DestoyHandle() error {

	errCode := Err(int32(C.MV_CODEREADER_DestroyHandle(d.handle)))

	return errCode
}

func (d *Device) OpenDevice() error {
	return Err(int32(C.MV_CODEREADER_OpenDevice(d.handle)))

}

func (d *Device) CloseDevice() error {
	return Err(int32(C.MV_CODEREADER_CloseDevice(d.handle)))
}

func (d *Device) StartGrabbing() error {
	return Err(int32(C.MV_CODEREADER_StartGrabbing(d.handle)))
}

func (d *Device) StopGrabbing() error {
	return Err(int32(C.MV_CODEREADER_StopGrabbing(d.handle)))
}

func (d *Device) GetOneFrameTimeout(pData *[]byte, timeOut uint32) (error, MVFrameOutInfo) {
	var pFrameInfo MVFrameOutInfo

	var p = (**C.uchar)(unsafe.Pointer(pData))

	code := Err(int32(C.MV_CODEREADER_GetOneFrameTimeout(d.handle, p, (*C.MV_CODEREADER_IMAGE_OUT_INFO)(unsafe.Pointer(&pFrameInfo)), C.uint(timeOut))))
	return code, pFrameInfo
}

func (d *Device) GetOneFrameTimeoutEx(pData *[MaxFrameSize]byte, timeOut uint32) (error, MVFrameOutInfoEx) {
	var pFrameInfo MVFrameOutInfoEx
	bb := make([]byte, MaxFrameSize)
	b := C.CBytes(bb)

	var p = (**C.uchar)(unsafe.Pointer(&b))

	code := Err(int32(C.MV_CODEREADER_GetOneFrameTimeoutEx(d.handle, p, (*C.MV_CODEREADER_IMAGE_OUT_INFO_EX)(unsafe.Pointer(&pFrameInfo)), C.uint(timeOut))))

	pd := *((*[]byte)(unsafe.Pointer(p)))

	log.Printf("%v", pd)

	return code, pFrameInfo
}

func (d *Device) GetOneFrameTimeoutEx2(pData *[MaxFrameSize]byte, timeOut uint32) (error, MVFrameOutInfo) {
	var pFrameInfo MVFrameOutInfo

	var p = (**C.uchar)(unsafe.Pointer(pData))

	code := Err(int32(C.MV_CODEREADER_GetOneFrameTimeoutEx2(d.handle, p, (*C.MV_CODEREADER_IMAGE_OUT_INFO_EX2)(unsafe.Pointer(&pFrameInfo)), C.uint(timeOut))))
	return code, pFrameInfo
}

func (d *Device) SetEnumValue(key string, value uint32) error {
	code := Err(int32(C.MV_CODEREADER_SetEnumValue(d.handle, C.CString(key), C.uint(value))))
	return code
}

func (d *Device) RegisterImageCallBack(path string) error {

	cstr := C.CString(path)
	pstr := unsafe.Pointer(cstr)

	code := Err(int32(C.MV_CODEREADER_RegisterImageCallBack(d.handle, (*[0]byte)(C.CallBackGetOneFrameTimeoutEx2), pstr)))

	return code
}

func (d *Device) RegisterImageCallBackEx(path string) error {

	cstr := C.CString(path)
	pstr := unsafe.Pointer(cstr)

	code := Err(int32(C.MV_CODEREADER_RegisterImageCallBackEx(d.handle, (*[0]byte)(C.CallBackGetOneFrameTimeoutEx2), pstr)))

	return code
}

func (d *Device) RegisterImageCallBackEx2(path string) error {

	cstr := C.CString(path)
	pstr := unsafe.Pointer(cstr)

	code := Err(int32(C.MV_CODEREADER_RegisterImageCallBackEx2(d.handle, (*[0]byte)(C.CallBackGetOneFrameTimeoutEx2), pstr)))

	return code
}
