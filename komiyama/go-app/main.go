package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gocql/gocql"
)

var session *gocql.Session

func main() {
	// Cassandra接続
	cluster := gocql.NewCluster("cassandra")
	cluster.Keyspace = "testkeyspace"
	cluster.Consistency = gocql.Quorum
	var err error
	session, err = cluster.CreateSession()
	if err != nil {
		log.Fatalf("Cassandra接続失敗: %v", err)
	}
	defer session.Close()

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/add", addHandler)
	http.HandleFunc("/delete", deleteHandler)
	fmt.Println("サーバー起動: http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	rows := ""
	iter := session.Query("SELECT id, value FROM items").Iter()
	var id gocql.UUID
	var value string
	for iter.Scan(&id, &value) {
		rows += fmt.Sprintf("<tr><td>%s</td><td>%s</td><td><form method='POST' action='/delete'><input type='hidden' name='id' value='%s'><button type='submit'>削除</button></form></td></tr>", id, value, id)
	}
	iter.Close()
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, `<html><head><meta charset="UTF-8"><title>Cassandra管理</title></head><body><h1>アイテム一覧</h1><form method='POST' action='/add'><input name='value' placeholder='値'><button type='submit'>追加</button></form><table border='1'><tr><th>ID</th><th>値</th><th>操作</th></tr>%s</table></body></html>`, rows)
}

func addHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		value := r.FormValue("value")
		id := gocql.TimeUUID()
		err := session.Query("INSERT INTO items (id, value) VALUES (?, ?)", id, value).Exec()
		if err != nil {
			log.Printf("追加失敗: %v", err)
		}
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		idStr := r.FormValue("id")
		id, err := gocql.ParseUUID(idStr)
		if err == nil {
			err = session.Query("DELETE FROM items WHERE id = ?", id).Exec()
			if err != nil {
				log.Printf("削除失敗: %v", err)
			}
		}
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
