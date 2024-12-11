{{define "profile.yaml"}}
    connection: 'default'
    materialization: 'table'
    persist_inputs: true
{{end}}

SELECT 
    address_id,
    LENGTH(wallet_address)
from {{ Ref "mart2.mart_addresses" }}


