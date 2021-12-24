
CREATE TABLE logs(      id INTEGER primary key AUTOINCREMENT,      tabela VARCHAR NOT NULL,      tipo VARCHAR NOT NULL,     pk VARCHAR NOT NULL,      error bool default false,      message VARCHAR default '',      created_at datetime default current_timestamp    , cliente varchar)

insert into logs (tabela,tipo,pk,error,message) VALUES('PRD','I','123',false,'')

select * from logs


drop table logs






