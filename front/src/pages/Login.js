import React, { useState, useEffect } from 'react';
import axios from 'axios';
 
function Login() {
    const [mail, setMail] = useState('')
    const [password, setPassword] = useState('')
 
    const onChangeMailHandler = (e) => {
        setMail(e.target.value)
    }
 
    const onChangePasswordHandler = (e) => {
        setPassword(e.target.value)
    }
 
    const login = () =>{
        let body = {
            mail: mail,
            password: password,
        }

        axios
        .post("/auth", body)
        .then( res => {
            alert(res)
        })
        .catch(e => {
            alert(e)
        })
    
    }

    const onClickLoginHandler = () => {
        console.log('try to login')
        login()
    }


    // useEffect(() => {
    //     login()
    // },
    // // 페이지 호출 후 처음 한번만 호출될 수 있도록 [] 추가
    // [])
 
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