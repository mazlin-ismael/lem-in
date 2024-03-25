var dataView = document.getElementById("datasView")
var view = document.querySelector(".view")
var rooms = dataView.getElementsByClassName("room")
var relations = dataView.getElementsByClassName("relation")
var farm = document.getElementById("farm")


function setFarmVizualiser() {
    farm.replaceChildren()
    for (let index = 0; index < rooms.length; index++) {
        const room = rooms[index]
    
        const x = room.querySelector(".x")
        const y = room.querySelector(".y")
        const name = room.querySelector(".name")
    
        room.style.left = x.textContent+"%"
        room.style.top = y.textContent+"%"
    
        const newRoom = document.createElement("div")
        newRoom.style.left = x.textContent+"%"
        newRoom.style.top = y.textContent+"%"
        newRoom.id = name.textContent
        newRoom.classList.add("room")    
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
        console.log(x1, y1, x2, y2)
        traceLine(parseFloat(x1), parseFloat(x2), parseFloat(y1), parseFloat(y2))
    }
}

function traceLine(x1, x2, y1, y2) {
    console.log(y2)
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