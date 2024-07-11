package assets

import (	
	"github.com/go-teal/teal/pkg/models"
	"github.com/go-teal/teal/pkg/configs"
	"github.com/go-teal/teal/pkg/processing"
)

const RAW_SQL_STAGING_TRANSACTIONS = `

select
      id,
      created_on as tx_created_on,
      tx_hash,
      currency,
      raw_amount as amount,
      wallet_address,
      index as tx_index,
      is_suspicious
 from  read_csv('store/transactions.csv',
    delim = ',',
    header = true,
    columns = {
      'id': 'INT',
      'created_on': 'DATE',
      'tx_hash': 'VARCHAR',
      'currency': 'VARCHAR',
      'raw_amount': 'VARCHAR',
      'wallet_address': 'VARCHAR',
      'index': 'INT',
      'is_suspicious':'BOOL'}
    )

`
const SQL_STAGING_TRANSACTIONS_CREATE_TABLE = `
create table staging.transactions as (
select
      id,
      created_on as tx_created_on,
      tx_hash,
      currency,
      raw_amount as amount,
      wallet_address,
      index as tx_index,
      is_suspicious
 from  read_csv('store/transactions.csv',
    delim = ',',
    header = true,
    columns = {
      'id': 'INT',
      'created_on': 'DATE',
      'tx_hash': 'VARCHAR',
      'currency': 'VARCHAR',
      'raw_amount': 'VARCHAR',
      'wallet_address': 'VARCHAR',
      'index': 'INT',
      'is_suspicious':'BOOL'}
    )
)
`
const SQL_STAGING_TRANSACTIONS_INSERT = `
insert into staging.transactions ({{ ModelFields }}) (
select
      id,
      created_on as tx_created_on,
      tx_hash,
      currency,
      raw_amount as amount,
      wallet_address,
      index as tx_index,
      is_suspicious
 from  read_csv('store/transactions.csv',
    delim = ',',
    header = true,
    columns = {
      'id': 'INT',
      'created_on': 'DATE',
      'tx_hash': 'VARCHAR',
      'currency': 'VARCHAR',
      'raw_amount': 'VARCHAR',
      'wallet_address': 'VARCHAR',
      'index': 'INT',
      'is_suspicious':'BOOL'}
    )
)
`
const SQL_STAGING_TRANSACTIONS_DROP_TABLE = `
drop table staging.transactions
`
const SQL_STAGING_TRANSACTIONS_TRUNCATE = `
truncate table staging.transactions
`

var stagingTransactionsModelDescriptor = &models.SQLModelDescriptor{
	Name: 				"staging.transactions",
	RawSQL: 			RAW_SQL_STAGING_TRANSACTIONS,
	CreateTableSQL: 	SQL_STAGING_TRANSACTIONS_CREATE_TABLE,
	InsertSQL: 			SQL_STAGING_TRANSACTIONS_INSERT,
	DropTableSQL: 		SQL_STAGING_TRANSACTIONS_DROP_TABLE,
	TruncateTableSQL: 	SQL_STAGING_TRANSACTIONS_TRUNCATE,	
	Upstreams: []string {
	},
	Downstreams: []string {
		"dds.fact_transactions",
	},
	ModelProfile:  &configs.ModelProfile{
		Name: 				"transactions",
		Stage: 				"staging",
		Connection: 		"default",
		Materialization: 	"table",
		IsDataFramed: 		false,
		PersistInputs: 		false,
	},
}

var stagingTransactionsAsset processing.Asset = processing.InitSQLModelAsset(stagingTransactionsModelDescriptor)