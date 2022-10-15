
import { Route, Routes } from "react-router-dom";
import Home from "./pages/Home"
import Login from "./pages/Login"
import NotFound from "./pages/NotFound"
import Page from "./components/Page";
import SignUp from "./pages/SignUp";
import PrivateRouter from "./components/PrivateRoute";
import MakeCouple from "./pages/MakeCouple";

function App() {
  return (
    <>
      <Page/>
      <PrivateRouter/>
      <Routes>
        <Route exact path="/" element={<Home/>} />
        <Route exact path="/login" element={<Login/>}/>
        <Route exact path="/signup" element={<SignUp/>}/>
        <Route exact path="/couple/make" element={<MakeCouple/>}/>
        <Route exact path="*" element={<NotFound/>}/>
      </Routes>
      
    </>
  );
}

export default App;
