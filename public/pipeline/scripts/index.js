document.addEventListener('DOMContentLoaded', () => {
  console.log("has loaded")
  console.log("stinger injecting outside assets")
  let data = {
    "css": {
      "html-logo": "./static/assets/css/html-logo.css",
    },
    "js": {
      "app": "/static/assets/js/application.js",
    },
    "html": {
      "test": {
        "content": "<div><p>Hello from stinger<p></div>",
      },
      "font": {
        "src": 'https://fonts.googleapis.com/css?family=Questrial&display=swap',
        "type": 'stylesheet',
        "element": 'link'
      }
    }
  };

  console.log("response: ", data);
  var tag = document.createElement('script');
  tag.src = data.js.app;
  var firstScriptTag = document.getElementsByTagName('script')[0]
  firstScriptTag.parentNode.insertBefore(tag, firstScriptTag)
  tag = document.createElement(data.html.font.element)
  tag.href = data.html.font.src
  tag.setAttribute("rel", data.html.font.type)
  let secondScriptTag = document.getElementsByTagName('script')[1]
  secondScriptTag.parentNode.insertBefore(tag, firstScriptTag)

})