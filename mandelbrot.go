package mandelbrot

import (
    "math/big"
)

func NewFloat(f float64) *big.Float {
    return new(big.Float).SetPrec(128).SetFloat64(f)
}

func Atof(a string) *big.Float {
    f := NewFloat(0.0)
    f.Parse(a, 10)
    return f
}

func FirstSetting(height, width *big.Float) (*big.Float, *big.Float, *big.Float, *big.Float) {
    acc := NewFloat(-1.0)
    frameInitX := NewFloat(0.0)
    frameInitY := NewFloat(0.0)
    frameEndX := NewFloat(0.0)
    frameEndY := NewFloat(0.0)

    // frameInitX = - width / height
    acc.Neg(width)
    acc.Quo(acc, height)
    frameInitX.Copy(acc)
    frameInitY = NewFloat(-1.0)
    // frameEndX = width / height
    acc.Quo(width, height)
    frameEndX.Copy(acc)
    frameEndY = NewFloat(1.0)

    if (height.Cmp(width) == 1) {  // height > width
        frameInitX = NewFloat(-1.0)
        // frameInitY = - height / width
        acc = acc.Neg(height)
        acc.Quo(acc, width)
        frameInitY.Copy(acc)
        frameEndX = NewFloat(1.0)
        // frameEndY = height / width
        acc.Quo(height, width)
        frameEndY.Copy(acc)
    }

    return frameInitX, frameInitY, frameEndX, frameEndY
}

func Mandelbrot(frameInitX, frameInitY, frameEndX, frameEndY, height, width *big.Float) [][]bool {
    acc := NewFloat(0.0)
    dx := NewFloat(0.0)
    dy := NewFloat(0.0)

    // dx = (frameEndX - frameInitX) / width
    acc.Sub(frameEndX, frameInitX)
    acc.Quo(acc, width)
    dx.Copy(acc)

    // dy = (frameEndY - FrameInitY) / height
    acc.Sub(frameEndY, frameInitY)
    acc.Quo(acc, height)
    dy.Copy(acc)

    intWidth, _ := width.Int64()
    intHeight, _ := height.Int64()
    fractal := make([][]bool, intWidth)
    var i int64
    for i = 0; i < intWidth; i++ {
        fractal[i] = make([]bool, intHeight)
    }

    y := NewFloat(0.0)
    x := NewFloat(0.0)
    for y.Copy(frameInitY); y.Cmp(frameEndY) == -1; y.Add(y, dy) {
        for x.Copy(frameInitX); x.Cmp(frameEndX) == -1; x.Add(x, dx) {
            escapes := Escapes(x, y)
            fractal = setPixel(fractal, frameInitX, frameInitY, frameEndX, frameEndY, height, width, x, y, escapes)
        }
    }

    return fractal
}

func Escapes(cx, cy *big.Float) bool {
    zx := NewFloat(0.0)
    zy := NewFloat(0.0)
    two := NewFloat(2.0) // TODO turn this into a constant
    limit := two.Copy(two)
    maxItr := 1024  // TODO turn this into a constant


    for i := 0; i < maxItr && Pythagoras(zx, zy).Cmp(limit) == -1; i++ {
        acc := NewFloat(0.0)
        temp := NewFloat(0.0)
        zx2 := NewFloat(0.0)
        zy2 := NewFloat(0.0)

        // zx = zx*zx - zy*zy + cx
        zx2.Mul(zx, zx)
        zy2.Mul(zy, zy)
        acc.Sub(zx2, zy2)
        acc.Add(acc, cx)
        temp.Copy(acc)

        // zy = 2 * zx * zy + cy
        acc.Mul(two, zx)
        acc.Mul(acc, zy)
        acc.Add(acc, cy)

        zx.Copy(temp)
        zy.Copy(acc)
    }

    return Pythagoras(zx, zy).Cmp(two) == 1
}

func Pythagoras(x, y *big.Float) *big.Float {
    x2 := NewFloat(0.0)
    y2 := NewFloat(0.0)

    x2.Mul(x, x)
    y2.Mul(y, y)
    x2.Add(x2, y2)

    return x2.Sqrt(x2)
}

func setPixel(fractal [][]bool, frameInitX, frameInitY, frameEndX, frameEndY, height, width, px, py *big.Float, state bool) [][]bool {
    accA := NewFloat(0.0)
    accB := NewFloat(0.0)
    x := NewFloat(0.0)
    y := NewFloat(0.0)

    // float posX = ((px - FRAME_INIT_X) * width) / (FRAME_END_X - FRAME_INIT_X);
    accA.Sub(px, frameInitX)
    accA.Mul(accA, width)
    accB.Sub(frameEndX, frameInitX)
    accA.Quo(accA, accB)
    x.Copy(accA)

    // float posY = ((py - FRAME_INIT_Y) * height) / (FRAME_END_Y - FRAME_INIT_Y);
    accA.Sub(py, frameInitY)
    accA.Mul(accA, height)
    accB.Sub(frameEndY, frameInitY)
    accA.Quo(accA, accB)
    y.Copy(accA)

    i, _ := x.Int64()
    j, _ := y.Int64()
    fractal[i][j] = state

    return fractal
}

func ZoomAt(posX, posY, width, height, frameInitX, frameInitY, frameEndX, frameEndY *big.Float) (*big.Float, *big.Float, *big.Float, *big.Float) {
    factor := NewFloat(0.9)  // TODO turn this into a constant
    two := NewFloat(2.0)  // TODO turn this into a constant
    acc := NewFloat(0.0)
    centerX := NewFloat(0.0)
    centerY := NewFloat(0.0)
    distX := NewFloat(0.0)
    distY := NewFloat(0.0)

    // float centerX = FRAME_INIT_X + (posX * (FRAME_END_X - FRAME_INIT_X)) / (width);
    acc.Sub(frameEndX, frameInitX)
    acc.Mul(acc, posX)
    acc.Quo(acc, width)
    acc.Add(frameInitX, acc)
    centerX.Copy(acc)

    // float centerY = FRAME_INIT_Y - (posY * (FRAME_END_Y - FRAME_INIT_Y)) / (height);
    acc.Sub(frameEndY, frameInitY)
    acc.Mul(acc, posY)
    acc.Quo(acc, height)
    acc.Add(frameInitY, acc)
    centerY.Copy(acc)

    // float distX = FACTOR * (FRAME_END_X - FRAME_INIT_X) / 2;
    acc.Sub(frameEndX, frameInitX)
    acc.Mul(factor, acc)
    acc.Quo(acc, two)
    distX.Copy(acc)

    // float distY = FACTOR * (FRAME_END_Y - FRAME_INIT_Y) / 2;
    acc.Sub(frameEndY, frameInitY)
    acc.Mul(factor, acc)
    acc.Quo(acc, two)
    distY.Copy(acc)

    frameInitX.Sub(centerX, distX)
    frameInitY.Sub(centerY, distY)
    frameEndX.Add(centerX, distX)
    frameEndY.Add(centerY, distY)

    return frameInitX, frameInitY, frameEndX, frameEndY
}

func Follow(targetX, targetY, factor, frameInitX, frameInitY, frameEndX, frameEndY *big.Float) (*big.Float, *big.Float, *big.Float, *big.Float) {
    two := NewFloat(2.0)  // TODO turn this into a constant
    acc := NewFloat(0.0)
    halfDistX := NewFloat(0.0)
    halfDistY := NewFloat(0.0)
    centerX := NewFloat(0.0)
    centerY := NewFloat(0.0)
    newCenterX := NewFloat(0.0)
    newCenterY := NewFloat(0.0)

    // halfDistX = (frameEndX - frameInitX) / 2
    acc.Sub(frameEndX, frameInitX)
    halfDistX.Quo(acc, two)
    // halfDistY = (frameEndY - frameInitY) / 2
    acc.Sub(frameEndY, frameInitY)
    halfDistY.Quo(acc, two)
    // centerX = halfDistX + frameInitX
    centerX.Add(halfDistX, frameInitX)
    // centerY = halfDistY + frameInitY
    centerY.Add(halfDistY, frameInitY)
    // newCenterX = targetX + factor * (centerX - targetX)
    acc.Sub(centerX, targetX)
    acc.Mul(factor, acc)
    newCenterX.Add(targetX, acc)
    // newCenterY = targetY + factor * (centerY - targetY)
    acc.Sub(centerY, targetY)
    acc.Mul(factor, acc)
    newCenterY.Add(targetY, acc)
    // frameInitX = newCenterX - halfDistX
    frameInitX.Sub(newCenterX, halfDistX)
    // frameInitY = newCenterY - halfDistY
    frameInitY.Sub(newCenterY, halfDistY)
    // frameEndX = newCenterX + halfDistX
    frameEndX.Add(newCenterX, halfDistX)
    // frameEndY = newCenterY + halfDistY
    frameEndY.Add(newCenterY, halfDistY)

    return frameInitX, frameInitY, frameEndX, frameEndY
}
