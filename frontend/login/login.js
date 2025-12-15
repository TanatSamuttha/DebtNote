document.getElementById("login-form").addEventListener("submit", login)

function login(event){
    event.preventDefault()
    const username = document.forms["login-form"]["username"].value;
    const password = document.forms["login-form"]["password"].value;
    axios({
        method: 'post',
        url: 'http://127.0.0.1:8080/login',
        data: {
            username: username,
            password: password
        }
    })
    .then(res => {
        console.log(res.data)
        if(res.data != "OK"){
            const informationBox = document.querySelector(".information-block");
            if (!informationBox) return;
            informationBox.querySelectorAll("p").forEach(p => p.remove());
            const p = document.createElement("p");
            p.textContent = res.data;
            informationBox.appendChild(p);
        }
    })
}