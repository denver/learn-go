package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
)

var (
	debug         = flag.Bool("debug", false, "enable debugging")
	password      = flag.String("password", "", "the database password")
	port     *int = flag.Int("port", 1434, "the database port")
	server        = flag.String("server", "localhost", "the database server")
	user          = flag.String("user", "", "the database user")
	database      = flag.String("database", "", "the database name")
)

func main() {
	flag.Parse()

	if *debug {
		fmt.Printf(" password:%s\n", *password)
		fmt.Printf(" port:%d\n", *port)
		fmt.Printf(" server:%s\n", *server)
		fmt.Printf(" user:%s\n", *user)
	}

	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s", *server, *user, *password, *port, *database)
	if *debug {
		fmt.Printf(" connString:%s\n", connString)
	}
	conn, err := sql.Open("mssql", connString)
	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
	}
	defer conn.Close()

	rows, err := conn.Query(`SELECT * FROM tblUser where name = 'name'`)
	if err != nil {
		fmt.Println("Failed to run query", err)
		return
	}

	accountNames := []string{}
	for rows.Next() {
		var fname string
		rows.Scan(&fname)
		accountNames = append(accountNames, fname)
	}
	fmt.Println(accountNames)
}
