package modeltests

import (	
	"github.com/go-teal/teal/pkg/models"
	"github.com/go-teal/teal/pkg/configs"
	"github.com/go-teal/teal/pkg/processing"
)

const RAW_SQL_DDS_TEST_TRANSACTIONS = `

SELECT pk_id, count(pk_id) as c from dds.fact_transactions group by pk_id HAVING c > 1
`

const COUNT_TEST_SQL_DDS_TEST_TRANSACTIONS = `
select count(*) as test_count from 
(

SELECT pk_id, count(pk_id) as c from dds.fact_transactions group by pk_id HAVING c > 1
) having test_count > 0 limit 1
`


var ddsTestTransactionsTestDescriptor = &models.SQLModelTestDescriptor{
	Name: 				"dds.test_transactions",
	RawSQL: 			RAW_SQL_DDS_TEST_TRANSACTIONS,
	CountTestSQL: 		COUNT_TEST_SQL_DDS_TEST_TRANSACTIONS,
	TestProfile: 		&configs.TestProfile {
		Name: 				"dds.test_transactions",
		Stage: 				"dds",
		Connection: 		"default",
	},
}

var ddsTestTransactionsSimpleTestCase = processing.InitSQLModelTesting(ddsTestTransactionsTestDescriptor)