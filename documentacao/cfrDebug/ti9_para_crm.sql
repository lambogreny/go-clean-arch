-- PG
select * from tb_logs

delete from tb_logs


select * from cfr where codigo_pessoa = '007564'

update cfr set caixa_postal = 'teste' where codigo_pessoa = '007564'

select * from tb_crm_sincroniza where tabela = 'CFR'

-- MYSQL

select * from epcrm_decorlit.account where id = '007564' limit 10;
select * from epcrm_decorlit.account limit 10;

select created_at from epcrm_decorlit.account limit 10;