import axios from 'axios'
import { useEffect, useState } from 'react'
import { Row, Col, Button } from 'react-bootstrap'
import { useNavigate } from 'react-router'

const ConnectCoupleSender = (props) => {
    const userId = localStorage.getItem('userid')
    const navigate = useNavigate()
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

    // 빈 객체일 때
    if (JSON.stringify(coupleInfo) == JSON.stringify({})) {
        return (
            <>
                <Row>
                    <Col>
                        <p
                            style={{
                                fontWeight: 'bold',
                                padding: '5%',
                                fontSize: '1.2rem'
                            }}
                        >
                            couple 등록이 안되어 있습니다. couple 등록을
                            해주세요
                        </p>
                    </Col>
                </Row>
                <Row>
                    <Col>
                        <Button
                            style={{ minWidth: '25%' }}
                            variant="primary"
                            type="submit"
                            onClick={() => {
                                navigate('/couple/make')
                            }}
                        >
                            커플 등록하기
                        </Button>
                    </Col>
                </Row>
            </>
        )
    }

    if (coupleInfo.phase == 'awaiting') {
        return (
            <Row>
                <Col>
                    <p
                        style={{
                            fontWeight: 'bold',
                            padding: '5%',
                            fontSize: '1.2rem'
                        }}
                    >
                        {coupleInfo.receiver.name} 님에게 커플 요청을
                        보냈습니다.
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
