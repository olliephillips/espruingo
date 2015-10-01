// Sample script - 2
// Based on http://www.espruino.com/Quick+Start

var on = false;
function toggle1() {
 on = !on;
 digitalWrite(LED1, on);
 
}
function toggle2() {
 on = !on;
 digitalWrite(LED2, on);
}
var i = setInterval(toggle1, 500);
var i2 = setInterval(toggle2, 200);
console.log("Disco..");