create table students
(
    id    bigint unsigned auto_increment
        primary key,
    name  varchar(10)      not null,
    age   tinyint unsigned not null,
    grade varchar(10)      not null
) charset utf8mb4;

insert into students(name, age, grade) values ('张三', 20, '三年级');
select * from students where age > 18;
update students set grade = '四年级' where name = '张三';
delete from students where age < 15;