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
		res = append(res, *(*MvCodeReaderDeviceInfo)(unsafe.Pointer(pstDevList.pDeviceInfo[i])))
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

	num := int(uint32(pstDevList.nDeviceNum))
	res := make([]MvCodeReaderDeviceInfo, num)
	for i := 0; i < num; i++ {
		res[i] = *(*MvCodeReaderDeviceInfo)(unsafe.Pointer(pstDevList.pDeviceInfo[i]))
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
		res = append(res, *(*MvCodeReaderDeviceInfo)(unsafe.Pointer(pstDevList.pDeviceInfo[i])))
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

func (d *Device) SetIntValue(key string, value int64) error {
	code := Err(int32(C.MV_CODEREADER_SetIntValue(d.handle, C.CString(key), C.longlong(value))))
	return code
}

func (d *Device) GetIntValue(key string) (int64, error) {

	r := &MvCodeReaderIntValueEx{}

	err := Err(int32(C.MV_CODEREADER_GetIntValue(d.handle, C.CString(key), (*C.MV_CODEREADER_INTVALUE_EX)(unsafe.Pointer(r)))))

	if err != nil {
		return -1, err
	}

	return r.CurValue, err
}

func (d *Device) SetEnumValue(key string, value uint32) error {
	code := Err(int32(C.MV_CODEREADER_SetEnumValue(d.handle, C.CString(key), C.uint(value))))
	return code
}

func (d *Device) GetStringValue(key string) (string, error) {

	r := &MvCodeReaderStringValue{}

	err := Err(int32(C.MV_CODEREADER_GetStringValue(d.handle, C.CString(key), (*C.MV_CODEREADER_STRINGVALUE)(unsafe.Pointer(r)))))

	if err != nil {
		return "", err
	}

	return string(r.CurValue[:r.MaxLength]), err
}

func (d *Device) SetStringValue(key string, value string) error {
	code := Err(int32(C.MV_CODEREADER_SetStringValue(d.handle, C.CString(key), C.CString(value))))
	return code
}

func (d *Device) GetBoolValue(key string) (bool, error) {

	var res bool

	err := Err(int32(C.MV_CODEREADER_GetBoolValue(d.handle, C.CString(key), (*C.bool)(unsafe.Pointer(&res)))))

	return res, err
}

func (d *Device) SetBoolValue(key string, value bool) error {

	var i int8

	i = 0

	if value {
		i = 1
	}

	code := Err(int32(C.MV_CODEREADER_SetBoolValue(d.handle, C.CString(key), C.bool(i))))
	return code
}

func (d *Device) SetCommandValue(key string) error {

	code := Err(int32(C.MV_CODEREADER_SetCommandValue(d.handle, C.CString(key))))
	return code
}

func (d *Device) RegisterImageCallBack() (chan CallBackResultEx2, error) {

	i, ch := callback.NewCallback()

	code := Err(int32(C.MV_CODEREADER_RegisterImageCallBack(d.handle, (*[0]byte)(C.CallBackGetOneFrameTimeoutEx2), unsafe.Pointer(&i))))

	return ch, code
}

func (d *Device) RegisterImageCallBackEx() (chan CallBackResultEx2, error) {

	i, ch := callback.NewCallback()

	code := Err(int32(C.MV_CODEREADER_RegisterImageCallBackEx(d.handle, (*[0]byte)(C.CallBackGetOneFrameTimeoutEx2), unsafe.Pointer(&i))))

	return ch, code
}

func (d *Device) RegisterImageCallBackEx2() (chan CallBackResultEx2, error) {

	i, ch := callback.NewCallback()

	code := Err(int32(C.MV_CODEREADER_RegisterImageCallBackEx2(d.handle, (*[0]byte)(C.CallBackGetOneFrameTimeoutEx2), unsafe.Pointer(&i))))

	return ch, code
}
