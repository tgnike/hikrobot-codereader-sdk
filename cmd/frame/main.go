package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/tgnike/hikrobot-codereader-sdk/MvCodeReaderSDK"
)

func main() {

	var b [MvCodeReaderSDK.MaxFrameSize]byte

	dir := filepath.Dir(os.Args[0])
	imath := filepath.Join(dir, "test.jpg")

	vers := MvCodeReaderSDK.GetSDKVersion()

	log.Printf("sdk version %v", vers)

	devs, err := MvCodeReaderSDK.EnumDevices(MvCodeReaderSDK.MvLayerGigeDevice)
	log.Printf("err %v devs %v", err, len(devs))

	devs, err = MvCodeReaderSDK.EnumDevices(MvCodeReaderSDK.MvLayerVirGigeDevice)
	log.Printf("err %v virt devs %v", err, len(devs))

	devs, err = MvCodeReaderSDK.EnumDevices(MvCodeReaderSDK.MvLayerVirUsbDevice)
	log.Printf("err %v virt usb devs %v", err, len(devs))

	devs, err = MvCodeReaderSDK.EnumDevices(MvCodeReaderSDK.MvLayerUnknowDevice)
	log.Printf("err %v virt usb devs %v", err, len(devs))

	iddevs, err := MvCodeReaderSDK.EnumIdDevices()
	log.Printf("err %v iddevs %v", err, len(iddevs))

	crdevs, err := MvCodeReaderSDK.EnumCodeReader()
	log.Printf("err %v crdevs %v", err, len(crdevs))

	// if len(crdevs) == 0 {
	// 	log.Print("no devs")
	// 	return
	// }

	// devInfo := crdevs[0]

	// if !MvCodeReaderSDK.IsDeviceAccessible(devInfo, MvCodeReaderSDK.MvAccessExclusive) {
	// 	log.Fatalf("device is not accessible %v", devInfo.SpecialInfo.MvGigeDeviceInfo.ModelName)
	// 	return
	// }

	handle, err := MvCodeReaderSDK.CreateHandleBySerialNumber("Device Control")

	if err != nil {
		log.Printf("err create handle %v", err)
		return
	}

	err = handle.OpenDevice()

	if err != nil {
		log.Printf("err open device %v", err)
	}

	defer handle.CloseDevice()
	defer handle.DestoyHandle()

	log.Print("SetEnumValue")
	handle.SetEnumValue(MvCodeReaderSDK.TriggerMode, 0)

	log.Print("StartGrabbing")
	handle.StartGrabbing()

	log.Print("GetOneFrameTimeout")
	err, info := handle.GetOneFrameTimeoutEx(&b, 500)
	//errcode = handle.RegisterImageCallBackEx(imath)

	log.Print("StopGrabbing")
	handle.StopGrabbing()

	if err != nil {
		log.Printf("err GetOneFrameTimeout %v", err)
		return
	}

	log.Printf("frame len %v, %s", info.FrameLen, imath)

	// err := os.WriteFile(imath, b[:info.FrameLen], 0644)

	// if err != nil {
	// 	log.Printf("WriteFile %v", err)
	// }

}
