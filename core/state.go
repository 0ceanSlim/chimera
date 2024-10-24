package core

import (
	"encoding/json"
	"os"
)

type GameState struct {
    PlayerX, PlayerY float64
    CurrentMap       string // Add map name if multiple maps exist
}

func SaveState(player *Player, currentMap string) error {
    state := GameState{
        PlayerX:   player.X,
        PlayerY:   player.Y,
        CurrentMap: currentMap,
    }

    file, err := os.Create("save.json")
    if err != nil {
        return err
    }
    defer file.Close()

    encoder := json.NewEncoder(file)
    return encoder.Encode(state)
}

func LoadState(player *Player) (*GameState, error) {
    file, err := os.Open("save.json")
    if err != nil {
        return nil, err
    }
    defer file.Close()

    var state GameState
    decoder := json.NewDecoder(file)
    err = decoder.Decode(&state)
    if err != nil {
        return nil, err
    }

    // Set the player's position
    player.X = state.PlayerX
    player.Y = state.PlayerY

    return &state, nil
}
