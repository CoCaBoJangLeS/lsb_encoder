import '../App.css'
import Upload from './Upload'
import Text from './Text'
import SecondUpload from './SecondUpload'
import Dropdown from './Dropdown'
import Button from './Button'

function Embed() {
  return (
    <div className="App">
      <h1>Embed a Message</h1>
      <Upload />
      <Text />
      <SecondUpload />
      <Dropdown />
      <Button />
    </div>
  )
}

export default Embed
