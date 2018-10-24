package lexer

import "monkey/token"

type Lexer struct {
    input           string
    position        int     // current position in input (points to current char)
    readPosition    int     // current reading position in input (after current char)
    ch              byte    // current char under examination
}

func New(input string) *Lexer {
    l := &Lexer{ input: input }
    l.readChar()
    return l
}

func (l *Lexer) readChar() {
    if l.readPosition >= len(l.input) {
        l.ch = 0
    } else {
        l.ch = l.input[l.readPosition]
    }
    l.position = l.readPosition
    l.readPosition += 1
}

func (l *Lexer) NextToken() (tok token.Token) {
    l.skipWhitespace()

    switch l.ch {
    case '=':
        tok = newToken(token.ASSIGN, l.ch)
    case '+':
        tok = newToken(token.PLUS, l.ch)
    case ',':
        tok = newToken(token.COMMA, l.ch)
    case ';':
        tok = newToken(token.SEMICOLON, l.ch)
    case '(':
        tok = newToken(token.LPAREN, l.ch)
    case ')':
        tok = newToken(token.RPAREN, l.ch)
    case '{':
        tok = newToken(token.LBRACE, l.ch)
    case '}':
        tok = newToken(token.RBRACE, l.ch)
    default:
        if isLetter(l.ch) {
            tok.Literal = l.readIdentifier();
            tok.Type = token.LookupIdent(tok.Literal)
            return
        } else if isDigit(l.ch) {
            tok.Type = token.INT
            tok.Literal = l.readNumber()
            return
        } else {
            tok = newToken(token.ILLEGAL, l.ch)
        }
    }
    l.readChar()
    return
}

func (l *Lexer) readIdentifier() string {
    pos := l.position
    for isLetter(l.ch) {
        l.readChar()
    }
    return l.input[pos:l.position]
}

func (l *Lexer) readNumber() string {
    pos := l.position
    for isDigit(l.ch) {
        l.readChar()
    }
    return l.input[pos:l.position]
}

func (l *Lexer) skipWhitespace() {
    for isWhitespace(l.ch) {
        l.readChar()
    }
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
    return token.Token{
        Type:       tokenType,
        Literal:    string(ch),
    }
}

func isLetter(ch byte) bool {
    return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
    return '0' <= ch && ch <= '9'
}

func isWhitespace(ch byte) bool {
    return ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r'
}

