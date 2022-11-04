#! /usr/local/bin/python3.11
#TODO: atualizar shebang p/ usar protable do pendrive
from util import *

@dataclass
class vid:
	def __init__(this, vidname:str, choices:list[str]):
		this.vidname = vidname
		this.playname = vidname
		this.choices = choices
	def __call__(this, choice:str) -> str:
		return "./vids/"+this.choice[choice]

def main() -> int:
	vids = []
	with open("./vids.txt", 'r') as f:
		fvids = ';'.join(list(map(lambda x:(x.strip()), f.readlines()))).split(';;')
	fvids = [tuple(x.split(';')) for x in fvids]
	print(fvids)
	return 0

main()
