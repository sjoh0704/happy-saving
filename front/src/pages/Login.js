import React, { useState, useEffect } from 'react'
import axios from 'axios'
import { useNavigate } from 'react-router'
import { Button, Col, Container, Form, Row } from 'react-bootstrap'
import Title from '../components/Title'

const style = {
    // display: 'inline-block',
    // width: '100px',
    // height: '100px',
    // boder: '1px solid black',
    // background: 'orange'
}

function Login() {
    const [mail, setMail] = useState('')
    const [password, setPassword] = useState('')
    const navigate = useNavigate()

    const onChangeMailHandler = (e) => {
        setMail(e.target.value)
    }

    const onChangePasswordHandler = (e) => {
        setPassword(e.target.value)
    }

    const login = async () => {
        let body = {
            mail: mail,
            password: password
        }

        try {
            let res = await axios.post('/auth', body)
            alert(res.data.message)

            let payload = res.data.payload
            localStorage.setItem('mail', payload.mail)
            localStorage.setItem('userid', payload.id)
            localStorage.setItem('name', payload.name)
            localStorage.setItem('gender', payload.gender)
            navigate('/')
            return
        } catch (err) {
            alert(err.response.data.message)
            console.log(err.response.data.message)
        }
    }

    const onClickLoginHandler = () => {
        console.log('try to login')
        login()
    }

    return (
        <Container>
            <Row>
                <Col
                    lg={{ span: 8, offset: 2 }}
                    sm={{ span: 10, offset: 1 }}
                    xs={{ span: 12, offset: 0 }}
                >
                    <Title title="로그인" />
                </Col>
            </Row>
            <br />

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
                                onChange={onChangeMailHandler}
                            />
                            <Form.Text className="text-muted">
                                We'll never share your email with anyone else.
                            </Form.Text>
                        </Form.Group>
                    </Form>
                </Col>
            </Row>
            <br />
            <Row>
                <Col
                    lg={{ span: 8, offset: 2 }}
                    sm={{ span: 10, offset: 1 }}
                    xs={{ span: 12, offset: 0 }}
                >
                    <Form.Group className="mb-3" controlId="formPassword">
                        <Form.Label>비밀번호</Form.Label>
                        <Form.Control
                            type="password"
                            placeholder="Password"
                            name="password"
                            value={password}
                            onChange={onChangePasswordHandler}
                        />
                    </Form.Group>
                </Col>
            </Row>
            <Row>
                <Col
                    lg={{ span: 8, offset: 2 }}
                    sm={{ span: 10, offset: 1 }}
                    xs={{ span: 12, offset: 0 }}
                >
                    <Button
                        style={{ minWidth: '25%' }}
                        variant="primary"
                        type="submit"
                        onClick={onClickLoginHandler}
                    >
                        Login
                    </Button>
                </Col>
            </Row>
        </Container>
    )
}

export default Login
