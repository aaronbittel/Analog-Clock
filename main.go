package main

import (
	"fmt"
	"math"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	WIDTH  = 600
	HEIGHT = WIDTH

	RADIUS   = WIDTH * 0.45
	CENTER_X = WIDTH / 2
	CENTER_Y = HEIGHT / 2

	FONTSIZE = 32
)

func main() {
	rl.InitWindow(WIDTH, HEIGHT, "Analog-Clock")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	showDigitalTime := false

	for !rl.WindowShouldClose() {
		now := time.Now()

		if rl.GetCharPressed() == rune('?') {
			showDigitalTime = !showDigitalTime
		}

		rl.BeginDrawing()

		rl.ClearBackground(rl.Black)
		rl.DrawCircle(CENTER_X, CENTER_Y, RADIUS, rl.RayWhite)

		DrawMinuteMarkers()
		DrawHourNumbers()
		DrawHourHand(now.Hour(), now.Minute(), now.Hour())
		DrawMinuteHand(now.Minute(), now.Second())
		DrawSecondHand(now.Second())

		if showDigitalTime {
			text := now.Format(time.TimeOnly)
			size := rl.MeasureText(text, FONTSIZE)
			rl.DrawText(text, CENTER_X-size/2, CENTER_Y-0.4*RADIUS-FONTSIZE/2, FONTSIZE, rl.DarkGray)
		}

		rl.EndDrawing()
	}
}

func DrawMinuteMarkers() {
	const (
		BORDER_OFFSET     = 5
		LENGTH            = 20
		HOUR_EXTRA_LENGTH = 15
		THINKNESS         = 3

		OUTER_DISTANCE = RADIUS - BORDER_OFFSET
	)

	isHourMarker := func(i int) bool {
		return i%(60/12) == 0
	}

	for angle, i := 0, 0; angle <= 360; angle, i = angle+360/60, i+1 {
		outerX := float32(OUTER_DISTANCE * math.Cos(float64(angle)*rl.Deg2rad))
		outerY := float32(OUTER_DISTANCE * math.Sin(float64(angle)*rl.Deg2rad))

		inner_distance := OUTER_DISTANCE - LENGTH
		if isHourMarker(i) {
			inner_distance -= HOUR_EXTRA_LENGTH
		}

		innerX := float32(inner_distance * math.Cos(float64(angle)*rl.Deg2rad))
		innerY := float32(inner_distance * math.Sin(float64(angle)*rl.Deg2rad))

		startV := rl.NewVector2(CENTER_X+outerX, CENTER_Y-outerY)
		endV := rl.NewVector2(CENTER_X+innerX, CENTER_Y-innerY)

		rl.DrawLineEx(startV, endV, THINKNESS, rl.Black)
	}

}

func DrawSecondHand(second int) {
	const (
		RADIUS_PERC = 0.98
		THINKNESS   = 5
	)

	angle := -(360*second/60 - 90)
	x := float32((RADIUS * RADIUS_PERC) * math.Cos(float64(angle)*rl.Deg2rad))
	y := float32((RADIUS * RADIUS_PERC) * math.Sin(float64(angle)*rl.Deg2rad))

	endV := rl.NewVector2(CENTER_X+x, CENTER_Y-y)
	startV := rl.NewVector2(CENTER_X, CENTER_Y)

	rl.DrawLineEx(startV, endV, THINKNESS, rl.Green)
}

func DrawMinuteHand(minute, second int) {
	const RADIUS_PERC = 0.85

	secondProgress := 360 / 60 * float32(second) / 60

	angle := -(360*float32(minute)/60 - 90) - secondProgress
	x := float32((RADIUS * RADIUS_PERC) * math.Cos(float64(angle)*rl.Deg2rad))
	y := float32((RADIUS * RADIUS_PERC) * math.Sin(float64(angle)*rl.Deg2rad))

	endV := rl.NewVector2(CENTER_X+x, CENTER_Y-y)
	startV := rl.NewVector2(CENTER_X, CENTER_Y)

	rl.DrawLineEx(startV, endV, 5, rl.Blue)
}

func DrawHourHand(hour, minute, second int) {
	const RADIUS_PERC = 0.65

	minuteProgress := 360 / 60 * float32(minute) / 60
	secondProgress := 360 / 60 * float32(minute) / 60

	angle := -(360*float32(hour)/12 - 90) - minuteProgress - secondProgress
	x := float32((RADIUS * RADIUS_PERC) * math.Cos(float64(angle)*rl.Deg2rad))
	y := float32((RADIUS * RADIUS_PERC) * math.Sin(float64(angle)*rl.Deg2rad))

	endV := rl.NewVector2(CENTER_X+x, CENTER_Y-y)
	startV := rl.NewVector2(CENTER_X, CENTER_Y)

	rl.DrawLineEx(startV, endV, 5, rl.Red)
}

func DrawHourNumbers() {
	for angle, hour := -60, 0; angle <= 270; angle, hour = angle+360/12, hour+1 {
		x := int32(RADIUS * 0.75 * math.Cos(float64(angle)*rl.Deg2rad))
		y := int32(RADIUS * 0.75 * math.Sin(float64(angle)*rl.Deg2rad))
		hourStr := fmt.Sprintf("%d", hour+1)
		size := rl.MeasureText(hourStr, FONTSIZE)
		rl.DrawText(hourStr, CENTER_X+x-size/2, CENTER_Y+y-FONTSIZE/2, FONTSIZE, rl.Black)
	}

}
