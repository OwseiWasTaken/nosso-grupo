package main

include "gutil"

func main(){
	InitGu()
	var FILE []string
	var OUTFILE string
	var line string
	var FilenameOut = "./sitelib.c"
	var FilenameTemplate = "./template.c"
	var FilenameIn = "../site.html"
	var temp []string

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
		line = "HTML+="+line+";\n"
		OUTFILE += line
	}

	FILE = strings.Split(ReadFile(FilenameTemplate), "\n")

	for i:=0;i<len(FILE);i++ {
		if FILE[i] == "////HTMLMAKE" {
			temp = FILE[i+1:]
			FILE = FILE[:i+1]
			FILE[i] = "\n"+OUTFILE
			FILE = append(FILE, temp...)
		}
	}

	OUTFILE = strings.Join(FILE, "\n")

	WriteFile(FilenameOut, OUTFILE)

	exit(0)
}
