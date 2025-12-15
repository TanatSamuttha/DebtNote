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
}