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

