import login from "../fetch/login.js"

$(document).ready(function(){
  const form = $(".form-login")

  form.on('submit', function(e){
    e.preventDefault()
    const data = new FormData(e.target)
    const newData = new URLSearchParams(data)
    login(newData)
  })

})


