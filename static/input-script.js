(function () {
    if (sessionStorage.getItem('googleUserCreds') == null) {
        //Redirect to login page, no user entity available in sessionStorage
        window.location.href = 'homepage';
    }
})();

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

async function autoGenerate() {
    const apiUrl = "http://localhost:3005/handphone/auto"
    await sendFetchRequest(apiUrl, "POST")
}

function encryptAes(str) {
    const key = CryptoJS.enc.Utf8.parse("key-jti-techtest");
    const iv = CryptoJS.enc.Utf8.parse("1234567890123456");

    return CryptoJS.AES.encrypt(str, key, { iv, padding: CryptoJS.pad.Pkcs7 }).toString();
}

async function submit() {
    const noHandphone = document.getElementById("noHandphone");
    const provider = document.getElementById("provider")

    if (!noHandphone.value || !provider.value) {
        alert("Harap isi data yang kosong")
        return
    }

    const pattern = /^08[1-9][0-9]{6,9}$/
    if (!pattern.test(noHandphone.value)) {
        alert("Format no handphone harus menggunakan angka depan 0, tanpa spasi, dan memiliki panjang 10 - 13 karakter")
        return
    }

    const encryptedHandphone = encryptAes(noHandphone.value)
    const encryptedProvider = encryptAes(provider.value)

    const payload = {
        no_handphone: encryptedHandphone, provider: encryptedProvider
    }
    const apiUrl = "http://localhost:3005/handphone"
    const headers = {
        "Content-Type": "application/json"
    }
    await sendFetchRequest(apiUrl, "POST", headers, payload)

    noHandphone.value = ""
    provider.value = ""
}

function logout() {
    sessionStorage.clear();
    window.location.href = 'homepage';
}
