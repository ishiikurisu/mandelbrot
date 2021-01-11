package processing

import (
    "image"
    "image/color"
)

func FractalToImage(fractal [][]bool) image.Image {
    // TODO turn the following into constants
    white := color.RGBA{
        0xff,
        0xff,
        0xff,
        0xff,
    }
    black := color.RGBA{
        0x00,
        0x00,
        0x00,
        0xff,
    }

    width := len(fractal)
    height := len(fractal[0])
    upLeft := image.Point{0, 0}
    lowRight := image.Point{width, height}
    img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

    for x := 0; x < width; x++ {
        for y := 0; y < height; y++ {
            if fractal[x][y] {
                img.Set(x, y, white)
            } else {
                img.Set(x, y, black)
            }
        }
    }

    return img
}
