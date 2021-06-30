package test

import (
	"log"
	"lxh027.com/xml-dbms/dbserver/parser/tokenizer"
	"testing"
)

var sqls map[string]string

func init()  {
	sqls = map[string]string{
		"createDataBaseSql": "create database test",
		"dropDataBaseSql": "drop database test",
		"createTableSql": "create table test ( aa int, bb string )",
		"dropTableSql": "drop table test",
		"createViewSql": "create view test (select * from testT)",
		"dropViewSql": "drop view test",
	}
}

func TestTokenParse(t *testing.T)  {
	for name, sql := range sqls {
		log.Printf("name: %v \t sql: %v \n", name, sql)
		sqlToken, err := tokenizer.ParseSql(sql)
		if err != nil {
			log.Println(err.Error())
		} else {
			for _, token := range sqlToken.Tokens {
				log.Printf("token: %v \t type: %v\n", token.Name, token.Type)
			}
		}
	}
}