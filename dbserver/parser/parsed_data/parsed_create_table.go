package parsed_data

import (
	"lxh027.com/xml-dbms/dbserver/data/model"
	"lxh027.com/xml-dbms/dbserver/parser/tokenizer"
)

type ParsedCreateTable struct {
	ParsedBasicData
	Columns 	[]model.Column
}

func (parsed *ParsedCreateTable) SetBasicInfo(op tokenizer.TokenType, target tokenizer.TokenType, name string)  {
	parsed.Operation = op
	parsed.Target = target
	parsed.Name = name
}