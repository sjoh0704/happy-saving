import React from 'react'
import { Container, Row, Col } from 'react-bootstrap'
import Nav from 'react-bootstrap/Nav'

const Page = () => {
    return (
        <Container>
            <br/>
            <Row>
                <Col>
                    <Nav variant="pills" defaultActiveKey="/home">
                        <Nav.Item>
                            <Nav.Link href="/">홈</Nav.Link>
                        </Nav.Item>
                        <Nav.Item>
                            <Nav.Link href="/login">로그인</Nav.Link>
                        </Nav.Item>
                        <Nav.Item>
                            <Nav.Link href="/signup">회원가입</Nav.Link>
                        </Nav.Item>
                    </Nav>
                </Col>
            </Row>
            <br/>
        </Container>
    )
}

export default Page
