package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"study/internal/handler"
	"study/internal/repository"
	"study/internal/service"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	ctx := context.Background()
	r := mux.NewRouter()
	postgres := os.Getenv("DB_URL")
	pool, err := pgxpool.New(ctx, postgres)
	if err != nil {
		fmt.Println("pool err:", err)
		return
	}
	if err := pool.Ping(ctx); err != nil {
		fmt.Println("pool err:", err)
		return
	}
	fmt.Println("BD OK")
	repo := repository.NewUserRepository(pool)
	svc := service.NewUserService(repo)
	h := handler.NewUserHandler(svc)
	r.HandleFunc("/users", h.Create).Methods(http.MethodPost)
	r.HandleFunc("/users", h.Get).Methods(http.MethodGet)
	r.HandleFunc("/users/{id}", h.GetById).Methods(http.MethodGet)

	if err := http.ListenAndServe(":8089", r); err != nil {
		fmt.Println(err)
		return
	}
}
