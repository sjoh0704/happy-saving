import {React, useEffect, useState} from "react";
import axios from 'axios';
import { useNavigate } from "react-router";

const MakeCouple =() => {
    const navigate = useNavigate()
    const senderId = parseInt(localStorage.getItem("userid"))
    const [mail, setMail] = useState("")
    const [candidate, setCandidate] = useState(null)

    // const getCouplebyUserId = async (userId) =>{
    //     if (userId == null){
    //         return
    //     }
    //     try {
    //         let res = await axios.get("/apis/v1/couples?userid=" + userId)
    //         console.log(res)
    //         let payload = res.data.payload
    //         console.log(payload)
            
    //         setCoupleInfo(payload.phase)

    //     } catch (err) {
    //         console.log(err)
    //     }
    // }

    const onChangeHandler = (e) =>{
        setMail(e.target.value)
    }

    const getUserByEmail = async(mail) =>{
        try {
            let res = await axios.get("/apis/v1/users?mail=" + mail)
            let payload = res.data.payload
            // alert(payload.name)
            if (senderId == payload.id){
                alert("자신의 이메일은 조회할 수 없습니다.")
                return
            }

            setCandidate({
                id: parseInt(payload.id),
                name: payload.name,
                mail: payload.mail
            })
        } catch (err) {
            setCandidate(null)
            console.log(err)
        }
    }

    const MakeCouple = async(candidate) =>{
        if (candidate == null){
            console.log("잘못된 요청")
            return
        }
        try {
            let body = {
                send_id: senderId,
                recv_id: candidate.id
            }
            await axios.post("/apis/v1/couples", body)
            alert("상대방에게 커플 요청을 보냈습니다. 상대방이 수락하기전까지 기다려주세요")
            navigate("/")

        } catch (err) {
            if (err instanceof axios.AxiosError){
                alert(err.response.data.message)
            }
            
        }
    }

    const onClickHandler = () =>{
        getUserByEmail(mail)

    }

    const onClickMakeCoupleHandler = () =>{
        MakeCouple(candidate)
    }

    // useEffect(() => {
    //     // getCouplebyUserId(userId)
    // },
    // [])
 
    return (
        <div>
            <h2>사용자 검색</h2>
            <div>
                <label htmlFor='mail'>사용자 이메일: </label>
                <input type='text' name='mail' value={mail} onChange={onChangeHandler} />
            </div>

            <div>
                <button type='button' onClick={onClickHandler}>검색</button>
            </div>

            <div>
                {candidate == null ?
                <div>
                    없음
                </div>:
                <div>
                    <p>
                        email: {candidate.mail}    
                    </p>
                    <p>
                        name: {candidate.name}    
                    </p>

                     <button type='button' onClick={onClickMakeCoupleHandler}>커플 신청하기</button>
                    
                </div>}

            </div>
        </div>
    );
}

export default MakeCouple;