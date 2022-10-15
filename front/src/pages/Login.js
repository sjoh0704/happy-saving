import React, { useState, useEffect } from 'react';
import axios from 'axios';
import { useNavigate } from 'react-router';
function Login() {
    const [mail, setMail] = useState('')
    const [password, setPassword] = useState('')
    const navigate = useNavigate()
 
    const onChangeMailHandler = (e) => {
        setMail(e.target.value)
    }
 
    const onChangePasswordHandler = (e) => {
        setPassword(e.target.value)
    }
 
    const login = async() =>{
        let body = {
            mail: mail,
            password: password,
        }

        try {
            let res = await axios.post("/auth", body)
            alert(res.data.message)

            let payload = res.data.payload
            localStorage.setItem("mail", payload.mail)
            localStorage.setItem("userid", payload.id)
            localStorage.setItem("name", payload.name)
            localStorage.setItem("gender", payload.gender)
            navigate("/")
            return
        } catch (err) {
            alert(err.response.data.message)
            console.log(err.response.data.message)
        }
        
    }

    const onClickLoginHandler = () => {
        console.log('try to login')
        login()
    }

    return(
        <div>
            <h2>Login</h2>
            <div>
                <label htmlFor='input_id'>EMAIL : </label>
                <input type='text' name='input_id' value={mail} onChange={onChangeMailHandler} />
            </div>
            <div>
                <label htmlFor='input_pw'>PW : </label>
                <input type='password' name='input_pw' value={password} onChange={onChangePasswordHandler} />
            </div>
            <div>
                <button type='button' onClick={onClickLoginHandler}>Login</button>
            </div>
        </div>
    )
}
 
export default Login;