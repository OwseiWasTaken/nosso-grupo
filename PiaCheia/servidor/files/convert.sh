#! /bin/bash

set -e

# if argc==0
if [[ $# = '0' ]]
then
	# do pages/./file.md
	files=$(echo ./pages/**/*.md)
else
	# do args
	files=$@
fi


for md_file in $files
do
	# remove .md from filename
	page_path=$(echo $md_file | sed 's|\(.*\)\.md|\1|')
	article_path=$(echo $md_file | sed 's|\pages/\(.*\)\.md|articles/\1.html|')
	# remove everything until last /
	page_name=$(basename $page_path)
	echo -ne "<!DOCTYPE html>
<html>
<head>
	<title>$page_name</title>
	<link rel=\"stylesheet\" href=\"/files/css/page-style.css\">
</head>
<body>
" > $article_path

	cat $md_file \
	| sed 's|<|\&lt;|g;s|>|\&gt;|g' \
	| sed 's|^{$|<div>|' \
	| sed 's|^{#\(.*\)$|<div id="\1">|g;s|^}$|</div>|' \
	| sed 's|^\(\w\+\)``$|<code class="\1">|' \
	| sed 's|^``|</code>|g' \
	| sed 's|^\[\[\(.*\)\]\]|<a href="/article/\1.html">\1</a>|g' \
	| sed 's|^\[\(.*\?\)\](\(.*\?\))|<a href="\/article/\2.html">\1</a>|g' \
	| perl -pe 's|\[(.*?)\]=\((.*?)\)|<a href\="\2">_\1_</a>|g' \
	| perl -pe 's|\b(?<!\\)_(.*?)_\((.*?)\)|<span class="popup" explanation="\2"><i>\1</i></span>|g' \
	| perl -pe 's|\b(?<!\\)_(.*?)_\b|<i>\1</i>|g' \
	| perl -pe 's|(?<!\S)\*(.*?)\*|<strong>\1</strong>|g' \
	| sed 's|###### \(.*\?\) ######|<h6>\1</h6>|g' \
	| sed 's|##### \(.*\?\) #####|<h5>\1</h5>|g' \
	| sed 's|#### \(.*\?\) ####|<h4>\1</h4>|g' \
	| sed 's|### \(.*\?\) ###|<h3>\1</h3>|g' \
	| sed 's|## \(.*\?\) ##|<h2>\1</h2>|g' \
	| perl -pe 's|(?<!\\)# (.*?) #|<h1>\1</h1>|g' \
	| sed 's|\\\*|*|g' \
	| sed 's|\\_|_|g' \
	| sed 's|\\#|#|g' \
	| sed 's|#\/|#|g' \
	| sed 's|^\(\w*\)\?`\(.*\?\)`|<code class="\1">\2</code>|g' \
	| perl -pe 's|(?<!\\){{(.*?)( .*?)?}}|<\1\2>|g' \
	| sed 's|\\{{\(.*\?\)\( .*\?\)\?}}|{{\1\2}}|g' \
	| perl -pe 's|(?<!\\);$|<br>|' \
	| sed 's|\\;$|;|' \
	| sed 's|;;$|;<br>|' \
	| sed 's|^$|<br>|' \
	| sed 's|\
|<br>|' >> $article_path

	echo -ne "</body></html>" >> $article_path

	echo "converted $md_file -> $article_path"
done

