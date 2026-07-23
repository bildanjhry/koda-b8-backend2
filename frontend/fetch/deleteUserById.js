export default async function deleteUserById(id) {
	try {
		const url = `http://localhost:8080/user/delete/${id}`
		const res = await fetch(url, {
			method: "DELETE",
			headers: {
				'Content-Type': 'application/json',
				'Authorization': 'hello',
			},
		})
		if (!res.ok) {
			throw new Error(res.error)
		}
		const data = await res.json()
		return data

	} catch (err) {
		console.error(err.Message)
	}
}