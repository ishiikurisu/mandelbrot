package main

import (
    "fmt"
    "github.com/ishiikurisu/mandelbrot"
    "github.com/ishiikurisu/mandelbrot/processing"
)

func main() {
    height := mandelbrot.NewFloat(900.0)
    width := mandelbrot.NewFloat(1600.0)
    var noIterations int = 100
    frameInitX, frameInitY, frameEndX, frameEndY := mandelbrot.FirstSetting(height, width)

    for i := 0; i < noIterations; i++ {
        // saving image
        fractal := mandelbrot.Mandelbrot(frameInitX, frameInitY, frameEndX, frameEndY, height, width)
        img := processing.FractalToImage(fractal)
        fn := fmt.Sprintf("./out/m%03d.png", i)
        processing.SaveImage(img, fn)

        // zooming
        targetWidth := mandelbrot.NewFloat(10.0)
        targetHeight := mandelbrot.NewFloat(2.0)
        targetWidth.Quo(width, targetWidth)
        targetHeight.Quo(height, targetHeight)
        frameInitX, frameInitY, frameEndX, frameEndY = mandelbrot.ZoomAt(
            targetWidth, targetHeight, width, height,frameInitX, frameInitY, frameEndX, frameEndY)
    }
}
