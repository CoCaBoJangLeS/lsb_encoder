import React from 'react'

const Text = () => {
  return (
    <div className="upload">
      <h6>your secret message could either be some text</h6>
      <div className="row">
        <div className="col-1">
          <h2>2</h2>
        </div>
        <div className="col-11 text">
          <div className="row">
            <div className="col-6">
              <input 
                value="put some cool text here"
                className="inputText"
                type="text"

                >
              </input>
            </div> 
          </div>
        </div>
      </div>
    </div>

  )
}

export default Text