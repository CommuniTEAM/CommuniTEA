import React from 'react'
import '../../App.css'
import useResponsiveHeight from '../../Hooks/useResponsiveHeight'
import PouringTea from '../../assets/PouringTea.png'
import PeopleDrinkingTea from '../../assets/PeopleDrinkingTea.jpg'
import FriendsDrinkingTea from '../../assets/FriendsDrinkingTea.jpg'

export default function ValueProp (): JSX.Element {
  const responsiveHeight = useResponsiveHeight() // custom hook

  return (
    <div
      style={{
        backgroundColor: '#FFFFF0',
        display: 'flex',
        height: responsiveHeight
      }}
    >
      <div style={{ width: '50vw' }}>
        <div style={{ transform: 'translate(7vw, 2vw)' }}>
          <img
            src={PouringTea}
            alt="Pouring Tea"
            style={{ width: '20vw', height: 'auto', borderRadius: '50px' }}
          />
        </div>
        <div style={{ transform: 'translate(26vw, -12vw)' }}>
          <img
            src={PeopleDrinkingTea}
            alt="People Drinking Tea"
            style={{ width: '20vw', height: 'auto', borderRadius: '50px' }}
          />
        </div>
        <div style={{ transform: 'translate(15vw, -17vw)' }}>
          <img
            src={FriendsDrinkingTea}
            alt="Friends Drinking Tea"
            style={{ width: '21vw', height: 'auto', borderRadius: '50px' }}
          />
        </div>
      </div>
      <div
        style={{
          width: '50vw',
          display: 'flex',
          flexDirection: 'column'
        }}
      >
        <div style={{ width: '80%' }}>
          <h1
            style={{
              fontFamily: 'Roboto Slab',
              fontSize: '3vw',
              fontWeight: 300,
              color: '#416543'
            }}
          >
            CommuniTEA is your gateway to a vibrant world of tea exploration.
          </h1>
          <p
            style={{
              fontFamily: 'Roboto',
              fontSize: '1vw',
              fontWeight: 300,
              color: '#416543',
              overflowX: 'hidden'
            }}
          >
            Whether you are a seasoned connoisseur or just starting your tea journey, our platform
            brings together the finest tea varieties and the best local businesses, creating a tea
            community like no other.
          </p>
        </div>
      </div>
    </div>
  )
}
