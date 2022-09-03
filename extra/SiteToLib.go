package main

include "gutil"

func main(){
	InitGu()
	var FILE []string
	var OUTFILE string
	var line string
	// TODO(3): get a proper name
	var FilenameOut = "out.c"
	var FilenameIn = "../site.html"

	FILE = strings.Split(ReadFile(FilenameIn), "\n")

	for i:=0;i<len(FILE);i++ {
		line = FILE[i]
		// pular linha
		if len(line) <= 1 {
			continue
		}

		//remover \n
		if line[len(line)-1] == 13 {
			line = line[:len(line)-1]
		}
		line = strings.Replace(line, "\n", "", -1)
		line = strings.Replace(line, "\"", "\\\"", -1)
		line = "\""+line+"\""
		line = "HTML+="+line+"\n"
		OUTFILE += line
	}

	WriteFile(FilenameOut, OUTFILE+"\n")

	exit(0)
}
