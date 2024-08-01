SELECT 
    u.name,
    o.amount,
    o.created_at
FROM 
    users u
JOIN 
    dblink('dbname=second', 'SELECT user_id, amount, created_at FROM orders') 
    AS o(user_id INT, amount NUMERIC(10,2), created_at TIMESTAMP)
ON 
    u.id = o.user_id;
