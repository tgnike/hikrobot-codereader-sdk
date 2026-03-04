package main

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"log"
	"math/big"
	"net"
	"os"
	"path/filepath"
	"time"

	"github.com/tgnike/hikrobot-codereader-sdk/MvCodeReaderSDK"
)

func IP4toInt(IPv4Address net.IP) int64 {
	IPv4Int := big.NewInt(0)
	IPv4Int.SetBytes(IPv4Address.To4())
	return IPv4Int.Int64()
}

func main() {

	//var b [MvCodeReaderSDK.MaxFrameSize]uint8

	dir := filepath.Dir(os.Args[0])

	vers := MvCodeReaderSDK.GetSDKVersion()

	log.Printf("sdk version %v", vers)

	// errcode, devs := MvCodeReaderSDK.EnumDevices(MvCodeReaderSDK.MvLayerGigeDevice)
	// log.Printf("err %v devs %v", errcode, len(devs))

	// errcode, iddevs := MvCodeReaderSDK.EnumIdDevices()
	// log.Printf("err %v iddevs %v", errcode, len(iddevs))

	crdevs, err := MvCodeReaderSDK.EnumDevices(MvCodeReaderSDK.MvAccessExclusive)
	log.Printf("err %v crdevs %v", err, len(crdevs))

	devInfo := crdevs[0]

	if !MvCodeReaderSDK.IsDeviceAccessible(devInfo, MvCodeReaderSDK.MvAccessExclusive) {
		log.Fatalf("device is not accessible %v", devInfo.SpecialInfo.MvGigeDeviceInfo.ModelName)
		return
	}

	handle, err := MvCodeReaderSDK.CreateHandle(devInfo)

	if err != nil {
		log.Printf("err create handle %v", err)
	}

	log.Printf("handle %v", handle)

	err = handle.OpenDevice()

	if err != nil {
		log.Printf("err open device %v", err)
	}

	defer handle.CloseDevice()
	defer handle.DestoyHandle()

	// log.Print("SetEnumValue")
	// handle.SetEnumValue(MvCodeReaderSDK.TriggerMode, 0)

	roiIndex, err := handle.GetIntValue("AlgoRegionIndex")

	if err != nil {
		log.Printf("Get roi %v", err)
	}
	log.Printf("roiIndex= %v", roiIndex)

	roi, err := handle.GetIntValue("AlgoRegionLeftX")

	if err != nil {
		log.Printf("Get roi %v", err)
	}
	log.Printf("roi= %v", roi)

	handle.SetIntValue("AlgoRegionIndex", 2)

	roi1, err := handle.GetIntValue("AlgoRegionLeftX")

	if err != nil {
		log.Printf("Get roi %v", err)
	}
	log.Printf("roi1= %v", roi1)

	roi2, err := handle.GetIntValue("AlgoRegionLeftX")

	if err != nil {
		log.Printf("Get roi %v", err)
	}
	log.Printf("roi2= %v", roi2)

	userId, err := handle.GetStringValue("DeviceUserID")

	if err != nil {
		log.Printf("Get DeviceUserId %v", err)
	}
	log.Printf("UserId= %s", userId)

	// log.Print("GetOneFrameTimeout")
	// errcode, info := handle.GetOneFrameTimeoutEx(&b, 500)

	ch, err := handle.RegisterImageCallBackEx2()

	ctx, c := context.WithCancel(context.Background())
	defer c()

	go func(ctx context.Context) {
		c := 0
		for {
			select {
			case <-ctx.Done():
				return
			case cbr := <-ch:
				c++

				if cbr.FrameInfo == nil {
					continue
				}
				log.Printf("frame: %d framenum: %v length: %v", c, cbr.FrameInfo.FrameNum, cbr.FrameInfo.FrameLen)

				if cbr.FrameInfo.CodeListEx == nil {
					continue
				}

				codeNum := cbr.FrameInfo.CodeListEx.CodeNum
				log.Printf("codenum: %d", codeNum)

				r := bytes.NewReader(cbr.Image)

				img, err := jpeg.Decode(r)
				if err != nil {
					log.Print("no images")
					continue
				}
				rect1 := img.Bounds()
				rgba := image.NewRGBA(img.Bounds())
				draw.Draw(rgba, img.Bounds(), img, image.Point{0, 0}, draw.Src)

				for i := 0; i < int(codeNum); i++ {
					codeInfo := cbr.FrameInfo.CodeListEx.BcrInfoEx[i]
					codeLength := codeInfo.Len
					codeValue := codeInfo.Code[:codeLength]
					log.Printf("code %d length: %d, value: %s", i+1, codeLength, codeValue)
					log.Printf("x1 %d y1: %d, x2: %d, y2: %d", codeInfo.Pt[0].X, codeInfo.Pt[0].Y, codeInfo.Pt[1].X, codeInfo.Pt[1].Y)

					rect2 := image.Rect(int(codeInfo.Pt[0].X), int(codeInfo.Pt[0].Y), int(codeInfo.Pt[2].X), int(codeInfo.Pt[2].Y))

					if !rect2.In(rect1) {
						err = fmt.Errorf("error: rectangle outside image")
						return
					}

					DrawThickFrame(rgba, rect2.Min.X, rect2.Min.Y, rect2.Max.X, rect2.Max.Y, 10, color.White)

					// for x := rect1.Min.X; x <= rect1.Max.X; x++ {
					// 	for y := rect1.Min.Y; y <= rect1.Max.Y; y++ {
					// 		p := image.Pt(x, y)
					// 		if p.In(rect2) {
					// 			rgba.Set(x, y, color.White)
					// 			// } else {
					// 			// 	rgba.Set(x, y, img.At(x, y))
					// 		}
					// 	}
					// }

				}

				w, err := os.Create(filepath.Join(dir, fmt.Sprintf("testQW-%v.jpg", c)))
				log.Print("save show bc %v", err)

				err = jpeg.Encode(w, rgba, nil)
				log.Print("save encode bc %v", err)
				w.Close()

				saveImage(filepath.Join(dir, fmt.Sprintf("test-%v.jpg", c)), cbr.Image)

				log.Print("----------------------------------")
			}
		}
	}(ctx)

	log.Print("StartGrabbing")
	handle.StartGrabbing()

	t := time.NewTimer(time.Duration(time.Second * 10))
	<-t.C

	log.Print("StopGrabbing")
	handle.StopGrabbing()

	if err != nil {
		log.Printf("err GetOneFrameTimeout %v", err)
		return
	}

	// log.Printf("frame len %v", info.FrameLen)

	// imath := filepath.Join(dir, "test.jpg")

	// err := os.WriteFile(imath, b[:info.FrameLen], 0644)

	// if err != nil {
	// 	log.Printf("WriteFile %v", err)
	// }

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

// DrawThickFrame draws a rectangular frame inside the given image.
func DrawThickFrame(img *image.RGBA, x1, y1, x2, y2, thickness int, col color.Color) {
	// Draw top and bottom borders
	for t := 0; t < thickness; t++ {
		for x := x1; x <= x2; x++ {
			img.Set(x, y1+t, col) // Top border
			img.Set(x, y2-t, col) // Bottom border
		}
	}

	// Draw left and right borders
	for t := 0; t < thickness; t++ {
		for y := y1; y <= y2; y++ {
			img.Set(x1+t, y, col) // Left border
			img.Set(x2-t, y, col) // Right border
		}
	}
}

// HLine draws a horizontal line
func HLine(img *image.RGBA, x1, y, x2 int, c color.Color) {
	for ; x1 <= x2; x1++ {
		img.Set(x1, y, c)
	}
}

// VLine draws a veritcal line
func VLine(img *image.RGBA, x, y1, y2 int, c color.Color) {
	for ; y1 <= y2; y1++ {
		img.Set(x, y1, c)
	}
}

// Rect draws a rectangle utilizing HLine() and VLine()
func Rect(img *image.RGBA, x1, y1, x2, y2 int, c color.Color) {
	HLine(img, x1, y1, x2, c)
	HLine(img, x1, y2, x2, c)
	VLine(img, x1, y1, y2, c)
	VLine(img, x2, y1, y2, c)
}
