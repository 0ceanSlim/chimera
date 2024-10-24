package core

import "github.com/hajimehoshi/ebiten/v2"

type GameMap struct {
    Image *ebiten.Image
}

func NewGameMap(image *ebiten.Image) *GameMap {
    return &GameMap{
        Image: image,
    }
}
