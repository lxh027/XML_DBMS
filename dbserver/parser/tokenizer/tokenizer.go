package tokenizer

import (
	"errors"
	"fmt"
	"strings"
)

type SqlTokens struct {
	Sql 	string
	Tokens 	[]Token
}
var ok bool

func ParseSql(sql string) (*SqlTokens, error) {
	sqlTokens := SqlTokens{
		Sql: strings.ToLower(sql),
		Tokens: make([]Token, 0),
	}
	err := sqlTokens.parseMove(0)
	return &sqlTokens, err
}

// parse Create Drop
func (sqlTokens *SqlTokens) parseMove(pos int) error  {
	if pos, ok = sqlTokens.hasNext(pos); !ok {
		return errors.New(fmt.Sprintf("Syntax Error in %v", pos))
	}
	sql := sqlTokens.Sql

	if sqlTokens.equals(pos, Create) {
		sqlTokens.Tokens = append(sqlTokens.Tokens, Token{Name: sql[pos:pos+len(Create)], Type: Create})
		return sqlTokens.parseTarget(pos+ len(Create))
	} else if sqlTokens.equals(pos, Drop) {
		sqlTokens.Tokens = append(sqlTokens.Tokens, Token{Name: sql[pos:pos+len(Drop)], Type: Drop})
		return sqlTokens.parseTarget(pos+ len(Drop))
	} else if sqlTokens.equals(pos, Use) {
		sqlTokens.Tokens = append(sqlTokens.Tokens, Token{Name: sql[pos:pos+len(Use)], Type: Use})
		pos += len(Use)
		return sqlTokens.parseName(&pos)
	}  else if sqlTokens.equals(pos, Show) {
		sqlTokens.Tokens = append(sqlTokens.Tokens, Token{Name: sql[pos:pos+len(Show)], Type: Show})
		pos += len(Use)
		return sqlTokens.parseName(&pos)
	}
	return errors.New(fmt.Sprintf("Syntax Error in %v", pos))
}

// parse Database Table View
func (sqlTokens *SqlTokens) parseTarget(pos int) error {
	if pos, ok = sqlTokens.hasNext(pos); !ok {
		return errors.New(fmt.Sprintf("Syntax Error in %v", pos))
	}

	sql := sqlTokens.Sql
	if sqlTokens.equals(pos, Database) {
		sqlTokens.Tokens = append(sqlTokens.Tokens, Token{Name: sql[pos:pos+len(Database)], Type: Database})
		pos += len(Database)
		return sqlTokens.parseName(&pos)
	} else if sqlTokens.equals(pos, Table) {
		sqlTokens.Tokens = append(sqlTokens.Tokens, Token{Name: sql[pos:pos+len(Table)], Type: Table})
		pos += len(Table)
		if sqlTokens.Tokens[0].Type == Create {
			return sqlTokens.parseTableInfo(pos)
		} else {
			return sqlTokens.parseName(&pos)
		}
	} else if sqlTokens.equals(pos, View) {
		sqlTokens.Tokens = append(sqlTokens.Tokens, Token{Name: sql[pos:pos+len(View)], Type: View})
		pos += len(View)
		if sqlTokens.Tokens[0].Type == Create {
			return sqlTokens.parseViewInfo(pos)
		} else {
			return sqlTokens.parseName(&pos)
		}
	} else if sqlTokens.equals(pos, Index) {
		return errors.New(fmt.Sprintf("index operation is not supported"))
	}
	return errors.New(fmt.Sprintf("Syntax Error in %v", pos))
}



// parse Table Info
func (sqlTokens *SqlTokens) parseTableInfo(pos int) error {
	if err := sqlTokens.parseName(&pos); err != nil {
		return err
	}
	if err := sqlTokens.parseTokenType(&pos, LeftCell); err != nil {
		return err
	}
	sql := sqlTokens.Sql
	// parse Table Columns Info
	for pos < len(sql) && !sqlTokens.equals(pos, RightCell) {
		if err := sqlTokens.parseName(&pos); err != nil {
			return err
		}
		if errInt := sqlTokens.parseTokenType(&pos, Int); errInt != nil {
			if errStr := sqlTokens.parseTokenType(&pos, String); errStr != nil {
				return errStr
			}
		}
		if sqlTokens.equals(pos, RightCell) {
			break
		}
		if err := sqlTokens.parseTokenType(&pos, Comma); err != nil {
			return err
		}
	}

	if err := sqlTokens.parseTokenType(&pos, RightCell); err != nil {
		return err
	}
	return nil
}

// parse View Info
func (sqlTokens *SqlTokens) parseViewInfo(pos int) error{
	if err := sqlTokens.parseName(&pos); err != nil {
		return err
	}
	if err := sqlTokens.parseTokenType(&pos, LeftCell); err != nil {
		return err
	}
	sql := sqlTokens.Sql
	// skip space
	if pos, ok = sqlTokens.hasNext(pos); !ok {
		return errors.New(fmt.Sprintf("Syntax Error in %v", pos))
	}
	// TODO parse select info view
	l := 0
	for pos+l < len(sql) && !sqlTokens.equals(pos+l, RightCell) {
		l++
	}
	sqlTokens.Tokens = append(sqlTokens.Tokens, Token{Name: sql[pos:pos+l], Type: Sql})
	pos += l
	if err := sqlTokens.parseTokenType(&pos, RightCell); err != nil {
		return err
	}
	return nil
}

// parse DataType Info
func (sqlTokens *SqlTokens) parseTokenType(pos *int, tokenType TokenType) error {
	if *pos, ok = sqlTokens.hasNext(*pos); !ok {
		return errors.New(fmt.Sprintf("Syntax Error in %v", *pos))
	}
	sql := sqlTokens.Sql
	if sqlTokens.equals(*pos, tokenType) {
		sqlTokens.Tokens = append(sqlTokens.Tokens, Token{Name: sql[*pos:*pos+len(tokenType)], Type: tokenType})
		*pos += len(tokenType)
		return nil
	} else {
		return errors.New(fmt.Sprintf("Syntax Error in %v", pos))
	}
}

// parse Variable Info
func (sqlTokens *SqlTokens) parseName(pos *int) error {
	if *pos, ok = sqlTokens.hasNext(*pos); !ok {
		return errors.New(fmt.Sprintf("Syntax Error in %v", *pos))
	}
	sql := sqlTokens.Sql
	if l, ok := sqlTokens.isVariable(*pos); ok {
		sqlTokens.Tokens = append(sqlTokens.Tokens, Token{Name: sql[*pos:*pos+l], Type: Variable})
		*pos += l
		return nil
	}
	return errors.New(fmt.Sprintf("Syntax Error in %v", *pos))
}

func (sqlTokens *SqlTokens) hasNext(pos int) (int, bool) {
	sql := sqlTokens.Sql
	for pos < len(sql) && letterSpace(sql[pos]){
		pos++
	}
	return pos, pos < len(sql)
}

func (sqlTokens *SqlTokens) equals(pos int, token TokenType) bool {
	if pos, ok = sqlTokens.hasNext(pos); !ok {
		return false
	}
	sql := sqlTokens.Sql
	if len(sql)-pos < len(token) {
		return false
	}
	for i := 0; i < len(token); i++ {
		if token[i] != sql[pos+i] {
			return false
		}
	}
	return true
}

func (sqlTokens *SqlTokens) isVariable(pos int) (int, bool) {
	if pos, ok = sqlTokens.hasNext(pos); !ok {
		return 0, false
	}
	sql := sqlTokens.Sql
	// first char should be _ or A, a..
	if !isLetter(sql[pos]) && !isUnderline(sql[pos]) {
		return 0, false
	}
	l := 1
	for pos+l < len(sql) {
		if letterOk(sql[pos+l]) {
			l++
		} else if letterSpace(sql[pos+l]) || letterSymbol(sql[pos+l]) {
			return l, true
		} else {
			break
		}
	}
	if pos+l == len(sql) {
		return l, true
	}
	return 0, false
}

func letterSymbol (letter byte) bool {
	return letter == '(' || letter == ')' || letter == ','
}

func letterSpace (letter byte) bool {
	return letter == ' ' || letter == '\t' || letter == '\n' || letter == '\r'
}

func letterOk(letter byte) bool {
	return isLetter(letter) || isNumber(letter) || isUnderline(letter)
}

func isLetter(letter byte) bool {
	return (letter >= 'A' && letter <= 'Z') || (letter >= 'a' && letter <= 'z')
}

func isNumber(letter byte) bool {
	return letter >= '0' && letter <= '9'
}

func isUnderline(letter byte) bool {
	return letter == '_'
}