import React from 'react'

const GetBackURL = () => {
    return process.env.REACT_APP_BACK_BASE_URL
}

const changeTimeFormat = (date) => {
    try {
        const d = date.split('.')[0]
        const array = d.split('T')
        const day = array[0]
        const time = array[1].slice(0, -3)
        // const day = array[0].replace("-", "/")
        // console.log(day)
        return day + '-' + time
    } catch (err) {
        console.log(err)
        return date
    }
}

export { GetBackURL, changeTimeFormat }
