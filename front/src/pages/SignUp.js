import React, { useState, useEffect } from 'react'
import axios from 'axios'
import { useNavigate } from 'react-router'
import { Button, Col, Container, Form, Row } from 'react-bootstrap'
function SignUp() {
    const navigate = useNavigate()
    const [userData, setUserData] = useState({
        name: '',
        mail: '',
        password: '',
        password_check: '',
        gender: ''
    })

    const { name, mail, password, password_check, gender } = userData

    const onChangeUserHandler = (e) => {
        const { name, value } = e.target
        setUserData({
            ...userData,
            [name]: value
        })
    }

    const signUp = async () => {
        try {
            if (password != password_check) {
                alert('비밀번호가 일치하지 않습니다.')
                return
            }
            await axios.post('/apis/v1/users', userData)
            alert('회원가입이 정상적으로 이루어졌습니다.')
            navigate('/login')
        } catch (err) {
            console.log(err.response.data.message)
            alert(err.response.data.message)
        }
    }

    const onClickSignUpHandler = () => {
        console.log('try to sign up')
        signUp()
    }

    // useEffect(() => {
    //     login()
    // },
    // // 페이지 호출 후 처음 한번만 호출될 수 있도록 [] 추가
    // [])

    return (
        <Container>
            <Row>
                <Col>
                    <h2>Sign Up</h2>
                </Col>
            </Row>
            <br />
            <Row>
                <Col>
                    <Form>
                        <Form.Group className="mb-3" controlId="formEmail">
                            <Form.Label>Email address</Form.Label>
                            <Form.Control
                                type="email"
                                placeholder="Enter email"
                                name="mail"
                                value={mail}
                                onChange={onChangeUserHandler}
                            />
                            <Form.Text className="text-muted">
                                We'll never share your email with anyone else.
                            </Form.Text>
                        </Form.Group>
                    </Form>
                </Col>
            </Row>
            <Row>
                <Col>
                    <Form>
                        <Form.Group className="mb-3" controlId="formName">
                            <Form.Label>User Name</Form.Label>
                            <Form.Control
                                type="email"
                                placeholder="Enter name"
                                name="name"
                                value={name}
                                onChange={onChangeUserHandler}
                            />
                            <Form.Text className="text-muted">
                                사용할 애칭이나 이름을 적어주세요
                            </Form.Text>
                        </Form.Group>
                    </Form>
                </Col>
            </Row>
            <Row>
                <Col>
                    <Form.Group className="mb-3" controlId="formPassword">
                        <Form.Label>비밀번호</Form.Label>
                        <Form.Control
                            type="password"
                            placeholder="Password"
                            name="password"
                            value={password}
                            onChange={onChangeUserHandler}
                        />
                    </Form.Group>
                </Col>
            </Row>
            <Row>
                <Col>
                    <Form.Group className="mb-3" controlId="formPasswordCheck">
                        <Form.Label>비밀번호 확인</Form.Label>
                        <Form.Control
                            type="password"
                            placeholder="Password check"
                            name="password_check"
                            value={password_check}
                            onChange={onChangeUserHandler}
                        />
                    </Form.Group>
                </Col>
            </Row>
            <Row>
                <Col>
                    <Form>
                        <Form.Group className="mb-3" controlId="formGender">
                            <Form.Label>성별(male/female)</Form.Label>
                            <Form.Control
                                type="email"
                                placeholder="male"
                                name="gender"
                                value={gender}
                                onChange={onChangeUserHandler}
                            />
                        </Form.Group>
                    </Form>
                </Col>
            </Row>
            <Row>
                <Col>
                    <Button
                        variant="primary"
                        type="submit"
                        onClick={onClickSignUpHandler}
                    >
                        OK
                    </Button>
                </Col>
            </Row>
        </Container>
    )
}

export default SignUp
