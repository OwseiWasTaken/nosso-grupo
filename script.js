"use strict";

const Replaceble = [
    ["<", "&lt;"],
    [">", "&gt;"],
    [/\n{\n/, "\n<div>\n" ],
	[/\n{#(.*)\n/, "\n<div id=\"$1\">\n"],
    [/\n}\n/, "\n</div>\n"],
	[/\n(\w+)``\n/, "\n<code class=\"$1\">\n"],
	[/\n``\n/, "\n</code>\n"],
	[/\s\[\[(.*)\]\]\s/, "<a href=\"/article/$1.html\">$1</a>" ],
	[/\s\[(.*?)\]\((.*?)\)\s/, "<a href=\"\/article/$2.html\">$1</a>" ],
	[/\[(.*?)\]=\((.*?)\)/, "<a href\=\"$2\">_$1_</a>"],
	[/\b(?<!\\)_(.*?)_\((.*?)\)/, "<span class=\"popup\" explanation=\"$2\"><i>$1</i></span>" ],
	[/\b(?<!\\)_(.*?)_\b/, "<i>$1</i>"],
	/*
	[/(?<!\S)\*(.*?)\*|<strong>\1</strong>|g' ],
	[/###### (.*?) ######|<h6>\1</h6>|g' ],
	[/##### (.*?) #####|<h5>\1</h5>|g' ],
	[/#### (.*?) ####|<h4>\1</h4>|g' ],
	[/### (.*?) ###|<h3>\1</h3>|g' ],
	[/## (.*?) ##|<h2>\1</h2>|g' ],
	[/(?<!\\)# (.*?) #|<h1>\1</h1>|g' ],
	[/\\\*|*|g' ],
	[/\\_|_|g' ],
	[/\\#|#|g' ],
	[/#\/|#|g' ],
	[/^\(\w*\)\?`\(.*\?\)`|<code class="\1">\2</code>|g' ],
	[/(?<!\\){{(.*?)( .*?)?}}|<\1\2>|g' ],
	[/\\{{\(.*\?\)\( .*\?\)\?}}|{{\1\2}}|g' ],
	[/(?<!\\);$|<br>|' ],
	[/\\;$|;|' ],
	[/;;$|;<br>|' ],
	[/^$|<br>|' ],
    */
];

function TranslateMd(MdText) {
    MdText="\n"+MdText+"\n"
    Replaceble.forEach((rnr)=>{
        MdText=MdText.replace(rnr[0], rnr[1])
    })
    return MdText;
}

window.onload = () => {
    const codeArea = document.getElementById("code");
    const htmlArea = document.getElementById("result")?.contentDocument.querySelector("html");

    codeArea.addEventListener("keyup", (e)=>{
        htmlArea.innerHTML=TranslateMd(codeArea.value)
    })
}
