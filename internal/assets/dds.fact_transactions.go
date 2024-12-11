
package assets

import (	
	"github.com/go-teal/teal/pkg/models"
	"github.com/go-teal/teal/pkg/configs"
	"github.com/go-teal/teal/pkg/processing"
)

const RAW_SQL_DDS_FACT_TRANSACTIONS = `


with source as (
    select 
        sha256( t.tx_hash || t.currency || t.wallet_address || t.tx_index) as pk_id,
        t.amount::HUGEINT as amount,
        t.tx_created_on as tx_created_on,
        date_trunc('day',  t.tx_created_on) as tx_date,
        date_trunc('hour', t.tx_created_on) as tx_hour,
        t.tx_hash as tx_hash,
        sha256(t.wallet_address || t.currency) as fk_address_id,        
        t.tx_index as tx_index,
    from staging.transactions as t
    inner join dds.dim_addresses as a 
        on a.wallet_address = t.wallet_address 
        and  a.currency = t.currency 
)

select 
    pk_id,
    amount,
    tx_created_on,
    tx_date,
    tx_hour,
    tx_hash,
    fk_address_id,
    tx_index
 from source

 {{- if IsIncremental }}
    where tx_date > (select max(tx_created_on) from dds.fact_transactions)
 {{- end}}
`
const SQL_DDS_FACT_TRANSACTIONS_CREATE_TABLE = `
create table dds.fact_transactions 
as (

with source as (
    select 
        sha256( t.tx_hash || t.currency || t.wallet_address || t.tx_index) as pk_id,
        t.amount::HUGEINT as amount,
        t.tx_created_on as tx_created_on,
        date_trunc('day',  t.tx_created_on) as tx_date,
        date_trunc('hour', t.tx_created_on) as tx_hour,
        t.tx_hash as tx_hash,
        sha256(t.wallet_address || t.currency) as fk_address_id,        
        t.tx_index as tx_index,
    from staging.transactions as t
    inner join dds.dim_addresses as a 
        on a.wallet_address = t.wallet_address 
        and  a.currency = t.currency 
)

select 
    pk_id,
    amount,
    tx_created_on,
    tx_date,
    tx_hour,
    tx_hash,
    fk_address_id,
    tx_index
 from source

 {{- if IsIncremental }}
    where tx_date > (select max(tx_created_on) from dds.fact_transactions)
 {{- end}});

`
const SQL_DDS_FACT_TRANSACTIONS_INSERT = `
insert into dds.fact_transactions ({{ ModelFields }}) (

with source as (
    select 
        sha256( t.tx_hash || t.currency || t.wallet_address || t.tx_index) as pk_id,
        t.amount::HUGEINT as amount,
        t.tx_created_on as tx_created_on,
        date_trunc('day',  t.tx_created_on) as tx_date,
        date_trunc('hour', t.tx_created_on) as tx_hour,
        t.tx_hash as tx_hash,
        sha256(t.wallet_address || t.currency) as fk_address_id,        
        t.tx_index as tx_index,
    from staging.transactions as t
    inner join dds.dim_addresses as a 
        on a.wallet_address = t.wallet_address 
        and  a.currency = t.currency 
)

select 
    pk_id,
    amount,
    tx_created_on,
    tx_date,
    tx_hour,
    tx_hash,
    fk_address_id,
    tx_index
 from source

 {{- if IsIncremental }}
    where tx_date > (select max(tx_created_on) from dds.fact_transactions)
 {{- end}})
`
const SQL_DDS_FACT_TRANSACTIONS_DROP_TABLE = `
drop table dds.fact_transactions
`
const SQL_DDS_FACT_TRANSACTIONS_TRUNCATE = `
delete from dds.fact_transactions where true;
truncate table dds.fact_transactions;
`

var ddsFactTransactionsModelDescriptor = &models.SQLModelDescriptor{
	Name: 				"dds.fact_transactions",
	RawSQL: 			RAW_SQL_DDS_FACT_TRANSACTIONS,
	CreateTableSQL: 	SQL_DDS_FACT_TRANSACTIONS_CREATE_TABLE,
	InsertSQL: 			SQL_DDS_FACT_TRANSACTIONS_INSERT,
	DropTableSQL: 		SQL_DDS_FACT_TRANSACTIONS_DROP_TABLE,
	TruncateTableSQL: 	SQL_DDS_FACT_TRANSACTIONS_TRUNCATE,	
	Upstreams: []string {
		"staging.transactions",
		"dds.dim_addresses",
	},
	Downstreams: []string {
		"mart.mart_wallet_report",
	},
	ModelProfile:  &configs.ModelProfile{
		Name: 				"fact_transactions",
		Stage: 				"dds",
		Connection: 		"default",
		Materialization: 	"incremental",
		IsDataFramed: 		false,
		PersistInputs: 		false,
		Tests: []*configs.TestProfile {
		},
	},
}

var ddsFactTransactionsAsset processing.Asset = processing.InitSQLModelAsset(ddsFactTransactionsModelDescriptor)