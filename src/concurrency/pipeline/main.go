package main

import "math"

func main() {
	s := []float64{7, 1, 3, 2, 9}

	p := &pipeline{
		start: sq,
		end:   display,
	}

	p.stages = append(p.stages, divideByTwo, subtractOne, round)

	p.execute(s)
}

func sq(s []float64) <-chan float64 {
	ch := make(chan float64)

	go func() {
		for _, v := range s {
			v *= v
			ch <- v
		}
		close(ch)
	}()

	return ch
}

type start func(s []float64) <-chan float64

type stage func(ch <-chan float64) <-chan float64

type end func(ch <-chan float64)

type pipeline struct {
	start
	stages []stage
	end
}

func (p pipeline) execute(s []float64) {
	ch := p.start(s)
	for _, stage := range p.stages {
		ch = stage(ch)
	}
	p.end(ch)
}

func divideByTwo(ch <-chan float64) <-chan float64 {
	divided := make(chan float64)

	go func() {
		for v := range ch {
			divided <- v / 2
		}
		close(divided)
	}()

	return divided
}

func subtractOne(ch <-chan float64) <-chan float64 {
	subtracted := make(chan float64)

	go func() {
		for v := range ch {
			subtracted <- v - 1
		}
		close(subtracted)
	}()

	return subtracted
}

func round(ch <-chan float64) <-chan float64 {
	roundedCh := make(chan float64)

	go func() {
		for v := range ch {
			roundedCh <- Round(v)
		}
		close(roundedCh)
	}()

	return roundedCh
}

func Round(val float64) float64 {
	pow10 := math.Pow(10, 2)
	digit := pow10 * val
	rounded := math.Round(digit) / pow10
	return rounded
}

func display(c <-chan float64) {
	for v := range c {
		println(int(v))
	}
}
