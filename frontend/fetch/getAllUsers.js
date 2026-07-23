export default async function getAllUsers(params) {
  try {
    const token = window.localStorage.getItem("token_user") || ''
    const url = `http://localhost:8080/users/all?${params}`
    const response = await fetch(url, {
      headers: {
        'Content-Type': 'application/json',
        'Authorization': `Bearer ${token}`,
      },
    })
    if(!response.ok) {
      throw new Error(response.error)
    }
		const data = await response.json()
		return data
    
	} catch(err) {
			console.error(err.Message)
	}
}