package main

import (
	"time"

	"forest_fire/models"
)

func main() {
	forest := models.NewForest(20, 20, 1)

	// forest.Draw()

	ticker := time.NewTicker(forest.Tick * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		if forest.CanBurn() {
			ticker.Stop()
			models.GameOver()
			break
		}
		models.Clear()
		forest.UpdateForest()
		forest.Draw()
	}
}
