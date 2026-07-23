export default async function editUserEmailById(id, datas) {

	console.log(datas.toString())
	try {
		const url = `http://localhost:8080/user/edit/${id}`
		const response = await fetch(url, {
			method: "PATCH",
			headers: {
				'Content-Type': 'application/x-www-form-urlencoded',
				'Authorization': 'hello',
			},
			body: datas.toString()

		})

		if (!response.ok) {
			throw new Error(response.error)
		}
		const data = await response.json()
		console.log(data)
		return data

	} catch (err) {
		console.error(err.Message)
	}

}