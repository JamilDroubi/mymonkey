package lexer

import (
    // "testing"

    "mymonkey/token"

)

type Lexer struct {
    input        string
    position     int
    readPosition int
    ch           byte
}

func (l *Lexer) readChar() {
    if l.readPosition >= len(l.input) {
        l.ch = 0
        // Set it to EOF if readPosition is >= the length of the input
    } else {
        l.ch = l.input[l.readPosition]
    }
    l.position = l.readPosition
    l.readPosition++
}

func New(input string) *Lexer {
    l := &Lexer{input: input}
    l.readChar()
    return l
}

// func TestNextToken(t *testing.T) {
//     input := `=+(){},;`
//
//     tests := []struct {
//         expectedType     token.TokenType
//         expectedLiteral string
//     }{
//         {token.ASSIGN,    "="},
//         {token.PLUS,      "+"},
//         {token.LPAREN,    "("},
//         {token.RPAREN,    ")"},
//         {token.LBRACE,    "{"},
//         {token.RBRACE,    "}"},
//         {token.COMMA,     ","},
//         {token.SEMICOLON, ";"},
//         {token.EOF,        ""},
//     }
//
//     l := New(input)
//     for i, tt := range tests {
//         tok := l.NextToken()
//         if tok.Type != tt.expectedType {
//             t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
//             i, tt.expectedType, tok.Type)
//         }
//     }
// }

func (l *Lexer) NextToken() token.Token {
    var tok token.Token
    switch l.ch {
    case '=': tok = newToken(token.ASSIGN, l.ch)
    case ';': tok = newToken(token.SEMICOLON, l.ch)
    case '(': tok = newToken(token.LPAREN, l.ch)
    case ')': tok = newToken(token.RPAREN, l.ch)
    case ',': tok = newToken(token.COMMA, l.ch)
    case '+': tok = newToken(token.PLUS, l.ch)
    case '{': tok = newToken(token.LBRACE, l.ch)
    case '}': tok = newToken(token.RBRACE, l.ch)
    case  0 : tok.Literal = ""
              tok.Type = token.EOF
    default :
        if isLetter(l.ch) {
            tok.Literal = l.readIdentifier()
            tok.Type    = token.LookupIdent(tok.Literal)
            return tok
        } else {
            tok = newToken(token.ILLEGAL, l.ch)
        }
    }

    l.readChar()
    return tok
}

func (l *Lexer) readIdentifier() string {
    position := l.position
    for isLetter(l.ch) {
        l.readChar()
    }
    return l.input[position:l.position]
}

func isLetter(ch byte) bool {
    return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'z' || ch == '_'
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
    return token.Token{Type: tokenType, Literal: string(ch)}
}