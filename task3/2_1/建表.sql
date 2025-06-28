create table employees
(
    id         bigint unsigned auto_increment
        primary key,
    name       varchar(20)  not null comment '名字',
    department varchar(20)  not null comment '部门',
    salary     int unsigned not null comment '工资'
) charset utf8mb4;

INSERT INTO employees (id, name, department, salary) VALUES (1, '小A', '技术部', 10);
INSERT INTO employees (id, name, department, salary) VALUES (2, '小B', '技术部', 40);
INSERT INTO employees (id, name, department, salary) VALUES (3, '小C', '信息部', 30);
INSERT INTO employees (id, name, department, salary) VALUES (4, '小D', '技术部', 40);