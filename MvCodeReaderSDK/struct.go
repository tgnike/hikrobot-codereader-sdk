package MvCodeReaderSDK

import "C"
import "unsafe"

const MvMaxDeviceNum = 256
const InfoMaxBufferSize = 64
const MaxBcrCodeLen = 256
const MaxBcrCodeLenEx = 4096

type MvCodeReaderDeviceInfo struct {
	MajorVer     uint16
	MinorVer     uint16
	MacAddrHigh  uint32
	MacAddrLow   uint32
	TLayerType   uint32
	DevTypeInfo  uint32
	SelectDevice bool
	Reserved     [2]uint32
	SpecialInfo  sSpecialInfo
	GigEInfo     MvGigeDeviceInfo
	Usb3VInfo    MvUsb3DeviceInfo
}

func (i MvCodeReaderDeviceInfo) MajorVersion() uint16 {

	return i.MajorVer

}

type sSpecialInfo struct {
	MvGigeDeviceInfo MvGigeDeviceInfo ///< [OUT] \~chinese GigE设备信息              \~english GigE Device Info
	MvUsb3DeviceInfo MvUsb3DeviceInfo ///< [OUT] \~chinese USB设备信息               \~english USB Device Info
}

type MvGigeDeviceInfo struct {
	IpCfgOption              uint32
	IpCfgCurrent             uint32
	CurrentIp                uint32
	CurrentSubNetMask        uint32
	DefultGateWay            uint32
	ManufacturerName         [32]byte
	ModelName                [32]byte
	DeviceVersion            [32]byte
	ManufacturerSpecificInfo [48]byte
	SerialNumber             [16]byte
	UserDefinedName          [16]byte
	NetExport                uint32
	Reserved                 [4]uint32
}

type MvUsb3DeviceInfo struct {
	CrtlInEndPoint   uint8
	CrtlOutEndPoint  uint8
	StreamEndPoint   uint8
	EventEndPoint    uint8
	IdVendor         uint16
	IdProduct        uint16
	DeviceNumber     uint32
	DeviceGUID       [InfoMaxBufferSize]byte
	VendorName       [InfoMaxBufferSize]byte
	ModelName        [InfoMaxBufferSize]byte
	FamilyName       [InfoMaxBufferSize]byte
	DeviceVersion    [InfoMaxBufferSize]byte
	ManufacturerName [InfoMaxBufferSize]byte
	SerialNumber     [InfoMaxBufferSize]byte
	UserDefinedName  [InfoMaxBufferSize]byte
	BcdUSB           uint32
	DeviceAddress    uint32
	Reserved         [2]uint32
}

type MVFrameOutInfo struct {
	Width     uint16
	Height    uint16
	PixelType uint32

	FrameNum         uint32
	DevTimeStampHigh uint32
	DevTimeStampLow  uint32
	Reserved0        uint32
	HostTimeStamp    int64

	FrameLen uint32

	SecondCount uint32
	CycleCount  uint32
	CycleOffset uint32

	Gain              float32
	ExposureTime      float32
	AverageBrightness uint32

	Red   uint32
	Green uint32
	Blue  uint32

	FrameCounter uint32
	TriggerIndex uint32

	Input  uint32
	Output uint32

	OffsetX     uint16
	OffsetY     uint16
	ChunkWidth  uint16
	ChunkHeight uint16

	LostPacket uint32

	UnparsedChunkNum uint32

	UnparsedChunkContent unsafe.Pointer // This is the void* pUnparsedChunkContent in the union
	nAligning            int64          ///< [OUT] \~chinese 校准                   \~english Aligning

	ExtendWidth  uint32
	ExtendHeight uint32

	IsGetCode bool
	CodeList  MvResultBcr

	Reserved [34]uint32
}

type MVFrameOutInfoEx struct {
	Width     uint16
	Height    uint16
	PixelType uint32

	TriggerIndex  uint32
	FrameNum      uint32
	FrameLen      uint32
	TimeStampHigh uint32
	TimeStampLow  uint32
	FlaseTrigger  uint32
	FocusScore    uint32
	IsGetCode     bool
	CodeList      *MvResultBcr
	WaybillList   MvWaybill

	EventID   uint32
	ChannelID uint32
	ImageCost uint32

	UnparsedOcrList []uint

	WholeFlag uint16
	Res       uint16
	Reserved  [3]uint32
}

type MVImageOutInfoEx2 struct {
	Width     uint16
	Height    uint16
	PixelType uint32

	TriggerIndex  uint32
	FrameNum      uint32
	FrameLen      uint32
	TimeStampHigh uint32
	TimeStampLow  uint32
	FlaseTrigger  uint32
	FocusScore    uint32
	IsGetCode     bool
	CodeListEx    *MvResultBcrEx
	WaybillList   MvWaybill

	EventID   uint32
	ChannelID uint32
	ImageCost uint32

	UnparsedBcrList []uint
	UnparsedOcrList []uint

	WholeFlag uint16
	Res       uint16
	Reserved  [25]uint
}

type MvWaybill struct {
}

type MvResultBcr struct {
	CodeNum   uint32
	BcrInfo   [200]MvBcrInfo
	NoReadNum uint16
	Res       uint16
	Reserved  [3]uint32
}

type MvResultBcrEx struct {
	CodeNum   uint32
	BcrInfoEx [200]MvBcrInfoEx
	NoReadNum uint16
	Res       uint16
	Reserved  [7]uint32
}

type MvBcrInfoEx struct {
	ID          uint32
	Code        [MaxBcrCodeLen]byte
	Len         uint32
	BarType     uint32
	Pt          [4]MvPoint
	CodeQuality MvCodeInfo

	Angle         int32
	MainPackageId uint32
	SubPackageId  uint32
	AppearCount   uint16
	PPM           uint16
	AlgoCost      uint16
	Sharpness     uint16

	IsGetQuality       bool
	IDRScore           uint32
	D1IsGetQuality     uint32
	TotalProcCost      uint32
	TriggerTimeTvHigh  uint32
	TriggerTimeTvLow   uint32
	TriggerTimeUtvHigh uint32
	TriggerTimeUtvLow  uint32
	PollingIndex       uint16
	Res                uint16
	Reserved           [23]uint32
}

type MvCodeInfo struct{}

type MvBcrInfo struct {
	ID      uint32
	Code    [MaxBcrCodeLen]byte
	Len     uint32
	BarType uint32
	pt      [4]MvPoint

	Angle         int32
	MainPackageId uint32
	SubPackageId  uint32
	AppearCount   uint16
	PPM           uint16
	AlgoCost      uint16
	Sharpness     uint16
}

type MvPoint struct {
	X int32
	Y int32
}

type MvFrameOut struct {
	Addr      unsafe.Pointer
	FrameInfo MVFrameOutInfo
	Res       [16]uint32
}

// MV_CODEREADER_STRINGVALUE
type MvCodeReaderStringValue struct {
	CurValue  [256]byte
	MaxLength int64
	Res       [2]uint32
}

// MV_CODEREADER_INTVALUE_EX
type MvCodeReaderIntValueEx struct {
	CurValue int64
	Max      int64
	Min      int64
	Inc      int64
	Res      [16]uint32
}
