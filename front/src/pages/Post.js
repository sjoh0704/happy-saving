import axios from 'axios'
import React, { useEffect, useState } from 'react'
import { Container, Row, Col } from 'react-bootstrap'
import { useMatch, useNavigate } from 'react-router'
import PostCard from '../components/Postcard'

const Post = () => {
    const match = useMatch('/post/:id')
    const navigate = useNavigate()
    const { id } = match.params
    const [post, setPost] = useState(null)

    const getPostByPostId = async (id) => {
        try {
            let res = await axios.get('/apis/v1/posts/' + id)
            let payload = res.data.payload
            console.log(payload)
            setPost(payload)
        } catch (err) {
            console.log(err)
            // navigate('/notfound')
        }
    }

    useEffect(() => {
        getPostByPostId(id)
    }, [])
    if (post == null) {
        return (
            <Container>
                <Row>
                    <Col>
                        <h1>Loading</h1>
                    </Col>
                </Row>
            </Container>
        )
    } else {
        return (
            <Container>
                <Row>
                    <Col><PostCard postInfo={post}/></Col>
                </Row>
            </Container>
        )
    }
}

export default Post
