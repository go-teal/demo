
package assets

import "github.com/go-teal/teal/pkg/processing"

var ProjectAssets = map[string] processing.Asset{
	"staging.addresses":stagingAddressesAsset,
	"staging.transactions":stagingTransactionsAsset,
	"staging.wallets":stagingWalletsAsset,
	"dds.dim_addresses":ddsDimAddressesAsset,
	"dds.fact_transactions":ddsFactTransactionsAsset,
	"mart.mart_address_len":martMartAddressLenAsset,
	"mart.mart_wallet_report":martMartWalletReportAsset,
	"mart2.mart_addresses":mart2MartAddressesAsset,
}

var DAG = [][]string{
	{
		"staging.addresses",
		"staging.transactions",
		"staging.wallets",			
	},
	{
		"dds.dim_addresses",
		"mart2.mart_addresses",			
	},
	{
		"dds.fact_transactions",
		"mart.mart_address_len",			
	},
	{
		"mart.mart_wallet_report",			
	},
}