package database

import (
	"database/sql"
	"log"
    "fmt"
    "bytes"
    "os"
)

const DB_FILE = "./db/chatter.sqlite3"
const UPTDB_FILE = "./db/db.sql"

func Open() *sql.DB {
	db, err := sql.Open("sqlite", DB_FILE)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func InitDB() {
	db, err := sql.Open("sqlite", DB_FILE)
	defer db.Close()
	if err != nil {
		panic(fmt.Sprintf("WE DONT HAVE A DATABASE, BYE: %s", err))
	}

	uptfile, err := os.ReadFile(UPTDB_FILE)
	if err != nil {
		log.Printf("error init db: failed reading update file: %s", err)
		return
	}

	cmds := bytes.Split(uptfile, []byte{';'})
    cmds = cmds[:len(cmds)-1]
	for _, cmd := range cmds {
        log.Printf("DB: Executing: %s\n", cmd)
        _, err := db.Exec(string(cmd))
        if err != nil {
            log.Printf("Execution failed: %s\n", err)
        }
	}
}
