select * from tb_crm_sincroniza where tabela = 'PRD'

insert into prd (codigo_produto) values ('AUGUSTO_1')

select * from prd where codigo_produto = 'AUGUSTO_1'

update prd set data_hora_alteracao = current_timestamp where codigo_produto = 'AUGUSTO_1'

