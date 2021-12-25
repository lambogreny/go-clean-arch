-- Rodar no banco do postgresql


select * from tb_logs

--delete from tb_logs

SELECT table_name,column_name,is_nullable,data_type,character_maximum_length
FROM information_schema.columns
WHERE table_schema = 'public'
  AND table_name = 'cfr'
  AND data_type = 'character varying'
  AND character_maximum_length = 6
ORDER BY ordinal_position;