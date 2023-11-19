import type { ReactElement } from 'react'
import NavBar from './components/Landing Page/Navbar'
import './App.css'
import HeroBanner from './components/Landing Page/HeroBanner'
import ValueProp from './components/Landing Page/ValueProp'
import FeaturedTeas from './components/Landing Page/FeaturedTeas'

function App (): ReactElement {
  return (
    <div>
      <NavBar />
      <HeroBanner />
      <ValueProp />
      <FeaturedTeas />
    </div>
  )
}

export default App
