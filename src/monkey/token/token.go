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

func LookupIdent(ident string) TokenType {
    switch ident {
    case "fn":
        return FUNCTION
    case "let":
        return LET
    default:
        return IDENT
    }
}

