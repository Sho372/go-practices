package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gocraft/dbr"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "github.com/lib/pq" // PostgreSQL driver
)

/*
	参考ドキュメント: https://echo.labstack.com/docs/binding
*/

type (

	/*
		query - query parameter
		param - path parameter (also called route)
		header - header parameter
		json - request body. Uses builtin Go json package for unmarshalling.
		xml - request body. Uses builtin Go xml package for unmarshalling.
		form - form data. Values are taken from query and request body. Uses Go standard library form parsing.
	*/

	User struct {
		Name    string  `json:"name,omitempty" query:"name"`
		Email   string  `json:"email" query:"email"`
		Hobbies Hobbies `json:"hobbies"`
	}

	Hobby struct {
		HobbyID   int
		HobbyName string
	}

	Hobbies []Hobby

	//ビジネスロジックに渡す型。これはバインドしない
	UserDTO struct {
		Name    string
		Email   string
		Hobbies Hobbies
		IsAdmin bool
	}
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", hello)
	// /show?team=x-men&member=wolverine
	e.GET("/show", show)
	// /user?user=shohei&email=foo@example.com
	e.GET("/user", getUser)
	//
	e.POST("/postUser", PostUser)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func show(c echo.Context) error {
	// Get team and member from the query string
	team := c.QueryParam("team")
	member := c.QueryParam("member")
	return c.String(http.StatusOK, "team:"+team+", member:"+member)
}

func PostUser(c echo.Context) error {

	u := new(User)
	if err := c.Bind(u); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	// bindする構造体とビジネスロジックに投げる構造体は分けるべき
	// Load into separate struct for security
	user := UserDTO{
		Name:    u.Name,
		Email:   u.Email,
		Hobbies: u.Hobbies,
		IsAdmin: false, // avoids exposing field that should not be bound バインドすべきでないフィールドはexportすべきではない
	}

	executeSomeBusinessLogic(user)

	return c.JSON(http.StatusOK, u)
}

func getUser(c echo.Context) error {
	u := new(User)
	qp := c.QueryParams()
	fmt.Printf("QueryParamas: %s\n", qp)
	qs := c.QueryString()
	fmt.Printf("QueryString: %s\n", qs)
	pn := c.ParamNames()
	fmt.Printf("ParamNames: %s\n", pn)
	fmt.Printf("BEFORE BIND User.name: [%s], User.email: [%s]\n", u.Name, u.Email)
	if err := c.Bind(u); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, u)
}

func executeSomeBusinessLogic(user UserDTO) {

	//空配列の場合は, len:0, cap:0のHobbiesの空配列が来る
	fmt.Printf("👍user: %v\n", user)

	ids := extractIDs(user.Hobbies)
	fmt.Printf("👍ids: %v\n", ids)

	//データ取得
	executeStmt(ids)
}

func extractIDs(hobbies Hobbies) []int {

	ids := make([]int, 0, len(hobbies))

	for _, hobby := range hobbies {
		ids = append(ids, hobby.HobbyID)
	}

	return ids
}

func executeStmt(ids []int) {

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

	selectStmt := session.Select("actor_id").From("actor").
		Where(
			dbr.And(
				dbr.Eq("first_name", "NICK"),
				dbr.Neq("actor_id", ids), // NOT IN
			),
		)

	toRawQuery(*selectStmt)
	
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

func toRawQuery(stmt dbr.SelectStmt) {

	buf := dbr.NewBuffer()
	_ = stmt.Build(stmt.Dialect, buf)
	fmt.Println(buf.String())
}
