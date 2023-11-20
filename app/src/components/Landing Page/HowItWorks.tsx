import React from 'react'
import { Grid } from '@mui/material'
import useResponsiveHeight from '../../Hooks/useResponsiveHeight'
import One from '../../assets/One.svg'
import Two from '../../assets/Two.svg'
import Three from '../../assets/Three.svg'
import '../../App.css'

export default function HowItWorks (): JSX.Element {
  const responsiveHeight = useResponsiveHeight() // custom hook

  return (
    <div style={{ height: responsiveHeight }}>
      <h1
        style={{
          fontFamily: 'Roboto Slab',
          fontSize: '2vw',
          fontWeight: 300,
          color: '#416543',
          textAlign: 'center'
        }}
      >
        How does it work?
      </h1>
      <Grid container spacing={2} justifyContent="center">
        <Grid item xs={12} sm={6} md={4} style={{ textAlign: 'center' }}>
          <img
            src={One}
            alt="Step one"
            style={{ width: '15vw', height: 'auto', paddingTop: '5vh' }}
          />
          <h1
            style={{
              fontFamily: 'Roboto Slab',
              fontSize: '1.5vw',
              fontWeight: 300,
              color: '#416543'
            }}
          >
            Discover
          </h1>
          <p style={{ fontFamily: 'Roboto', fontSize: '1vw', fontWeight: 300, color: '#416543' }}>
            Find your favorite teas from curated local businesses
          </p>
        </Grid>
        <Grid item xs={12} sm={6} md={4} style={{ textAlign: 'center' }}>
          <img
            src={Two}
            alt="Step two"
            style={{ width: '15vw', height: 'auto', paddingTop: '5vh' }}
          />
          <h1
            style={{
              fontFamily: 'Roboto Slab',
              fontSize: '1.5vw',
              fontWeight: 300,
              color: '#416543'
            }}
          >
            Connect
          </h1>
          <p style={{ fontFamily: 'Roboto', fontSize: '1vw', fontWeight: 300, color: '#416543' }}>
            Explore tea-related events and connect enthusiasts.
          </p>
        </Grid>
        <Grid item xs={12} sm={6} md={4} style={{ textAlign: 'center' }}>
          <img
            src={Three}
            alt="Step three"
            style={{ width: '15vw', height: 'auto', paddingTop: '5vh' }}
          />
          <h1
            style={{
              fontFamily: 'Roboto Slab',
              fontSize: '1.5vw',
              fontWeight: 300,
              color: '#416543'
            }}
          >
            Experience
          </h1>
          <p style={{ fontFamily: 'Roboto', fontSize: '1vw', fontWeight: 300, color: '#416543' }}>
            Transform your tea-ppreciation into a collective experience with CommuniTEA
          </p>
        </Grid>
      </Grid>
    </div>
  )
}
