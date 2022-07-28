-- a_id,  b_id
CREATE TABLE IF NOT EXISTS "patient" (
	"id"	INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	"id_a"	INTEGER,
	"id_b"	INTEGER,
    unique("id_a","id_b")
);
-- 1,1
-- 1,2
-- 1,1 (false) ,karena kombinasi id_a,id_b = (1,1) sudah exist