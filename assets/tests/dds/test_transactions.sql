{{- define "profile.yaml" }}
    connection: 'default'    
{{-  end }}
SELECT pk_id, count(pk_id) as c from {{ Ref "dds.fact_transactions" }} group by pk_id HAVING c > 1