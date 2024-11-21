// main.js
const input1 = document.getElementById("numer1");
const input2 = document.getElementById("numer2");

const button = document.querySelector("button");

button.addEventListener("click", (e) => {
    e.preventDefault();
    // Convert strings to numbers using parseInt or + operator
    const value1 = input1.value;
    const value2 = input2.value;

    const res = value1 + value2;

    document.getElementById("result").innerHTML = res;
});