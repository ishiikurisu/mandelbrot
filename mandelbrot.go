package mandelbrot

import (
    "math"
)

func FirstSetting(height, width float64) (float64, float64, float64, float64) {
    frameInitX := -width / height
    frameInitY := -1.0
    frameEndX := width / height
    frameEndY := 1.0

    if (height > width) {
        frameInitX = -1.0
        frameInitY = -height / width
        frameEndX = 1.0
        frameEndY = height / width
    }

    return frameInitX, frameInitY, frameEndX, frameEndY
}

func Mandelbrot(frameInitX, frameInitY, frameEndX, frameEndY, height, width float64) [][]bool {
    dx := (frameEndX - frameInitX) / width
    dy := (frameEndY - frameInitY) / height

    fractal := make([][]bool, int(width))
    for i := 0; i < int(width); i++ {
        fractal[i] = make([]bool, int(height))
    }

    for y := frameInitY; y < frameEndY; y += dy {
        for x := frameInitX; x < frameEndX; x += dx {
            fractal = setPixel(fractal, frameInitX, frameInitY, frameEndX, frameEndY, height, width, x, y, Escapes(x, y))
        }
    }

    return fractal
}

func Escapes(cx, cy float64) bool {
    zx := 0.0
    zy := 0.0
    maxItr := 10000  // TODO turn this into a constant

    for i:= 0; i < maxItr && math.Hypot(zx, zy) <= 2.0; i++ {
        temp := zx*zx - zy*zy + cx
        zy = 2*zx*zy + cy
        zx = temp
    }

    return math.Hypot(zx, zy) > 2
}

func setPixel(fractal [][]bool, frameInitX, frameInitY, frameEndX, frameEndY, height, width, px, py float64, state bool) [][]bool {
    x := ((px - frameInitX) * width) / (frameEndX - frameInitX)
    y := ((py - frameInitY) * height) / (frameEndY - frameInitY)

    i := int(math.Floor(x))
    j := int(math.Floor(y))

    fractal[i][j] = state

    return fractal
}
