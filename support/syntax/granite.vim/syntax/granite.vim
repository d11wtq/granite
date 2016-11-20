if exists('b:current_syntax')
    finish
endif

syntax keyword graniteKeyword import
syntax keyword graniteKeyword function
syntax keyword graniteKeyword record
highlight link graniteKeyword Keyword

syntax keyword graniteCondition case
syntax keyword graniteCondition else
highlight link graniteCondition Conditional

syntax region graniteString start='"' skip='\\.' end='"'
syntax region graniteString start="'" skip='\\.' end="'"
highlight link graniteString String

syntax region graniteComment start=";" end="\n"
highlight link graniteComment Comment

syntax match graniteIdentifier '\v[A-Za-z][A-Za-z0-9_]*[?!]?'
highlight link graniteIdentifier Identifier

syntax match graniteConstant '\v[A-Z][A-Za-z0-9_]*[?!]?'
highlight link graniteConstant Constant

syntax match graniteNumber '\v[0-9]+'
syntax match graniteNumber '\v0x[0-9A-Fa-f]+'
syntax match graniteNumber '\v0d[0-9]+'
syntax match graniteNumber '\v0o[0-8]+'
syntax match graniteNumber '\v0b[0-1]+'
highlight link graniteNumber Number

syntax match graniteOperator '='
syntax match graniteOperator ':'
syntax match graniteOperator '*'
syntax match graniteOperator '/'
syntax match graniteOperator '+'
syntax match graniteOperator '-'
syntax match graniteOperator '\.'
highlight link graniteOperator Operator

let b:current_syntax = 'granite'
