import React from 'react'
import upload from '../media/upload.png'

const Upload = () => {
  return (
    <div className="upload">
      <h6>add your message with this upload</h6>
      <div className="row">
        <div className="col-1">
          <h2>1</h2>
        </div>
        <div className="col-11">
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

export default Upload