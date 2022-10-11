drop table if exists resources;

create table resources(
	id serial primary key,
	resid varchar(200) unique not null,
	path varchar(1000) not null,
    hashval bytea not null,
	lastupdate timestamptz not null
);
set timezone TO 'Asia/Ho_Chi_Minh';
-- insert into resources(resid, path, lastupdate) values (
-- 	'493423',
-- 	'/user',
-- 	current_timestamp
-- )


-- select * from resources

-- drop database postgres

-- select * from resources