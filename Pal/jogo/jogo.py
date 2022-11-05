#! /usr/local/bin/python3.11
from util import *

@dataclass
class Video:
	def __init__(this, vidname:str, choices:dict[str, Any]): # any -> str | Video
		this.vidname = vidname
		this.playname = vidname+".mp4"
		this.choices = choices
	def __call__(this, choice:str) -> str:
		return "./vids/"+this.choice[choice]
	def __str__(this):
		return "@"+this.vidname
	def __repr__(this):
		return f"{this.vidname}:{[k+'->'+str(v) for k, v in this.choices.items()]}"

def MakeVids() -> list[Video]:
	vids = []
	with open("./vids.txt", 'r') as f:
		fvids = ';'.join(list(map(
			lambda x:(x.strip()),
			f.readlines()
			))).split(';;')
	fvids = [tuple(x.split(';')) for x in fvids]
	cvids = []
	for vid in fvids:
		d = {}
		if len(vid) != 1:
			for n in vid[1:]:
				k, v = n.split("->")
				d[k] = v
		cvids.append(Video(vid[0], d))

	return cvids

def LinkVids(vids: list[Video]) -> list[Video]:
	d = {x.vidname:x for x in vids}
	for vid in vids:
		vid.choices = {k:d[v] for k, v in vid.choices.items()}
	return vids

def main() -> str:
	vids = MakeVids()
	LinkVids(vids)
	for vid in vids:
		print(repr(vid))
	return ""

#while True:
if True:
	if x:=main():
		dprint(stderr, "ERROR", x+'\n')
