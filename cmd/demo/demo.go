package main

import (
	"fmt"

	_ "github.com/marcboeker/go-duckdb"

	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/go-teal/demo/internal/assets"
	modeltests "github.com/go-teal/demo/internal/model_tests"
	"github.com/go-teal/teal/pkg/core"
	"github.com/go-teal/teal/pkg/dags"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	log.Info().Msg("demo")
	core.GetInstance().Init("config.yaml", ".")
	config := core.GetInstance().Config
	dag := dags.InitChannelDag(assets.DAG, assets.PorjectAssets, config, "instance 1")
	wg := dag.Run()
	result := <-dag.Push("TEST", nil, make(chan map[string]interface{}))
	fmt.Println(result)
	dag.Stop()
	wg.Wait()

	modeltests.TestAll()
}
