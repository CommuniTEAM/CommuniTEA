import type { ReactElement } from 'react'
import TestButton from './components/TestButton'
import NavBar from './components/Landing Page/Navbar'
import './App.css'

function App (): ReactElement {
  return (
    <>
      <div>
        <NavBar />
      </div>
      <div>
        Hello CommuniTEAM Members! Welcome to TypeScript :)
        <br />
        <TestButton />
      </div>
    </>
  )
}

export default App
