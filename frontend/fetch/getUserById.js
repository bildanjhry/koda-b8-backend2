export default async function getUserById(id) {
	try {
		const token = window.localStorage.getItem("token_user") || ''
		const url = `http://localhost:8080/users/detail/${id}`
		const res = await fetch(url, {
			headers: {
				'Authorization': `Bearer ${token}`,
			}
		})
		if(!res.ok){
			throw new Error(res.error)
		}
		const data = await res.json()
		console.log(data)
		return data

	} catch (err) {
		console.error(err.message)
	}
}