//main.js
function changeBackground() {
    var rnum = Math.floor(Math.random()*256);
    var gnum = Math.floor(Math.random()*256);
    var bnum = Math.floor(Math.random()*256);
    var bgColor = "rgb(" + rnum + "," + gnum + "," +  bnum + ")";
    console.log(bgColor);
    document.body.style.backgroundColor  = bgColor;
}
