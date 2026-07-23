export default async function uploadUserPicById(id, datas) {
	try {
		const token = window.localStorage.getItem("token_user") || ''
		const url = `http://localhost:8080/user/upload-pic/${id}`
		const response = await fetch(url, {
			method: "PATCH",
			headers: {
				'Authorization': `Bearer ${token}`,
			},
			body: datas
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