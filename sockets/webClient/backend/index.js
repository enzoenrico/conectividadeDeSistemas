const apiUrl = 'https://fluffy-space-funicular-rp6vrvx4vrp264x-9999.app.github.dev/';
const outputDiv = document.getElementById('output');


function fetchData() {
  fetch(apiUrl, {method: "GET"})
    .then(response => {
      if (!response.ok) {
        throw new Error('Network response was not ok');
      }
      return response.json();
    })
    .then(data => {
      console.log(data);
      outputDiv.innerText = data
    })
    .catch(error => {
      console.error('Error:', error);
    });
    outputDiv.innerText = "ball"
}