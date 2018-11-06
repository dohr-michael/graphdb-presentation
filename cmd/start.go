package cmd

import (
	"github.com/spf13/cobra"
	"net/http"
	"log"
	"github.com/dohr-michael/graphdb-presentation/pokeapi/importer"
	"encoding/json"
)

var startCmd = &cobra.Command{
	Use: "start",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("Start graphdb-presentation")
		server := http.NewServeMux()

		r := func(fn func() interface{}) http.HandlerFunc {
			return func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				res, _ := json.Marshal(fn())
				w.Write(res)
			}
		}

		server.HandleFunc("/@/health", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`{"status": "OK"}`))
		})
		server.HandleFunc("/importer/all", r(importer.All))
		server.HandleFunc("/importer/generations", r(importer.LoadGenerations))
		server.HandleFunc("/importer/items", r(importer.LoadItems))
		server.HandleFunc("/importer/item-categories", r(importer.LoadItemCategories))
		server.HandleFunc("/importer/pockets", r(importer.LoadPockets))
		server.HandleFunc("/importer/types", r(importer.LoadTypes))
		server.HandleFunc("/importer/moves", r(importer.LoadMoves))
		server.HandleFunc("/importer/move-methods", r(importer.LoadMoveMethods))
		server.HandleFunc("/importer/evolution-triggers", r(importer.LoadEvolutionTriggers))
		server.HandleFunc("/importer/pokemons", r(importer.LoadPokemons))

		return http.ListenAndServe(":8080", server)
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
