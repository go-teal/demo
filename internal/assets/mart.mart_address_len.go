
package assets

import (	
	"github.com/go-teal/teal/pkg/models"
	"github.com/go-teal/teal/pkg/configs"
	"github.com/go-teal/teal/pkg/processing"
)

const RAW_SQL_MART_MART_ADDRESS_LEN = `


SELECT 
    address_id,
    LENGTH(wallet_address)
from tmp_mart2_mart_addresses



`
const SQL_MART_MART_ADDRESS_LEN_CREATE_TABLE = `
create table mart.mart_address_len 
as (

SELECT 
    address_id,
    LENGTH(wallet_address)
from tmp_mart2_mart_addresses


);

`
const SQL_MART_MART_ADDRESS_LEN_INSERT = `
insert into mart.mart_address_len ({{ ModelFields }}) (

SELECT 
    address_id,
    LENGTH(wallet_address)
from tmp_mart2_mart_addresses


)
`
const SQL_MART_MART_ADDRESS_LEN_DROP_TABLE = `
drop table mart.mart_address_len
`
const SQL_MART_MART_ADDRESS_LEN_TRUNCATE = `
delete from mart.mart_address_len where true;
truncate table mart.mart_address_len;
`

var martMartAddressLenModelDescriptor = &models.SQLModelDescriptor{
	Name: 				"mart.mart_address_len",
	RawSQL: 			RAW_SQL_MART_MART_ADDRESS_LEN,
	CreateTableSQL: 	SQL_MART_MART_ADDRESS_LEN_CREATE_TABLE,
	InsertSQL: 			SQL_MART_MART_ADDRESS_LEN_INSERT,
	DropTableSQL: 		SQL_MART_MART_ADDRESS_LEN_DROP_TABLE,
	TruncateTableSQL: 	SQL_MART_MART_ADDRESS_LEN_TRUNCATE,	
	Upstreams: []string {
		"mart2.mart_addresses",
	},
	Downstreams: []string {
	},
	ModelProfile:  &configs.ModelProfile{
		Name: 				"mart_address_len",
		Stage: 				"mart",
		Connection: 		"default",
		Materialization: 	"table",
		IsDataFramed: 		false,
		PersistInputs: 		true,
		Tests: []*configs.TestProfile {
		},
	},
}

var martMartAddressLenAsset processing.Asset = processing.InitSQLModelAsset(martMartAddressLenModelDescriptor)