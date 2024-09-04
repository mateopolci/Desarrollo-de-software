const textInput = document.getElementById("textInput");
const acronym = document.getElementById("acronym");
const acronymbtn = document.getElementById("acronymbtn");
const acronymcleaner = document.getElementById("acronymcleaner");

acronymbtn.addEventListener("click", () => {
    
    const userInput = textInput.value.toUpperCase().trim();
    if (userInput === "") {
        acronym.innerHTML = `
            <div class="border">
                <p> Please enter a phrase </p>
            </div>
            `;
        return;
    } else if (userInput.split(" ").length === 1) {
        acronym.innerHTML = `
            <div class="border">
                <p> Please enter more than a word </p>
            </div>
            `;
        return;
    }
    const words = userInput.split(" ");
    const firstLetters = words.map((el) => el.charAt(0));
    acronym.innerHTML = `
        <div class="border">
            <p> Generated acronym: ${firstLetters.join(".")} </p>
        </div>
        `;
});

acronymcleaner.addEventListener("click", () => {
    acronym.innerHTML = "";
    textInput.value="";
})