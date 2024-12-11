{{ define "profile.yaml" }}
    connection: 'test_pg'
    materialization: 'table'  
    persist_inputs: true
    is_data_framed: true
{{ end }}

SELECT 
    encode(sha256( (wallet_address || currency)::bytea),'hex') as pk_id,    
    id as address_id,
    wallet_address as wallet_address,
    currency
 from {{ Ref "staging.addresses" }}