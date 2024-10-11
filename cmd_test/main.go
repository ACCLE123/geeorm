package main

import (
	"fmt"
	"geeorm"
)


func main() {
	engine, _  := geeorm.NewEngine("sqlite3", "gee.db")
	defer engine.Close()

	s := engine.NewSession()
	s.Raw("drop table if exists User;").Exec()
	s.Raw("create table User(Name text);").Exec()
	s.Raw("create table User(Name text);").Exec()

	result, _ := s.Raw("insert into User(`name`) values (?), (?)", "Tom", "Sam").Exec()
	count, _ := result.RowsAffected()

	fmt.Printf("Exec success, %d affercted\n", count)
}