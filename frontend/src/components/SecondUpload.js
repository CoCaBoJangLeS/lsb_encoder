import React from 'react'
import upload from '../media/upload.png'

const SecondUpload = () => {
  return (
    <div className="upload">
      <div className="row">
        <div className="col-1">
          <h2>3</h2>
        </div>
        <div className="col-11">
        <h6>or another upload</h6>
          <div className="uploadBox">
            <div className="row">
            <div className="col-6">
            <img 
              src ={upload}
              alt="upload"
              width='70px' 
            />
            <p>Drag and Drop or</p>
            </div> 
            <div className="col-6 select">
              <button>Select from your computer</button><br />
              <small>please use .jpeg .png or .bmp</small>
            </div>
            </div>
          </div>
        </div>
      </div>
    </div>

  )
}

export default SecondUpload