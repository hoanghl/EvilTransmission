drop table if exists resources;

create table resources(
	id 					serial 			primary key
	, res_type 			varchar(6) 	 	not null
	, path 				varchar(1000)	not null
    , belong_res_name 	varchar(1000)	not null
	, belong_res_type 	varchar(6) 	 	not null
	, last_update 	timestamptz 		not null
);
set timezone to 'Europe/Helsinki';
-- insert into resources(resid, path, lastupdate) values (
-- 	'493423',
-- 	'/user',
-- 	current_timestamp
-- )


-- select * from resources

-- drop database postgres

-- select * from 




