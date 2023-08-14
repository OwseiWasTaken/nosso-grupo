#! /bin/bash

set -e

echo -ne '<DOCTYPE html>
<html>
<head>
	<link rel="stylesheet" href="../style.css">
</head>
<body>
'

cat $1 \
| perl -pe 's|\b_(.*?)_\b|<i>\1</i>|g' \
\
| sed 's|###### \(.*\?\) ######|<h6>\1</h6>|g' \
| sed 's|##### \(.*\?\) #####|<h5>\1</h5>|g' \
| sed 's|#### \(.*\?\) ####|<h4>\1</h4>|g' \
| sed 's|### \(.*\?\) ###|<h3>\1</h3>|g' \
| sed 's|## \(.*\?\) ##|<h2>\1</h2>|g' \
| sed 's|# \(.*\?\) #|<h1>\1</h1>|g' \
\
| sed 's|\(\w*\)\?`\(.*\?\)`|<code class="\1">\2</code>|g' \
\
| sed 's|;;|<br>|' \
| sed 's|;$|<br>|' \
| sed 's|^$|<br>|' \
| sed 's|\
|<br>|'

echo -ne "</body></html>"
