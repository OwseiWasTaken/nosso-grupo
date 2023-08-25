"use strict";

class Dir {
	constructor(filename) {
		this.ShrinkName = ShrinkPath(filename)
		this.isDir = true;
		this.fsEndpoint = "/fs/"+this.ShrinkName
		this.filesEndpoint = "/files/"+this.ShrinkName
		this.filename = filename
		this.fullname = BaseName(filename)
		this.name = filename.split(".").slice(0, -1).join(".")
		this.path = PathName(filename)
		this.files = [];
		this.fileMap = {};
		this.size = 4096;
		this.modTime = new Date();
	}
	push(file) {
		this.files.push(file.fullname)
		this.fileMap[file.fullname] = file
	}
	get(filename) {
		return this.fileMap[filename]
	}
	forEach(fnc) {
		this.files.forEach(fl=>{
			fnc(this.fileMap[fl])
		})
	}
	filter(fnc) {
		return this.files.filter(fl=>{
			return fnc(this.fileMap[fl])
		}).map((name)=>this.fileMap[name])
	}
}

class File {
	constructor(filename) {
		this.isDir = false;
		this.fsEndpoint = "/fs/"+filename
		this.filesEndpoint = "/files/"+filename
		this.filename = filename
		this.fullname = BaseName(filename)
		this.name = BaseName(filename).split(".").slice(0, -1).join(".")
		this.extension = BaseName(filename).split(".").at(-1)
		this.path = PathName(filename)
		this.size = this.fullname.length;
		this.modTime = new Date();
	}
}

function _FS_walk(tree, path) {
	if (path.length === 0) return tree;
	//tree = tree[path[0]];
	tree = tree.fileMap[path[0]];
	return _FS_walk(tree, path.slice(1));
}

function IsDir(path) {
	return path.substr(path.length - 1) === '/';
}

function StripLastSlash(path) {
	if (IsDir(path)) {
		path = path.substr(0,path.length-1);
	}
	return path;
}

// ./pages/./here/../there -> pages/there
// "./" -> ""
function ShrinkPath(path) {
	return path
		.replaceCertain(/\/\w*\/\.\.\//g, "/")
		.replaceCertain(/\/\.\//g, "/")
		.replace(/^\.?\//, "")
		.replace(/\/$/, "")
}

function FS_aproach(tree, path) {
	path = ShrinkPath(path)
	path = path.split("/");
	path.length--;
	return _FS_walk(tree, path);
}

function FS_walk(tree, path) {
	path = ShrinkPath(path)
	if (!path) return tree;
	path = path.split("/");
	return _FS_walk(tree, path);
}

function PathName(path) {
	path = StripLastSlash(path)
	path = path.split("/");
	path.length--;
	return path;
}

function BaseName(path) {
	path = StripLastSlash(path)
	path = path.split("/");
	return path.at(-1);
}

function BuildFilsSystem(PathRes) {
	const ROOT = new Dir("/");
	PathRes.forEach((path)=>{
		if (IsDir(path)) {
			FS_aproach(ROOT, path).push(new Dir(path));
		} else {
			FS_aproach(ROOT, path).push(new File(path));
		}
	})
	return ROOT;
}

//TODO get this list from /fs/
const PATHRES = [
	"pages/",
		"pages/IO/",
			"pages/IO/fopen.html",
			"pages/IO/sync.html",
			"pages/IO/fclose.html",
			"pages/IO/md.html",
		"pages/OS/",
			"pages/OS/alert.html",
			"pages/OS/assert.html",
			"pages/OS/nice.html",
]

function SelectFile(callback) {
	const FS = BuildFilsSystem(PATHRES)
	let CWD = "";
	FS_redraw( null, FS, CWD, callback, {
		createFile:true,
	})
}

window.onload = () => {
	SelectFile(log);
}

function GoToFile(CWD="") {
	let callback = (file)=>{
		window.location = file.filesEndpoint;
	}
	const FS = BuildFilsSystem(PATHRES)
	FS_redraw(null, FS, CWD, callback)
}

const FS_default_draw_options = {
	//TODO: implement selecting folder & disabling file selection
	selectFile:true,
	selectFolder:false,
	createFile:false,
	createFolder:false,
}

function FS_redraw(
	oldDialog, FS, CWD,
	callback, options=FS_default_draw_options
) {

	// define unset option 'a as FS_default_draw_options['a]
	options = {...FS_default_draw_options, ...options};
	let Dialog = null;
	if (oldDialog === null) {
		Dialog = hCreateElement("dialog", {
			class:"FS-file-select-dialog"
		}, [
			createElement("h2", "At /"+CWD, {
				style:{
					position:"relative",
					display:"inline-block",
					bottom:"20px",
					margin:0,
					right:"10px",
				}
			}),
			hCreateElement("table", [
				hCreateElement("thead", [
					hCreateElement("tr", [
						createElement("th", "Nome"),
						createElement("th", "Endpoint"),
						createElement("th", "Tamanho"),
						createElement("th", "Última Modificação"),
					]),
				]),
				createElement("tbody"),
			]),
			hCreateElement("div", {id:"FS-bottom"}, [
				createElement("input", {id:"FS-search"}),
				createElement("button", {id:"FS-select"}, "Selecionar"),
				createElement("button", {id:"FS-create-file"}, "Novo Arquivo"),
				createElement("button", {id:"FS-create-folder"}, "Nova Pasta"),
			]),
		])
		document.body.appendChild(Dialog)
		Dialog.showModal()
	} else {
		Dialog = oldDialog;
		cquery(oldDialog, "tbody").innerHTML = "";
	}

	const tbody = cquery(Dialog, "tbody")
	const makeFile = cquery(Dialog, "#FS-create-file")
	const makeFolder = cquery(Dialog, "#FS-create-folder")

	if (!options.createFile) makeFile.disabled = true;
	if (!options.createFolder) makeFolder.disabled = true;

	makeFile.addEventListener("click", (event)=>{
		FS_walk(FS, CWD).push(new File("new-file.txt"))
	})

	makeFolder.addEventListener("click", (event)=>{
		FS_walk(FS, CWD).push(new Dir("new-folder/"))
	});

	//TODO only do folders.push(new Dir("..")) when ShrinkPath is implemented
	if (CWD) {
		const upDir = new Dir("../")
		debugger;
		//upDir.filesEndpoint = "/files/"+PathName(CWD).join("/")
		//upDir.filename = PathName(CWD).join("/")
		tbody.appendChild( FS_TROW(upDir, ()=>{
			FS_redraw(Dialog, FS, upDir.filename, callback, options)
		}))
	}

	FS_walk(FS, CWD).filter(({isDir})=>isDir).forEach(folder=>{
		tbody.appendChild( FS_TROW(folder, ()=>{
			FS_redraw(Dialog, FS, folder.filename, callback, options)
		}))
	})

	FS_walk(FS, CWD).filter(({isDir})=>!isDir).forEach(file=>{
		tbody.appendChild( FS_TROW(file, (e)=>{callback(file, e)}))
	})

}

function FS_TROW(fileOrDir, Click) {
	const name = createElement("button", fileOrDir.fullname, {
		class:"FS-btn-like-link",
		style:{
			width:"120%",
			position:"relative",
			right:"10px",
		}
	})
	name.classList.add(fileOrDir.isDir?"FS-dir":"FS-file")
	name.addEventListener("click", Click)

	return hCreateElement("tr", [
		name,
		createElement("td", fileOrDir.filesEndpoint),
		createElement("td", fileOrDir.size),
		createElement("td", UnixFormatDate(fileOrDir.modTime)),
	])
}

function MakeFile(fsEndpoint, name, content="") {
	log(`Make File ${name} @ ${fsEndpoint} "${content}" `)
}

function MakeFolder(fsEndpoint, name) {
	log(`Make Folder ${name} @ ${fsEndpoint}`)
}

