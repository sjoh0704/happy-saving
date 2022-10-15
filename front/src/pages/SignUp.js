import React, { useState, useEffect } from 'react';
import axios from 'axios';
import { useNavigate } from 'react-router';
 
function SignUp() {
    const navigate = useNavigate()
    const [userData, setUserData] = useState({
        name: "",
        mail: "",
        password: "",
        gender: ""
    })

    const {name, mail, password, gender} = userData
    
    const onChangeUserHandler = (e) => {
        const {name, value} = e.target
        setUserData(
            {
                ...userData,
                [name]: value
            }
        )
    }
 
    const signUp = () =>{
        axios
        .post("/apis/v1/users", userData)
        .then( res => {
            alert("회원가입이 정상적으로 이루어졌습니다.")
            navigate('/login')
        })
        .catch(e => {
            console.log(e.response.data.message)
            alert(e.response.data.message)
        })
    
    }

    const onClickSignUpHandler = () => {
        console.log('try to sign up')
        signUp()
    }


    // useEffect(() => {
    //     login()
    // },
    // // 페이지 호출 후 처음 한번만 호출될 수 있도록 [] 추가
    // [])
 
    return(
        <div>
            <h2>Sign Up</h2>
            <div>
                <label htmlFor='input_id'>EMAIL : </label>
                <input type='text' name='mail' value={mail} onChange={onChangeUserHandler} />
            </div>
            <div>
                <label htmlFor='input_id'>NAME : </label>
                <input type='text' name='name' value={name} onChange={onChangeUserHandler} />
            </div>
            <div>
                <label htmlFor='input_pw'>PW : </label>
                <input type='password' name='password' value={password} onChange={onChangeUserHandler} />
            </div>
            <div>
                <label htmlFor='input_id'>GENDER(male/female) : </label>
                <input type='text' name='gender' value={gender} onChange={onChangeUserHandler} />
            </div>
            <div>
                <button type='button' onClick={onClickSignUpHandler}>OK</button>
            </div>
        </div>
    )
}
 
export default SignUp;