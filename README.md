# mysql-banchmarks

Test InnoDB in MySQL with 40 million rows in table (read, write) 

- You need to run this SQL code in database

```
  CREATE TABLE users (
  id INT NOT NULL AUTO_INCREMENT,
  firstname VARCHAR(50) NOT NULL,
  lastname VARCHAR(50) NOT NULL,
  email VARCHAR(255) NOT NULL,
  password VARCHAR(255) NOT NULL,
  date_of_birth DATE NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id)
  );

```

- Generate 40 million rows in table `users`

```
DELIMITER //

CREATE PROCEDURE generate_users()
BEGIN
  DECLARE i INT DEFAULT 1;

  WHILE i <= 40000000 DO
    INSERT INTO users (firstname, lastname, email, password, date_of_birth)
    VALUES (
      CONCAT('Firstname', i),
      CONCAT('Lastname', i),
      CONCAT('email', i, '@example.com'),
      MD5(RAND()),
      DATE_ADD('1980-01-01', INTERVAL FLOOR(RAND() * 365 * 30) DAY)
    );
    
    SET i = i + 1;
    
    IF i % 10000 = 0 THEN
      SELECT CONCAT('Generated ', i, ' rows') AS status;
    END IF;
  END WHILE;
  
  SELECT 'Done' AS status;
END //

DELIMITER ;

CALL generate_users();
```

## Result of benchmarks

### READ with filter by `date_of_birth`

`FOR HASH index i changed property - innodb_adaptive_hash_index` 

`gr than 2000-01-01`

|                                 | **Time, sec.** | **Rows** |
|:-------------------------------:|:------:|:------:|
|        **Without index**        |  30.8  |  38505723 |
|       **With BTREE index**      | 27.24 | 38505723 |
|     **With HASH index**         |  24.48 | 38505723  |

`less than 2000-01-01`

|                                 | **Time, sec.** | **Rows** |
|:-------------------------------:|:------:|:------:|
|        **Without index**        |  38.5  |  38505723 |
|       **With BTREE index**      | 35.25 | 38505723 |
|     **With HASH index**         |  34.7 | 38505723  |

`equal 2000-01-01`

|                                 | **Time, sec** | **Rows** |
|:-------------------------------:|:------:|:------:|
|        **Without index**        |  21.72  |  38505723 |
|       **With BTREE index**      | 0.05 | 3034 |
|     **With HASH index**         |  0.03 | 3034  |

`between 2000-01-01 and 2001-01-01`

|                                 | **Time, sec** | **Rows** |
|:-------------------------------:|:------:|:------:|
|        **Without index**        |  25.61  |  38505723 |
|       **With BTREE index**      | 15.2 | 2228480 |
|     **With HASH index**         |  14.04 | 2228480  |


### INSERT with different `innodb_flush_log_at_trx_commit` - 100 requests
 
`200 concurent`

|                                 | **AVG Response Time, sec** |
|:-------------------------------:|:------:|
|        **0**        |  0.04  |
|       **1**      | 0.10 |
|     **2**         |  0.05 |

`400 concurent`

|                                 | **AVG Response Time, sec** |
|:-------------------------------:|:------:|
|        **0**        |  0.33  |
|       **1**      | 0.48 |
|     **2**         |  0.41 |

`600 concurent`

|                                 | **AVG Response Time, sec** |
|:-------------------------------:|:------:|
|        **0**        |  0.43  |
|       **1**      | 0.46 |
|     **2**         |  0.47 |
