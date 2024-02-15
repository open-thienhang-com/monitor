package inventory

import (
	"fmt"
	"os"

	"log"
	"net/http"

	"os/signal"

	"github.com/spf13/cobra"

	"github.com/gorilla/mux"
	"mono.thienhang.com/pkg/engine"
	_ "mono.thienhang.com/pkg/framework/gorilla"
	"mono.thienhang.com/pkg/plugins"
	"mono.thienhang.com/pkg/plugins/admin"
)

var cmdServer = &cobra.Command{
	Use:   "server",
	Short: "Launch the HTTP server",
	Run: func(cmd *cobra.Command, args []string) {
		// Init engine
		app := mux.NewRouter()
		//
		eng := engine.Default()
		eng.AddConfigFromJSON("config.json")

		plgs := []plugins.Plugin{
			admin.NewAdmin(),
		}

		eng.AddPluginList(plgs)

		if err := eng.Use(app); err != nil {
			panic(err)
		}

		// DEBUG
		fmt.Println("Registered routes:")
		app.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
			pathTemplate, err := route.GetPathTemplate()
			if err != nil {
				return err
			}
			fmt.Printf("%s %s\n", "P", pathTemplate)
			return nil
		})

		// Start Service & Wait for terminal
		go func() {
			_ = http.ListenAndServe(":80", app)
		}()

		// Graceful shutdown
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, os.Interrupt)
		<-quit
		log.Print("closing database connection")
	},
}
