import './App.css'
import Upload from './components/Upload'
import Text from './components/Text'
import SecondUpload from './components/SecondUpload'
import Dropdown from './components/Dropdown'
import Button from './components/Button'

function App() {
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

export default App
