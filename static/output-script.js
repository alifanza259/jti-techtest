(function () {
    if (sessionStorage.getItem('googleUserCreds') == null) {
        //Redirect to login page, no user entity available in sessionStorage
        window.location.href = 'homepage';
    }
})();

const socket = new WebSocket("ws://localhost:3005/ws")
socket.onmessage = (async ({ data }) => {
    const dataObj = JSON.parse(data)
    if (dataObj.type === "input") {
        const audio = new Audio('https://bucket-alif.s3.ap-southeast-1.amazonaws.com/sfx.mp3')
        audio.play()

        const dataEl = document.getElementById("table-body")
        const { id, no_handphone: noHandphone, provider } = dataObj.handphone
        const isNewDataOdd = noHandphone % 2 === 1

        if (      
            isNewDataOdd
            ? window.oddIndex >= window.evenIndex
            : window.evenIndex >= window.oddIndex
            ) {
            const tbodyEl = document.getElementById("table-body")

            let addElement;
            if (isNewDataOdd) {
                addElement = `
                    <tr>
                        <td class="userdata" onmouseover="changeColor(this, true);"
                            onmouseout="changeColor(this, false);" onclick="saveState(this);">${`${id}-${noHandphone}-${provider}`}</td>
                        <td></td>
                    </tr>
                    `
                window.oddIndex += 1
            } else {
                addElement = `
                    <tr>
                        <td></td>
                        <td class="userdata" onmouseover="changeColor(this, true);"
                            onmouseout="changeColor(this, false);" onclick="saveState(this);">${`${id}-${noHandphone}-${provider}`}</td>
                    </tr>
                    `
                window.evenIndex += 1
            }
            tbodyEl.innerHTML = tbodyEl.innerHTML + addElement
        } else {
            let tdEl;
            if (isNewDataOdd) {
                tdEl = dataEl.rows[window.oddIndex].getElementsByTagName("td")[0]
                window.oddIndex += 1
            } else {
                tdEl = dataEl.rows[window.evenIndex].getElementsByTagName("td")[1]
                window.evenIndex += 1
            }

            tdEl.classList.add("userdata")
            tdEl.innerHTML = `${id}-${noHandphone}-${provider}`

            tdEl.addEventListener("mouseover", function () {
                changeColor(tdEl, true)
            })
            tdEl.addEventListener("mouseout", function () {
                changeColor(tdEl, false)
            })
            tdEl.addEventListener("click", function () {
                saveState(tdEl)
            })
        }
    }
})

let selectedTableRow = null
let selectedId = 0;

function changeColor(tableRow, highLight) {
    if (highLight) {
        tableRow.style.backgroundColor = '#dcfac9';
    }
    else if (selectedId !== tableRow.innerHTML.split("-")[0]) {
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

async function sendFetchRequest(apiUrl, method, headers, payload) {
    try {
        const response = await fetch(apiUrl, {
            method,
            headers,
            ...(!!payload ? { body: JSON.stringify(payload) } : {})
        })
        if (!response.ok) throw new Error(response.statusText)

        alert("Sukses")
    } catch (error) {
        console.error(error.message)
        alert("Maaf terjadi kesalahan")
    }
}

async function edit() {
    const newNoHandphone = document.getElementById("edit-field").value
    if (!newNoHandphone) {
        alert("Nomor baru tidak boleh kosong")
        return
    }
    const pattern = /^08[1-9][0-9]{6,9}$/
    if (!pattern.test(newNoHandphone)) {
        alert("Format no handphone harus menggunakan angka depan 0, tanpa spasi, dan memiliki panjang 10 - 13 karakter")
        return
    }

    const apiUrl = "http://localhost:3005/handphone"
    const payload = {
        id: selectedId,
        no_handphone: newNoHandphone
    }
    const headers = {
        "Content-type": "application/json"
    }

    await sendFetchRequest(apiUrl, "PATCH", headers, payload)
    location.reload()
}

async function deleteEntry() {
    const apiUrl = "http://localhost:3005/handphone/" + selectedId
    await sendFetchRequest(apiUrl, "DELETE")
    location.reload()
}


function logout() {
    sessionStorage.clear();
    window.location.href = 'homepage';
}