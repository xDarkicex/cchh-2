console.log("Forgin script loaded...");

const Stinger = {}

window.Stinger = Stinger

Stinger.Get = function Get() {
  let xhr = new XMLHttpRequest()

  xhr.onreadystatechange = function () {
    if (this.readyState == 4 && this.status == 200) {
      Stinger.Remote_Payload = JSON.parse(this.responseText);
    }
  };
  xhr.open('GET', 'http://localhost:3001/home-health/reviews/json', true);
  xhr.setRequestHeader('Content-Type', 'application/json');

  xhr.send();
};