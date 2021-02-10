package processing

import (
    "fmt"
    "os"
    "image"
    "image/color"
    "image/png"
)

func FractalToImage(fractal [][]int, maxItr int) image.Image {
    width := len(fractal)
    height := len(fractal[0])
    upLeft := image.Point{0, 0}
    lowRight := image.Point{width, height}
    img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

    for x := 0; x < width; x++ {
        for y := 0; y < height; y++ {
            img.Set(x, y, assignColor(fractal[x][y], maxItr))
        }
    }

    return img
}

func SaveImage(img image.Image, fn string) error {
    fp, _ := os.Create(fn)
    defer fp.Close()
    return png.Encode(fp, img)
}

func assignColor(noItr, maxItr int) color.RGBA {
    black := color.RGBA{
        0x00,
        0x00,
        0x00,
        0xff,
    }

    if noItr == -1 {
        return black
    } else {
        c := uint8(float64(0xff) * (float64(maxItr) - float64(noItr)) / float64(maxItr))
        fmt.Printf("%d %d => %d %d\n", noItr, maxItr, c, 0xff)
        return color.RGBA{
            c,
            c,
            c,
            0xff,
        }
    }
}
