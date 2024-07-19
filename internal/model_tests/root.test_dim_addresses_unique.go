package modeltests

import (	
	"github.com/go-teal/teal/pkg/models"
	"github.com/go-teal/teal/pkg/configs"
	"github.com/go-teal/teal/pkg/processing"
)

const RAW_SQL_ROOT_TEST_DIM_ADDRESSES_UNIQUE = `

select count(pk_id) as c from dds.dim_addresses group by pk_id having c > 1
`

const COUNT_TEST_SQL_ROOT_TEST_DIM_ADDRESSES_UNIQUE = `
select count(*) as test_count from 
(

select count(pk_id) as c from dds.dim_addresses group by pk_id having c > 1
) having test_count > 0 limit 1
`


var rootTestDimAddressesUniqueTestDescriptor = &models.SQLModelTestDescriptor{
	Name: 				"root.test_dim_addresses_unique",
	RawSQL: 			RAW_SQL_ROOT_TEST_DIM_ADDRESSES_UNIQUE,
	CountTestSQL: 		COUNT_TEST_SQL_ROOT_TEST_DIM_ADDRESSES_UNIQUE,
	TestProfile: 		&configs.TestProfile {
		Name: 				"root.test_dim_addresses_unique",
		Stage: 				"root",
		Connection: 		"default",
	},
}

var rootTestDimAddressesUniqueSimpleTestCase = processing.InitSQLModelTesting(rootTestDimAddressesUniqueTestDescriptor)