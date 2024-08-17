package main

import (
	"math/rand/v2"

	"github.com/eihigh/miniten"
)

var (
	x             = 200.0
	y             = 150.0
	vy            = 0.0 // Velocity of y (速度のy成分) の略
	g             = 0.1 // Gravity (重力加速度) の略
	jump          = -4.0
	frames        = 0       // 経過フレーム数
	interval      = 120     // 壁の追加間隔
	wallStartX    = 640     // 壁の初期X座標
	wallXs        = []int{} // 壁のX座標
	wallWidth     = 20      // 壁の幅
	wallHeight    = 360     // 壁の高さ
	holeYs        = []int{} // 穴のY座標
	holeYMax      = 150     // 穴のY座標の最大値
	holeHeight    = 240     // 穴のサイズ（高さ）
	playerWidth   = 100
	playerHeight  = 100
	scene         = "title"
	score         = 0
	isPrevClicked = false
	isJustClicked = false
)

func main() {
	miniten.Run(draw)
}

func draw() {
	isJustClicked = miniten.IsClicked() && !isPrevClicked
	isPrevClicked = miniten.IsClicked()

	switch scene {
	case "title":
		drawTitle()
	case "game":
		drawGame()
	case "gameOver":
		drawGameOver()
	}
}

func drawTitle() {
	miniten.DrawImage("public/sky.png", 0, 0)
	miniten.Println("クリックしてスタート")
	miniten.DrawImage("public/player.png", int(x), int(y))
	if isJustClicked {
		scene = "game"
	}
}

func drawGame() {
	miniten.DrawImage("public/sky.png", 0, 0)
	for i, wallX := range wallXs {
		if wallX < int(x) {
			score = i + 1
		}
	}
	miniten.Println("Score: ", score)

	if miniten.IsClicked() {
		vy = jump
	}
	vy += g
	y += vy
	miniten.DrawImage("public/player.png", int(x), int(y))

	frames += 1
	// interval フレームごとに壁を追加
	if frames%interval == 0 {
		wallXs = append(wallXs, wallStartX)
		holeYs = append(holeYs, rand.N(holeYMax))
	}

	// 壁を左に移動
	for i := range wallXs {
		wallXs[i] -= 2
	}

	for i := range wallXs {
		wallX := wallXs[i]
		holeY := holeYs[i]
		miniten.DrawImage("public/wall.png", wallX, holeY-wallHeight)

		miniten.DrawImage("public/wall.png", wallX, holeY+holeHeight)

		playerLeft := int(x)
		playerTop := int(y)
		playerRight := int(x) + playerWidth
		playerBottom := int(y) + playerHeight

		aboveWallLeft := wallX
		aboveWallTop := holeY - wallHeight
		aboveWallRight := wallX + wallWidth
		aboveWallBottom := holeY

		if playerLeft < aboveWallRight &&
			playerRight > aboveWallLeft &&
			playerTop < aboveWallBottom &&
			playerBottom > aboveWallTop {
			scene = "gameOver"
		}

		belowWallLeft := wallX
		belowWallTop := holeY + holeHeight
		belowWallRight := wallX + wallWidth
		belowWallBottom := holeY + holeHeight + wallHeight

		if playerLeft < belowWallRight &&
			playerRight > belowWallLeft &&
			playerTop < belowWallBottom &&
			playerBottom > belowWallTop {

			scene = "gameOver"
		}
	}

	if y < 10 {

		scene = "gameOver"
	}

	if y > 360 {

		scene = "gameOver"
	}
}

func drawGameOver() {
	miniten.DrawImage("public/sky.png", 0, 0)

	miniten.DrawImage("public/player.png", int(x), int(y))

	for i := range wallXs {
		wallX := wallXs[i]
		holeY := holeYs[i]

		miniten.DrawImage("public/wall.png", wallX, holeY-wallHeight)
		miniten.DrawImage("public/wall.png", wallX, holeY+holeHeight)
	}

	miniten.Println("Game Over")
	miniten.Println("Score: ", score)

	if isJustClicked {
		scene = "title"
		x = 200.0
		y = 150.0
		vy = 0.0
		frames = 0
		wallXs = []int{}
		holeYs = []int{}
		score = 0
	}
}
