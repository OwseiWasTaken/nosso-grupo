package main

include "gutil"

func ReadBin (s string) (int) {
	var c int = 0
	for i:=0;i<len(s);i++ {
		c=c<<1
		if s[i] == '1' {
			c++
		}
	}
	return c
}

func main(){
	InitGu()
	var fl []string = strings.Split(ReadFile("Vids.txt"), "\n")
	var VidFiles []string = make([]string, len(fl))
	var line string
	var sl []string
	var c int

	for i:=0;i<len(fl);i++ {
		line = fl[i]
		if len(line) < 6 { continue }
		sl = strings.Split(line, ":")
		if len(sl) != 2 { continue }
		c = ReadBin(sl[0])
		VidFiles[c-1] = sl[1]
	}
	PS(VidFiles)
	exit(0)
}
