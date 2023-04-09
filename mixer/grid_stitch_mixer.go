package mixer

import (
	"errors"
	"github.com/earayu/photo_manager/common"
	"image"
	"image/color"
	"image/draw"
)

type GridStitchMixerCreator struct {
	Grid [][]int
}

func (m *GridStitchMixerCreator) Create() *common.Mixer {
	return common.CreateMixer(func(imagePool []*image.Image) (*image.Image, error) {
		// Make sure we have at least two images to stitch
		if len(imagePool) < 2 {
			return nil, errors.New("need at least two images to stitch")
		}

		PhotoCountInRowSide := len(m.Grid[0])
		PhotoCountInColumnSide := len(m.Grid)

		//shuffle the images
		originImagePool := imagePool
		imagePool = common.Shuffle(imagePool)

		// Get the size of the first image and calculate the size of the stitched image
		w := (*imagePool[0]).Bounds().Max.X
		h := (*imagePool[0]).Bounds().Max.Y
		for _, img := range imagePool[1:] {
			bounds := (*img).Bounds()
			if bounds.Max.X > w {
				w = bounds.Max.X
			}
			if bounds.Max.Y > h {
				h = bounds.Max.Y
			}
		}

		// Create the output image
		outImg := image.NewRGBA(image.Rect(0, 0, w*PhotoCountInRowSide, h*PhotoCountInColumnSide))
		backgroundColor := color.RGBA{246, 129, 129, 60}

		for i := 0; i < PhotoCountInColumnSide; i++ {
			for j := 0; j < PhotoCountInRowSide; j++ {
				x := j * w
				y := i * h
				if m.Grid[i][j] == 0 {
					//draw background color
					r := image.Rect(x, y, x+w, y+h)
					draw.Draw(outImg, r, &image.Uniform{backgroundColor}, image.Point{}, draw.Src)
					continue
				}
				//poll out the first image from imagePool
				img := imagePool[0]
				imagePool = imagePool[1:]
				if len(imagePool) == 0 {
					imagePool = common.Shuffle(originImagePool)
				}
				// Draw each image onto the output image
				bounds := (*img).Bounds()
				r := image.Rect(x, y, x+bounds.Max.X, y+bounds.Max.Y)
				draw.Draw(outImg, r, *img, bounds.Min, draw.Src)
			}
		}

		// Return the new image and nil error to indicate a successful mix
		return common.RgbaToImage(outImg), nil
	})
}
