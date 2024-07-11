
package assets

import "github.com/go-teal/teal/pkg/processing"

var PorjectAssets = map[string] processing.Asset{
	"staging.addresses":stagingAddressesAsset,
	"staging.transactions":stagingTransactionsAsset,
	"staging.wallets":stagingWalletsAsset,
	"dds.dim_addresses":ddsDimAddressesAsset,
	"dds.fact_transactions":ddsFactTransactionsAsset,
	"mart.mart_wallet_report":martMartWalletReportAsset,
}

var DAG = [][]string{
	{
		"staging.addresses",
		"staging.transactions",
		"staging.wallets",			
	},
	{
		"dds.dim_addresses",			
	},
	{
		"dds.fact_transactions",			
	},
	{
		"mart.mart_wallet_report",			
	},
}