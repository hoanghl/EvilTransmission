DROP TABLE IF EXISTS RESOURCES;

CREATE TABLE RESOURCES(
	ID 			SERIAL 			PRIMARY KEY
	, RESID 	VARCHAR(200) 	UNIQUE 		NOT NULL
	, PATH 	VARCHAR(1000) 				NOT NULL
    , HASHVAL 	BYTEA 						NOT NULL
	, LASTUPDATE 	TIMESTAMPTZ 			NOT NULL
);
SET TIMEZONE TO 'ASIA/HO_CHI_MINH';
-- INSERT INTO RESOURCES(RESID, PATH, LASTUPDATE) VALUES (
-- 	'493423',
-- 	'/USER',
-- 	CURRENT_TIMESTAMP
-- )


-- SELECT * FROM RESOURCES

-- DROP DATABASE POSTGRES

-- SELECT * FROM 

select * from resources;



