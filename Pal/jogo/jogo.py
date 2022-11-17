#! /usr/local/bin/python3.11
import lib
from lib.util import *
from lib.util import GetCh as gtk
from lib.keys import KeyDict as keys
import subprocess

def move(y, x):
	print("\033[%d;%dH" % (y, x), end="")


if OS == "linux":
	def clear():
		ss("clear")
	def GetCh():
		return subprocess.run(
			("./lib/gtk", "--once", "--python"), capture_output=True
		).stdout[12:-3]
else:
	def clear():
		ss("cls")
	def GetCh():
		return ''.join(list(map(chr, gtk())))

def GetKey():
	x = GetCh()
<<<<<<< HEAD
	return keys.get(x, x)
=======
	#print(len(x))
>>>>>>> master
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
		return x

gamin = not get("--jogo").exists

@dataclass
class Video:
	def __init__(this, vidname:str, choices:dict[str, Any]): # any -> str | Video
		this.vidname = vidname
		this.playname = "./vids/"+vidname
		if exists(x:=this.playname+".mp4"):
			this.playname = x
		elif exists(x:=this.playname+".txt"):
			this.playname = x
		else:
			if gamin:
				print(f"arquivo {this.playname}.mp4 (ou .txt) nao existe")
				exit(2)
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
	return vids, d, d["intro"]

def Show(ops, off):
	stdout.write(f"\x1B[{1+off};1H")
	stdout.write("Você vai:\n")
	for i in r(ops):
		stdout.write("( )"+ops[i]+'\n')

def CMD(y, playname):
	stdout.write(pos(y)+f"[CMD]: vlc {playname}")

<<<<<<< HEAD
def CMD(playname):
	if gamin:
		if playname.endswith(".mp4"):
			ss("vlc {playname}")
			return 0
		else:
			with open(playname, 'r') as f:
				pf = f.readlines()
			stdout.write("\x1B[1;1H")
			stdout.write('+'+'-'*(mx-2)+'+') # add 1 to y

			stdout.write("\x1B[1;4H") # inline, +3
			# com espaco
			stdout.write(' '+playname[7:-4]+' ') # 7: to cut ./vids/
			# sem espaco
			#stdout.write(playname[7:-4]) # 7: to cut ./vids/

			for i in r(pf):
				stdout.write(f"\x1B[{2+i};1H|")
				stdout.write(pf[i])
				stdout.write(f"\x1B[{2+i};{mx}H|")
			stdout.write('+'+'-'*(mx-2)+'+') # add 1 to y


			stdout.flush()
			return len(pf)+2
	else:
		stdout.write(pos(my-2)+f"[not gamin]: vlc {playname}")
		return 0

def statusline(atual):
	stdout.write(pos(my-1)+f"selected {atual}")
=======
>>>>>>> master

def main() -> str:
	vids, dic, atual = LinkVids(MakeVids())
	stdout.flush()
	y = 0
	mx, my = GetTerminalSize()
	ops = list(atual.choices.keys())
	#TESTZONE
	#TESTZONE

	clear()
<<<<<<< HEAD
	c = CMD(atual.playname)
	Show(ops, c)
	statusline(atual)
	while len(ops)-1:
=======
	Show(ops)
	stdout.write(pos(my-1)+f"selected {atual}")
	CMD(my-2, atual.playname)
	while len(ops):
>>>>>>> master
		# mover cursor
		Show(ops, c)
		stdout.write("\x1B[%i;2H@" % (y+2+c))
		move(my, 0)

		stdout.flush()
		k = GetKey()
		if k == "up":
			if y != 0:
				y -=1
		elif k == "down":
			if y != len(ops)-1:
				y+=1
		elif k in ("space", "enter"):
<<<<<<< HEAD
			# reset
			atual = atual.choices[ops[y]]
			statusline(atual)
			clear()
			c = CMD(atual.playname)
=======
			clear()
			stdout.write(pos(my-1)+f"selected {atual.choices[ops[y]]}")
			# reset
			atual = atual.choices[ops[y]]
			CMD(my-2, atual.playname)
>>>>>>> master
			ops = list(atual.choices.keys())
			y = 0
			Show(ops, c)
		else:
			if not gamin:
				stdout.write(pos(my-3)+f"{k} key")
	clear()
	c = 0
	if atual.playname.endswith(".txt"):
		c = CMD(atual.playname)
	stdout.write(f"\x1B[{c+1};{1}H")
	stdout.write("fim.\n")
	return ""

#while True:
if True:
	if x:=main():
		dprint(stderr, "ERROR", x+'\n')
