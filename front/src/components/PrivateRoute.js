import {React, useState, useEffect} from "react"
import { Navigate } from "react-router"
import axios from 'axios';

// login 하지 않았거나 token이 만료되었다면 login으로 리다이렉션
const PrivateRouter = () =>{
    const mail = localStorage.getItem("mail")

    const [status, setStatus] = useState(0)

    if (mail == null){
        return(
            <Navigate to="/login"/>
        )
    }
    // mail이 있을 때    
    // const checkCookie = () => {
    //     axios
    //     .get("/apis/v1/users?mail=" + mail)
    //     .then( res => {
    //         console.log(res.data)
    //     })
    //     .catch(e => {
    //         // 토큰이 없거나 유저가 없을 때
    //         console.log(e.response.status)
    //         setStatus(e.response.status)
    //     })

    //     if (status == 400 || status == 403){
    //         return(
    //             <Navigate to="/login"/>
    //         )
    //     }
    // }


    // useEffect(() => {
    //     checkCookie()
    // },
    // [status, mail])

}

export default PrivateRouter