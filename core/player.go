package core

import "github.com/hajimehoshi/ebiten/v2"

type Player struct {
    X, Y  float64
    Image *ebiten.Image
}

func NewPlayer(image *ebiten.Image) *Player {
    return &Player{
        X:     320, // Initial position (center of a tile)
        Y:     240,
        Image: image,
    }
}

func (p *Player) Update() {
    // Handle movement with arrow keys
    speed := 2.0
    if ebiten.IsKeyPressed(ebiten.KeyUp) {
        p.Y -= speed
    }
    if ebiten.IsKeyPressed(ebiten.KeyDown) {
        p.Y += speed
    }
    if ebiten.IsKeyPressed(ebiten.KeyLeft) {
        p.X -= speed
    }
    if ebiten.IsKeyPressed(ebiten.KeyRight) {
        p.X += speed
    }
}

func (p *Player) Draw(screen *ebiten.Image, offsetX, offsetY, zoom float64) {
    opts := &ebiten.DrawImageOptions{}
    opts.GeoM.Scale(zoom, zoom)
    opts.GeoM.Translate(p.X*zoom + offsetX, p.Y*zoom + offsetY)
    screen.DrawImage(p.Image, opts)
}
