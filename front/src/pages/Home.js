import { React, useEffect, useId, useState } from 'react'
import axios from 'axios'
import { Form, Container, Button, Row, Col, Alert } from 'react-bootstrap'
import ConnectCoupleSender from '../components/ConnectCoupleSender'
import ConnectCoupleReceiver from '../components/ConnectCoupleReceiver'
const Home = () => {
    const userId = localStorage.getItem('userid')
    const name = localStorage.getItem('name')
    const [senderCoupleInfo, setSenderCoupleInfo] = useState(null)
    const [receiverCoupleInfo, setReceiverCoupleInfo] = useState(null)

    console.log(userId)
    const getCouplebySenderId = async (useId) => {
        if (userId == null) {
            return
        }
        try {
            let res = await axios.get('/apis/v1/couples?userid=' + userId)
            if (res.status == 204) {
                setSenderCoupleInfo({})
                return
            }
            let payload = res.data.payload
            setSenderCoupleInfo(payload)
        } catch (err) {
            console.log(err)
        }
    }

    // receiver id를 통해서 sender id를 조회 
    const getAllCoupleRequestByRecvId = async (useId) => {
        if (userId == null) {
            return
        }
        try {
            let res = await axios.get(
                '/apis/v1/couples/senders?userid=' + useId
            )
            if (res.status == 204) {
                setReceiverCoupleInfo([])
                return
            }
            let payload = res.data.payload
            setReceiverCoupleInfo(payload)
        } catch (err) {
            console.log(err)
        }
    }

    useEffect(() => {
        getCouplebySenderId(userId)
        getAllCoupleRequestByRecvId(userId)
    }, [])

    if (senderCoupleInfo == null || receiverCoupleInfo == null) {
        return (
            <Container>
                <Row>
                    <Col>
                        <h1>Loading</h1>
                    </Col>
                </Row>
            </Container>
        )
    } else {
        return (
            <Container>
                <Row>
                    <Col
                        sm={{ span: 6, offset: 2 }}
                        lg={{ span: 4, offset: 3 }}
                    >
                        <h3>안녕하세요! {name} 님!</h3>
                    </Col>
                </Row>

                <Alert key="to" variant="primary">
                    내가 보낸 요청
                </Alert>
                <ConnectCoupleSender coupleInfo={senderCoupleInfo} />
                <br />

                <Alert key="from" variant="primary">
                    내가 받은 요청
                </Alert>
                <ConnectCoupleReceiver coupleInfo={receiverCoupleInfo} />
            </Container>
        )
    }
}

export default Home
