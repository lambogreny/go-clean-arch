-- Creation of product table
CREATE TABLE IF NOT EXISTS product (
  product_id INT NOT NULL,
  name varchar(250) NOT NULL,
  PRIMARY KEY (product_id)
);

CREATE TABLE IF NOT EXISTS teste (
  teste INT NOT NULL,
  name varchar(250) NOT NULL,
  PRIMARY KEY (teste)
);

-- Criando a tabela de logs
CREATE TABLE IF NOT EXISTS tb_logs
(
    id integer serial,
    tabela character varying(30) NOT NULL,
    tipo character varying(2) NOT NULL,
    pk character varying(50) NOT NULL,
    error boolean NULL,
    message text NULL,
    created_at timestamp without time zone NULL,
    cliente character varying(30) NULL
);

-- Criando a tabela de log_details
CREATE TABLE  IF NOT EXISTS tb_logs_details
(
    id integer serial,
    tabela character varying(30) NOT NULL,
    pk character varying(50) NOT NULL,
    querystring text NULL,
    dbresponse text NULL,
    responsetype character varying(50) NULL,
    created_at timestamp without time zone NULL
);

-- Criando tabela de correlação entre as pks
CREATE TABLE IF NOT EXISTS crm_pk_corr
(
    id integer serial,
    version character varying(10) NOT NULL,
    tabela character varying(30) NOT NULL,
    pk character varying(50) NOT NULL
);


