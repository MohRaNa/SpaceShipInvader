package main

import (
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Bullet struct{
	posX int32
	posY int32
	velocity int32
	radius float32
	Draw bool
	Color rl.Color
}

type Enemy struct{
	posX int32
	posY int32
	image_down bool
	draw bool

}

func main(){

	screenWidth := int32 (600)
	screenHeight := int32 (900)

	var xCoords int32 = 2
	var yCoords int32 = 800 

	bullets := []Bullet{}

	var shoot bool = true

	rl.InitWindow(screenWidth,screenHeight,"Space Invaders")

	ShipImg := rl.LoadImage("assets/Ship.png")
	Ship := rl.LoadTextureFromImage(ShipImg)

	Enemies := []Enemy{}
	Enemy1 := rl.LoadImage("assets/Enemy.png")
	Enemy2  := rl.LoadImage("assets/Enemy2.png")
	EnemyUp := rl.LoadTextureFromImage(Enemy1)
	EnemyDown := rl.LoadTextureFromImage(Enemy2)
	
	var enemyInt int = 7
	var currentX int = 10
	for enemyInt != 0{
		enemyInt--
		currentEnemy := Enemy{int32(currentX),100,true,true}
		Enemies = append(Enemies,currentEnemy)
		currentX+=75
	}

	//var EnemySpeed int32 = 1


	for !rl.WindowShouldClose(){

		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)
		rl.DrawTexture(Ship,xCoords,yCoords,rl.White)

		for index,currentEnemy := range Enemies {
			if(Enemies[index].draw){
				if(currentEnemy.image_down){
					rl.DrawTexture(EnemyDown,currentEnemy.posX,currentEnemy.posY,rl.White)
					Enemies[index].image_down = false
				}else{
					rl.DrawTexture(EnemyUp,currentEnemy.posX,currentEnemy.posY,rl.White)
					Enemies[index].image_down = true
				}
				time.Sleep(50000000)
			}
		}

		for index1,currentBullet := range bullets{
			if(currentBullet.Draw){
				bullets[index1].posY = bullets[index1].posY - currentBullet.velocity
				rl.DrawCircle(currentBullet.posX-16, currentBullet.posY,currentBullet.radius,currentBullet.Color)
				if(currentBullet.posY<0 || currentBullet.posY>screenHeight){
					bullets[index1].Draw = false
					shoot = true
				}
			}

		}
		if rl.IsKeyDown(rl.KeySpace)&&shoot{
			currentBullet := Bullet{int32(xCoords+50),int32(yCoords+25),5,float32(10),true,rl.White}
			bullets = append(bullets, currentBullet)
			shoot = false
		}
		if rl.IsKeyDown(rl.KeyD)|| rl.IsKeyDown(rl.KeyRight){
			if(xCoords+1 >= screenWidth){
			}else{
				xCoords+=5
			}
		}
		if rl.IsKeyDown(rl.KeyA)|| rl.IsKeyDown(rl.KeyLeft){
			if(xCoords+1 >= screenWidth){
			}else{
				xCoords-=5
			}
		}

		rl.EndDrawing()

	}
	rl.CloseWindow()

}