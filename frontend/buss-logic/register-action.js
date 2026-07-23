import register from "../fetch/register.js"

$(document).ready( async function(){
  const formRegister = $(".form-register")
    
  formRegister.on('submit', async function (e){
    e.preventDefault()
    const data = new FormData(e.target)
    const newData = new URLSearchParams(data)
    const res = await register(newData)
    alert(res.Message)
	})

})


