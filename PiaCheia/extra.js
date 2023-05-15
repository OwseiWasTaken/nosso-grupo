
var LastEnum = 0;
function ContinuosEnum(inverted, ...args) {
	let S = {};
	for (arg of args) {
		if (inverted) inverted.push(arg);
		S[arg] = LastEnum++;
	}
	Object.freeze(inverted);
	return Object.freeze(S);
}

function Enum(inverted, ...args) {
	let S = {};
	for (arg of args) {
		if(inverted) inverted.push(arg);
		S[arg] = Symbol(arg);
	}
	Object.freeze(inverted);
	return Object.freeze(S);
}

