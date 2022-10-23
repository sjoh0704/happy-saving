import { Col, Row } from 'react-bootstrap'
import Button from 'react-bootstrap/Button'
import Card from 'react-bootstrap/Card'
import { changeTimeFormat } from '../util/util'
function PostCard(props) {
    const { title, content, image_url, created_at, updated_at } = props.postInfo

    return (
        <Card style={{ width: 'auto' }}>
            <Card.Img variant="top" src="img/logo512.png" />
            <Card.Body>
                <Card.Title>{title}</Card.Title>
                <Row>
                    <Col>
                        <Card.Text>생성 날짜: {changeTimeFormat(created_at)}/{changeTimeFormat(updated_at)}</Card.Text>
                    </Col>
                    <Col>
                        <Card.Text>수정 날짜: </Card.Text>
                    </Col>
                </Row>
                <br/>
                <Card.Text>{content}</Card.Text>
       

                <Button variant="primary">Go somewhere</Button>
            </Card.Body>
        </Card>
    )
}

export default PostCard
