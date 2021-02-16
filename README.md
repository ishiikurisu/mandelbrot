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

Adjust the parallel execution settings for
your context by tweaking the `config.json` file.

``` json
{
    "Height": "900.0",
    "Width": "1600.0",
    "TargetX": "-1.62917",
    "TargetY": "-0.0203968",
    "Factor": "0.9",
    "MaxItr": 1024,
    "NumberOfIterations": 1000,
    "NumberOfSkips": 30,
    "TotalGroups": 3
}
```

Now, just execute the project:

``` sh
make video
```

It takes a lot of time to generate one frame, so I recommend you grab some
coffee and donuts while you wait for the task to complete.
