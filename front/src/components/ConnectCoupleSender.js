import axios from 'axios'
import { useEffect, useState } from 'react'
import { Row, Col, Button } from 'react-bootstrap'

const ConnectCoupleSender = (props) => {
    const userId = localStorage.getItem('userid')
    // const [coupleInfo, setCoupleInfo] = useState(props.coupleInfo)
    let coupleInfo = props.coupleInfo

    const onClickCoupleCancleHandler = async () => {
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

    if (coupleInfo == null) {
        return (
            <Row>
                <Col>
                    <p>couple 등록이 안되어 있습니다.</p>
                    <p>couple 등록을 해주세요</p>

                    <a href="/couple/make">couple 등록 신청</a>
                </Col>
            </Row>
        )
    }

    if (coupleInfo.phase == 'awaiting') {
        return (
            <Row>
                <Col>
                    <p>
                        {coupleInfo.sender.name} 님에게 커플 요청을 보냈습니다.
                    </p>
                </Col>
                <Col>
                    <Button
                        variant="primary"
                        onClick={onClickCoupleCancleHandler}
                    >
                        신청 취소하기
                    </Button>
                </Col>
            </Row>
        )
    } else if (coupleInfo.phase == 'denyed') {
        return (
            <Row>
                <Col>
                    <p>
                        {coupleInfo.sender.name} 님에게 커플 요청을 거절
                        당했거나 등록을 취소하셨습니다.
                    </p>
                    <p>
                        <a href="/couple/make">새로운 couple 등록 신청</a>
                    </p>
                </Col>
            </Row>
        )
    } else {
        return (
            <Row>
                <Col>승인됨</Col>
            </Row>
        )
    }
}

export default ConnectCoupleSender
