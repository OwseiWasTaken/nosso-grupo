function die(txt){
	alert("ERROR PARSING HTML");
	_._();
}

class Parser{
	constructor(string) {
		this.buffer = string
			.replaceAll("\n", " ")
			.replaceAll("\t", " ")
		while (this.buffer.indexOf("  ")!=-1) {
			this.buffer = this.buffer.replaceAll("  ", " ");
		}
		this.buffer = this.buffer.split(">").map(x=>x+">")
		this.buffer.length--; // remove "last" tag

		this.tags = [];
	}

	get next() {
		return this.buffer[++this.index];
	}

	peep(jump=1) {
		return this.buffer[this.index+jump]
	}

	parse() {
		this.buffer.forEach(this.tag.bind(this))
	}

	tag(text, index) {
		if (this.peep()[0] != '<'){
			hasContent = true;
		}
		console.log(text)
	}
}

window.onload = () => {
	button = document.getElementById("read")
	textarea = document.getElementById("text")
	button.onclick = ()=>{
		x = new Parser(textarea.value).parse()
	}
}
