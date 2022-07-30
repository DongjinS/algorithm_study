package main

import (
	"fmt"
	"math"
)

type planetS struct {
	x             float64
	y             float64
	r             float64
	distFromStart float64
	distFromEnd   float64
}

func (p *planetS) distCal(sx, sy, ex, ey float64) {
	p.distFromStart = math.Sqrt((p.x-sx)*(p.x-sx) + (p.y-sy)*(p.y-sy))
	p.distFromEnd = math.Sqrt((p.x-ex)*(p.x-ex) + (p.y-ey)*(p.y-ey))
	return
}

func (p *planetS) check() bool {
	if p.distFromStart < p.r && p.distFromEnd < p.r {
		return false
	}
	if p.distFromStart > p.r && p.distFromEnd > p.r {
		return false
	}
	return true
}

func main() {
	var n, planetN int
	var startX, startY, endX, endY float64
	fmt.Scanf("%d", &n)

	answer := make([]int, n)
	for i := 0; i < n; i++ {
		func() {
			fmt.Scanf("%f %f %f %f", &startX, &startY, &endX, &endY)
			if startX < endX {
				tmpX := startX
				tmpY := startY
				startX = endX
				startY = endY
				endX = tmpX
				endY = tmpY
			}
			fmt.Scanf("%d", &planetN)
			for j := 0; j < planetN; j++ {
				p := &planetS{}
				fmt.Scanf("%f %f %f", &p.x, &p.y, &p.r)
				p.distCal(startX, startY, endX, endY)
				if p.check() {
					answer[i] += 1
				}
			}
			return
		}()
	}
	for i := 0; i < n; i++ {
		fmt.Println(answer[i])
	}
}
