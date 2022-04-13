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

/**
*   Return the second largest number in the array.
*   @param {Number[]} nums - An array of numbers.
*   @return {Number} The second largest number in the array.
**/
function getSecondLargest(nums) {
    const n = nums.length;
    var largest, secondLargest;
    if (n > 0) {
        largest = nums[0];
        for (var i = 1; i < n; i++) {
            if (largest < nums[i])
                largest = nums[i];
        }
        for (var i = 0; i < n; i++) {
            if (nums[i] < largest) {
                if (secondLargest == undefined) 
                    secondLargest = nums[i];
                else if (nums[i] > secondLargest)
                    secondLargest = nums[i];
            }
        }
    }
    return secondLargest;
}


function main() {
    const n = +(readLine());
    const nums = readLine().split(' ').map(Number);
    
    console.log(getSecondLargest(nums));
}
