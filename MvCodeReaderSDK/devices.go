package MvCodeReaderSDK

// Device Transport Layer Protocol Type
const (
	MvLayerUnknowDevice          = 0x00000000 // MvUnknowDevice: \~chinese 未知设备类型，保留意义 \~english Unknown Device Type, Reserved
	MvLayerGigeDevice            = 0x00000001 // MvGigeDevice: \~chinese GigE设备 \~english GigE Device
	MvLayer1394Device            = 0x00000002 // Mv1394Device: \~chinese 1394-a/b 设备 \~english 1394-a/b Device
	MvLayerUsbDevice             = 0x00000004 // MvUsbDevice: \~chinese USB 设备 \~english USB Device
	MvLayerCameralinkDevice      = 0x00000008 // MvCameralinkDevice: \~chinese CameraLink设备 \~english CameraLink Device
	MvLayerVirGigeDevice         = 0x00000010 // MvVirGigeDevice: \~chinese 虚拟GigE设备 \~english Virtual GigE Device
	MvLayerVirUsbDevice          = 0x00000020 // MvVirUsbDevice: \~chinese 虚拟USB设备 \~english Virtual USB Device
	MvLayerGentlGigeDevice       = 0x00000040 // MvGentlGigeDevice: \~chinese 自研网卡下GigE设备 \~english GenTL GigE Device
	MvLayerGentlCameralinkDevice = 0x00000080 // MvGentlCameralinkDevice: \~chinese CameraLink相机设备 \~english GenTL CameraLink Camera Device
	MvLayerGentlCxpDevice        = 0x00000100 // MvGentlCxpDevice: \~chinese CoaXPress设备 \~english GenTL CoaXPress Device
	MvLayerGentlXofDevice        = 0x00000200 // MvGentlXofDevice: \~chinese XoF设备 \~english GenTL XoF Device
)
