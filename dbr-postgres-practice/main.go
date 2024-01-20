package main

import (
	"fmt"
	"log"

	"github.com/gocraft/dbr/v2"
	_ "github.com/lib/pq" // PostgreSQL driver
)

func main() {
	// PostgreSQL接続情報
	connectionInfo := "user=postgres password=postgres dbname=sakila sslmode=disable"
	// 接続
	conn, err := dbr.Open("postgres", connectionInfo, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// dbrセッションの生成
	session := conn.NewSession(nil)

	// SELECT文を構築

	/*
		sakila=# select * from actor where first_name = 'NICK';
		 actor_id | first_name | last_name |     last_update
		----------+------------+-----------+---------------------
		        2 | NICK       | WAHLBERG  | 2006-02-15 04:34:33
		       44 | NICK       | STALLONE  | 2006-02-15 04:34:33
		      166 | NICK       | DEGENERES | 2006-02-15 04:34:33
		(3 rows)
	*/

	ids := []int64{2, 166}

	// IN
	// selectStmt := session.Select("actor_id").From("actor").Where(dbr.Eq("actor_id", ids)) // sliceの場合は、INになる
	// selectStmt := session.Select("actor_id").From("actor").
	// 	Where(
	// 		dbr.And(
	// 			dbr.Eq("first_name", "NICK"),
	// 			dbr.Eq("actor_id", ids),
	// 		),
	// 	)

	// NOT IN
	//  When value is a slice, it will be translated to `NOT IN`
	// selectStmt := session.Select("actor_id").From("actor").Where(dbr.Neq("actor_id", ids)) // sliceの場合は、NOT INになる
	selectStmt := session.Select("actor_id").From("actor").
		Where(
			dbr.And(
				dbr.Eq("first_name", "NICK"),
				dbr.Neq("actor_id", ids), // NOT IN
			),
		)

	// =
	// selectStmt := session.Select("actor_id").From("actor").Where("actor_id = ?", 10)

	// raw statment
	// selectStmt := session.SelectBySql(`select actor_id from actor where actor_id = ?`, 60)

	// クエリの実行と結果の格納
	var results []int64
	_, err = selectStmt.Load(&results)
	if err != nil {
		log.Fatal(err)
	}

	// 結果の表示
	for _, d := range results {
		fmt.Printf("d %v\n", d)
	}

}
