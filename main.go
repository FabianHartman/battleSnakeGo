package main

import (
	"battleSnakeGo/models"
	"encoding/json"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/snake", snake)
	http.HandleFunc("/snake/start", startGame)
	http.HandleFunc("/snake/move", makeMove)
	http.HandleFunc("/snake/end", endGame)
	http.HandleFunc("/snake/playedGames", playedGames)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

func snake(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case "GET":
		getSnake(w)
	case "PUT":
		updateSnake(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func getSnake(w http.ResponseWriter) {
	err := json.NewEncoder(w).Encode(models.CurrentSnake)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

type UpdateSnakeReq struct {
	Color string `json:"color"`
	Head  string `json:"head"`
	Tail  string `json:"tail"`
}

func updateSnake(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "failed to read body", http.StatusBadRequest)

		return
	}
	defer r.Body.Close()

	var req UpdateSnakeReq
	err = json.Unmarshal(body, &req)
	if err != nil {
		http.Error(w, "invalid JSON: "+err.Error(), http.StatusBadRequest)

		return
	}

	models.CurrentSnake.Customizations.Color = req.Color
	models.CurrentSnake.Customizations.Head = req.Head
	models.CurrentSnake.Customizations.Tail = req.Tail

	err = json.NewEncoder(w).Encode(models.CurrentSnake)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func startGame(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "failed to read body", http.StatusBadRequest)

		return
	}
	defer r.Body.Close()

	var req models.GameRequest
	err = json.Unmarshal(body, &req)
	if err != nil {
		http.Error(w, "invalid JSON: "+err.Error(), http.StatusBadRequest)

		return
	}

	game := models.NewGame(req.Game.Id, req.Game.Map, req.Board.Height, req.Board.Width)

	game.Moves = append(game.Moves, *req.You.Head)
}

type MoveResponse struct {
	Move string `json:"move"`
}

func makeMove(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "failed to read body", http.StatusBadRequest)

		return
	}
	defer r.Body.Close()

	var req models.GameRequest
	err = json.Unmarshal(body, &req)
	if err != nil {
		http.Error(w, "invalid JSON: "+err.Error(), http.StatusBadRequest)

		return
	}

	game := models.SearchGameByID(req.Game.Id)
	if game == nil {
		http.Error(w, "game not found", http.StatusNotFound)

		return
	}

	move := req.GenerateMove()

	game.Moves = append(game.Moves, *req.You.Head)

	err = json.NewEncoder(w).Encode(MoveResponse{Move: move})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func endGame(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "failed to read body", http.StatusBadRequest)

		return
	}
	defer r.Body.Close()

	var req models.GameRequest
	err = json.Unmarshal(body, &req)
	if err != nil {
		http.Error(w, "invalid JSON: "+err.Error(), http.StatusBadRequest)

		return
	}

	game := models.SearchGameByID(req.Game.Id)
	if game == nil {
		http.Error(w, "game not found", http.StatusNotFound)

		return
	}

	game.Finished = true

	game.Score = len(req.You.Body)

	err = json.NewEncoder(w).Encode(game)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func playedGames(w http.ResponseWriter, r *http.Request) {
	err := json.NewEncoder(w).Encode(models.AllGames)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
