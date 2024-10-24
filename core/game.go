package core

import (
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
    Player       *Player
    Map          *GameMap
    Zoom         float64
    MaxZoom      float64
    MinZoom      float64
}

func (g *Game) Update() error {
    g.Player.Update() // Handle player movement

    // Capture both x and y scroll values
    _, scrollY := ebiten.Wheel()

    // Handle zooming with the vertical scroll wheel value
    if scrollY > 0 {
        g.Zoom += 0.1
    } else if scrollY < 0 {
        g.Zoom -= 0.1
    }

    // Clamp the zoom level between MinZoom and MaxZoom
    g.Zoom = math.Max(g.MinZoom, math.Min(g.Zoom, g.MaxZoom))

    return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
    // Use Bounds().Dx() and Bounds().Dy() to get screen width and height
    screenWidth := screen.Bounds().Dx()
    screenHeight := screen.Bounds().Dy()

    // Calculate offset to center the player
    offsetX := float64(screenWidth)/2 - g.Player.X * g.Zoom
    offsetY := float64(screenHeight)/2 - g.Player.Y * g.Zoom

    // Draw the map with zoom
    opts := &ebiten.DrawImageOptions{}
    opts.GeoM.Scale(g.Zoom, g.Zoom)
    opts.GeoM.Translate(offsetX, offsetY)
    screen.DrawImage(g.Map.Image, opts)

    // Draw the player on the map
    g.Player.Draw(screen, offsetX, offsetY, g.Zoom)
}


func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
    // Fixed game resolution (1920x1080)
    return 1920, 1080
}

func RunGame() {
    // Load map and player images
    playerImage, _, err := ebitenutil.NewImageFromFile("core/assets/character.png")
    if err != nil {
        log.Fatal(err)
    }

    mapImage, _, err := ebitenutil.NewImageFromFile("core/assets/map.png")
    if err != nil {
        log.Fatal(err)
    }

    // Initialize player and map
    player := NewPlayer(playerImage)
    gameMap := NewGameMap(mapImage)

    // Initialize game with default zoom settings
    game := &Game{
        Player:   player,
        Map:      gameMap,
        Zoom:     1.0,  // Default zoom level
        MaxZoom:  6.0,  // Upper zoom limit
        MinZoom:  3.0,  // Lower zoom limit
    }

    if err := ebiten.RunGame(game); err != nil {
        log.Fatal(err)
    }
}
