package gin

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
)

var once sync.Once

type RouteDump struct {
	Methods []string `json:"methods"`
	Path    string   `json:"path"`
	Handler string   `json:"handler"`
}

func DumpRoutes(engine *Engine) {
	once.Do(func() {
		outputFile := "routes.jsonl"
		file, err := os.Create(outputFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to create dump file: %v\n", err)
			os.Exit(1)
		}
		defer file.Close()

		routes := engine.Routes()

		for _, r := range routes {
			dump := RouteDump{
				Methods: []string{r.Method},
				Path:    r.Path,
				Handler: r.Handler,
			}
			data, _ := json.Marshal(dump)
			fmt.Fprintln(file, string(data))
		}

		fmt.Printf("Routes dumped to %s, exiting...\n", outputFile)
		os.Exit(0)
	})
}
