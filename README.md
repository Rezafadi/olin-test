# OLIN Assignment

## TASK: Golang

### Steps to Run the Project

1. Clone the repository:

   ```bash
   git clone https://github.com/Rezafadi/olin-test.git
   ```

2. Change into the project directory:

   ```bash
   cd olin-test/
   ```

3. Tidy up the dependencies:

   ```bash
   go mod tidy
   ```

4. Run the main file:

   ```bash
   go run main.go
   ```

5. Open your browser and navigate to the URL:
   ```URL
   http://localhost:8080/v1
   ```

### API Endpoints

- **Question 1:** [http://localhost:8080/v1/test1](http://localhost:8080/v1/test1)
- **Question 2:** [http://localhost:8080/v1/test2](http://localhost:8080/v1/test2)
- **Question 3:** [http://localhost:8080/v1/test3](http://localhost:8080/v1/test3)

## TASK: PostgreSQL

### Query for Question 1:

```sql
SELECT
    u.name AS user_name,
    SUM(o.amount) AS total_spent
FROM
    users u
JOIN
    orders o ON u.id = o.user_id
WHERE
    o.created_at >= '2022-01-01'
GROUP BY
    u.name
HAVING
    SUM(o.amount) >= 1000.00;
```

### Query for Question 2:

```sql
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

```
