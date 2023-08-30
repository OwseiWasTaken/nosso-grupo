const immutable = Object.freeze;
const HTTPOK=(s)=>~~(s/100)==2&&s!=218;
const {log, assert} = console;
const id = document.getElementById.bind(document);
const query = document.querySelector.bind(document);
const queryAll = document.querySelectorAll.bind(document);
const cquery = (doc, query)=>doc.querySelector(query);
const cqueryAll = (doc, query)=>doc.querySelectorAll(query);
const insertBefore = (el, newnode) => el.parentElement.insertBefore(newnode, el);
const insertAfter = (el, newnode) => el.parentElement.insertBefore(newnode, el.nextSibling);
const remove = (arr, item) => delete arr[arr.indexOf(item)];
const urlVars = immutable(parseUrlVars(window.location.search));
const LANGS = immutable([
	"C",
	"PHP",
	"JS",
	"CSS",
	"MD",
	"SQL",
])

function hCreateElement(name, elements=[], attributes=null) {
	if (!Array.isArray(elements)) {
		[elements, attributes] = [attributes, elements];
		if (elements === null) {elements = []};
	}
	const el = document.createElement(name);
	for (const attr in attributes) {
		if (attr === "style") {
			for (const stl in attributes[attr]) {
				el.style[stl] = attributes[attr][stl];
			}
			continue;
		}
		el.setAttribute(attr, attributes[attr]);
	}
	el.append(...elements)
	return el;
}

function createElement(name, value="", attributes=null) {
	if (typeof (value) == "object") {
		[value, attributes] = [attributes, value];
		if (value === null) {value = ""};
	}
	const el = document.createElement(name);
	for (const attr in attributes) {
		if (attr === "style") {
			for (const stl in attributes[attr]) {
				el.style[stl] = attributes[attr][stl];
			}
			continue;
		}
		el.setAttribute(attr, attributes[attr]);
	}
	el.innerHTML = value;
	return el;
}

function zip(keys,values,out={}){
	keys.forEach( (val,idx)=>out[val] = values[idx] );
	return out;
}

//const monthNames = ["January", "February", "March", "April", "May", "June",
//  "July", "August", "September", "October", "November", "December"
//];

const monthNames = ["Janeiro", "Fevereiro", "Março", "Abril", "Maio", "Junho",
	"Julho", "Agosto", "Setembro", "Outubro", "Novembro", "Dezembro"
];

function UnixFormatDate(date) {
	const Mon = monthNames[date.getMonth()].substr(0,3);
	const day = date.getDate();//.toString().padStart(2, 0);
	const hour = date.getHours();
	const min = date.getMinutes();
	return `${day} ${Mon} ${hour}:${min}`
}

//TODO implement text query dialog

String.prototype.replaceCertain = function (searchPattern, replacePattern="") {
	let txt = this.replaceAll(searchPattern, replacePattern);
	if (txt.search(searchPattern) != -1) {
		txt = txt.replaceCertain(searchPattern, replacePattern)
	}
	return txt
}

function RemoveListeners(oldEl) {
	const newEl = oldEl.cloneNode();
	newEl.innerHTML = oldEl.innerHTML;
	oldEl.parentElement.replaceChild(newEl, oldEl);
	return newEl;
}

function skipSubstr(str, substr, howmany = 1, eatlast = true) {
  for (let i = 0; i < howmany; i++) {
    if (i == howmany - 1 && !eatlast) {
      str = str.substr(str.indexOf(substr))
    } else {
      str = str.substr(str.indexOf(substr) + substr.length)
    }
  }
  return str;
}

function MdToArticle(MdLink) {
	log(MdLink)
	return MdLink.replace(
		/files\/pages\/(.*).md/,
		"articles/$1.html",
	);
}

function ArticleToMd(ArticleLink) {
	log(ArticleLink)
	return ArticleLink.replace(
		/articles\/(.*)\.html/,
		"files/pages/$1.md",
	)
}

function parseUrlVars(text) {
	let vars = {};
	if (text[0] == '?') {text = text.substr(1)}

	text.split("&").forEach(def => {
		const [key, value] = def.split("=");
		vars[key] = value;
	})
	return vars;
}

function IinA(item, arr) {
  return arr.flat().indexOf(item) != -1;
}

function IinAv(item, ...arr) {
  return arr.flat().indexOf(item) != -1;
}

function fIinA(item, arr) {
  return arr.indexOf(item) != -1;
}

function fIinAv(item, ...arr) {
  return arr.indexOf(item) != -1;
}

function setPrefix(str, prefix) {
  if (!str.startsWith(prefix)) {
    str = prefix + str;
  }
  return str;
}

function setSuffix(str, suffix) {
  if (!str.endsWith(suffix)) {
    str += suffix;
  }
  return str;
}

function removePrefix(str, prefix) {
  if (str.startsWith(prefix)) {
    str = str.substr(prefix.length)
  }
  return str;
}

function removeSuffix(str, suffix) {
  if (str.endsWith(suffix)) {
    str = str.substr(0, str.length - suffix.length)
  }
  return str;
}
