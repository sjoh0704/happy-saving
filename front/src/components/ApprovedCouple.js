import React from 'react'
import { Accordion, Badge, Button, Col, Row, Table } from 'react-bootstrap'
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
            <Accordion.Item eventKey={index} key={index}>
                <Accordion.Header>
                    #{postSize - index} {post.title}
                    <span>
                        ........
                    </span>
                    <Badge bg="primary">
                        {findAuthorName(post.author_id, coupleInfo)}
                    </Badge>
                </Accordion.Header>
                <Accordion.Body>
                    <p>{changeTimeFormat(post.created_at)}</p>

                    <p>{post.content}</p>
                </Accordion.Body>
            </Accordion.Item>
            // <tr key={index}>
            //     <td>{postSize - index}</td>
            //     <td>
            //         <a href={'/post/' + post.id}>{post.title}</a>
            //     </td>
            //     <td>{findAuthorName(post.author_id, coupleInfo)}</td>
            //     <td>{changeTimeFormat(post.created_at)}</td>
            // </tr>
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
            <Accordion>{prettyPosts}</Accordion>
            <br />
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
