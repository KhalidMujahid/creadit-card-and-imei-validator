const message = document.querySelector(".message");
const input = document.querySelector("input");
const button = document.querySelector("button");

button.addEventListener("click", (e) => {
  e.preventDefault();

  if (input.value === "" || !input.value) return;

  // check if value is a number
  if (!parseInt(input.value)) return alert("Please enter a valid number");

  sendInput(input.value);
});

async function sendInput(value) {
  await fetch("/verify", {
    method: "POST",
    body: JSON.stringify({
      value,
    }),
    headers: {
      "content-type": "application/json",
    },
  })
    .then((res) => res.json())
    .then((data) => {
      message.innerHTML = data.value;
    })
    .catch((error) => console.log(error));
}
