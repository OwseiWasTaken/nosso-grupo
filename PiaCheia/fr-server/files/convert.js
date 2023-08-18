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
	[/\B\[\[(.*)\]\]\B/gm, "<a href=\"/article/$1.html\">$1</a>" ],
	[/\B\[(.*?)\]\((.*?)\)\B/gm, "<a href=\"\/article/$2.html\">$1</a>" ],
	[/\[(.*?)\]=\((.*?)\)/gm, "<a href\=\"$2\">$1</a>"],
	[/\b(?<!\\)_(.*?)_\((.*?)\)/gm, "<span class=\"popup\" explanation=\"$2\"><i>$1</i></span>" ],
	[/\b(?<!\\)_(.*?)_\b/gm, "<i>$1</i>"],
	[/\B(?<!\\)\~(.*?)\~\B/gm, "<strike>$1</strike>"],
	[/(?<!\S)\*(.*?)\*/gm, "<strong>$1</strong>"],
	[/^$/gm, "<br>"],
	[/\B(?<!\\)###### (.*?) ######\B/gm, "<h6>$1</h6>"],
	[/\B(?<!\\)##### (.*?) #####\B/gm, "<h5>$1</h5>"],
	[/\B(?<!\\)#### (.*?) ####\B/gm, "<h4>$1</h4>"],
	[/\B(?<!\\)### (.*?) ###\B/gm, "<h3>$1</h3>"],
	[/\B(?<!\\)## (.*?) ##\B/gm, "<h2>$1</h2>"],
	[/\B(?<!\\)# (.*?) #\B/gm, "<h1>$1</h1>"],
	[/\\\*/gm, "*"],
	[/\\_/gm, "_"],
	[/\\#/gm, "#"],
	[/\\\[/gm, "["],
	[/\\\]/gm, "]"],
	[/^(\w*)?`(.*?)`/gm, "<code class=\"$1\">$2</code>"],
	[/(?<!\\){{(.*?)( .*?)?}}/gm, "<$1$2>"],
];

function TranslateMd(MdText) {
	MdText = "\n"+MdText
	Replaceble.forEach((rnr)=>{
		MdText=MdText.replaceAll(rnr[0], rnr[1])
	})
	MdText = "<link rel=\"stylesheet\" href=\"/files/page-style.css\">"+MdText;
	return MdText;
}

window.onload = () => {
	const codeArea = document.getElementById("code");
	const htmlArea = document.getElementById("result")?.contentDocument.querySelector("html");

	codeArea.addEventListener("keyup", (e)=>{
		htmlArea.innerHTML=TranslateMd(codeArea.value)
	})
}

