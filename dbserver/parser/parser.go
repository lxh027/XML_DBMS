package parser

import (
	"errors"
	"fmt"
	"lxh027.com/xml-dbms/dbserver/data/model"
	"lxh027.com/xml-dbms/dbserver/parser/parsed_data"
	"lxh027.com/xml-dbms/dbserver/parser/tokenizer"
)

func ParseSql(sql string) (interface{}, error) {
	tokens, err := tokenizer.ParseSql(sql)
	if err != nil {
		return nil, err
	}
	switch tokens.Tokens[0].Type {
	case tokenizer.Create:
		{
			switch tokens.Tokens[1].Type {
			case tokenizer.Database:
				{
					var parsed parsed_data.ParsedBasicData
					parsed.SetBasicInfo(tokens.Tokens[0].Type, tokens.Tokens[1].Type, tokens.Tokens[2].Name)
					return &parsed, nil
				}
			case tokenizer.View:
				{
					var parsed parsed_data.ParsedCreateView
					parsed.SetBasicInfo(tokens.Tokens[0].Type, tokens.Tokens[1].Type, tokens.Tokens[2].Name)
					parsed.Sql = tokens.Tokens[4].Name
					return &parsed, nil
				}
			case tokenizer.Table:
				{
					var parsed parsed_data.ParsedCreateTable
					parsed.SetBasicInfo(tokens.Tokens[0].Type, tokens.Tokens[1].Type, tokens.Tokens[2].Name)
					columns := make([]model.Column, 0)
					for i := 4; tokens.Tokens[i].Type != tokenizer.RightCell; i+=2 {
						columns = append(columns, model.Column{Name: tokens.Tokens[i].Name, Type: tokens.Tokens[i+1].Name})
						if tokens.Tokens[i+2].Type == tokenizer.Comma {
							i++
						}
					}
					parsed.Columns = columns
					return &parsed, nil
				}
			}
		}
	case tokenizer.Drop:
		{
			var parsed parsed_data.ParsedBasicData
			parsed.SetBasicInfo(tokens.Tokens[0].Type, tokens.Tokens[1].Type, tokens.Tokens[2].Name)
			return &parsed, nil
		}
	case tokenizer.Use:
		{
			var parsed parsed_data.ParsedBasicData
			parsed.SetBasicInfo(tokens.Tokens[0].Type, tokenizer.Use, tokens.Tokens[1].Name)
			return &parsed, nil
		}
	default:
		return nil, errors.New(fmt.Sprintf("parse sql error"))
	}
	return nil, errors.New(fmt.Sprintf("parse sql error"))
}