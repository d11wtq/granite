package parser

//go:generate -command yacc go tool yacc
//go:generate yacc -o bijou.go -p Bijou bijou.y
