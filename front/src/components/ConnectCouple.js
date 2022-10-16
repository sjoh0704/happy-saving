const ConnectCouple = (props) => {
    const userId = localStorage.getItem('userid')
    const coupleInfo = props.coupleInfo

    console.log(111, coupleInfo)
    if (coupleInfo == null) {
        return (
            <div>
                당신은 아직 couple 등록이 안되어 있습니다! couple 신청을
                해야합니다.
                <p>
                    <a href="/couple/make">couple 신청</a>
                </p>
            </div>
        )
    } else if (coupleInfo.send_id == userId) {
        if (coupleInfo.phase == 'awaiting') {
            return (
                <div>
                    <p>
                        {coupleInfo.receiver.name} 님의 수락을 기다리고
                        있습니다. 기다려 주세요.
                    </p>
                </div>
            )
        } else if (coupleInfo.phase == 'denyed') {
            return <div>거절됨</div>
        } else {
            return <div>승인됨</div>
        }
    } else if (coupleInfo.recv_id == userId) {
        return <div>현재 커플 요청이 들어와 있습니다. 요청을 확인해 주세요</div>
    } else {
        return <div>error</div>
    }
}

export default ConnectCouple
