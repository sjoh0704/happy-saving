import { useEffect, useState } from 'react'
import { Col, Row } from 'react-bootstrap'
import Button from 'react-bootstrap/Button'
import Card from 'react-bootstrap/Card'
import { changeTimeFormat } from '../util/util'
import axios from 'axios'
import { Navigate, useNavigate } from 'react-router'

function PostCard(props) {
    const [auth, setAuth] = useState(false)
    const navigate = useNavigate()
    const { id, couple_id, title, content, image_url, created_at, updated_at } =
        props.postInfo

    const checkCoupleAuth = async () => {
        try {
            let res = await axios.get('/apis/v1/posts/' + id)
            let payload = res.data.payload
            if (payload.couple_id == couple_id) {
                setAuth(true)
                return
            }
            setAuth(false)
            alert('잘못된 요청입니다.')
            navigate('/')
        } catch (err) {
            console.log(err)
            setAuth(false)
            alert('잘못된 요청입니다.')
            navigate('/')
        }
    }

    useEffect(() => {
        checkCoupleAuth()
    }, [])

    return (
        <>
            {auth ? (
                <Card style={{ width: 'auto' }}>
                    <Card.Img variant="top" src="img/logo512.png" />
                    <Card.Body>
                        <Card.Title>{title}</Card.Title>
                        <Row>
                            <Col>
                                <Card.Text>
                                    생성 날짜: {changeTimeFormat(created_at)}/
                                    {changeTimeFormat(updated_at)}
                                </Card.Text>
                            </Col>
                            <Col>
                                <Card.Text>수정 날짜: </Card.Text>
                            </Col>
                        </Row>
                        <br />
                        <Card.Text>{content}</Card.Text>

                        <Button variant="primary">Go somewhere</Button>
                    </Card.Body>
                </Card>
            ) : (
                <></>
            )}
        </>
    )
}

export default PostCard
