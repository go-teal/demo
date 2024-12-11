
package assets

import (	
	"github.com/go-teal/teal/pkg/models"
	"github.com/go-teal/teal/pkg/configs"
	"github.com/go-teal/teal/pkg/processing"
)

const RAW_SQL_MART2_MART_ADDRESSES = `


SELECT 
    encode(sha256( (wallet_address || currency)::bytea),'hex') as pk_id,    
    id as address_id,
    wallet_address as wallet_address,
    currency
 from tmp_staging_addresses
`
const SQL_MART2_MART_ADDRESSES_CREATE_TABLE = `
create table mart2.mart_addresses 
as (

SELECT 
    encode(sha256( (wallet_address || currency)::bytea),'hex') as pk_id,    
    id as address_id,
    wallet_address as wallet_address,
    currency
 from tmp_staging_addresses);

`
const SQL_MART2_MART_ADDRESSES_INSERT = `
insert into mart2.mart_addresses ({{ ModelFields }}) (

SELECT 
    encode(sha256( (wallet_address || currency)::bytea),'hex') as pk_id,    
    id as address_id,
    wallet_address as wallet_address,
    currency
 from tmp_staging_addresses)
`
const SQL_MART2_MART_ADDRESSES_DROP_TABLE = `
drop table mart2.mart_addresses
`
const SQL_MART2_MART_ADDRESSES_TRUNCATE = `
delete from mart2.mart_addresses where true;
truncate table mart2.mart_addresses;
`

var mart2MartAddressesModelDescriptor = &models.SQLModelDescriptor{
	Name: 				"mart2.mart_addresses",
	RawSQL: 			RAW_SQL_MART2_MART_ADDRESSES,
	CreateTableSQL: 	SQL_MART2_MART_ADDRESSES_CREATE_TABLE,
	InsertSQL: 			SQL_MART2_MART_ADDRESSES_INSERT,
	DropTableSQL: 		SQL_MART2_MART_ADDRESSES_DROP_TABLE,
	TruncateTableSQL: 	SQL_MART2_MART_ADDRESSES_TRUNCATE,	
	Upstreams: []string {
		"staging.addresses",
	},
	Downstreams: []string {
		"mart.mart_address_len",
	},
	ModelProfile:  &configs.ModelProfile{
		Name: 				"mart_addresses",
		Stage: 				"mart2",
		Connection: 		"test_pg",
		Materialization: 	"table",
		IsDataFramed: 		true,
		PersistInputs: 		true,
		Tests: []*configs.TestProfile {
		},
	},
}

var mart2MartAddressesAsset processing.Asset = processing.InitSQLModelAsset(mart2MartAddressesModelDescriptor)