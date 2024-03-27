var dataView = document.getElementById("datasView")
var view = document.querySelector(".view")
var rooms = dataView.getElementsByClassName("room")
var relations = dataView.getElementsByClassName("relation")
var farm = document.getElementById("farm")
var lis = document.getElementsByTagName("li")

var start = document.querySelector(".EndPointStart").textContent
var end = document.querySelector(".EndPointEnd").textContent

var comb = document.getElementsByClassName("pathsComb")
var antsPaths = document.getElementsByClassName("antsbyPaths")

function setFarmVizualiser() {
    farm.replaceChildren()
    for (let index = 0; index < rooms.length; index++) {
        const room = rooms[index]
    
        const x = room.querySelector(".x")
        const y = room.querySelector(".y")
        const name = room.querySelector(".name")
    
        const newRoom = document.createElement("div")
        newRoom.style.left = x.textContent+"%"
        newRoom.style.top = y.textContent+"%"
        newRoom.id = name.textContent
        
        if (name.textContent == start) {
            newRoom.classList.add("start")
                initStartAnts(x.textContent, y.textContent)
        }
        else if (name.textContent == end) {
            newRoom.classList.add("end")
        } else {
            newRoom.style.backgroundColor = "rgb(98, 0, 255)"
        }
        newRoom.classList.add("room")
        newRoom.addEventListener("mouseover", associedNameOn)
        newRoom.addEventListener("mouseout", associedNameOff)
        farm.appendChild(newRoom)
    }
    
    for (let index = 0; index < relations.length; index++) {
        const relation = relations[index]
        const firstRoom = relation.querySelector(".firstRoom")
        const secondRoom = relation.querySelector(".secondRoom")
    
        var roomOne = document.getElementById(firstRoom.textContent)
        var roomTwo = document.getElementById(secondRoom.textContent)
    
        var x1 = roomOne.offsetLeft
        var y1 = roomOne.offsetTop
        var x2 = roomTwo.offsetLeft
        var y2 = roomTwo.offsetTop
        traceLine(parseFloat(x1), parseFloat(x2), parseFloat(y1), parseFloat(y2))
    }
}

function traceLine(x1, x2, y1, y2) {
    distance = Math.sqrt(Math.pow(x1-x2, 2) + Math.pow(y1-y2, 2))
    xMid = (x1+x2) / 2
    yMid = (y1+y2) / 2

    slopeInRadian = Math.atan2(y1-y2, x1-x2)
    slopeInDegrees = slopeInRadian * 180 / Math.PI

    const newLink = document.createElement("div")
    newLink.classList.add("relation")
    newLink.style.width = distance+"px"
    newLink.style.top = yMid+"px"
    newLink.style.left = (xMid - distance/2) +"px"
    newLink.style.transform = "rotate(" + slopeInDegrees + "deg"
    farm.appendChild(newLink)
}

window.onload = setFarmVizualiser
window.onresize = setFarmVizualiser

function roomLightOn(ev) {
    let room = document.getElementById(ev.textContent)
    room.style.backgroundColor = "orange"
}

function roomLightOff(ev) {
    let room = document.getElementById(ev.textContent)
    if (ev.textContent == start) {
        room.style.backgroundColor = "#08f000"
    } else if (ev.textContent == end) {
        room.style.backgroundColor = "rgb(255, 0, 0)"
    } else {
        room.style.backgroundColor = "rgb(98, 0, 255)"
    }
}

function associedNameOn() {
    for (let index = 0; index < lis.length; index++) {
        const li = lis[index];
        if (li.textContent == this.id) {
            li.style.filter = "brightness(75%)"
        }
    }
}

function associedNameOff() {
    for (let index = 0; index < lis.length; index++) {
        const li = lis[index];
        if (li.textContent == this.id) {
            li.style.filter = "brightness(100%)"
        }
    }
}


function initStartAnts(x, y) {
    var ants = document.getElementsByClassName("antsForPath")
    var numberAnts = 0
    
    for (index = 0; index < ants.length; index++) {
        const ant = ants[index]
        numberAnts += parseInt(ant.textContent)
    }
    
    
    for (index = 0; index < numberAnts; index++) {
        var img = document.createElement("img")
        img.src = "/vizualizer/static/img/ethant.webp"
        img.alt = "ant" + (index + 1)
        img.className = "ant"
        img.style.left = x + "%"
        img.style.top = y + "%"
        farm.appendChild(img)
    }
}

function movingAntsOnPaths() {
    var ants = document.getElementsByClassName("ant")
    var paths = document.getElementsByClassName("path")
    var pathsForAnt = document.getElementsByClassName("antsForPath")

    var comb = []
    for (let index = 0; index < paths.length; index++) {
        const path = paths[index];
        var pathOfComb = []
        let rooms = path.getElementsByClassName("roomOfPath")
        for (let ind = 0; ind < rooms.length; ind++) {
            const room = rooms[ind].textContent;
            pathOfComb.push(room)
        }
        comb.push(pathOfComb)
    }

    var antsByPath = []
    for (let index = 0; index < pathsForAnt.length; index++) {
        const pathForAnt = pathsForAnt[index];
        antsByPath.push(parseInt(pathForAnt.textContent))
    }
    
    var posAnt = 0
    function newSalve() {
        var allAntsLaunched = true
        for (let index = 0; index < antsByPath.length; index++) {
            const antsOfPath = antsByPath[index];
            if (antsOfPath > 0) {
                antsByPath[index]--
                allAntsLaunched = false
                movingAnt(ants[posAnt], comb[index])
                posAnt++
            }
        }
        if (allAntsLaunched == false) {
            setTimeout(newSalve, 1100)
        }
    }
    newSalve()
}

async function movingAnt(ant, path) {
    var index = 0
    function moveToNextRoom() {
        if (index < path.length) {
            const room = path[index]
            const nextRoom = document.getElementById(room)
            ant.style.left = nextRoom.style.left
            ant.style.top = nextRoom.style.top
            index++
            setTimeout(moveToNextRoom, 1100)
        }
    }
    moveToNextRoom()
}
