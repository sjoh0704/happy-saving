import React from 'react'
import { Table } from 'react-bootstrap'

const ApprovedCouple = (props) => {
    let coupleInfo = props.coupleInfo
    let postSize = coupleInfo.posts.length
    let sender = coupleInfo.sender
    let receiver = coupleInfo.receiver
    let prettyPosts = coupleInfo.posts.map((post, index) => (
        <tr key={index}>
            <td>{postSize - index}</td>
            <td>{post.title}</td>
            <td>{findAuthorName(post.author_id, coupleInfo)}</td>
            <td>{post.created_at}</td>
        </tr>
    ))

    return (
        <>
            <h3>
                {sender.name} 하트 {receiver.name}
            </h3>
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
