// Module loading example
var s = require("servo").connect(C7);
var Clock = require("clock").Clock;

var clk=new Clock(2015,9,9,12,00,0,0);   // Initialise with specific date
var servoPos = [0,0];

function getPositions() {
	var t = getTime();
	servoPos = [ 0.5+Math.sin(t)*0.5, 0.5+Math.cos(t)*0.5 ];
	console.log(servoPos);
	var d1=clk.getDate(); 
   	console.log("Reading taken at: " + d1.toString());
}
function moveServos() {
	digitalPulse(A1,1,1+E.clip(servoPos[0],0,1));
	digitalPulse(A2,1,1+E.clip(servoPos[1],0,1));
}

setInterval("getPositions();moveServos()", 50);