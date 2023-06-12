package mml

import (
	"os"
	"io"
	"fmt"
	"strings"
)

func unpack[T any](v T, err error) (T) {
	if (err != nil) {panic(err)}
	return v
}

func bog[T any](expr bool, True, False T) (T) {
	if (expr) {return True}
	return False
}

func assert(expr bool, format string, rest... interface{}) {
	if (expr) {return}
	Die("Assert Failed!\n"+fmt.Sprintf(format, rest...))
}


// !idea Char struct to see what can init and keep it up
//type Char struct {
//	Type int
//	Acceptable []int,
//}

const (
	CHAR_TYPE = iota
	//CHAR_ID = Char{iota, []int{CHAR_NAME}} // #
	CHAR_ID // #
	CHAR_CLASS // .
	CHAR_STRING // "
	CHAR_OPEN // {
	CHAR_CLOSE // }
	CHAR_EQL // =
	CHAR_EOS // End Of Segment (; or \n)
	CHAR_MULTISTR // multi-line stirng (auto <br>) `
	CHAR_CHILD // > // div>section>p
	CHAR_SPACE
	CHAR_NAME
	CHAR_NUM
)

var (
	NAME_USERS = []int{
		CHAR_ID,
		CHAR_CLASS,
	}
	NAME_TYPES = []int{
		CHAR_NAME,
		CHAR_NUM,
	}
)

var (
	CODE_TO_NAME = map[int]string{
		CHAR_ID: "CHAR_ID",
		CHAR_CLASS: "CHAR_CLASS",
		CHAR_STRING: "CHAR_STRING",
		CHAR_OPEN: "CHAR_OPEN",
		CHAR_CLOSE: "CHAR_CLOSE",
		CHAR_EQL: "CHAR_EQL",
		CHAR_EOS: "CHAR_EOS",
		CHAR_MULTISTR: "CHAR_MULTISTR",
		CHAR_CHILD: "CHAR_CHILD",
		CHAR_SPACE: "CHAR_SPACE",
		CHAR_NAME: "CHAR_NAME",
		CHAR_NUM: "CHAR_NUM",
	}
)

func byteNotForName(char byte) (bool) {
	return IinA(char, []byte("#.{;!\n,()"))
}

func byteForClosingTag(char byte) (bool) {
	return char == '}'
}

func byteForOpenningTag(char byte) (bool) {
	return char == '{'
}

func IinA[T comparable](a T, arr []T) bool {
	for i := range arr {
		if arr[i] == a {
			return true
		}
	}
	return false
}

func IinV[T comparable](a T, arr ...T) bool {
	for i := range arr {
		if arr[i] == a {
			return true
		}
	}
	return false
}

func Die(message string) {
	fmt.Fprintln(os.Stderr, message)
	os.Exit(1)
}

func readFileBytes(filename string) []byte {
	file := unpack(os.Open(filename))
	defer file.Close()
	FILE := unpack(io.ReadAll(file))
	return FILE
}

func CleanBuffer(buffer []byte) ([]byte) {
	sbuffer := string(buffer)
	// remove tabs
	sbuffer = strings.Replace(sbuffer, "\t", "", -1)
	buffer = []byte(sbuffer)
	return buffer
}

//CTX
func ReadFile(FileName string) (*Context) {
	var ctx Context
	ctx.FileName = FileName
	ctx.Buffer = CleanBuffer(readFileBytes(FileName))
	ctx.TagStack.Init(&ctx, 20)
	return &ctx
}

type Stack[T any] struct {
	Stack []*T
	Index int
	ctx *Context
	Size int
}

func (S *Stack[T]) Init(ctx *Context, size int) {
	S.Size = size
	S.Stack = make([]*T, S.Size)
	S.ctx = ctx
}

func (S *Stack[T]) Expand(increase int) {
	t := make([]*T, S.Size)
	copy(t, S.Stack)
	S.Size+=increase
	S.Stack = make([]*T, S.Size)
	copy(S.Stack, t)
}

func (S *Stack[T]) Push(thing *T) {
	if (S.Index == S.Size) {
		S.Expand(S.Size) // double size
	}
	S.Stack[S.Index] = thing
	S.Index++
}

func (S *Stack[T]) Pop() (*T) {
	if (S.Index == 0) {
		S.ctx.die("pop() on empty stack")
		return nil
	}
	S.Index--
	ret := S.Stack[S.Index]
	S.Stack[S.Index] = nil // remove reference
	return ret
}

func (ctx *Context) die(format string, rest... interface{}) {
	fmt.Printf("Error occoured while reading file '%s'\nat %d:%d\n",
		ctx.FileName, ctx.Cursor.Y, ctx.Cursor.X)
	Die(fmt.Sprintf(format, rest...))
}

func (ctx *Context) Next() (byte) {
	ctx.Cursor.Raw++
	ctx.Cursor.X++
	if (ctx.Current() == '\n') {
		ctx.Cursor.Y++
		ctx.Cursor.X=0
	}
	return ctx.Current()
}

func (ctx *Context) Name() (text string) {
	var now byte = ctx.Current()
	for ctx.CanRead() && !byteNotForName(now) {
		// die on syntax error
		if IinV(now, '!', '}') { ctx.die("names can't contain the char '%c'", now) }
		text+=string(now)
		now = ctx.Next()
	}
	return text
}

func (ctx *Context) next() (byte) {
	return ctx.Buffer[ctx.Cursor.Raw+1]
}

func (ctx *Context) Current() (byte) {
	return ctx.Buffer[ctx.Cursor.Raw]
}

func (ctx Context) CanRead(i... int) (bool) {
	if len(i) == 0 {
		return ctx.Cursor.Raw+1 < len(ctx.Buffer)
	} else if len(i) == 1 {
		return ctx.Cursor.Raw+1+i[0] < len(ctx.Buffer)
	} else {
		ctx.die("CanRead executed with more than 0 or 1 arguments, %+d", i)
		return false
	}
}

func WhichChar(char byte) (CODE int) {
	switch (char) {
	case '#':
		return CHAR_ID // #
	case '.':
		return CHAR_CLASS // .
	case '"':
		return CHAR_STRING // "
	case '{':
		return CHAR_OPEN // {
	case '}':
		return CHAR_CLOSE // }
	case '=':
		return CHAR_EQL // =
	case '\r', '\n', ';':
		return CHAR_EOS // End Of Segment
	case '`':
		return CHAR_MULTISTR // multi-line stirng (auto <br>) `
	case '>':
		return CHAR_CHILD // > // div>section>p
	case ' ':
		return CHAR_SPACE
	case '1', '2', '3', '4', '5', '6', '7', '8', '9', '0':
		return CHAR_NUM
	default:
		return CHAR_NAME
	}
	return 0
}

func ColapseChars(chars []byte) ([]Word) {
	var (
		Words = []Word{}
		last = -1 // Words.len
		Type int
		InString bool
		UsingName bool
		pos Position
	)
	for _, char := range chars {
		pos.Raw++
		pos.X++
		if char == '\n' {
			pos.Y++
			pos.X=0
		}

		if InString { // TODO: check chars[i-1] != \
			Words[last].Info = append(Words[last].Info, char)
			if WhichChar(char) == CHAR_STRING {InString = false}
			continue
		}
		if UsingName && IinA(WhichChar(char), NAME_TYPES) {
			Words[last].Info = append(Words[last].Info, char)
			continue
		}
		UsingName = false

		// no special char active
		if WhichChar(char) == Type {
			Words[last].Info = append(Words[last].Info, char)
		} else {
			Type = WhichChar(char)
			if Type == CHAR_STRING {
				InString = true
			} else if IinA(Type, NAME_USERS) {
				UsingName = true
			}
			Words = append(Words, Word{Type, []byte{char}, pos})
			last++
		}
	}
	return Words
}

type Position struct {
	Raw, X, Y int
}

// string with type
type Word struct {
	Type int
	Info []byte
	Pos Position
}

// sanatize word -> token
type Token struct {
	Type int
	// Text is raw info. Name & Value are sanatized division
	Text, Name, Value string
	Pos Position // add *Pos,*Span / *PosStart/*PosEnd after creation?
}

// removes " from strings
func cleanWordString(w Word) Word {
	assert(w.Type == CHAR_STRING, "cleanWordString only accepts words of type CHAR_STRING\n")
	w.Info = w.Info[1:len(w.Info)-1]
	return w
}

func getContent(w Word) string {
	if w.Type == CHAR_STRING {
		w = cleanWordString(w)
	}
	return string(w.Info)
}

// variadic internal ColapseWords
// vcolapse(type, ...words) -> token
func vcolapseWords(Type int, Words ...Word) (Token) {
	switch(Type) {
	case CHAR_EQL:
		name := string(Words[0].Info)
		value := string(Words[1].Info)
		return Token{CHAR_EQL, name+"="+value, name, value, Words[0].Pos}
	}
	return Token{}
}

/*
func WordsToTokens(Words []Word) {
	var lasti = len(Words)-1
	var Tokens = []Token{}
	var last = -1 // Tokens.len

	for i, word := range Words {
	switch (word.Type) {

	}
	}
}
*/

type Context struct {
	Cursor Position
	FileName string
	Buffer []byte
	TagStack Stack[Tag]
}

func DebugChar(char byte) {
	code := WhichChar(char)
	fmt.Printf("'%c' u%d %s [%d]\n",
		char, char,
		CODE_TO_NAME[code],
		code,
	)
}

func DebugWord(word Word) {
	fmt.Printf("%s %s [%d] @ %d,%d\n",
		//bog(word.Type==CHAR_EOS, "EOS", "'"+string(word.Info)+"'"),
		word.Info,
		CODE_TO_NAME[word.Type],
		word.Type,
		word.Pos.Y,
		word.Pos.X,
	)
}

func DebugToken(tkn Token) {
	fmt.Printf(`%s: {
- "%s",
- %s:%s,
- @ %d,%d,
}
`, CODE_TO_NAME[tkn.Type],
		tkn.Text,
		tkn.Name,
		tkn.Value,
		tkn.Pos.Y,
		tkn.Pos.X,
	)
}

func (ctx *Context) Parse() () {
	//words := ColapseChars(ctx.Buffer)
	//for _, word := range words {
	//	DebugWord(word)
	//}
	p:=Position{}
	DebugToken(
		vcolapseWords(CHAR_EQL,
			Word{CHAR_NAME, []byte("FileName"), p},
			Word{CHAR_STRING, []byte("\"out.csv\"") ,p},
		))
}

