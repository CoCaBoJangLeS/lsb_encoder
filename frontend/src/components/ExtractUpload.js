import React from 'react'
import upload from '../media/upload.png'
import { embedRequest } from '../helpers/api'

const ExtractUpload = () => {
  const handleFileUpload = (e) => {
    let file = e.target.files[0]

    let formData = new FormData()
    formData.append('file', file)
    formData.append('message', 'this is a test')
    embedRequest(formData)
  }
  return (
    <div className='upload'>
      <div className='row'>
        <div className='col-1'>
          <h2>1</h2>
        </div>
        <div className='col-11'>
          <h3>Upload an image where your message will be extracted from</h3>
          <div className='uploadBox'>
            <div className='row'>
              <div className='col-6'>
                <img src={upload} alt='upload' width='70px' />
                <p>Drag and Drop or</p>
              </div>
              <div className='col-6 select'>
                <label className='uploadButton'>
                  <input type='file' onChange={handleFileUpload} required />
                  <span>Select From Your Computer</span>
                </label>
                <br />
                <small>please use .jpeg .png or .bmp</small>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  )
}

export default ExtractUpload