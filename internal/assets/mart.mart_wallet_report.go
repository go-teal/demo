package assets

import (	
	"github.com/go-teal/teal/pkg/models"
	"github.com/go-teal/teal/pkg/configs"
	"github.com/go-teal/teal/pkg/processing"
)

const RAW_SQL_MART_MART_WALLET_REPORT = `


SELECT * from dds.fact_transactions



`
const SQL_MART_MART_WALLET_REPORT_CREATE_VIEW = `
create view mart.mart_wallet_report as (

SELECT * from dds.fact_transactions


)
`
const SQL_MART_MART_WALLET_REPORT_DROP_VIEW = `
drop view mart.mart_wallet_report
`

var martMartWalletReportModelDescriptor = &models.SQLModelDescriptor{
	Name: 				"mart.mart_wallet_report",
	RawSQL: 			RAW_SQL_MART_MART_WALLET_REPORT,
	CreateViewSQL: 		SQL_MART_MART_WALLET_REPORT_CREATE_VIEW,
	DropViewSQL: 		SQL_MART_MART_WALLET_REPORT_DROP_VIEW,	
	Upstreams: []string {
		"dds.fact_transactions",
	},
	Downstreams: []string {
	},
	ModelProfile:  &configs.ModelProfile{
		Name: 				"mart_wallet_report",
		Stage: 				"mart",
		Connection: 		"default",
		Materialization: 	"view",
		IsDataFramed: 		false,
		PersistInputs: 		false,
	},
}

var martMartWalletReportAsset processing.Asset = processing.InitSQLModelAsset(martMartWalletReportModelDescriptor)