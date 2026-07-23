export default async function deleteUserById(id) {
	try {
		const token = window.localStorage.getItem("token_user")
		const url = `http://localhost:8080/users/delete/${id}`
		const res = await fetch(url, {
			method: "DELETE",
			headers: {
				'Content-Type': 'application/json',
				'Authorization': `Bearer ${token}`,
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