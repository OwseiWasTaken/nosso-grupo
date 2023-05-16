package mml

import (
	"os"
	"io"
	"fmt"
	"errors"
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
	panic(errors.New(message))
}

func ReadFileBytes(filename string) []byte {
	file := unpack(os.Open(filename))
	defer file.Close()
	FILE := unpack(io.ReadAll(file))
	return FILE
}

func ReadFile(FileName string) (Context) {
	var ctx Context
	ctx.FileName = FileName
	ctx.Buffer = ReadFileBytes(FileName)
	//for ctx.CanRead() {
	//	fmt.Println(ctx.Next(), ctx.Cursor)
	//}
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
	var now byte
	text+=string(ctx.Current()) // add letter already found
	// keep addind until syntax error, or end of string
	for ctx.CanRead() && !IinA(now, '#', '.', '{', ';') {
		now = ctx.Next()
		// syntax error
		if IinA(now, '!', '}') { ctx.die("names can't contain the char '%c'", now) }
		text+=string(now)
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

type Context struct {
	Cursor struct {Raw, X, Y int}
	FileName string
	Buffer []byte
}
