package parsed_data

import (
	"lxh027.com/xml-dbms/dbserver/parser/tokenizer"
)

type ParsedBasicData struct {
	Operation 	tokenizer.TokenType
	Target 		tokenizer.TokenType
	Name 		string

}

func (parsed *ParsedBasicData) SetBasicInfo(op tokenizer.TokenType, target tokenizer.TokenType, name string)  {
	parsed.Operation = op
	parsed.Target = target
	parsed.Name = name
}