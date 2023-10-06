"use strict";

const Replaceble = [
	["<", "&lt;"],
	[">", "&gt;"],
	[/^{$/gm, "<div>" ],
	[/^{#(.*?)(\.(.*))?$/gm, "<div id=\"$1\" class=\"$3\">"],
	[/^{\.(.*?)#(.*)$/gm, "<div class=\"$1\" id=\"$2\">"],
	[/^{\.(.*)$/gm, "<div class=\"$1\">"],
	[/^}$/gm, "</div>"],
	[/^(\w+)``$/gm, "<code class=\"$1\">"],
	[/^``$/gm, "</code>"],
	[/\B\[\[(C|PHP|JS|CSS|MD|SQL)\/(.*)\]\]\B/gm, "<a href=\"/articles/$1/$2.html\">$2</a>" ],
	[/\B\[(.*?)\]\((.*?)\)\B/gm, "<a href=\"\/articles/$2.html\">$1</a>" ],
	[/\[(.*?)\]=\((.*?)\)/gm, "<a href\=\"$2\">$1</a>"],
	[/\b(?<!\\)_(.*?)_\((.*?)\)/gm, "<span class=\"popup\" explanation=\"$2\"><i>$1</i></span>" ],
	[/\b(?<!\\)_(.*?)_\b/gm, "<i>$1</i>"],
	[/\B(?<!\\)\~(.*?)\~\B/gm, "<strike>$1</strike>"],
	[/\B(?<!\\)\*(.*?)\*/gm, "<strong>$1</strong>"],
	[/^$/gm, "<br>"],
	[/\B(?<!\\)###### (.*?) ######\B/gm, "<h6>$1</h6>"],
	[/\B(?<!\\)##### (.*?) #####\B/gm, "<h5>$1</h5>"],
	[/\B(?<!\\)#### (.*?) ####\B/gm, "<h4>$1</h4>"],
	[/\B(?<!\\)### (.*?) ###\B/gm, "<h3>$1</h3>"],
	[/\B(?<!\\)## (.*?) ##\B/gm, "<h2>$1</h2>"],
	[/\B(?<!\\)# (.*?) #\B/gm, "<h1>$1</h1>"],
	[/^(\w*)?`(.*?)`/gm, "<code class=\"$1\">$2</code>"],
	[/(?<!\\){{(.*?)( .*?)?}}/gm, "<$1$2>"],
	[/\\\*/gm, "\*"],
	[/\\_/gm, "_"],
	[/\\#/gm, "#"],
	[/\\\[/gm, "["],
	[/\\\]/gm, "]"],
	[/\\\{/gm, "{"],
	[/\\\}/gm, "}"],
];

function TranslateMd(MdText) {
	MdText = "\n"+MdText
	Replaceble.forEach((rnr)=>{
		MdText=MdText.replaceAll(rnr[0], rnr[1])
	})
	MdText = "<link rel=\"stylesheet\" href=\"/files/css/page-style.css\">"+MdText;
	return MdText;
}

function followPage(anchor, e) {
	// "remove http[s]://.../"
	const path = skipSubstr(anchor.href, "/", 3)
	if (path.startsWith("articles")) {
		window.location.search = "file="+ArticleToMd(path)
	}

	e.preventDefault();
}

function MdReplacer(read, write) {
	return (e)=>{
		write.innerHTML=TranslateMd(read.value)

		//TODO maybe delegate !click for hole page and check e.target
		cqueryAll(write, "a").forEach(anchor=>{
			anchor.addEventListener('click', (event)=>{
				followPage(anchor, event)
			});
		})
	}
}

function GetFile(callback, folder="", path="", opt={}) {
	const FS = BuildFilsSystem(FS_MIRROR_PATH())
	return FS_redraw( null, FS_walk(FS, folder), path, callback, {
		createFile:true, createFolder:true, ...opt
	})
}

window.onload = async () => {
	await FS_REMIRROR()
	const FS = BuildFilsSystem(FS_MIRROR_PATH())

	const openFile = urlVars["file"]??"files/markdowns/MD/CheatSheet.md"
	const fileCont = LoadFile(openFile);

	const lang = removePrefix(openFile, "/").split("/", 3).at(-1);
	const inlang = lang != openFile && fIinA(lang, LANGS);

	const codeArea = id("code");
	const htmlArea = id("result")?.contentDocument.querySelector("html");
	codeArea.addEventListener("input", MdReplacer(codeArea, htmlArea));

	id("-file")
	id("open-file").addEventListener("click", ()=>{
		GetFile( ({filesEndpoint}) =>{
			window.location.search = "file="+filesEndpoint
		}, "markdowns", inlang?lang:"")
	})
	id("-file")
	id("-file")

	codeArea.value = await fileCont;
	MdReplacer(codeArea, htmlArea)()
}

