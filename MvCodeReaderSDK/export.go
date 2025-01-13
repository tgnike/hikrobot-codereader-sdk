package MvCodeReaderSDK

/*
#cgo CFLAGS: -I../include
#cgo windows amd64 LDFLAGS: -L${SRCDIR}/../lib/win/64 -lMvCodeReaderCtrl -static
#include "MvCodeReaderCtrl.h"
*/
import "C"
import (
	"bufio"
	"log"
	"os"
	"unsafe"
)

//export go_callback_output
func go_callback_output(pData *C.uchar, pstFrameInfo *C.MV_CODEREADER_IMAGE_OUT_INFO_EX, pUser unsafe.Pointer) {

	stru := ((*MVFrameOutInfoEx)(unsafe.Pointer(pstFrameInfo)))

	for i, v := range stru.CodeList.stBcrInfo {

		if v.Len == 0 {
			continue
		}

		code := v.Code[:v.Len]
		log.Printf("%v code %v", i, string(code))

	}

	path := "C:/code/go/hikrobot-mvcodereader/cmd/code/d2.jpg"
	data := *((*[]byte)(unsafe.Pointer(&pData)))
	lenth := int(stru.FrameLen)
	image, err := getImageBytes(data, lenth)

	if err != nil {
		log.Print(err)
	}

	barcodes(pstFrameInfo)
	if image != nil {
		saveImage(path, image)
	}

	//err := os.WriteFile("C:/code/go/hikrobot-mvcodereader/cmd/code/d2.jpg", []byte(pk), 0644)

}

func barcodes(pstFrameInfo *C.MV_CODEREADER_IMAGE_OUT_INFO_EX) {

	codes := ((int)(pstFrameInfo.pstCodeList.nCodeNum))

	log.Printf("codes %v", codes)

}

func getImageBytes(data []byte, lenth int) ([]byte, error) {

	return data[:lenth], nil

}

func saveImage(path string, data []byte) {

	f, err := os.Create(path)

	if err != nil {
		log.Printf("Create %v", err)
	}

	w := bufio.NewWriter(f)
	_, err = w.Write(data)

	if err != nil {
		log.Printf("WriteFile %v", err)
	}

	w.Flush()
	f.Close()
}
