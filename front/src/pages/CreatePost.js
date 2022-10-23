import axios from 'axios'
import React, { useState } from 'react'
import {
    Button,
    Col,
    Container,
    FloatingLabel,
    Form,
    Row
} from 'react-bootstrap'
import { useNavigate } from 'react-router'

const CreatePost = () => {
    const navigate = useNavigate()
    const couple_id = parseInt(localStorage.getItem('couple_id'))
    const author_id = parseInt(localStorage.getItem('userid'))
    const [postData, setPostData] = useState({
        title: '',
        content: '',
        couple_id: couple_id,
        author_id: author_id,
        image_url: ''
    })
    let { title, content, image_url } = postData

    const onChangePostHandler = (e) => {
        const { name, value } = e.target
        setPostData({
            ...postData,
            [name]: value
        })
        console.log(postData)
    }

    const onClickCreatePost = async (postData) => {
        try {
            let body = postData
            let res = await axios.post('/apis/v1/posts', body)
            navigate('/')
        } catch (err) {
            console.log(err)
            alert('문제가 발생했습니다. \n관리자에게 문의하세요')
        }
    }

    return (
        <>
            <Container>
                <Row>
                    <Col>
                        <h3>post 작성하기</h3>
                    </Col>
                </Row>
                <br />

                <Row>
                    <Col>
                        <FloatingLabel
                            controlId="floatingTextarea"
                            label="제목을 작성하세요"
                            className="mb-3"
                        >
                            <Form.Control
                                as="textarea"
                                name="title"
                                value={title}
                                onChange={onChangePostHandler}
                                placeholder="Leave a comment here"
                            />
                        </FloatingLabel>
                    </Col>
                </Row>

                <Row>
                    <Col>
                        <FloatingLabel
                            controlId="floatingTextarea2"
                            label="내용을 작성하세요"
                        >
                            <Form.Control
                                as="textarea"
                                name="content"
                                value={content}
                                onChange={onChangePostHandler}
                                placeholder="Leave a comment here"
                                style={{ height: '200px' }}
                            />
                        </FloatingLabel>
                    </Col>
                </Row>
                <br />
                <Row>
                    <Col>
                        <Button onClick={() => onClickCreatePost(postData)}>
                            작성 완료
                        </Button>
                    </Col>
                </Row>
            </Container>
        </>
    )
}

export default CreatePost
