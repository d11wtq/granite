package parser

//go:generate -command yacc go tool yacc
//go:generate yacc -o granite.go -p Granite granite.y
