# lem-in | imazlin | bdesouza

### VIDEO PROJET PRESENTATION
👉 [Lem-in Demo](https://streamable.com/pt8kti)

## REQUIRED

- Before starting this project you will need of go and an IDE
### Installation Golang & IDE Advice
- Linux Terminal
```
sudo apt-get update
sudo apt-get install golang

sudo apt-get update
sudo apt-get install code
```
- Windows Mac  
Go -> [GOLANG](https://go.dev/doc/install)  
Go -> [VSCODE](https://code.visualstudio.com/download)

### Launching Program
After launching the project with vscode you can open
a new terminal in the top of page  
1. Write in the terminal 
```
go mod init lem-in
```
2. On the same terminal execute the command
```
go run . FILES/<fileName.txt>
```
the visualizers 2d and 3d are in theses addresses
```
http://localhost:2030/2d
http://localhost:2030/3d
```


### PROGRAM

This project is meant to make you code a digital version of an ant farm.

Create a program lem-in that will read from a file (describing the ants and the colony) given in the arguments.

Upon successfully finding the quickest path, lem-in will display the content of the file passed as argument and each move the ants make from room to room.

How does it work?

You make an ant farm with tunnels and rooms.
You place the ants on one side and look at how they find the exit.

You need to find the quickest way to get n ants across a colony (composed of rooms and tunnels).

At the beginning of the game, all the ants are in the room ##start. The goal is to bring them to the room ##end with as few moves as possible.
The shortest path is not necessarily the simplest.
Some colonies will have many rooms and many links, but no path between ##start and ##end.
Some will have rooms that link to themselves, sending your path-search spinning in circles. Some will have too many/too few ants, no ##start or ##end, duplicated rooms, links to unknown rooms, rooms with invalid coordinates and a variety of other invalid or poorly-formatted input. In those cases the program will return an error message ERROR: invalid data format. If you wish, you can elaborate a more specific error message (example: ERROR: invalid data format, invalid number of Ants or ERROR: invalid data format, no start room found).

You must display your results on the standard output in the following format :

number_of_ants
the_rooms
the_links

Lx-y Lz-w Lr-o ...

x, z, r represents the ants numbers (going from 1 to number_of_ants) and y, w, o represents the rooms names.

A room is defined by "name coord_x coord_y", and will usually look like "Room 1 2", "nameoftheroom 1 6", "4 6 7".

The links are defined by "name1-name2" and will usually look like "1-2", "2-5".

Here is an example of this in practice :
```
##start
1 23 3
2 16 7
#comment
3 16 3
4 16 5
5 9 3
6 1 5
7 4 8
##end
0 9 5
0-4
0-6
1-3
4-3
5-2
3-5
#another comment
4-2
2-1
7-6
7-2
7-4
6-5
```
Which corresponds to the following representation :
```
        _________________
       /                 \
  ____[5]----[3]--[1]     |
 /            |    /      |
[6]---[0]----[4]  /       |
 \   ________/|  /        |
  \ /        [2]/________/
  [7]_________/
```
### Instructions

You need to create tunnels and rooms.
A room will never start with the letter L or with # and must have no spaces.
You join the rooms together with as many tunnels as you need.
A tunnel joins only two rooms together never more than that.
A room can be linked to multiple rooms.
Two rooms can't have more than one tunnel connecting them.
Each room can only contain one ant at a time (except at ##start and ##end which can contain as many ants as necessary).
Each tunnel can only be used once per turn.
To be the first to arrive, ants will need to take the shortest path or paths. They will also need to avoid traffic jams as well as walking all over their fellow ants.
You will only display the ants that moved at each turn, and you can move each ant only once and through a tunnel (the room at the receiving end must be empty).
The rooms names will not necessarily be numbers, and in order.
Any unknown command will be ignored.
The program must handle errors carefully. In no way can it quit in an unexpected manner.
The coordinates of the rooms will always be int.
Your project must be written in Go.


### ALGORITHM
- after have check is the file was in the good format and respect the sensitives cases, we linked the rooms to the others rooms by the links  
which be a slice of room pointers, we find all paths possible, a path have a start and a end and if he have room between them, the rooms cannot be in doublon  
when we have all possible paths, we find the combinations of possible paths, a combination of paths cannot have twice the same room except start and end,  
the numbers of path of the combination is the smallest number between the rooms linked to the start, the rooms linked to the end and the numbers of ants  
we calculate also the combinations which be lower in numbers of path and we see what is the shortest combination,   
when we have the shortest combination, we send the ants   


## FILEST
-To launch the filetest In terminal root of project in directory algo or errFile
```
go test
```

- The filetest do the units tests of functions which return errors and check every possibles cases

## DOCKER

- To launch the docker In terminal root of project
```
cd docker
sh lem-in.sh
```

The docker is launched in port 2030