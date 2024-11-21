import React from 'react'

function Fetch() {

    const endpoint = "https://pokeapi.co/api/v2/pokemon/ditto";

    fetch(endpoint)

  return (
    <div>
        <h1>Fetch</h1>
    </div>
  )
}

export default Fetch