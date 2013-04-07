package core

import (

)

type LexerMap map[string]int

type Lexer interface{
	//StringTokenizer
	Consume();
	ConsumeK(k int)
    HasMoreChars() bool
    IsAlpha(ch byte) bool
    IsDigit(ch byte) bool
    IsHexDigit(ch byte) bool
    LookAhead() (byte, error)
    LookAheadK(k int) (byte, error)
    PeekLine() string
    GetLine() string
    GetNextTokenByDelim(delim byte) string
    
	//Lexer
	SetLexerName(lexerName string)
	GetLexerName() string
	AddKeyword(name string, value int)
    AddLexer(lexerName string) LexerMap
    ByteStringNoComma() string
    ByteStringNoSemicolon() string
    CharAsString(nchars int) string
    Comment() (s string, ParseException error)
    CurrentLexer() LexerMap
    GetBuffer() string
    GetNextId() string
    GetNextToken() *Token
    GetPtr() int
    GetRest() string
    GetString(c byte) (s string, ParseException error)
    LookupToken(value int) string
    MarkInputPosition() int
    Match(tok int) (t *Token, ParseException error)
    Number() (s string, ParseException error)
    PeekNextId() string
    PeekNextToken() *Token
    PeekNextTokenK(ntokens int) []*Token
    QuotedString() (s string, ParseException error)
    RewindInputPosition(position int)
    SPorHT()
    SelectLexer(lexerName string)
    StartsId() bool
    Ttoken() string
    TtokenAllowSpace() string
}