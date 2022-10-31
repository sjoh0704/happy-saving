import { Col, Row } from 'react-bootstrap'
import Spinner from 'react-bootstrap/Spinner'

function Loading() {
    return (
        <>
            <br />
            <Row>
                <Col
                    lg={{ span: 3, offset: 4 }}
                    sm={{ span: 4, offset: 1 }}
                    xs={{ span: 5, offset: 1 }}
                >
                    <p style={{ fontSize: '1.8rem', fontWeight: 'bold' }}>
                        Loading...
                    </p>
                </Col>
                <Col
                    lg={{ span: 4, offset: 0 }}
                    sm={{ span: 6, offset: 0 }}
                    xs={{ span: 6, offset: 0 }}
                >
                    <Spinner animation="border" size="xxl"></Spinner>
                </Col>
            </Row>

            <br />
        </>
    )
}

export default Loading
