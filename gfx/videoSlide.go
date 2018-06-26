package gfx

import (
	_ "image/jpeg" // Import JPEG Decoding
	_ "image/png"  // Import PNG Decoding
)

//VideoSlide A Simple VideoSlide element
type VideoSlide struct {
	*MediaSlide
	video *Video

	gotNewFrame   chan bool
	finishedVideo chan bool
}

//NewSlideForVideo Create a new Slide for Video content
func NewSlideForVideo(path, uid string) *VideoSlide {

	s := createSlide(uid, true)

	vs := &VideoSlide{
		gotNewFrame:   make(chan bool),
		finishedVideo: make(chan bool),
	}
	vs.MediaSlide = s

	vs.video = CreateVideo(path, vs)
	return vs
}

func (s *VideoSlide) BackgroundThread() {
	go func() {
		defer s.video.CleanUP()
		s.video.LoopPlay()
	}()
}

func (s *VideoSlide) Play() {
	if s.IsVideo {
		select {
		case newFrame := <-s.gotNewFrame:
			if newFrame == true {
				s.Tex.SetDefaultImage(s.img)
			}
		default:
		}
	}
}

func (s *VideoSlide) GoToNextSlide(currentDuration float64) bool {
	select {
	case stop := <-s.finishedVideo:
		if stop == true {
			return true
		}
	default:
	}

	return false
}
