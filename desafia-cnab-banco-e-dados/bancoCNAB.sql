SELECT name FROM sys.databases;
SELECT local_net_address, local_tcp_port FROM sys.dm_exec_connections WHERE local_net_address IS NOT NULL;


create table CNAB (

	id int primary key identity(1,1),
	Tipo varchar(50) not null,
	Natureza varchar(50) not null,
	Sinal varchar(50) not null,
	Data date not null,
	Valor float not null,
	Cpf char(15) not null,
	Cartao varchar(50) not null,
	Hora time not null,
	Dono_loja varchar(50) not null,
	Nome_loja varchar(50) not null

);