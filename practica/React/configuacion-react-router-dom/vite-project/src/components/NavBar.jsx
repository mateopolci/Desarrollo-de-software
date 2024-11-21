import React from "react";
import { Link } from "react-router-dom";

function NavBar() {
    return (
        <div>
            <nav>
                <ul>
                    <li>
                        <Link to="/BotonesLocos">Botones Locos</Link>
                    </li>
                    <li>
                        <Link to="/">Inicio</Link>
                    </li>
                </ul>
            </nav>
        </div>
    );
}

export default NavBar;
