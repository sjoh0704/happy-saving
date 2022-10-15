import {React, useEffect, useState} from "react";
import axios from 'axios';
import ConnectCouple from "../components/ConnectCouple";

const Home =() => {


    const userId = localStorage.getItem("userid")
    const name = localStorage.getItem("name")
    const [coupleInfo, setCoupleInfo] = useState(null)
    
    console.log(userId)
    const getCouplebyUserId = async (userId) =>{
        if (userId == null){
            return
        }
        try {
            let res = await axios.get("/apis/v1/couples?userid=" + userId)
            let payload = res.data.payload
            setCoupleInfo(payload)

        } catch (err) {
            console.log(err)
        }
    }

    useEffect(() => {
        getCouplebyUserId(userId)
    },
    [])
 

    return (
    <div>
        <h3>
            안녕하세요! {name} 님!
        </h3>
        <div>
            <ConnectCouple coupleInfo={coupleInfo}/>
        </div>
    </div>
    );
}

export default Home;