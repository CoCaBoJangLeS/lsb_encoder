import '../App.css'

import before_after from '../media/before_after.png'

function About() {
  return (
    <div className='beforeAfter'>
      <h1>Can your eyes see the difference?</h1>
      <img src={before_after} alt='Before/After' width='100%' />
    </div>
  )
}

export default About
