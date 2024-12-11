package main

import (
	_ "github.com/marcboeker/go-duckdb"

	"os"

	"github.com/go-teal/demo/internal/assets"
	modeltests "github.com/go-teal/demo/internal/model_tests"
	"github.com/go-teal/teal/pkg/core"
	"github.com/go-teal/teal/pkg/dags"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	log.Info().Msg("Starting demo")
	core.GetInstance().Init("config.yaml", ".")
	defer core.GetInstance().Shutdown()
	config := core.GetInstance().Config
	// dag := dags.InitChannelDag(assets.DAG, assets.ProjectAssets, config, "demo")
	dag := dags.InitChannelDagWithTests(assets.DAG, assets.ProjectAssets, modeltests.ProjectTests, config, "demo")
	wg := dag.Run()
	result := <-dag.Push("demo", nil, make(chan map[string]interface{}))
	log.Info().Any("Result", result).Send()
	dag.Stop()
	wg.Wait()

	modeltests.TestAll()
	log.Info().Msg("Finishing demo")
}
