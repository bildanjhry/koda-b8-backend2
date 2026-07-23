import getUserById from "../fetch/getUserById.js"
import uploadUserPicById from "../fetch/uploadUserPicById.js"

$(document).ready( async function() {
	const uploadBtn = $(".upload-input")
	const queryParams = new URLSearchParams(window.location.search)
	const id = queryParams.get("id")
	const response = await getUserById(id)	
	const data = response.Results

	$(".user-email").text(data.email)
	$(".my-picture").attr("src", `http://localhost:8080/${data.picture}`)

	uploadBtn.on('change', async function(e){
		const picture = e.target.files[0]
		const data = new FormData()
		data.append("picture", picture)

		const response = await uploadUserPicById(id, data)
		if(response.Message){
			alert(response.Message)
			window.location.reload()
		}
	})
})