import React from 'react'
import { Button, Card } from 'react-bootstrap'

const UserCard = (props) => {
    let buttonValue = props.buttonValue
    let user = props.user
    let onClickHandler = props.onClick

    return (
        <>
            <Card style={{ width: 'auto' }}>
                <Card.Body>
                    <Card.Title>{user.name}</Card.Title>
                    <Card.Subtitle className="mb-2 text-muted">
                        {user.mail}
                    </Card.Subtitle>
                    <Card.Text>
                        성별: {user.gender == 'female' ? '여자' : '남자'}
                    </Card.Text>
                    <Card.Text>
                        가입일: {user.created_at}
                    </Card.Text>
                    <Button
                        variant="primary"
                        type="submit"
                        onClick={onClickHandler}
                    >
                        {buttonValue}
                    </Button>
                    {/* <Card.Link href="#">Card Link</Card.Link>
                                <Card.Link href="#">Another Link</Card.Link> */}
                </Card.Body>
            </Card>
        </>
    )
}

export default UserCard
