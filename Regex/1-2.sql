CREATE TABLE students(
   id INT PRIMARY KEY     NOT NULL,
   name           TEXT    NOT NULL,
   age            INT     NOT NULL,
   gender         INT     NOT NULL,
   current_class  VARCHAR(255) NOT NULL,
   freshman_class VARCHAR(255),
   zju_id           BIGINT NOT NULL UNIQUE
);
INSERT INTO students (id,name,gender,age,current_class,freshman_class,zju_id) VALUES ( 1,'小张',1,20,'计科1901','工信1918',3190101234);
INSERT INTO students (id,name,gender,age,current_class,zju_id) VALUES (2,'小胡',0,19,'工信2017',3200105678 );
INSERT INTO students (id,name,gender,age,current_class,freshman_class,zju_id) VALUES (3,'小王',1,21,'机械1802','机材1801',3180109012);

-- 1 for w 0 for m
SELECT * FROM students;

