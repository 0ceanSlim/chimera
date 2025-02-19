package main

import (
	"chimera/core"
	"chimera/src/handlers"
	"chimera/src/routes"
	"chimera/src/utils"

	"fmt"
	"net/http"
)

func main() {
	// Load Configurations
	cfg, err := utils.LoadConfig()
	if err != nil {
		fmt.Printf("Failed to load config: %v\n", err)
		return
	}

	core.RunGame() // Run the game loop

	mux := http.NewServeMux()
	// Login / Logout
	mux.HandleFunc("/login", routes.Login) // Login route
	mux.HandleFunc("/init-user", handlers.InitUser)
	mux.HandleFunc("/logout", handlers.LogoutHandler) // Logout process

	// Initialize Routes
	mux.HandleFunc("/", routes.Index)
	mux.HandleFunc("/settings", routes.Settings)

	// Function Handlers
	mux.HandleFunc("/send-signed-kind1", handlers.SendSignedKind1)

	// Serve Web Files
	// Serve specific files from the root directory
	mux.HandleFunc("/favicon.svg", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "web/favicon.svg")
	})
	// Serve static files from the /web/static directory at /static/
	staticDir := "web/static"
	mux.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir(staticDir))))

	// Serve CSS files from the /web/style directory at /style/
	styleDir := "web/style"
	mux.Handle("/style/", http.StripPrefix("/style", http.FileServer(http.Dir(styleDir))))

	fmt.Printf("Server is running on http://localhost:%d\n", cfg.Port)
	http.ListenAndServe(fmt.Sprintf(":%d", cfg.Port), mux)
}
