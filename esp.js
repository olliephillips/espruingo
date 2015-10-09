// Module loading example
var s = require("servo").connect(C7);
var Clock = require("clock").Clock;

var clk=new Clock(2015,9,9,12,00,0,0);   // Initialise with specific date
var servoPos = [0,0];

function getPositions() {
	var t = getTime();
	var pos1 = 0.5+Math.sin(t)*0.5;
	var pos2 = 0.5+Math.cos(t)*0.5;
	servoPos = [ pos1, pos2 ];
	console.log(servoPos);
	var d1=clk.getDate(); 
   	console.log("Reading taken at: " + d1.toString());
	if(pos1 < 0.1) {
		digitalWrite(LED1, 1);
	} else {
		digitalWrite(LED1, 0);
	}
	if(pos2 < 0.1) {
		digitalWrite(LED2, 1);
	} else {
		digitalWrite(LED2, 0);
	}
}
function moveServos() {
	digitalPulse(A1,1,1+E.clip(servoPos[0],0,1));
	digitalPulse(A2,1,1+E.clip(servoPos[1],0,1));
}

setInterval("getPositions();moveServos();", 5000);
save();