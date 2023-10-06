
// ARL_FOLDER or ARL_FILE should be placed under an ARL_ROOT
// or ARL_FOLDER that is under ARL_ROOT
ARL_ROOT = (content=[]) => hCreateElement("ol",
	[content].flat(), {class:"tree"}
)

ARL_FOLDER = (name, depth, files=[]) => hCreateElement("li", {
	class:"folder",
}, [
	createElement("label", {
		for:`ARL_${name}-${depth}`,
	}, name),
	createElement("input", {
		type:"checkbox", id:`ARL_${name}-${depth}`,
	}),
	hCreateElement("ol", files),
])

ARL_FILE = (fileName, filesEndpoint) => hCreateElement("li",
	{ class:"file" },
	[ createElement("a", fileName, {href:filesEndpoint}) ],
)

function ARL_MAKE_CHILD (fileOrDir, depth) {
	return (fileOrDir.isDir?
		ARL_MAKE_FOLDER(fileOrDir.fullname, fileOrDir, depth):
		ARL_FILE(fileOrDir.name, fileOrDir.filesEndpoint)
	)
}

function ARL_MAKE_FOLDER (folderName, fsDir, depth=0) {
	return ARL_FOLDER(folderName, depth,
		fsDir.files.map(
			a=>ARL_MAKE_CHILD(fsDir.fileMap[a], depth+1)
		)
	);
}

ARL_MAKE_HEADLESS_FOLDER = (folderName, fsDir, depth=0) => {
	const ofolder = hCreateElement("div", [
		createElement("h2", folderName),
		hCreateElement("ol", {class:"tree"}, [
			...fsDir.files.map(
				a=>ARL_MAKE_CHILD(fsDir.fileMap[a], depth+1)
			)
		])
	])
	return ofolder;
	const folder =  ARL_FOLDER(folderName, depth,
		fsDir.files.map(
			a=>ARL_MAKE_CHILD(fsDir.fileMap[a], depth+1)
		)
	);

	debugger;
	folder.classList.remove("folder");
	cquery(folder, "input").remove();
	cquery(folder, "label").remove();

	insertTop(folder, createElement("h2", folderName))
	return ARL_ROOT(folder);
}

function makeLang(FS, langPath, langName) {
	const langDir = FS_walk(FS, langPath)
	for (const articleName in langDir.fileMape) {
		log(langDir.fileMap[articleName])
	}
}

//TODO include aside in <article>
function AddArticleListing(container) {
	const FS = BuildFilsSystem(FS_MIRROR_PATH())
	container.appendChild(
		ARL_MAKE_HEADLESS_FOLDER(
			"articles", FS_walk(FS, "articles")
		)
	);
}
