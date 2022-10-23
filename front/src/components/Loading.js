import Spinner from 'react-bootstrap/Spinner'

function Loading() {
    return (
        <>
            <br />
            <span>
                <h3>Loading...</h3>
            </span>

            <Spinner animation="border" size="xxl"></Spinner>
            <br />
        </>
    )
}

export default Loading
