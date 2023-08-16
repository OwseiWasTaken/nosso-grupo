#! /bin/bash

set -e

for md_file in $@
do
	# remove .md from filename
	page_path=$(echo $md_file | sed 's|\(.*\)\.md|\1|')
	# remove everything until last /
	page_name=$(basename $page_path)
	html_file="$page_path.html"

	echo -ne "<!DOCTYPE html>

<html>
<head>
	<title>$page_name</title>
	<link rel=\"stylesheet\" href=\"/files/page-style.css\">
</head>
<body>
" > $html_file

	cat $md_file \
	| sed 's|<|\&lt;|g;s|>|\&gt;|g' \
	| sed 's|^{$|<div>|' \
	| sed 's|^{#\(.*\)$|<div id="\1">|g;s|^}$|</div>|' \
	| sed 's|^\(\w\+\)``$|<code class="\1">|' \
	| sed 's|^``|</code>|g' \
	| sed 's|^\[\[\(.*\)\]\]|<a href="/article/\1.html">\1</a>|g' \
	| sed 's|^\[\(.*\?\)\](\(.*\?\))|<a href="\/article/\2.html">\1</a>|g' \
	| perl -pe 's|\[(.*?)\]=\((.*?)\)|<a href\="\2">_\1_</a>|g' \
	| perl -pe 's|\b_(.*?)_\((.*?)\)|<span class="popup" explanation="\2"><i>\1</i></span>|g' \
	| perl -pe 's|\b_(.*?)_\b|<i>\1</i>|g' \
	| sed 's|\(\s\)\*\(.*\?\)\*\(\s\)|\1<strong>\2</strong>\3|g' \
	| sed 's|^###### \(.*\?\) ######|<h6>\1</h6>|g' \
	| sed 's|^##### \(.*\?\) #####|<h5>\1</h5>|g' \
	| sed 's|^#### \(.*\?\) ####|<h4>\1</h4>|g' \
	| sed 's|^### \(.*\?\) ###|<h3>\1</h3>|g' \
	| sed 's|^## \(.*\?\) ##|<h2>\1</h2>|g' \
	| sed 's|^# \(.*\?\) #|<h1>\1</h1>|g' \
	| sed 's|^\(\w*\)\?`\(.*\?\)`|<code class="\1">\2</code>|g' \
	| sed 's|^{{\(.*\)\( .*\)\?}}|<\1\2>|g' \
	| sed 's|;;|<br>|;s|;$|<br>|;s|^$|<br>|' \
	| sed 's|\
|<br>|' >> $html_file

	echo -ne "</body></html>" >> $html_file

	echo "converted $md_file -> $html_file"
done

