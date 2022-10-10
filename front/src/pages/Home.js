import {React, useEffect} from "react";
import axios from 'axios';

const Home =({ history }) => {


    const mail = localStorage.getItem("mail")
    
    console.log(mail)
    // const getUserbyMail = async (mail) =>{
    //     if (mail == null) {
    //         return 
    //     }
    //     axios
    //     .get("/apis/v1/users?mail=" + mail)
    //     .then( res => {
    //         console.log(res.data)
    //     })
    //     .catch(e => {
    //         console.log(e.response.data.message)
    //     })
    // }

    // useEffect(() => {
    //     getUserbyMail(mail)
    // },
    // [])
 

    return (
        <div>
            <h1>
                Home
            </h1>
        </div>
    );
}

export default Home;