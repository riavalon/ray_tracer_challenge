package main

import (
	"fmt"
	"time"

	tuples "github.com/riavalon/ray_tracer/tuples"
)

type environment struct {
	wind    *tuples.Tuple
	gravity *tuples.Tuple
}

type projectile struct {
	position *tuples.Tuple
	velocity *tuples.Tuple
}

func tick(env environment, proj projectile) projectile {
	position := proj.position.Add(proj.velocity)
	velocity := proj.velocity.Add(env.gravity.Add(env.wind))
	return projectile{position, velocity}
}

func main() {
	env := environment{
		gravity: tuples.CreateVector(0, -0.1, 0),
		wind:    tuples.CreateVector(-0.01, 0, 0),
	}

	p := projectile{
		position: tuples.CreatePoint(0, 1, 0),
		velocity: tuples.NormalizeVector(tuples.CreateVector(1, 1, 0)),
	}
	initProj := tick(env, p)

	for {
		initProj = tick(env, initProj)
		fmt.Printf("Projectile Position: %v\n", initProj.position)

		if initProj.position.Y <= 0.0 {
			break
		}
		time.Sleep(time.Duration(time.Millisecond * 300))
	}

	fmt.Printf("Finished launching projectile. Final position is: %v", initProj.position)
}
