import { React, useEffect, useState } from 'react'
import axios from 'axios'
import { useNavigate } from 'react-router'
import { Button, Col, Container, Form, Row } from 'react-bootstrap'
import UserCard from '../components/UserCard'
const MakeCouple = () => {
    const navigate = useNavigate()
    const senderId = parseInt(localStorage.getItem('userid'))
    const senderMail = localStorage.getItem('mail')
    const [mail, setMail] = useState('')
    const [candidate, setCandidate] = useState(null)

    const onChangeHandler = (e) => {
        setMail(e.target.value)
    }

    const getUserByEmail = async (mail) => {
        try {
            let res = await axios.get('/apis/v1/users?mail=' + mail)

            // 자신의 이메일을 조회할 때
            if (senderMail == mail) {
                alert('자신의 이메일은 조회할 수 없습니다.')
                return
            }
            // 존재하지 않는 사용자일때
            if (res.status == 204) {
                alert('없는 사용자 입니다. ')
                return
            }
            let payload = res.data.payload
            console.log(payload)
            setCandidate({
                id: parseInt(payload.id),
                name: payload.name,
                mail: payload.mail,
                gender: payload.gender,
                created_at: payload.created_at
            })
        } catch (err) {
            setCandidate(null)
            console.log(err)
        }
    }

    const MakeCouple = async (candidate) => {
        if (candidate == null) {
            console.log('잘못된 요청')
            return
        }
        try {
            let body = {
                send_id: senderId,
                recv_id: candidate.id
            }
            await axios.post('/apis/v1/couples', body)
            alert(
                '상대방에게 커플 요청을 보냈습니다. 상대방이 수락하기전까지 기다려주세요'
            )
            navigate('/')
        } catch (err) {
            if (err instanceof axios.AxiosError) {
                alert(err.response.data.message)
            }
        }
    }

    const onClickHandler = () => {
        getUserByEmail(mail)
    }

    const onClickMakeCoupleHandler = () => {
        MakeCouple(candidate)
    }

    // useEffect(() => {
    //     // getCouplebyUserId(userId)
    // },
    // [])

    return (
        <Container>
            <Row>
                <Col
                    lg={{ span: 8, offset: 2 }}
                    sm={{ span: 10, offset: 1 }}
                    xs={{ span: 12, offset: 0 }}
                >
                    <p style={{ fontWeight: 'bold', fontSize: '1.5rem' }}>
                        커플 등록하기
                    </p>
                </Col>
            </Row>

            <Row>
                <Col
                    lg={{ span: 8, offset: 2 }}
                    sm={{ span: 10, offset: 1 }}
                    xs={{ span: 12, offset: 0 }}
                >
                    <Form>
                        <Form.Group className="mb-3" controlId="formEmail">
                            <Form.Label>이메일 주소</Form.Label>
                            <Form.Control
                                type="email"
                                placeholder="Enter email"
                                name="mail"
                                value={mail}
                                onChange={onChangeHandler}
                            />
                            <Form.Text className="text-muted">
                                상대방 이메일을 검색하세요
                            </Form.Text>
                        </Form.Group>
                    </Form>
                </Col>
            </Row>
            <Row>
                <Col
                    lg={{ span: 8, offset: 2 }}
                    sm={{ span: 10, offset: 1 }}
                    xs={{ span: 12, offset: 0 }}
                >
                    <Button
                        variant="primary"
                        type="submit"
                        onClick={onClickHandler}
                        style={{ minWidth: '25%' }}
                    >
                        검색
                    </Button>
                </Col>
            </Row>
            <div style={{ height: '3rem' }}></div>

            <Row>
                <Col
                    lg={{ span: 8, offset: 2 }}
                    sm={{ span: 10, offset: 1 }}
                    xs={{ span: 12, offset: 0 }}
                >
                    {candidate == null ? (
                        <div></div>
                    ) : (
                        <UserCard
                            buttonValueLeft="커플 신청하기"
                            onClickLeft={onClickMakeCoupleHandler}
                            user={candidate}
                        />
                    )}
                </Col>
            </Row>
        </Container>
    )
}

export default MakeCouple
