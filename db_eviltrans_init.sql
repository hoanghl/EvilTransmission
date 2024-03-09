create table if not EXISTS resources(
	id 					serial 			primary key
	, res_type 			varchar(10) 	not null
	, file_name 		varchar(1000)	not null
	, thumbnail_id		integer			null				
	, last_update 		timestamptz 	not null
	-- , embedding 		vector(10)		not null
);
ALTER DATABASE eviltrans SET timezone TO 'Europe/Helsinki';