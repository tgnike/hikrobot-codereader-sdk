package MvCodeReaderSDK

/*
#cgo CFLAGS: -I../include
#include <stdio.h>
#include <stdlib.h>
#cgo windows amd64 LDFLAGS: -L${SRCDIR}/../lib/win/64 -lMvCodeReaderCtrl -static
#include "MvCodeReaderCtrl.h"
*/
import "C"
import (
	"log"
	"sync"
	"unsafe"
)

type CallBackResultEx2 struct {
	Lenth     int
	Image     []byte
	FrameInfo *MVImageOutInfoEx2
}

type CallBackRegister struct {
	results []chan CallBackResultEx2
	mutex   sync.Mutex
}

func (c *CallBackRegister) NewCallback() (int, chan CallBackResultEx2) {

	c.mutex.Lock()
	defer c.mutex.Unlock()

	ch := make(chan CallBackResultEx2, 10)
	c.results = append(c.results, ch)

	return len(c.results) - 1, ch

}

var callback = &CallBackRegister{results: make([]chan CallBackResultEx2, 0), mutex: sync.Mutex{}}

//export go_callback_output
func go_callback_output(pData *C.uchar, pstFrameInfo *C.MV_CODEREADER_IMAGE_OUT_INFO_EX2, pUser unsafe.Pointer) {

	id := *(*int)(unsafe.Pointer(pUser))
	rLen := len(callback.results)

	if id < 0 || id > rLen {
		return
	}

	ch := callback.results[id]

	frameInfo := NewMVImageOutInfoEx2(pstFrameInfo)
	lenth := int(frameInfo.FrameLen)

	if lenth == 0 {
		return
	}

	data := *((*[]byte)(unsafe.Pointer(&pData)))
	image, err := getImageBytes(data, lenth)

	if err != nil {
		log.Print(err)
	}

	ch <- CallBackResultEx2{Image: image, FrameInfo: frameInfo}

}

// Copy bytes with length
func getImageBytes(data []byte, length int) ([]byte, error) {

	i := data[:length]
	var copyField = make([]byte, length)
	copy(copyField, i)

	return copyField, nil

}

func CopyBarcodeResults(r *MVImageOutInfoEx2, s *C.MV_CODEREADER_IMAGE_OUT_INFO_EX2) {

	for i := 0; i < len(s.pstCodeListEx.stBcrInfoEx); i++ {
		r.CodeListEx.BcrInfoEx[i] = NewMvBcrInfoEx(s.pstCodeListEx.stBcrInfoEx[i])
	}

}

func NewMVImageOutInfoEx2(pstFrameInfo *C.MV_CODEREADER_IMAGE_OUT_INFO_EX2) *MVImageOutInfoEx2 {

	s := MVImageOutInfoEx2{}
	s.Width = uint16(pstFrameInfo.nWidth)
	s.Height = uint16(pstFrameInfo.nHeight)

	s.PixelType = pstFrameInfo.enPixelType
	s.TriggerIndex = uint32(pstFrameInfo.nTriggerIndex)
	s.FrameNum = uint32(pstFrameInfo.nFrameNum)
	s.FrameLen = uint32(pstFrameInfo.nFrameLen)
	s.TimeStampHigh = uint32(pstFrameInfo.nTimeStampHigh)
	s.TimeStampLow = uint32(pstFrameInfo.nTimeStampLow)
	s.FlaseTrigger = uint32(pstFrameInfo.bFlaseTrigger)
	s.FocusScore = uint32(pstFrameInfo.nFocusScore)
	s.IsGetCode = GoBool(pstFrameInfo.bIsGetCode)
	s.CodeListEx = NewMvResultBcrEx(pstFrameInfo.pstCodeListEx)
	// //s.WaybillList
	s.EventID = uint32(pstFrameInfo.nEventID)
	s.ChannelID = uint32(pstFrameInfo.nChannelID)
	s.ImageCost = uint32(pstFrameInfo.nImageCost)
	// //s.UnparsedBcrList
	// //s.UnparsedOcrList
	s.WholeFlag = uint16(pstFrameInfo.nWholeFlag)
	s.Res = uint16(pstFrameInfo.nRes)
	// //s.Reserved = pstFrameInfo.nReserved

	return &s
}

func NewMvResultBcrEx(c *C.MV_CODEREADER_RESULT_BCR_EX) *MvResultBcrEx {

	g := &MvResultBcrEx{}
	g.CodeNum = uint32(c.nCodeNum)
	//g.NoReadNum = uint16(c.nNoReadNum)
	// g.Res = uint16(c.nRes)

	maxCodeInfo := int(g.CodeNum)

	for i := 0; i < len(g.BcrInfoEx); i++ {

		if i == maxCodeInfo {
			break
		}

		g.BcrInfoEx[i] = NewMvBcrInfoEx(c.stBcrInfoEx[i])
	}

	return g

}

func NewMvBcrInfoEx(c C.MV_CODEREADER_BCR_INFO_EX) MvBcrInfoEx {
	g := MvBcrInfoEx{}
	g.ID = uint32(c.nID)
	g.Code = ([256]byte)(unsafe.Slice((*byte)(unsafe.Pointer(&c.chCode)), 256))

	g.Len = uint32(c.nLen)

	g.BarType = uint32(c.nBarType)
	//g.Pt = ([4]MvPoint)(unsafe.Slice((*MvPoint)(unsafe.Pointer(&c.pt)), 4))
	g.Pt = ([4]MvPoint)(make([]MvPoint, 4))

	for i, p := range c.pt {

		npt := MvPoint{X: int32(p.x), Y: int32(p.y)}

		g.Pt[i] = npt

	}
	//g.CodeQuality  =([4]MvCodeInfo)(unsafe.Slice((*C.CodeQuality)(unsafe.Pointer(&c.chCode)), 256))

	g.Angle = int32(c.nAngle)

	g.MainPackageId = uint32(c.nMainPackageId)
	g.SubPackageId = uint32(c.nSubPackageId)
	g.AppearCount = uint16(c.sAppearCount)
	g.PPM = uint16(c.sPPM)
	g.AlgoCost = uint16(c.sAlgoCost)
	g.Sharpness = uint16(c.sSharpness)

	g.IsGetQuality = GoBool(c.bIsGetQuality)
	g.IDRScore = uint32(c.nIDRScore)
	g.D1IsGetQuality = uint32(c.n1DIsGetQuality)
	g.TotalProcCost = uint32(c.nTotalProcCost)
	g.TriggerTimeTvHigh = uint32(c.nTriggerTimeTvHigh)
	g.TriggerTimeTvLow = uint32(c.nTriggerTimeTvLow)
	g.TriggerTimeUtvHigh = uint32(c.nTriggerTimeUtvHigh)
	g.TriggerTimeUtvLow = uint32(c.nTriggerTimeUtvLow)
	g.PollingIndex = uint16(c.sPollingIndex)
	//g.Res = uint16(c.sRes)
	//g.Reserved = c.nReserved
	return g

}

func GoBool(b C.bool) bool {
	return b != 0
}
