# Database Languages in SQL

# DDL - Data Definition Language -> create, alter, drop, truncate, rename
# DML - Data Manipulation Language -> select, insert, update, delete, explain
# DCL - Data Control Language -> grant, revoke
# TCL - Transactional Control Language -> roll back, commit, save point

# 
 - order by
    select * from car order by year esc;
 - group by
    select year from car group by year;

-- # aggregate queries/functions
 - max, min, count, sum, avg, distinct...

ALTER TABLE students ADD COLUMN gender gen;

update student set g='f' where lastname ilike '%va';
select * from where lastname ilike '%va';

 default gen_random_uuid()

create type gen as enum('male', 'female');

CREATE TABLE student (
    student_id UUID PRIMARY KEY default gen_random_uuid(),     
    name VARCHAR(50) NOT NULL,             
    lastname VARCHAR(50) NOT NULL,         
    phone VARCHAR(15) NOT NULL,            
    age INT NOT NULL,                      
    grade INT NOT NULL check(grade>59),     
    gender gen NOT NULL,     
    created_at TIMESTAMP NOT NULL default now(),     
    updated_at TIMESTAMP default now(),              
    deleted_at BIGINT DEFAULT 0         
);
-- one to many
create table course (
   id UUID PRIMARY KEY default gen_random_uuid(),     
   name varchar not null,
   started_date TIMESTAMP,
   tutor varchar,
   number int
);

create table student_course (
   id UUID PRIMARY KEY default gen_random_uuid(),     
   student_id uuid references student(student_id),
   course_id uuid references course(id)
);
-- student, course - CE, LOG, SE




INSERT INTO student (student_id, name, lastname, phone, age, grade, created_at, updated_at, deleted_at, gender)
VALUES
('94f0ca16-bb42-4b80-b878-5a4ca57fd4ca', 'Aziz', 'Yusupov', '+998901234567', 20, 85, '2024-05-14 12:49:39.543975', NULL, 0, 'male'),
('c6928344-ddd1-49b8-b70d-dfaa89c129b5', 'Dilnoza', 'Kamilova', '+998901234568', 21, 78, '2024-05-14 12:49:39.543975', NULL, 0, 'female'),
('581ad352-0d11-4c91-a4c0-6c280d174b66', 'Farrukh', 'Zokirov', '+998901234569', 22, 91, '2024-05-14 12:49:39.543975', NULL, 0, 'male'),
('2253f666-9640-4291-a62a-e34479d11ec9', 'Gulnara', 'Rakhimova', '+998901234570', 23, 75, '2024-05-14 12:49:39.543975', NULL, 0, 'female'),
('9ce54ed8-f321-4f48-a492-63e41d287e4d', 'Ibrohim', 'Abdullayev', '+998901234571', 24, 88, '2024-05-14 12:49:39.543975', NULL, 0, 'male');
('a67b3065-407c-46b0-b7e8-4e9ba4f71527', 'Jasur', 'Karimov', '+998901234572', 19, 73, '2024-05-14 12:49:39.543975', NULL, 0, 'male'),
('d9a7c912-7772-4ecd-aecc-e461a6e0fb97', 'Kamila', 'Safarova', '+998901234573', 20, 90, '2024-05-14 12:49:39.543975', NULL, 0, 'female'),
('cfc80055-02e3-4e1c-b510-ccf5832a008f', 'Lola', 'Tursunova', '+998901234574', 21, 92, '2024-05-14 12:49:39.543975', NULL, 0, 'female'),
('9e5916d9-7f64-449d-9cab-6fcbc5cc8b9a', 'Murod', 'Alimov', '+998901234575', 22, 81, '2024-05-14 12:49:39.543975', NULL, 0, 'male'),
('fd3e0268-7d0c-44f5-b508-8952e7cc6748', 'Nigora', 'Begimova', '+998901234576', 23, 84, '2024-05-14 12:49:39.543975', NULL, 0, 'female');
-- Existing 30 rows continue --
('a234f678-1122-4bbc-a119-7845d8124e4a', 'Nuriddin', 'Ahmedov', '+998901234591', 21, 76, '2024-05-14 12:49:39.543975', NULL, 0, 'male'),
('b873c9e5-2b57-4b21-88e9-c9c3c76870cd', 'Madina', 'Rustamova', '+998901234592', 22, 68, '2024-05-14 12:49:39.543975', NULL, 0, 'female'),
('c98e7b6a-3f67-43d4-b16b-2c71e9c42342', 'Jahongir', 'Husanov', '+998901234593', 20, 89, '2024-05-14 12:49:39.543975', NULL, 0, 'male'),
('d93a2176-3a9e-42f5-b7ec-27a112765ed2', 'Sardor', 'Alisherov', '+998901234594', 24, 93, '2024-05-14 12:49:39.543975', NULL, 0, 'male'),
('e83fd51c-7ba5-4b79-a51e-62726fb8126b', 'Shakhnoza', 'Islomova', '+998901234595', 23, 72, '2024-05-14 12:49:39.543975', NULL, 0, 'female'),
('f45a3876-b122-42c6-90a5-cf9831c234ed', 'Abdurahim', 'Nasimov', '+998901234596', 22, 95, '2024-05-14 12:49:39.543975', NULL, 0, 'male'),
('a56b739f-1ac4-4395-81f3-5fae687ea432', 'Sabina', 'Ulugbekova', '+998901234597', 21, 78, '2024-05-14 12:49:39.543975', NULL, 0, 'female'),
('a65c897f-7ea1-42c4-8a56-6fda75e7431b', 'Umid', 'Zafarov', '+998901234598', 23, 85, '2024-05-14 12:49:39.543975', NULL, 0, 'male'),
('b98d6a4e-4fa7-4c12-b5f9-6b21d8324e4f', 'Nilufar', 'Asadova', '+998901234599', 24, 91, '2024-05-14 12:49:39.543975', NULL, 0, 'female'),
('c47f8b2a-7de2-4d87-b9f2-4b98321d49f1', 'Timur', 'Sultonov', '+998901234600', 22, 87, '2024-05-14 12:49:39.543975', NULL, 0, 'male');
INSERT INTO student (student_id, name, lastname, phone, age, grade, created_at, updated_at, deleted_at, gender)
VALUES
('11e8b756-2c42-4c8f-b1a4-3e8128b4123d', 'Akmal', 'Shukurov', '+998901234601', 21, 82, '2024-05-14 12:49:39.543975', NULL, 0, 'male'),
('23f67a1c-8c9b-4e5e-bd7c-5e71876f4319', 'Nozima', 'Mahmudova', '+998901234602', 22, 79, '2024-05-14 12:49:39.543975', NULL, 0, 'female'),
('34b9a712-1c74-43d6-b8a7-6f912d374e8a', 'Javlon', 'Ergashev', '+998901234603', 23, 88, '2024-05-14 12:49:39.543975', NULL, 0, 'male'),
('45c6e214-5f8a-4e3b-b3f4-7d981234fb56', 'Gulshan', 'Ismoilova', '+998901234604', 24, 92, '2024-05-14 12:49:39.543975', NULL, 0, 'female'),
('56f7d381-2d42-43a9-b412-2e8912736c1f', 'Said', 'Abdurakhmanov', '+998901234605', 20, 85, '2024-05-14 12:49:39.543975', NULL, 0, 'male'),
('67a8b492-8c35-4b56-b413-3f912f765b12', 'Laylo', 'Sharifova', '+998901234606', 19, 77, '2024-05-14 12:49:39.543975', NULL, 0, 'female'),
('78d9c521-9c24-4f89-b6e7-4d901234fc9b', 'Olim', 'Nasriddinov', '+998901234607', 22, 94, '2024-05-14 12:49:39.543975', NULL, 0, 'male'),
('89e7a6b2-2f4c-4538-b5e9-5e9013456b8f', 'Shakhzod', 'Yuldashev', '+998901234608', 23, 73, '2024-05-14 12:49:39.543975', NULL, 0, 'male'),
('91b7c891-1d56-4c3b-b5d4-6f921243fc17', 'Farzona', 'Karimova', '+998901234609', 21, 80, '2024-05-14 12:49:39.543975', NULL, 0, 'female'),
('a2c9d461-2b34-4f8d-b7f6-7e9012358d2c', 'Mansur', 'Rashidov', '+998901234610', 24, 83, '2024-05-14 12:49:39.543975', NULL, 0, 'male'),
('b3d6a592-3e58-4f3b-b6f5-8f9012367b4e', 'Sevinch', 'Nuriddinova', '+998901234611', 19, 76, '2024-05-14 12:49:39.543975', NULL, 0, 'female'),
('c4e7a671-4b39-4c7a-b6e8-9f9012378c2a', 'Komil', 'Bekmirzaev', '+998901234612', 20, 89, '2024-05-14 12:49:39.543975', NULL, 0, 'male'),
('d5f8a732-5c28-4e7d-b8a9-af9123459d3b', 'Sabohat', 'Rustamova', '+998901234613', 22, 78, '2024-05-14 12:49:39.543975', NULL, 0, 'female'),
('e6c9a892-6d49-4b6f-b7f2-bf9132458e5f', 'Zokir', 'Nasimov', '+998901234614', 23, 91, '2024-05-14 12:49:39.543975', NULL, 0, 'male'),
('f7d6b953-7e58-4c8b-b7e3-cf9143469f6c', 'Madina', 'Islomova', '+998901234615', 21, 74, '2024-05-14 12:49:39.543975', NULL, 0, 'female'),
('08c7a613-8f49-4e9d-b6e4-df9153467a8f', 'Jasmina', 'Alimova', '+998901234616', 24, 86, '2024-05-14 12:49:39.543975', NULL, 0, 'female'),
('19d8b724-9e39-4f8c-b7f7-ef9163468b9a', 'Yusuf', 'Ishmatov', '+998901234617', 22, 81, '2024-05-14 12:49:39.543975', NULL, 0, 'male'),
('2ae7c835-0f28-4e9f-b8a2-ff9173479c1b', 'Nilufar', 'Usmonova', '+998901234618', 23, 82, '2024-05-14 12:49:39.543975', NULL, 0, 'female'),
('3bf8d946-1e49-4c7a-b5f4-1f9183480d2c', 'Sherzod', 'Rakhimov', '+998901234619', 20, 90, '2024-05-14 12:49:39.543975', NULL, 0, 'male'),
('4cd9e057-2f39-4e6f-b9e7-2f9193481e3d', 'Dildora', 'Sharifova', '+998901234620', 21, 79, '2024-05-14 12:49:39.543975', NULL, 0, 'female');


select b.id, b.name, page, a.name as Author from author as a left join book as b on a.name=b.author_name;

select s.name, lastname, age, grade, gender, c.name 
   from student s 
join student_course as sc on s.student_id=sc.student_id 
   join course as c on c.id=sc.course_id;