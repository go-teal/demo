package assets

import (	
	"github.com/go-teal/teal/pkg/models"
	"github.com/go-teal/teal/pkg/configs"
	"github.com/go-teal/teal/pkg/processing"
)

const RAW_SQL_STAGING_WALLETS = `


select
      id,    
      name,
      ledger_id,
      wallet_address,
      currency,
      ticker,
      contract_id,
      raw_balance,
      created_at,
      updated_at
    from read_csv('store/wallets.csv',
    delim = ',',
    header = true,
    columns = {
      'id': 'INT',
      'name': 'VARCHAR',
      'ledger_id': 'INT',
      'wallet_address': 'VARCHAR',
      'currency': 'VARCHAR',
      'ticker': 'VARCHAR',
      'contract_id': 'INT',
      'raw_balance': 'VARCHAR',
      'created_at':'DATE',
      'updated_at':'DATE'
      }
    )

`
const SQL_STAGING_WALLETS_CREATE_TABLE = `
create table staging.wallets as (

select
      id,    
      name,
      ledger_id,
      wallet_address,
      currency,
      ticker,
      contract_id,
      raw_balance,
      created_at,
      updated_at
    from read_csv('store/wallets.csv',
    delim = ',',
    header = true,
    columns = {
      'id': 'INT',
      'name': 'VARCHAR',
      'ledger_id': 'INT',
      'wallet_address': 'VARCHAR',
      'currency': 'VARCHAR',
      'ticker': 'VARCHAR',
      'contract_id': 'INT',
      'raw_balance': 'VARCHAR',
      'created_at':'DATE',
      'updated_at':'DATE'
      }
    )
)
`
const SQL_STAGING_WALLETS_INSERT = `
insert into staging.wallets ({{ ModelFields }}) (

select
      id,    
      name,
      ledger_id,
      wallet_address,
      currency,
      ticker,
      contract_id,
      raw_balance,
      created_at,
      updated_at
    from read_csv('store/wallets.csv',
    delim = ',',
    header = true,
    columns = {
      'id': 'INT',
      'name': 'VARCHAR',
      'ledger_id': 'INT',
      'wallet_address': 'VARCHAR',
      'currency': 'VARCHAR',
      'ticker': 'VARCHAR',
      'contract_id': 'INT',
      'raw_balance': 'VARCHAR',
      'created_at':'DATE',
      'updated_at':'DATE'
      }
    )
)
`
const SQL_STAGING_WALLETS_DROP_TABLE = `
drop table staging.wallets
`
const SQL_STAGING_WALLETS_TRUNCATE = `
truncate table staging.wallets
`

var stagingWalletsModelDescriptor = &models.SQLModelDescriptor{
	Name: 				"staging.wallets",
	RawSQL: 			RAW_SQL_STAGING_WALLETS,
	CreateTableSQL: 	SQL_STAGING_WALLETS_CREATE_TABLE,
	InsertSQL: 			SQL_STAGING_WALLETS_INSERT,
	DropTableSQL: 		SQL_STAGING_WALLETS_DROP_TABLE,
	TruncateTableSQL: 	SQL_STAGING_WALLETS_TRUNCATE,	
	Upstreams: []string {
	},
	Downstreams: []string {
	},
	ModelProfile:  &configs.ModelProfile{
		Name: 				"wallets",
		Stage: 				"staging",
		Connection: 		"default",
		Materialization: 	"table",
		IsDataFramed: 		false,
		PersistInputs: 		false,
	},
}

var stagingWalletsAsset processing.Asset = processing.InitSQLModelAsset(stagingWalletsModelDescriptor)