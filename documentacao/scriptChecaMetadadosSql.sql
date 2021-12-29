-- Rodar no banco do postgresql

select * from tb_logs
select * from tb_logs_details

delete from tb_logs
delete from tb_logs_details

select tb_logs_details.id,tb_logs_details.tabela,tb_logs.tipo,
       tb_logs.pk, tb_logs.error,tb_logs_details.queryString,tb_logs_details.dbResponse,tb_logs_details.created_at
from tb_logs INNER JOIN tb_logs_details on (tb_logs.pk = tb_logs_details.pk)

--delete from tb_logs

-- Checando os metadados no postgres
SELECT table_name,column_name,is_nullable,data_type,character_maximum_length
FROM information_schema.columns
WHERE table_schema = 'public'
  AND table_name = 'cfr'
  AND data_type = 'character varying'
  AND character_maximum_length = 6
ORDER BY ordinal_position;

-- Mostrando o create
SELECT
        'CREATE TABLE ' || relname || E'\n(\n' ||
  array_to_string(
    array_agg(
      '    ' || column_name || ' ' ||  type || ' '|| not_null
    )
    , E',\n'
  ) || E'\n);\n'
from
    (
        SELECT
            c.relname, a.attname AS column_name,
            pg_catalog.format_type(a.atttypid, a.atttypmod) as type,
            case
                when a.attnotnull
                    then 'NOT NULL'
                else 'NULL'
                END as not_null
        FROM pg_class c,
             pg_attribute a,
             pg_type t
        WHERE c.relname = 'tb_logs'
          AND a.attnum > 0
          AND a.attrelid = c.oid
          AND a.atttypid = t.oid
        ORDER BY a.attnum
    ) as tabledefinition
group by relname;

