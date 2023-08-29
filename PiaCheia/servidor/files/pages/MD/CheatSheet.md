# MarkDown #

## Tamanhos ##


md`\# texto \# ->` # texto #

md`#\# texto \#\# ->` ## texto ##

md`##\# texto \##\# ->` ### texto ###

md`###\# texto \###\# ->` #### texto ####

md`####\# texto \####\# ->` ##### texto #####

md`#####\# texto \#####\# ->` ###### texto ######


## Links ##


md`[[IO/sync]] ->`
[[IO/sync]]

md`[Esse link vai para a mesma página](IO/sync) ->`
[Esse link vai para a mesma página](IO/sync)


## Estrutura ##


md`{` isso cria uma div

md`}` isso fecha a div


md`{#essa-div` isso cria uma div com o id "essa-div"

md`}` isso ainda fecha a div


## Formatação ##


### Código ###

md``
 SQL``

 	SELECT * FROM accounts;;

 `` ->
``

SQL``
SELECT * FROM accounts;

``

md``

 SQL`INSERT INTO accounts (name) VALUES ("manse");`

`` ->
SQL`INSERT INTO accounts (name) VALUES ("manse");`


### formatação básica ###

md``
 \_Itálico_ _Itálico_

 \_Itálico_(com explicação) _Itálico_(com explicação)

 \*Negrito*  *Negrito*

``

## HTML ##


### tag manual ###


\{{img src="/files/ddg.png"}}

{{img src="/files/ddg.png"}}

### várias tags ###


\{{ol}}

\{{li}} Oi  \{{/li}}

\{{li}} Bom \{{/li}}

\{{li}} Dia \{{/li}}

\{{/ol}}


{{ol}}
{{li}} Oi  {{/li}}
{{li}} Bom {{/li}}
{{li}} Dia {{/li}}
{{/ol}}

\{{table}}

&nbsp;\{{thead}}

&nbsp;&nbsp;\{{tr}}

&nbsp;&nbsp;&nbsp;\{{th}} Nome  \{{/th}}

&nbsp;&nbsp;&nbsp;\{{th}} Idade \{{/th}}

&nbsp;&nbsp;\{{/tr}}

&nbsp;\{{/thead}}

&nbsp;\{{tbody}}

&nbsp;&nbsp;\{{tr}}

&nbsp;&nbsp;&nbsp;\{{td}} Manse \{{/td}}

&nbsp;&nbsp;&nbsp;\{{td}} 17 \{{/td}}

&nbsp;&nbsp;\{{/tr}}

&nbsp;&nbsp;\{{tr}}

&nbsp;&nbsp;&nbsp;\{{td}} Sof \{{/td}}

&nbsp;&nbsp;&nbsp;\{{td}} 19 \{{/td}}

&nbsp;&nbsp;\{{/tr}}

&nbsp;\{{/tbody}}

\{{/table}}


{{table}}
	{{thead}}
		{{tr}}
			{{th}} Nome  {{/th}}
			{{th}} Idade {{/th}}
		{{/tr}}
	{{/thead}}
	{{tbody}}
		{{tr}}
			{{td}} Manse {{/td}}
			{{td}} 17 {{/td}}
		{{/tr}}
		{{tr}}
			{{td}} Sof {{/td}}
			{{td}} 19 {{/td}}
		{{/tr}}
	{{/tbody}}
{{/table}}

