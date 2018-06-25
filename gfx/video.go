package gfx

import (
	"image"
	_ "image/jpeg" // Import JPEG Decoding
	_ "image/png"  // Import PNG Decoding
	"log"
	"time"

	"github.com/3d0c/gmf"

	"gocv.io/x/gocv"
)

//Video Simple Video structure
type Video struct {
	device *gocv.VideoCapture

	//decoding
	decFmt    *gmf.FmtCtx
	decStream *gmf.Stream
	decCodec  *gmf.CodecCtx
	decFrame  *gmf.Frame
	swsCtx    *gmf.SwsCtx
}

func (v *Video) CleanUP() {
	log.Println("Called Video Cleanup")
	if v.decFmt != nil {
		v.decFmt.CloseInputAndRelease()
		v.decFmt = nil
	}
	if v.decFrame != nil {
		v.decFrame.Release()
		v.decFrame = nil
	}
	if v.decCodec != nil {
		gmf.Release(v.decCodec)
		v.decCodec = nil
	}
	if v.swsCtx != nil {
		gmf.Release(v.swsCtx)
	}
}

func CreateVideo(srcFileName string) *Video {
	var err error

	v := &Video{}

	//Load Input file and put in AVFormatContext
	v.decFmt, err = gmf.NewInputCtx(srcFileName)

	v.decStream, err = v.decFmt.GetBestStream(gmf.AVMEDIA_TYPE_VIDEO)
	if err != nil {
		log.Println("Couldn't find stream information ", err)
		return nil
	}
	//v.decFmt.DumpAv()

	/*codec, err := gmf.FindDecoder(v.decStream.CodecCtx().Id())
	if err != nil {
		log.Println("Unsupported Codec:", err)
	}*/

	codec, err := gmf.FindEncoder(gmf.AV_CODEC_ID_RAWVIDEO)
	if err != nil {
		log.Println("Unsupported Codec:", err)
	}

	//kopiere den codec context vom original codec....
	v.decCodec = gmf.NewCodecCtx(codec)

	v.decCodec.
		SetPixFmt(gmf.AV_PIX_FMT_BGR32).
		SetWidth(v.decStream.CodecCtx().Width()).
		SetHeight(v.decStream.CodecCtx().Height()).
		SetTimeBase(gmf.AVR{Num: 1, Den: 1})

	if codec.IsExperimental() {
		v.decCodec.SetStrictCompliance(gmf.FF_COMPLIANCE_EXPERIMENTAL)
	}

	if err := v.decCodec.Open(nil); err != nil {
		log.Println("error open decCodec!:", err)
	}

	v.swsCtx = gmf.NewSwsCtx(v.decStream.CodecCtx(), v.decCodec, gmf.SWS_BILINEAR)

	v.decFrame = gmf.NewFrame().
		SetWidth(v.decCodec.Width()).
		SetHeight(v.decCodec.Height()).
		SetFormat(gmf.AV_PIX_FMT_BGR32) // see above

	if err := v.decFrame.ImgAlloc(); err != nil {
		log.Fatal("ImgAlloc: ", err)
	}

	return v
}

func (v *Video) Play(s *Slide) {

	log.Println("Start video playing->")
	for packet := range v.decFmt.GetNewPackets() {
		if packet.StreamIndex() != v.decStream.Index() {
			continue
		}

		stream, err := v.decFmt.GetStream(packet.StreamIndex())

		if err != nil {
			log.Println("Error on get stream from decFmt :", err)
		}

		for frame := range packet.Frames(stream.CodecCtx()) {
			v.swsCtx.Scale(frame, v.decFrame)

			if p, err := v.decFrame.Encode(v.decCodec); p != nil {

				width, height := frame.Width(), frame.Height()
				img := new(image.RGBA)
				img.Pix = p.Data()
				img.Stride = 4 * width // 4 bytes per pixel (RGBA), width times per row
				img.Rect = image.Rect(0, 0, width, height)
				time.Sleep(30 * time.Millisecond)
				s.imageMux.Lock()
				s.img = img
				s.imageMux.Unlock()
				s.gotNewFrame <- true

				defer gmf.Release(p)
			} else if err != nil {
				log.Println("error decoding frame:", err)
				panic("FEHLER")
			}

		}

		gmf.Release(packet)

	}

	log.Println("playing video is finished!.. simple try restart...")
	r := gmf.RescaleQ(0, gmf.AV_TIME_BASE_Q, v.decStream.TimeBase())

	err := v.decFmt.SeekFrameAt(r, v.decStream.Index())
	if err != nil {
		log.Println("error seeking!", err)
		panic("error seeking")
	} else {
		v.Play(s)
	}
}

func (v *Video) alloc() error {
	return nil
}

//GetFrame get the current Streaming Frame
func (v *Video) GetFrame() (image.Image, error) {
	return nil, nil
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
