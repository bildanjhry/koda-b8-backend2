import getAllUsers from "../fetch/getAllUsers.js"
import deleteUserById from "../fetch/deleteUserById.js"
import editUserEmailById from "../fetch/editUserEmailById.js"

$(document).ready( async function(){
  const form = $(".edit-form-wrapper")

  function checkLogin(){
    if(!window.localStorage.getItem("token_user"))
    window.location.href = "/index.html"
  }
	checkLogin()
  
	const data = await getAllUsers()
	const tableBody = $(".table-user")
  
  async function handleDelete(id){
    const res = await deleteUserById(id)
    if(res.Success){
      window.location.reload()
    }
  }
  
  function handleEdit(id, email){
    form.addClass("display-flex")
    const formEdit = $(".form-edit")
    const emailForm = $(".form-edit > div > input")
    emailForm.val(email)

    formEdit.on("submit", async function(e) {
      e.preventDefault()
      const data = new FormData(e.target)
      const newData = new URLSearchParams(data)

      const res = await editUserEmailById(id, newData)
      if(res.Success){
        alert(res.Message)
       form.removeClass("display-flex")
       window.location.reload()
      }
    })
  }

  function handleDetail(id){
    window.location.href = `/detail.html?id=${id}`
  }



  data?.Results.forEach((item, index) => {
    const tr = $("<tr>").css("border-bottom", "1px solid gray")
    const number = $("<td>").css("padding-left", "15px")
    const picture = $("<td>")
    const pictureVal = $("<img>")
    .attr("src", `http://localhost:8080/${item.picture}`)
    .addClass("profile-pic")
    const email = $("<td>")
    const buttonWrapper = $("<td>")
    const actionButton = $("<button>").addClass("btn-action")
		const detailButton  = $("<button>").addClass("btn-detail")
    const actionButtonDel = $("<button>").addClass("btn-action-del")

    number.text(index + 1)
    picture.append(pictureVal)
    email.text(item.email)
    actionButton.text("Edit")
    actionButton.on("click", () => {
      handleEdit(item.id, item.email)
    })
    detailButton.text("Detail")
    detailButton.on("click", () => {
      handleDetail(item.id)
    })
		actionButtonDel.text("Delete")
    actionButtonDel.on("click", () => {
      handleDelete(item.id)}
    )

    buttonWrapper.append(actionButton)
    buttonWrapper.append(detailButton)
		buttonWrapper.append(actionButtonDel)

    tr.append(number)
    tr.append(picture)
    tr.append(email)
    tr.append(buttonWrapper)

    tableBody.append(tr)
  })

  const btnExit = $(".btn-exit")
  btnExit.on("click", function(){
    form.removeClass("display-flex")
  })

  const logBtn = $(".btn-logout")
  logBtn.on("click", function(){
		window.localStorage.removeItem("token_user")
    window.location.href = "/index.html"
  })
})