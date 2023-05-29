package main

import (
	"fmt"
	"math/rand"
	"strconv"
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

	var Score = 0
	
	var enemyInt int = 7
	var currentX int = 10
	for enemyInt != 0{
		enemyInt--
		currentEnemy := Enemy{int32(currentX),100,true,true}
		Enemies = append(Enemies,currentEnemy)
		currentX+=75
	}
	var totalEnemies int = 7
	var EnemySpeed int32 = 1
	var Game bool = false

	for !rl.WindowShouldClose(){

		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)
		rl.DrawText("Score: "+ strconv.Itoa(Score),0,0,20,rl.LightGray)
		rl.DrawTexture(Ship,xCoords,yCoords,rl.White)

		if(totalEnemies == 0){
			Game = true
		}
		if(Game){
			Enemies = nil
			bullets = nil
			rl.UnloadTexture(Ship)
			xCoords = 1000000000
			rl.DrawText("Your Final Score: "+ strconv.Itoa(Score),30,80,40,rl.LightGray)
			rl.DrawText("Press enter to restart: "+ strconv.Itoa(Score),30,200,40,rl.LightGray)
			if rl.IsKeyPressed(rl.KeyEnter){
				rl.CloseWindow()
				main()
			}else{}

		}
		if rl.IsKeyPressed(rl.KeySpace)&&shoot{
			currentBullet := Bullet{int32(xCoords+30),int32(yCoords+25),5,float32(10),true,rl.White}
			bullets = append(bullets, currentBullet)
			shoot = false
		}
		if rl.IsKeyDown(rl.KeyD) || rl.IsKeyDown(rl.KeyRight){
			if(xCoords+1 >= screenWidth - 20){
			}else{
				xCoords+=5
			}
		}
		if rl.IsKeyDown(rl.KeyA)|| rl.IsKeyDown(rl.KeyLeft){
			if(xCoords-1 <= 0){
			}else{
				xCoords-=5
			}
		}

		for index,currentEnemy := range Enemies {
			if(Enemies[index].draw){
				var enemyShoot int32 = int32(rand.Intn(200))
				if(enemyShoot==10){
					currentBullet := Bullet{int32(currentEnemy.posX+32),int32(currentEnemy.posY),-5,float32(10),true,rl.Red}
					bullets = append(bullets, currentBullet)
				}
				if(currentEnemy.image_down){
					rl.DrawTexture(EnemyDown,currentEnemy.posX,currentEnemy.posY,rl.White)
					Enemies[index].image_down = false
				}else{
					rl.DrawTexture(EnemyUp,currentEnemy.posX,currentEnemy.posY,rl.White)
					Enemies[index].image_down = true
				}
				if(currentEnemy.posX==0 || currentEnemy.posX==screenWidth){
					for i,_ := range Enemies{
						Enemies[i].posY +=5
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
						if rl.CheckCollisionRecs(rl.NewRectangle(float32(currentBullet.posX), float32(currentBullet.posY), float32(currentBullet.radius), float32(currentBullet.radius)),
						rl.NewRectangle(float32(currentEnemy.posX), float32(currentEnemy.posY),float32(60),float32(32)))&&currentBullet.Color==rl.White&&currentEnemy.draw{
							Enemies[index].draw = false
							Score += 500
						}
						if rl.CheckCollisionRecs(rl.NewRectangle(float32(currentBullet.posX), float32(currentBullet.posY), float32(currentBullet.radius), float32(currentBullet.radius)),
						rl.NewRectangle(float32(xCoords), float32(yCoords),float32(60),float32(32)))&&currentBullet.Color==rl.Red{
							fmt.Println("GAME OVER")

							Game = true
						}
					}
				}
				if(currentEnemy.posX <= 0){
					EnemySpeed = 1
				}else if(currentEnemy.posX>=screenWidth - 20){
					EnemySpeed = -1
				}else{}

				Enemies[index].posX += EnemySpeed
				time.Sleep(5000000)
			}
		}

		rl.EndDrawing()
	}
	rl.CloseWindow()

}