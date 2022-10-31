import { React, useEffect, useId, useState } from 'react'
import axios from 'axios'
import { Form, Container, Button, Row, Col, Alert } from 'react-bootstrap'
import ConnectCoupleSender from '../components/ConnectCoupleSender'
import ConnectCoupleReceiver from '../components/ConnectCoupleReceiver'
import ApprovedCouple from '../components/ApprovedCouple'
import Loading from '../components/Loading'
const Home = () => {
    const userId = localStorage.getItem('userid')
    const name = localStorage.getItem('name')
    const [approvedCoupleInfo, setApprovedCoupleInfo] = useState(null)
    const [senderCoupleInfo, setSenderCoupleInfo] = useState(null)
    const [receiverCoupleInfo, setReceiverCoupleInfo] = useState(null)

    const getCouplebySenderId = async (useId) => {
        if (userId == null) {
            return
        }
        try {
            let res = await axios.get(
                '/apis/v1/couples/senders?phase=awaiting&userid=' + userId
            )
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
                '/apis/v1/couples/receivers/all?userid=' + useId
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

    const getAprrovedCoupleRelationbyUserId = async (useId) => {
        if (userId == null) {
            return
        }
        try {
            let res = await axios.get(
                '/apis/v1/couples?phase=approved&userid=' + userId
            )
            if (res.status != 204) {
                // 커플 연결이 성사된 상태
                let payload = res.data.payload

                setApprovedCoupleInfo(payload)
                setSenderCoupleInfo({})
                setReceiverCoupleInfo({})
                localStorage.setItem('couple_id', payload.id)
                return
            }
            setApprovedCoupleInfo({})
            getCouplebySenderId(userId)
            getAllCoupleRequestByRecvId(userId)
        } catch (err) {
            console.log(err)
        }
    }

    useEffect(() => {
        getAprrovedCoupleRelationbyUserId(userId)
    }, [])

    if (
        approvedCoupleInfo == null ||
        senderCoupleInfo == null ||
        receiverCoupleInfo == null
    ) {
        return (
            <Container>
                <Row>
                    <Col
                        lg={{ span: 8, offset: 2 }}
                        sm={{ span: 10, offset: 1 }}
                        xs={{ span: 12, offset: 0 }}
                    >
                        <Loading />
                    </Col>
                </Row>
            </Container>
        )
    } else {
        return (
            <Container>
                <Row>
                    <Col
                        lg={{ span: 8, offset: 2 }}
                        sm={{ span: 10, offset: 1 }}
                        xs={{ span: 12, offset: 0 }}
                    >
                        {JSON.stringify(approvedCoupleInfo) !=
                        JSON.stringify({}) ? (
                            <ApprovedCouple coupleInfo={approvedCoupleInfo} />
                        ) : (
                            <>
                                <Alert key="to" variant="primary">
                                    <p
                                        style={{
                                            padding: [0, 'auto'],
                                            fontWeight: 'bold',
                                            fontSize: '1.3rem'
                                        }}
                                    >
                                        내가 보낸 요청
                                    </p>
                                </Alert>
                                <ConnectCoupleSender
                                    coupleInfo={senderCoupleInfo}
                                />
                                <br />
                                <div
                                style={{height: '3rem'}}>

                                </div>

                                <Alert key="from" variant="primary">
                                    <p
                                        style={{
                                            padding: [0, 'auto'],
                                            fontWeight: 'bold',
                                            fontSize: '1.3rem'
                                        }}
                                    >
                                        내가 받은 요청
                                    </p>
                                </Alert>
                                <ConnectCoupleReceiver
                                    coupleInfo={receiverCoupleInfo}
                                />
                            </>
                        )}
                    </Col>
                </Row>
            </Container>
        )
    }
}

export default Home
