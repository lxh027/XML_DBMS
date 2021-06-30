package runtime

import (
	"encoding/xml"
	"io/ioutil"
	"log"
	"lxh027.com/xml-dbms/config"
	"lxh027.com/xml-dbms/dbserver/data/model"
)

var Server *model.Server
var Databases map[string]*model.DataBase
var Tables map[string]*model.AllData
//var indexes map[string]*model.Indexes

func LoadDataFromXML()  {
	// 加载根信息
	filePos := config.DbConfig.Root
	xmlData, err := ioutil.ReadFile(filePos)
	if err != nil {
		log.Panicf(err.Error())
	}
	err = xml.Unmarshal(xmlData, Server)
	if err != nil {
		log.Panicf(err.Error())
	}
	log.Printf("加载根信息成功: %v", Server)
	Databases = make(map[string]*model.DataBase)
	Tables = make(map[string]*model.AllData)
	//indexes = make(map[string]*model.Indexes)
	for _, databaseInfo := range Server.DataBases {
		xmlData, err := ioutil.ReadFile(databaseInfo.Location)
		if err != nil {
			log.Panicf(err.Error())
		}
		var database model.DataBase
		err = xml.Unmarshal(xmlData, &database)
		if err != nil {
			log.Panicf(err.Error())
		}
		log.Printf("加载数据库信息成功: %v", database)
		Databases[databaseInfo.Name] = &database
		// 加载表信息
		for _, tableInfo := range database.Tables {
			xmlData, err := ioutil.ReadFile(tableInfo.Location)
			if err != nil {
				log.Panicf(err.Error())
			}
			var data model.AllData
			err = xml.Unmarshal(xmlData, &data)
			if err != nil {
				log.Panicf(err.Error())
			}
			log.Printf("加载数据信息成功: %v", data)
			Tables[databaseInfo.Name+"."+tableInfo.Name] = &data
			// 构造索引
			//indexes[databaseInfo.Name+"."+tableInfo.Name] = buildIndex(databaseInfo.Name, &tableInfo)
		}
	}
}


func buildIndex(databaseName string, table *model.Table) *model.Indexes  {
	indexes := make(model.Indexes, 0)
	for _, index := range table.Indexes {
		for indexId, column := range table.Columns {
			if index.Name == column.Name {
				tableIndex := make(map[string]int)
				for i, row := range Tables[databaseName+"."+table.Name].Rows {
					tableIndex[row.Data[indexId]] = i
				}
				indexes = append(indexes, model.Index{Name: index.Name, IndexMap: tableIndex})
				break
			}
		}
	}
	return &indexes
}
