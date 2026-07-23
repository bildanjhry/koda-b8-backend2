export default async function register(datas) {
	try {
		const url = "http://localhost:8080/auth/register"
		const response = await fetch(url, {
			method: "POST",
			headers: {
				'Content-Type': 'application/x-www-form-urlencoded',
			},
			body: datas.toString()

		})

		if (!response.ok) {
			throw new Error(response.error)
		}
		const data = await response.json()
		return data

	} catch (err) {
		console.error(err.Message)
	}

}