import {React, useEffect, useState} from "react";
import axios from 'axios';

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
            console.log(res)
            let payload = res.data.payload
            console.log(payload)
            
            setCoupleInfo(payload.phase)

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
            {coupleInfo == null ? 
            <div>
            당신은 아직 couple 등록이 안되어 있습니다!
            couple 신청을 해야합니다. 
            <p>
                <a href="/couple/make">
                couple 신청     
                </a>
            </p>
            </div>:
            <div>
                {coupleInfo}
                
            </div>}
        </div>
    </div>
    );
}

export default Home;