import { Route, Routes } from 'react-router-dom'
import Home from './pages/Home'
import Login from './pages/Login'
import NotFound from './pages/NotFound'
import Page from './components/NavigationBar'
import SignUp from './pages/SignUp'
import PrivateRouter from './components/PrivateRoute'
import MakeCouple from './pages/MakeCouple'
import Post from './pages/Post'
import NavigationBar from './components/NavigationBar'
import CreatePost from './pages/CreatePost'

function App() {
    return (
        <>
            <NavigationBar />
            <PrivateRouter />
            <Routes>
                <Route exact path="/" element={<Home />} />
                <Route exact path="/login" element={<Login />} />
                <Route path="/post/:id" element={<Post />} />
                <Route path="/post/create" element={<CreatePost />} />
                <Route exact path="/signup" element={<SignUp />} />
                <Route exact path="/couple/make" element={<MakeCouple />} />
                <Route exact path="/notfound" element={<NotFound />} />
                <Route exact path="*" element={<NotFound />} />
            </Routes>
        </>
    )
}

export default App
