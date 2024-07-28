{{- define "profile.yaml" }}
    connection: 'default'    
{{-  end }}
select pk_id, count(pk_id) as c from {{ Ref "dds.fact_transactions" }} group by pk_id having c > 1