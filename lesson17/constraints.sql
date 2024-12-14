-- constraints

    -- rule.
-- primary key - unique, only 1 primary key in a table, faster select;
-- foreign key - for referencing
-- not null - 
-- unique
-- check

-- indexes

create table message(id serial, user_id int, message varchar, created_at timestamp);

create unique index message_index2 on message(id);

insert into message(user_id, message) values (2,'Hello'), (3,'Values');

create unique index message_index3 on message using hash(id, user_id) ;

-- Insert mock data into the 'message' table
INSERT INTO message (user_id, message, created_at)
SELECT 
    (random() * (100000 - 1) + 1)::int AS user_id, -- Random user_id between 1 and 100
    'Sample message ' || gs AS message,        -- Generate unique message text
    NOW() - (random() * 365 * 24 * 60 * 60)::int * INTERVAL '1 second' AS created_at -- Random timestamp within last year
FROM generate_series(1, 100000000) AS gs;          -- Adjust the range (1, 1000) for the number of mock entries


explain analyze select count(1) from message;
explain analyze select count(id) from message ;