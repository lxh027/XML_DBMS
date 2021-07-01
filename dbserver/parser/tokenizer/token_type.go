package tokenizer

type Token struct {
	Name	string
	Type 	TokenType
}

type TokenType string

const (
	Variable 	= "variable"
	Sql 		= "sql"
)

const (
	Create 		= "create"
	Drop 		= "drop"
	Use  		= "use"
	Show 		= "show"
)

const (
	Database 	= "database"
	Table 		= "table"
	View 		= "view"
	Index 		= "index"
)

const (
	LeftCell 	= "("
	RightCell 	= ")"
	Comma 		= ","
)

const (
	Int 		= "int"
	String 		= "string"
)
