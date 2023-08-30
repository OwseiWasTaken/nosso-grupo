
window.onload = async () => {
	await FS_REMIRROR()
	const FS = BuildFilsSystem(FS_MIRROR_PATH())
	for (const lang of LANGS) {
		FS_walk(FS, "/pages/"+lang)
	}
	log(FS.fileMap)
}

