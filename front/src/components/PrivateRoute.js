import { React, useState, useEffect, useCallback } from 'react'
import { Navigate, useNavigate } from 'react-router'
import axios from 'axios'

// login 하지 않았거나 token이 만료되었다면 login으로 리다이렉션
const PrivateRouter = () => {
    const mail = localStorage.getItem('mail')
    const navigate = useNavigate()

    const checkLogined = async (mail) => {
        if (
            window.location.pathname == '/login' ||
            window.location.pathname == '/signup'
        ) {
            return
        }

        if (mail == null) {
            console.log('user session expired')
            navigate('/login')
        }
    }

    const checkCookie = async (mail) => {
        if (
            window.location.pathname == '/login' ||
            window.location.pathname == '/signup' ||
            mail == null
        ) {
            return
        }

        try {
            await axios.get('/apis/v1/users?mail=' + mail)
        } catch (err) {
            if (err instanceof axios.AxiosError) {
                if (err.response.status == 403) {
                    console.log('user cookie expired')
                    navigate('/login')
                }
            } else {
                console.log(err)
            }
        }
    }

    useEffect(() => {
        checkLogined(mail)
        checkCookie(mail)
    }, [mail])
}

export default PrivateRouter
