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

func IinA[T comparable](a T, arr ...T) bool {
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
	// remove spaces
	sbuffer = strings.Replace(sbuffer, " ", "", -1)
	// remove tabs
	sbuffer = strings.Replace(sbuffer, "\t", "", -1)
	// replace \n with ;
	sbuffer = strings.Replace(sbuffer, "\n", ";", -1)
	buffer = []byte(sbuffer)
	return buffer
}

func ReadFile(FileName string) (Context) {
	var ctx Context
	ctx.FileName = FileName
	ctx.Buffer = CleanBuffer(readFileBytes(FileName))
	ctx.TagStack.Init()
	return ctx
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
	for ctx.CanRead() && !IinA(now, '#', '.', '{', ';') {
		// die on syntax error
		if IinA(now, '!', '}') { ctx.die("names can't contain the char '%c'", now) }
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
	} else {
		return ctx.Cursor.Raw+1+i[0] < len(ctx.Buffer)
	}
}

type Stack[T any] struct {
	stack []*T
	index int
}

func (S *Stack[T]) Init() {
	S.stack = make([]*T, 20)
}

func (S *Stack[T]) Push(thing *T) {
	S.stack[S.index] = thing
	S.index++
}

type Context struct {
	Cursor struct {Raw, X, Y int}
	FileName string
	Buffer []byte
	TagStack Stack[Tag]
}
