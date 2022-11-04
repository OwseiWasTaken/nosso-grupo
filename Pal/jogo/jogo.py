#! /usr/local/bin/python3.11
from util import *

@dataclass
class vid:
	def __init__(this, vidname:str, choices:dict[str, str]):
		this.vidname = vidname
		this.playname = vidname+".mp4"
		this.choices = choices
	def __call__(this, choice:str) -> str:
		return "./vids/"+this.choice[choice]

def main() -> str:
	vids = []
	with open("./vids.txt", 'r') as f:
		fvids = ';'.join(list(map(
			lambda x:(x.strip()),
			f.readlines()
			))).split(';;')
	fvids = [tuple(x.split(';')) for x in fvids]
	print(fvids)
	return ""

#while True:
if True:
	if x:=main():
		dprint(stderr, "ERROR", x+'\n')
