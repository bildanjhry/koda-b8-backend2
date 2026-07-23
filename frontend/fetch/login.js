export default async function login(datas) {
	try {
		const url = "http://localhost:8080/auth/login"
		const response = await fetch(url, {
			method: "POST",
			headers: {
				'Content-Type': 'application/x-www-form-urlencoded',
			},
			body: datas.toString()

		})

		const data = await response.json()
		if(data.Message){
			alert(data.Message)
		}
		if (data.Success) {
			window.localStorage.setItem("token_user", data?.Results?.Token)
			window.location.href = "/frontend/dashboard.html"
		} 

	} catch (err) {
		console.error(err.Message)
	}
}