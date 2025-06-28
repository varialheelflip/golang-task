create table books
(
    id     bigint unsigned auto_increment
        primary key,
    title  varchar(20)      not null comment '书名',
    author varchar(20)      not null comment '作者',
    price  decimal unsigned not null comment '价格'
) charset utf8mb4;

INSERT INTO books (id, title, author, price) VALUES (1, '月亮与六便士', '毛姆', 20);
INSERT INTO books (id, title, author, price) VALUES (2, '四千周', 'A', 60);
INSERT INTO books (id, title, author, price) VALUES (3, '放弃的艺术', 'B', 70);