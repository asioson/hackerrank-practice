var res = document.getElementById('res');
var btn0 = document.getElementById('btn0');
var btn1 = document.getElementById('btn1');
var btnClr = document.getElementById('btnClr');
var btnEql = document.getElementById('btnEql');
var btnSum = document.getElementById('btnSum');
var btnSub = document.getElementById('btnSub');
var btnMul = document.getElementById('btnMul');
var btnDiv = document.getElementById('btnDiv');

var arg1 = '';
var arg2 = '';
var oper;

btn0.onclick = function() {
    res.innerHTML += '0';
    if (oper == undefined)
        arg1 += '0'
    else
        arg2 += '0'
};

btn1.onclick = function() {
    res.innerHTML += '1';
    if (oper == undefined)
        arg1 += '1'
    else
        arg2 += '1'
};

btnClr.onclick = function() {
    res.innerHTML = '';
    oper = undefined;
    arg1 = '';
    arg2 = '';
};

btnEql.onclick = function() {
    if (oper == undefined) 
        return;
    var answer;
    let a = parseInt(arg1,2);
    let b = parseInt(arg2,2);
    switch (oper) {
        case '+':
            answer = ((a+b) >>> 0).toString(2);
            break;
        case '-':
            answer = ((a-b) >>> 0).toString(2);
            break;
        case '*':
            answer = ((a*b) >>> 0).toString(2);
            break;
        case '/':
            answer = ((a/b) >>> 0).toString(2);
            break;
    }
    res.innerHTML = answer;
    arg1 = answer;
    arg2 = '';
    oper = undefined;
};

btnSum.onclick = function() {
    if (oper == undefined) {
        arg1 = res.innerHTML;
        arg2 = '';
        res.innerHTML += '+'
        oper = '+';
    }
};

btnSub.onclick = function() {
    if (oper == undefined) {
        arg1 = res.innerHTML;
        arg2 = '';
        res.innerHTML += '-'
        oper = '-';
    }
};

btnMul.onclick = function() {
    if (oper == undefined) {
        arg1 = res.innerHTML;
        arg2 = '';
        res.innerHTML += '*'
        oper = '*';
    }
};

btnDiv.onclick = function() {
    if (oper == undefined) {
        arg1 = res.innerHTML;
        arg2 = '';
        res.innerHTML += '/'
        oper = '/';
    }
};
