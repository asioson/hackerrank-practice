'use strict';

process.stdin.resume();
process.stdin.setEncoding('utf-8');

let inputString = '';
let currentLine = 0;

process.stdin.on('data', inputStdin => {
    inputString += inputStdin;
});

process.stdin.on('end', _ => {
    inputString = inputString.trim().split('\n').map(string => {
        return string.trim();
    });
    
    main();    
});

function readLine() {
    return inputString[currentLine++];
}

/*
 * Complete the vowelsAndConsonants function.
 * Print your output using 'console.log()'.
 */
function vowelsAndConsonants(s) {
    var sv = "", sc = "";
    for (var i = 0; i < s.length; i++) {
        switch (s[i]) {
            case 'a': case 'A': 
            case 'e': case 'E': 
            case 'i': case 'I':
            case 'o': case 'O':
            case 'u': case 'U':
                sv += s[i] + "\n";
                break;
            default:
                sc += s[i] + "\n";
        }
    }
    console.log(sv + sc);
}


function main() {
    const s = readLine();
    
    vowelsAndConsonants(s);
}
