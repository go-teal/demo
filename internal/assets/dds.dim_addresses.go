package assets

import (	
	"github.com/go-teal/teal/pkg/models"
	"github.com/go-teal/teal/pkg/configs"
	"github.com/go-teal/teal/pkg/processing"
)

const RAW_SQL_DDS_DIM_ADDRESSES = `
SELECT 
    sha256(wallet_address || currency) as pk_id,    
    wallet_address as wallet_address,
    currency
 from staging.addresses
`
const SQL_DDS_DIM_ADDRESSES_CREATE_TABLE = `
create table dds.dim_addresses as (SELECT 
    sha256(wallet_address || currency) as pk_id,    
    wallet_address as wallet_address,
    currency
 from staging.addresses)
`
const SQL_DDS_DIM_ADDRESSES_INSERT = `
insert into dds.dim_addresses ({{ ModelFields }}) (SELECT 
    sha256(wallet_address || currency) as pk_id,    
    wallet_address as wallet_address,
    currency
 from staging.addresses)
`
const SQL_DDS_DIM_ADDRESSES_DROP_TABLE = `
drop table dds.dim_addresses
`
const SQL_DDS_DIM_ADDRESSES_TRUNCATE = `
truncate table dds.dim_addresses
`

var ddsDimAddressesModelDescriptor = &models.SQLModelDescriptor{
	Name: 				"dds.dim_addresses",
	RawSQL: 			RAW_SQL_DDS_DIM_ADDRESSES,
	CreateTableSQL: 	SQL_DDS_DIM_ADDRESSES_CREATE_TABLE,
	InsertSQL: 			SQL_DDS_DIM_ADDRESSES_INSERT,
	DropTableSQL: 		SQL_DDS_DIM_ADDRESSES_DROP_TABLE,
	TruncateTableSQL: 	SQL_DDS_DIM_ADDRESSES_TRUNCATE,	
	Upstreams: []string {
		"staging.addresses",
	},
	Downstreams: []string {
		"dds.fact_transactions",
	},
	ModelProfile:  &configs.ModelProfile{
		Name: 				"dim_addresses",
		Stage: 				"dds",
		Connection: 		"default",
		Materialization: 	"table",
		IsDataFramed: 		false,
		PersistInputs: 		false,
	},
}

var ddsDimAddressesAsset processing.Asset = processing.InitSQLModelAsset(ddsDimAddressesModelDescriptor)