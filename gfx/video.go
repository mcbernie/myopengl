package gfx

import (
	"image"

	"gocv.io/x/gocv"
)

//Video Simple Video structure
type Video struct {
	device *gocv.VideoCapture
}

//GetFrame get the current Streaming Frame
func (video *Video) GetFrame() (image.Image, error) {
	img := gocv.NewMat()
	defer img.Close()
	video.device.Read(&img)

	image, err := img.ToImage()
	return image, err
}

//InitVideo Initialize Video Streaming
func InitVideo() *Video {
	webcam, _ := gocv.VideoCaptureDevice(0)

	return &Video{
		device: webcam,
	}
}

func InitVideoFromFile() *Video {
	v, _ := gocv.VideoCaptureFile("/users/nico/Downloads/e0f162aabe867186e61dd087825ebaef_asset.mp4")

	return &Video{
		device: v,
	}
}

//Delete Close the current streaming Device
func (video *Video) Delete() {
	video.device.Close()

}
