## Dev Container TemplateでGo+Postgre（サンプルデータ付き）環境を構築する

Go & PostgreSQL (go-postgres)

[https://github.com/devcontainers/templates/tree/main/src/go-postgres](https://github.com/devcontainers/templates/tree/main/src/go-postgres)

The Sakila example database

[https://github.com/jOOQ/sakila](https://github.com/jOOQ/sakila)

jsonでポートフォワーディング有効にしても、効かなかった？から、VSCodeで直接追加したらローカルからアクセスできた

![Untitled](VSCode%20Dev%20Containers%E3%81%A6%E3%82%99Docker%E7%92%B0%E5%A2%83%E9%96%8B%E7%99%BA%20ae98071ad5144fd9a74eb14e7882be23/Untitled%201.png)

1. psqlコマンドでログイン確認
    
    ```bash
    psql -h localhost -p 5432 -U postgres
    ```
    
2. DB作成
    
    ```bash
    postgres=# CREATE DATABASE sakila;
    CREATE DATABASE
    postgres=# \l
                                     List of databases
       Name    |  Owner   | Encoding |  Collate   |   Ctype    |   Access privileges   
    -----------+----------+----------+------------+------------+-----------------------
     postgres  | postgres | UTF8     | en_US.utf8 | en_US.utf8 | 
     sakila    | postgres | UTF8     | en_US.utf8 | en_US.utf8 | 
     template0 | postgres | UTF8     | en_US.utf8 | en_US.utf8 | =c/postgres          +
               |          |          |            |            | postgres=CTc/postgres
     template1 | postgres | UTF8     | en_US.utf8 | en_US.utf8 | =c/postgres          +
               |          |          |            |            | postgres=CTc/postgres
    (4 rows)
    ```
    
3. テーブル作成
    
    ```sql
    psql -h localhost -p 5432 -U postgres -d sakila -a -f postgres-sakila-schema.sql
    ```
    
4. データ作成
    
    ```sql
    psql -h localhost -p 5432 -U postgres -d sakila -a -f postgres-sakila-insert-data.sql
    ```
    
5. データ確認
    
    ```sql
    % psql -h localhost -p 5432 -U postgres
    psql (14.10 (Homebrew), server 16.1 (Debian 16.1-1.pgdg120+1))
    WARNING: psql major version 14, server major version 16.
             Some psql features might not work.
    Type "help" for help.
    
    【DB確認】
    postgres=# \c sakila
    psql (14.10 (Homebrew), server 16.1 (Debian 16.1-1.pgdg120+1))
    WARNING: psql major version 14, server major version 16.
             Some psql features might not work.
    You are now connected to database "sakila" as user "postgres".
    
    【クエリ実行】
    sakila=# select * from actor where actor_id in (2,3,60);
     actor_id | first_name | last_name |     last_update     
    ----------+------------+-----------+---------------------
            2 | NICK       | WAHLBERG  | 2006-02-15 04:34:33
            3 | ED         | CHASE     | 2006-02-15 04:34:33
           60 | HENRY      | BERRY     | 2006-02-15 04:34:33
    (3 rows)
    ```