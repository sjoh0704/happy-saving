import React from 'react'
import { Button, Col, Row, Table } from 'react-bootstrap'
import { useNavigate } from 'react-router'
import { changeTimeFormat } from '../util/util'

const ApprovedCouple = (props) => {
    const navigate = useNavigate()
    let coupleInfo = props.coupleInfo
    let postSize = coupleInfo.posts.length
    let sender = coupleInfo.sender
    let receiver = coupleInfo.receiver
    let prettyPosts = coupleInfo.posts
        .slice(0)
        .reverse()
        .map((post, index) => (
            <tr key={index}>
                <td>{postSize - index}</td>
                <td>
                    <a href={'/post/' + post.id}>{post.title}</a>
                </td>
                <td>{findAuthorName(post.author_id, coupleInfo)}</td>
                <td>{changeTimeFormat(post.created_at)}</td>
            </tr>
        ))

    const onClickMoveAddPostPage = () => {
        navigate('/post/create')
    }

    return (
        <>
            <h3>
                {sender.name} ❤ {receiver.name}
            </h3>
            <br />
            <Table striped bordered hover>
                <thead>
                    <tr>
                        <th>#</th>
                        <th>Title</th>
                        <th>Username</th>
                        <th>CreatedAt</th>
                    </tr>
                </thead>
                <tbody>{prettyPosts}</tbody>
            </Table>
            <Row>
                <Col>
                    <Button onClick={onClickMoveAddPostPage}>추가</Button>{' '}
                    <Button disabled>업데이트</Button>{' '}
                    <Button disabled>삭제</Button>
                </Col>
            </Row>
        </>
    )
}

const findAuthorName = (authorId, coupleInfo) => {
    let sender = coupleInfo.sender
    let receiver = coupleInfo.receiver

    if (authorId == sender.id) {
        return sender.name
    } else if (authorId == receiver.id) {
        return receiver.name
    }
    return null
}

export default ApprovedCouple
