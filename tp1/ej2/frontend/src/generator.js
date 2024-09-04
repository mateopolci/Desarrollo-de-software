import './style.css';
import './app.css';

function generateUsername() {
    let username = Math.random().toString(36).slice(2);
    return username;
}
function generatePassword() {
    let password = Math.random().toString(36).slice(2);
    return password;
}

const generator = document.getElementById("generator");
const credentials = document.getElementById("credentials");
const cleaner = document.getElementById("cleaner");

generator.addEventListener("click", () => {
    credentials.innerHTML = `
    <div class="border">
        <p> Generated username: ${generateUsername()} </p>
        <p> Generated password : ${generatePassword()} </p>
    </div>
    `;
});

cleaner.addEventListener("click", () => {
    credentials.innerHTML = ""
})