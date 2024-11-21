import React from "react";
import { useState } from "react";

function BotonesLocos() {

	const [count, setCount] = useState(0)

    return (
        <>
            <div className="card">
                <div className="card">{count}</div>
                <button className="button" onClick={() => setCount(count + 1)}>
                    Sumar
                </button>
                <button className="button" onClick={() => setCount(count - 1)}>
                    Restar
                </button>
            </div>
        </>
    );
}

export default BotonesLocos;
