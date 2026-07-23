export default async function editUserEmailById(id, datas) {
	try {
		const token = window.localStorage.getItem("token_user")
		const url = `http://localhost:8080/users/edit/${id}`
		const response = await fetch(url, {
			method: "PATCH",
			headers: {
				'Content-Type': 'application/x-www-form-urlencoded',
				'Authorization': `Bearer ${token}`,
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