package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)
func main(){

	screenWidth := int32 (600)
	screenHeight := int32 (900)

	var xCoords int32 = 2
	var yCoords int32 = 800 

	rl.InitWindow(screenWidth,screenHeight,"Space Invaders")

	ShipImg := rl.LoadImage("assets/Ship.png")
	Ship := rl.LoadTextureFromImage(ShipImg)



	for !rl.WindowShouldClose(){

		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)
		rl.DrawTexture(Ship,xCoords,yCoords,rl.White)
		if rl.IsKeyDown(rl.KeyD){
			if(xCoords+1 >= screenWidth){
			}else{
				xCoords+=5
			}
		}
		if rl.IsKeyDown(rl.KeyA){
			if(xCoords+1 >= screenWidth){
			}else{
				xCoords+=5
			}
		}


		rl.EndDrawing()

	}
	rl.CloseWindow()

}