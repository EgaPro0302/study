package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"study/internal/handler"
	"study/internal/repository"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	ctx := context.Background()
	r := mux.NewRouter()
	BD := os.Getenv("BD_URL")
	pool, err := pgxpool.New(ctx, BD)
	if err != nil {
		fmt.Println("pool error:", err)
		return
	}
	if err := pool.Ping(ctx); err != nil {
		fmt.Println("error Ping:", err)
		return
	}
	fmt.Println("BD OK")
	repo := repository.NewNoteRepository(pool)
	hand := handler.NewNoteHandler(repo)
	r.HandleFunc("/notes", hand.Create).Methods("POST")
	r.HandleFunc("/notes", hand.GetAll).Methods("GET")
	r.HandleFunc("/notes/{id}", hand.GetById).Methods("GET")
	r.HandleFunc("/notes/{id}", hand.DeleteById).Methods("DELETE")
	if err := http.ListenAndServe(":8089", r); err != nil {
		fmt.Println("error:", err)
		return
	}
}
