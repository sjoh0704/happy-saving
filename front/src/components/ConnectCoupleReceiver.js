import axios from 'axios'
import { useEffect, useState } from 'react'
import { Row, Col, Button, Card } from 'react-bootstrap'
import UserCard from './UserCard'

const ConnectCoupleReceiver = (props) => {
    const userId = localStorage.getItem('userid')
    let coupleInfo = props.coupleInfo

    const onClickCoupleCancleHandler = async (coupleInfo) => {
        if (coupleInfo == null) {
            return
        }
        try {
            let body = {
                phase: 'denyed'
            }
            await axios.put('/apis/v1/couples/' + coupleInfo.id, body)
            alert('커플 등록이 취소 되었습니다.')
            window.location.reload()
            // setCoupleInfo(null)
        } catch (err) {
            console.log(err)
            alert('문제가 발생했습니다.\n 관리자에게 문의하세요')
        }
    }

    let coupleRequestList = coupleInfo.map((c, index) => (
        <Row key={index}>
            <Col>
                <UserCard
                    buttonValue="거절하기"
                    onClick={() => onClickCoupleCancleHandler(c)}
                    user={c.sender}
                />
            </Col>
        </Row>
    ))

    if (coupleInfo.size == 0) {
        return (
            <Row>
                <Col>들어온 couple 요청이 없습니다. </Col>
            </Row>
        )
    }

    return <>{coupleRequestList}</>
}

export default ConnectCoupleReceiver
