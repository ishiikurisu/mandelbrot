# Mandelbrot Fractal

This is just me having fun with the Mandelbrot fractal, my favorite piece
of maths ever.

Originally made using Processing, I switched to Go so I could faster arbitrary
precision arithmetics.

## Instructions

Create a folder named `out` in this repository's root:

``` sh
mkdir out
```

On the `main/main.go` file, adjust the parallel execution settings for
your machine. These are the variables you probably want to tweak:

``` go
height := mandelbrot.NewFloat(900.0)  // frame dimensions
width := mandelbrot.NewFloat(1600.0)
targetX := mandelbrot.Atof("-0.6567347481")  // target point
targetY := mandelbrot.Atof("0.3416144335")
factor := mandelbrot.NewFloat(0.9)  // zoom factor between frames
// ...
const noIterations int = 100  // how many frames you want
const noSkip int = 8  // how many frames to skip if some have already been generated
const totalGroups int = 6  // how many parallel executions should happen at once
```

Now, just execute the project:

``` sh
make video
```

It takes a lot of time to generate one frame, so I recommend you grab some
coffee and donuts while you wait for the task to complete.
