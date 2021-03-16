import React, {useState} from 'react'

const Text = () => {

  const [stuff, setStuff] = useState('say something')
  const [message, setMessage] = useState("")

  const handleOnFocus = () => {
    setStuff("")
  }

  const handleOnBlur = () => {
    setStuff("neat")
  }

  const handleOnChange = (e) => {
    let userTyped = e.target.value
    setMessage(userTyped)
  }

  return (
    <div className="upload">
      <div className="row">
        <div className="col-1">
          <h2>2</h2>
        </div>
        <div className="col-11">
        <h6>your secret message could either be some text</h6>
          <div className="row">
            <div className="col-6">
              <input 
                placeholder={stuff}
                value={message}
                className="inputText"
                type="text"
                onBlur={handleOnBlur}
                onFocus={handleOnFocus}
                onChange={handleOnChange}
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