
import { Route, Routes } from "react-router-dom";
import Home from "./components/Home"
import NotFound from "./components/NotFound"
function App() {
  return (
    <div>
      <Routes>
        <Route exact path="/" element={<Home/>} />
        <Route exact path="*" element={<NotFound/>}/>
      </Routes>
      
    </div>
  );
}

export default App;
