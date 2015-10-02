// Sample script LED flashing script
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
var i = setInterval(toggle1, 100);
var i2 = setInterval(toggle2, 40);
console.log("Disco..");

