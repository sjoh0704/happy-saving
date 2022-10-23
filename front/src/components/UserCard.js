import React from 'react'
import { Button, Card } from 'react-bootstrap'
import { changeTimeFormat } from '../util/util'

const UserCard = (props) => {
    let user = props.user
    let buttonValueLeft = props.buttonValueLeft
    let onClickHandlerLeft = props.onClickLeft
    let buttonValueRight = props.buttonValueRight
    let onClickHandlerRight = props.onClickRight

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
                    <Card.Text>가입일: {changeTimeFormat(user.created_at)}</Card.Text>
                    <Button
                        variant="primary"
                        type="submit"
                        onClick={onClickHandlerLeft}
                    >
                        {buttonValueLeft}
                    </Button>{' '}
                    {buttonValueRight == null ? (
                        <></>
                    ) : (
                        <Button
                            variant="primary"
                            type="submit"
                            onClick={onClickHandlerRight}
                        >
                            {buttonValueRight}
                        </Button>
                    )}
                    {/* <Card.Link href="#">Card Link</Card.Link>
                                <Card.Link href="#">Another Link</Card.Link> */}
                </Card.Body>
            </Card>
        </>
    )
}

export default UserCard
