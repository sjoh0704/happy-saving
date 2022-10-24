import React from 'react'
import { Container, Row, Col, Navbar } from 'react-bootstrap'
import Nav from 'react-bootstrap/Nav'

const NavigationBar = () => {
    return (
        <>
            <Navbar bg="primary" variant="dark">
                <Container>
                    <Navbar.Brand href="/">🌈 Happy Saving</Navbar.Brand>
                    <Nav className="me-auto">
                        <Nav.Link href="/">Home</Nav.Link>
                        <Nav.Link href="/login">로그인</Nav.Link>
                        <Nav.Link href="/signup">회원가입</Nav.Link>
                    </Nav>
                </Container>
            </Navbar>
            <br/>
        </>
    )
}

export default NavigationBar
