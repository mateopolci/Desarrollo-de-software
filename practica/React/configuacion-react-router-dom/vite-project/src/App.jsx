import "./App.css";
import { BrowserRouter, Routes, Route } from "react-router-dom";
import BotonesLocos from "./components/BotonesLocos";
import NavBar from "./components/NavBar";

function App() {
    return (
        <>
            <BrowserRouter>
                <NavBar />
                <Routes>
                    <Route path="/BotonesLocos" element={<BotonesLocos />}/>
                </Routes>
            </BrowserRouter>
        </>
    );
}

export default App;
