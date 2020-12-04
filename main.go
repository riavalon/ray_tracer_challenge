package main

import (
	"io/ioutil"
	"math"

	"github.com/riavalon/ray_tracer/canvas"
	tuples "github.com/riavalon/ray_tracer/tuples"
)

type environment struct {
	wind    *tuples.Tuple
	gravity *tuples.Tuple
}

type projectile struct {
	position *tuples.Tuple
	velocity *tuples.Tuple
	colour   canvas.Colour
}

func tick(env environment, proj projectile) projectile {
	position := proj.position.Add(proj.velocity)
	velocity := proj.velocity.Add(env.gravity.Add(env.wind))
	return projectile{position: position, velocity: velocity, colour: proj.colour}
}

func main() {
	c := canvas.NewCanvas(900, 550)

	env := environment{
		gravity: tuples.CreateVector(0, -0.1, 0),
		wind:    tuples.CreateVector(-0.01, 0, 0),
	}

	p := projectile{
		position: tuples.CreatePoint(0, 1, 0),
		velocity: tuples.NormalizeVector(tuples.CreateVector(1, 1, 0)).MultiplyByScalar(9.95),
		colour:   canvas.NewColour(1, 0, 0),
	}
	initProj := tick(env, p)

	for {
		initProj = tick(env, initProj)

		// get the x y coordinates and map them to the canvas
		projY := int(math.Round(initProj.position.Y))
		x := int(math.Round(initProj.position.X))
		y := c.Height - projY

		c.WritePixel(x, y, canvas.NewColour(1, 0, 0))

		if initProj.position.Y <= 0.0 {
			break
		}
	}

	data := []byte(c.ToPPM())
	err := ioutil.WriteFile("./test.ppm", data, 0644)
	if err != nil {
		panic(err)
	}
}
