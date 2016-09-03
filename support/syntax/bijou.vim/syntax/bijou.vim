if exists('b:current_syntax')
    finish
endif

syntax keyword bijouKeyword import
syntax keyword bijouKeyword function
syntax keyword bijouKeyword record
highlight link bijouKeyword Keyword

syntax keyword bijouCondition case
syntax keyword bijouCondition else
highlight link bijouCondition Conditional

syntax region bijouString start='"' skip='\\.' end='"'
syntax region bijouString start="'" skip='\\.' end="'"
highlight link bijouString String

syntax region bijouComment start=";" end="\n"
highlight link bijouComment Comment

syntax match bijouIdentifier '\v[A-Za-z][A-Za-z0-9_]*[?!]?'
highlight link bijouIdentifier Identifier

syntax match bijouConstant '\v[A-Z][A-Za-z0-9_]*[?!]?'
highlight link bijouConstant Constant

syntax match bijouNumber '\v[0-9]+'
syntax match bijouNumber '\v0x[0-9A-Fa-f]+'
syntax match bijouNumber '\v0d[0-9]+'
syntax match bijouNumber '\v0o[0-8]+'
syntax match bijouNumber '\v0b[0-1]+'
highlight link bijouNumber Number

syntax match bijouOperator '='
syntax match bijouOperator ':'
syntax match bijouOperator '*'
syntax match bijouOperator '/'
syntax match bijouOperator '+'
syntax match bijouOperator '-'
syntax match bijouOperator '\.'
highlight link bijouOperator Operator

let b:current_syntax = 'bijou'
