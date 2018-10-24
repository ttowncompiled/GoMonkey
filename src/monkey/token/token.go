package token

type TokenType string

type Token struct {
    Type        TokenType
    Literal     string
}

const (
    // Metacharacters
    ILLEGAL     = "ILLEGAL" // unrecognized character
    EOF         = "EOF"     // end-of-file

    // Identifiers + Literals
    IDENT       = "IDENT"   // add, foobar, x, y, ...
    INT         = "INT"     // 1343456
    LET         = "let"
    FUNCTION    = "fn"

    // Operators
    ASSIGN      = "="
    PLUS        = "+"

    // Delimiters
    COMMA       = ","
    SEMICOLON   = ";"

    // Collections + Scopes
    LPAREN      = "("
    RPAREN      = ")"
    LBRACE      = "{"
    RBRACE      = "}"
)

var keywords = map[string]TokenType {
    "fn":       FUNCTION,
    "let":      LET,
}

func LookupIdent(ident string) TokenType {
    if tok, ok := keywords[ident]; ok {
        return tok
    }
    return IDENT
}

