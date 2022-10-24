import Alert from 'react-bootstrap/Alert'

function Title(props) {
    const title = props.title
    return (
        <>
            <Alert variant="primary">
                <Alert.Heading>{title}</Alert.Heading>
                {/* <p>
                    Aww yeah, you successfully read this important alert
                    message. This example text is going to run a bit longer so
                    that you can see how spacing within an alert works with this
                    kind of content.
                </p> */}
            </Alert>
            <hr />
        </>
    )
}

export default Title
