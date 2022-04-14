var btn = document.getElementById('btn');
var counter = Number(btn.innerHTML);
btn.onclick = function() {
    counter++;
    btn.innerHTML = counter;
};
