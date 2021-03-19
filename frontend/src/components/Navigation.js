import React from 'react'
import { BrowserRouter as Router, Switch, Route, Link } from 'react-router-dom'
import Extract from './Extract'
import About from './About'
import Embed from './Embed'

const Navigation = () => {
  return (
    <Router>
      <>
        <nav>
          <ul>
            <li>
              <Link to='/about'>About</Link>
            </li>
            <li>
              <Link to='/extract'>Extract</Link>
            </li>
            <li>
              <Link to='/embed'>Embed</Link>
            </li>
          </ul>
        </nav>
        <Switch>
          <Route path='/extract'>
            {/* <h1>this is the contact page</h1> */}
            <Extract />
          </Route>
          <Route path='/about'>
            {/* <h1>this is the about page</h1> */}
            <About />
          </Route>
          <Route path='/embed'>
            {/* <h1>this is the embed page</h1> */}
            <Embed />
          </Route>
        </Switch>
      </>
    </Router>
  )
}

export default Navigation
