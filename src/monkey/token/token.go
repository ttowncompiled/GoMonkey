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

