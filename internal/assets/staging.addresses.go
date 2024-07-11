package assets

import (	
	"github.com/go-teal/teal/pkg/models"
	"github.com/go-teal/teal/pkg/configs"
	"github.com/go-teal/teal/pkg/processing"
)

const RAW_SQL_STAGING_ADDRESSES = `


select
    id,
    wallet_id,
    wallet_address,
    currency
    from read_csv('store/addresses.csv',
    delim = ',',
    header = true,
    columns = {
        'id': 'INT',
        'wallet_id': 'VARCHAR',
        'wallet_address': 'VARCHAR',
        'currency': 'VARCHAR'}
    )
`
const SQL_STAGING_ADDRESSES_CREATE_TABLE = `
create table staging.addresses as (

select
    id,
    wallet_id,
    wallet_address,
    currency
    from read_csv('store/addresses.csv',
    delim = ',',
    header = true,
    columns = {
        'id': 'INT',
        'wallet_id': 'VARCHAR',
        'wallet_address': 'VARCHAR',
        'currency': 'VARCHAR'}
    ))
`
const SQL_STAGING_ADDRESSES_INSERT = `
insert into staging.addresses ({{ ModelFields }}) (

select
    id,
    wallet_id,
    wallet_address,
    currency
    from read_csv('store/addresses.csv',
    delim = ',',
    header = true,
    columns = {
        'id': 'INT',
        'wallet_id': 'VARCHAR',
        'wallet_address': 'VARCHAR',
        'currency': 'VARCHAR'}
    ))
`
const SQL_STAGING_ADDRESSES_DROP_TABLE = `
drop table staging.addresses
`
const SQL_STAGING_ADDRESSES_TRUNCATE = `
truncate table staging.addresses
`

var stagingAddressesModelDescriptor = &models.SQLModelDescriptor{
	Name: 				"staging.addresses",
	RawSQL: 			RAW_SQL_STAGING_ADDRESSES,
	CreateTableSQL: 	SQL_STAGING_ADDRESSES_CREATE_TABLE,
	InsertSQL: 			SQL_STAGING_ADDRESSES_INSERT,
	DropTableSQL: 		SQL_STAGING_ADDRESSES_DROP_TABLE,
	TruncateTableSQL: 	SQL_STAGING_ADDRESSES_TRUNCATE,	
	Upstreams: []string {
	},
	Downstreams: []string {
		"dds.dim_addresses",
	},
	ModelProfile:  &configs.ModelProfile{
		Name: 				"addresses",
		Stage: 				"staging",
		Connection: 		"default",
		Materialization: 	"table",
		IsDataFramed: 		true,
		PersistInputs: 		false,
	},
}

var stagingAddressesAsset processing.Asset = processing.InitSQLModelAsset(stagingAddressesModelDescriptor)