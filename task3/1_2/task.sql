create table accounts
(
    id      bigint unsigned auto_increment
        primary key,
    balance DECIMAL unsigned not null default 0
)
    charset = utf8mb4;

create table transactions
(
    id              bigint unsigned auto_increment
        primary key,
    from_account_id bigint unsigned  not null,
    to_account_id   bigint unsigned  not null,
    amount          DECIMAL unsigned not null
)
    charset = utf8mb4;

DELIMITER //
CREATE PROCEDURE conditional_transaction()
BEGIN
    START TRANSACTION;
    SET @result_value = (select balance from accounts where id = 1 for update);
    IF @result_value >= 100 THEN
        update accounts set balance = balance - 100 where id = 1;
        update accounts set balance = balance + 100 where id = 2;
        insert into transactions(from_account_id, to_account_id, amount) VALUES (1, 2, 100);
        COMMIT;  -- 满足条件则提交
    ELSE
        ROLLBACK;  -- 不满足条件则回滚
    END IF;
END //
DELIMITER ;

CALL conditional_transaction();
drop procedure conditional_transaction;