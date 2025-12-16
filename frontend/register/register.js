const { default: axios } = require("axios");

document.getElementById("register-form").addEventListener("submit", register)

function register(event){
    event.preventDefault()
    const username = document.forms["register-form"]["username"].value;
    const password = document.forms["register-form"]["password"].value;
    axios({
        method: 'post',
        url: 'http://127.0.0.1:8080/register',
        data: {
            username: username,
            password: password
        }
    })
    .then(res => {
        console.log(res.data)
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
        })
    })
}