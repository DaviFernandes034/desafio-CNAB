SELECT name FROM sys.databases;
SELECT local_net_address, local_tcp_port FROM sys.dm_exec_connections WHERE local_net_address IS NOT NULL;
