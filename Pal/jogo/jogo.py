#! /usr/local/bin/python3.11
import lib
from lib.util import *
from lib.keys import KeyDict as keys
import subprocess

if OS == "linux":
	def GetCh():
		return subprocess.run(
			("./lib/gtk", "--once", "--python"), capture_output=True
		).stdout[8:-3]
else:
	def GetCh():
		return subprocess.run(
			("./lib/gtk.exe", "--once", "--python"), capture_output=True
		).stdout[8:-3]

def GetKey():
	x = GetCh()
	#print(len(x))
	for k in keys.keys():
		#print(len(k))
		if len(k) != len(x):continue
		for i in r(k):
			if ord(k[i]) != x[i]:break
			#print(f"{ord(k[i])} == {x[i]}: {ord(k[i]) == x[i]}")
		else:
			#print(f"you pressed {keys[k]}")
			return keys[k]
	else:
		return "NULL"
print(GetKey())
exit(0)

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
	# open file, join lines with ; and split lines with ;; (\n\n)
	with open("./vids.txt", 'r') as f:
		fvids = ';'.join(list(map(
			lambda x:(x.strip()),
			f.readlines()
			))).split(';;')
	# make tuple of each video (based on ;)
	fvids = [tuple(x.split(';')) for x in fvids]
	cvids = []
	# set choice -> name for each vid line (except self's name)
	for vid in fvids:
		d = {}
		if len(vid) != 1:
			for n in vid[1:]:
				k, v = n.split("->")
				d[k] = v
		cvids.append(Video(vid[0], d))

	return cvids

def LinkVids(vids: list[Video]) -> list[Video]:
	# make lookup table obj.name->obj
	d = {x.vidname:x for x in vids}
	# for each vif
	for vid in vids:
		# set value (name) of choice[k] as obj (based on name->obj)
		vid.choices = {k:d[v] for k, v in vid.choices.items()}
	return vids

def main() -> str:
	vids = LinkVids(MakeVids())
	for vid in vids:
		print(repr(vid))
	return ""

#while True:
if True:
	if x:=main():
		dprint(stderr, "ERROR", x+'\n')
