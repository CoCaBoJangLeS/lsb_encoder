import React from 'react'

const Dropdown = () => {
  return (
    <div className="upload">
      <div className="row">
        <div className="col-1">
          <h2 className="four">4</h2>
        </div>
        <div className="col-11">
        <h6>encode in any of the following</h6>
          <div className="row">
            <div className="col-6">
            <select name="encoder" id="encoder">
              <option value="rot13">Rot13</option>
              <option value="base16">Base16</option>
              <option value="base32">Base32</option>
              <option value="base64">Base64</option>
              <option value="base85">Base85</option>
            </select>
            </div> 
          </div>
        </div>
      </div>
    </div>

  )
}

export default Dropdown