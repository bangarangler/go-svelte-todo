package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/bangarangler/go-svelte-todo/gqlgen"
	"github.com/bangarangler/go-svelte-todo/pg"
)

func main() {
	db, err := pg.Open(pg.PgConnStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	repo := pg.NewRepository(db)

	mux := http.NewServeMux()
	mux.Handle("/", gqlgen.NewPlaygroundHandler("/query"))
	mux.Handle("/query", gqlgen.NewHandler(repo))

	port := ":8080"
	fmt.Fprintf(os.Stdout, "🚀 Server ready at http://localhost%s\n", port)
	fmt.Fprintln(os.Stderr, http.ListenAndServe(port, mux))
}
