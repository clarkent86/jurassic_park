package main

import (
	"encoding/json"
	"net/http"

	"github.com/clarkent86/jurassic_park/internal/cage"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	sugar := logger.Sugar()

	sugar.Infow("Starting Gorilla/Mux Jurassic Park app")

	var park cage.Park

	r := mux.NewRouter()

	r.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	})

	park.AddDinosaurToCageHandler("/add/dinosaur").AddRoute(r)

	park.NewCageHandler("/api/v1/add/cage").AddRoute(r)
	park.ToggleCageHandler("/api/v1/togglePower").AddRoute(r)

	sugar.Fatal(http.ListenAndServe(":8080", r))
}
