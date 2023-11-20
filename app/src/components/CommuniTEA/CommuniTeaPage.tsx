import Footer from '../Landing Page/Footer'
import NavBar from '../Landing Page/Navbar'
import Filters from './Filters'
import LocationSearch from './LocationSearch'

export default function CommuniTeaPage (): JSX.Element {
  return (
    <div>
      <NavBar />
      <LocationSearch />
      <Filters />
      <Footer />
    </div>
  )
}
