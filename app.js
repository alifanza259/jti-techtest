const socket = new WebSocket("ws://localhost:3005/ws")
socket.onmessage = (({ data }) => {
    const dataObj = JSON.parse(data)
    if (dataObj.type === "input") {
        const audio = new Audio('./sfx.mp3')
        audio.play()

        const dataEl = document.getElementById("data")
        const { id, no_handphone: noHandphone, provider } = dataObj.user


        if (noHandphone % 2 == 1) {
            if (window.oddIndex >= window.evenIndex) {
                const tbodyEl = document.getElementById("data")
                tbodyEl.innerHTML = tbodyEl.innerHTML + `
                <tr>
                    <td class="userdata" onmouseover="changeColor(this, true);"
                        onmouseout="changeColor(this, false);" onclick="saveState(this);">${`${id}-${noHandphone}-${provider}`}</td>
                    <td></td>
                </tr>
                `
                window.oddIndex += 1
            } else {
                const y = dataEl.rows[window.oddIndex]
                const z = y.getElementsByTagName("td")[0]

                z.classList.add("userdata")
                z.innerHTML = `${dataObj.user.id}-${dataObj.user.no_handphone}-${dataObj.user.provider}`

                z.addEventListener("mouseover", function () {
                    changeColor(z, true)
                })
                z.addEventListener("mouseout", function () {
                    changeColor(z, false)
                })
                z.addEventListener("click", function () {
                    saveState(z)
                })

                window.oddIndex += 1
            }
        } else {
            if (window.evenIndex >= window.oddIndex) {
                const tbodyEl = document.getElementById("data")
                tbodyEl.innerHTML = tbodyEl.innerHTML + `
                <tr>
                    <td></td>
                    <td class="userdata" onmouseover="changeColor(this, true);"
                        onmouseout="changeColor(this, false);" onclick="saveState(this);">${`${id}-${noHandphone}-${provider}`}</td>
                </tr>
                `
                window.evenIndex += 1
            } else {
                const y = dataEl.rows[window.evenIndex]
                const z = y.getElementsByTagName("td")[1]

                z.classList.add("userdata")
                z.innerHTML = `${dataObj.user.id}-${dataObj.user.no_handphone}-${dataObj.user.provider}`

                z.addEventListener("mouseover", function () {
                    changeColor(z, true)
                })
                z.addEventListener("mouseout", function () {
                    changeColor(z, false)
                })
                z.addEventListener("click", function () {
                    saveState(z)
                })

                window.evenIndex += 1
            }
        }

    }
    console.log("atas")
})

let selectedTableRow = null
let selectedId = 0;

function changeColor(tableRow, highLight) {
    if (highLight) {
        tableRow.style.backgroundColor = '#dcfac9';
    }
    else {
        if (selectedId !== tableRow.innerHTML.split("-")[0])
            tableRow.style.backgroundColor = 'white';
    }
}

function saveState(tableRow) {
    if (selectedId === tableRow.innerHTML.split("-")[0]) return

    if (selectedTableRow) {
        selectedTableRow.style.backgroundColor = 'white'
    }
    selectedId = tableRow.innerHTML.split("-")[0]
    selectedTableRow = tableRow

    document.getElementById("delete-btn").disabled = false
    document.getElementById("edit-btn").disabled = false
}

async function edit() {
    const apiUrl = "http://localhost:3005/user"
    try {
        const newNoHandphone = document.getElementById("edit-field").value
        if (!newNoHandphone) {
            alert("Nomor baru tidak boleh kosong")
            return
        }
        const response = await fetch(apiUrl, {
            method: "PATCH",
            headers: {
                "Content-type": "application/json"
            },
            body: JSON.stringify({
                id: selectedId,
                no_handphone: newNoHandphone
            })
        })
        if (!response.ok) throw new Error(response)
        alert("Sukses")
        location.reload()
    } catch (error) {
        console.error(error)
        alert("Maaf terjadi kesalahan")
    }
}

async function deleteEntry() {
    const apiUrl = "http://localhost:3005/user/" + selectedId
    try {
        const response = await fetch(apiUrl, {
            method: "DELETE",
        })
        if (!response.ok) throw new Error(response)
        alert("Sukses")
        location.reload()
    } catch (error) {
        console.error(error)
        alert("Maaf terjadi kesalahan")
    }
}
