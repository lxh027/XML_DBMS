package parsed_data

import (
	"lxh027.com/xml-dbms/dbserver/parser/tokenizer"
)

type ParsedCreateView struct {
	ParsedBasicData
	Sql 	string
}

func (parsed *ParsedCreateView) SetBasicInfo(op tokenizer.TokenType, target tokenizer.TokenType, name string)  {
	parsed.Operation = op
	parsed.Target = target
	parsed.Name = name
}