package modeltests

import (
	"github.com/go-teal/teal/pkg/processing"	
	"github.com/rs/zerolog/log"
)

var PorjectTests = map[string] processing.ModelTesting{
	"root.test_dim_addresses_unique":rootTestDimAddressesUniqueSimpleTestCase,
	"dds.test_transactions":ddsTestTransactionsSimpleTestCase,
}


func TestAll() {
	for _, testCase := range PorjectTests {
		status, testName, err := testCase.Execute()
		if status {
			log.Info().Str("Test Case", testName).Msg("Success")
		} else {
			log.Error().Str("Test Case", testName).Err(err).Msg("Failed")
		}
	}
}