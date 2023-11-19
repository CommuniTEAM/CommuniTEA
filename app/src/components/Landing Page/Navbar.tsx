import React, { useState } from 'react'
import {
  AppBar,
  Box,
  Toolbar,
  IconButton,
  Typography,
  Menu,
  Avatar,
  Button,
  Tooltip,
  MenuItem,
  Container,
  Drawer,
  List,
  ListItem,
  ListItemText
} from '@mui/material'
import MenuIcon from '@mui/icons-material/Menu'
import { useNavigate } from 'react-router-dom'
import CommuniteaLogo from '../../assets/CommuniteaLogo.svg'
import '../../App.css'

function NavBar (): JSX.Element {
  const navigate = useNavigate()
  const [anchorElUser, setAnchorElUser] = useState<null | HTMLElement>(null)
  const [drawerOpen, setDrawerOpen] = useState(false)
  const [settingsDrawerOpen, setSettingsDrawerOpen] = useState(false)

  // Placeholder navigation pages and settings
  const pages = ['About Us', 'WikiTEAdia', 'CommuniTEA']
  const settings = ['Profile', 'Account', 'Dashboard', 'Logout']

  // Handlers
  const handleHomeNavigation = (): void => {
    navigate('/')
  }

  const handleOpenUserMenu = (event: React.MouseEvent<HTMLElement>): void => {
    setAnchorElUser(event.currentTarget)
  }

  const handleCloseUserMenu = (): void => {
    setAnchorElUser(null)
  }

  // Toggles the state of the drawer
  const toggleDrawer = (open: boolean) => (event: React.KeyboardEvent | React.MouseEvent) => {
    if (
      event.type === 'keydown' &&
      ((event as React.KeyboardEvent).key === 'Tab' ||
        (event as React.KeyboardEvent).key === 'Shift')
    ) {
      return
    }
    setDrawerOpen(open)
  }

  const toggleSettingsDrawer =
    (open: boolean) => (event: React.KeyboardEvent | React.MouseEvent) => {
      if (
        event.type === 'keydown' &&
        ((event as React.KeyboardEvent).key === 'Tab' ||
          (event as React.KeyboardEvent).key === 'Shift')
      ) {
        return
      }
      setSettingsDrawerOpen(open)
    }

  return (
    <AppBar position="static" sx={{ backgroundColor: '#FFF5E1' }}>
      <Container maxWidth="xl">
        <Toolbar>
          {/* Logo and title */}
          <Box
            onClick={handleHomeNavigation}
            sx={{
              display: 'flex',
              alignItems: 'center',
              marginRight: 'auto'
            }}
          >
            <img
              src={CommuniteaLogo}
              alt="Communitea Logo"
              style={{ width: 50, marginRight: 16 }}
            />
            <Typography
              variant="h6"
              noWrap
              component="a"
              href="#"
              sx={{
                fontFamily: 'Montserrat',
                fontWeight: 700,
                letterSpacing: '.3rem',
                color: 'inherit',
                textDecoration: 'none',
                display: { xs: 'none', md: 'flex', color: 'black' }
              }}
            >
              COMMUNITEA
            </Typography>
          </Box>

          {/* Navigation pages for large screens */}
          <Box sx={{ flexGrow: 1, display: { xs: 'none', md: 'flex' }, justifyContent: 'center' }}>
            {pages.map((page) => (
              <Button
                key={page}
                sx={{ mx: 5, color: 'black', display: 'block', fontFamily: 'Montserrat' }}
              >
                {page}
              </Button>
            ))}
          </Box>

          {/* Hamburger menu for small screens */}
          <Box sx={{ flexGrow: 1, display: { xs: 'flex', md: 'none' } }}>
            <IconButton
              size="large"
              edge="start"
              aria-label="menu"
              sx={{ mr: 2, display: { md: 'none' } }}
              onClick={toggleDrawer(true)}
            >
              <MenuIcon />
            </IconButton>
            <Drawer anchor="left" open={drawerOpen} onClose={toggleDrawer(false)}>
              <Box
                sx={{ width: 250 }}
                role="presentation"
                onClick={toggleDrawer(false)}
                onKeyDown={toggleDrawer(false)}
              >
                <List>
                  {pages.map((text) => (
                    <ListItem key={text}>
                      <ListItemText primary={text} />
                    </ListItem>
                  ))}
                </List>
              </Box>
            </Drawer>
          </Box>

          {/* User settings menu */}
          <Box sx={{ flexGrow: 0, marginLeft: 16 }}>
            <Tooltip title="Open settings">
              <IconButton onClick={handleOpenUserMenu} sx={{ p: 0 }}>
                <Avatar alt="User Avatar" src="/static/images/avatar/2.jpg" />
              </IconButton>
            </Tooltip>
            <Menu
              id="menu-appbar"
              anchorEl={anchorElUser}
              anchorOrigin={{ vertical: 'bottom', horizontal: 'right' }}
              keepMounted
              transformOrigin={{ vertical: 'top', horizontal: 'right' }}
              open={Boolean(anchorElUser)}
              onClose={handleCloseUserMenu}
            >
              {settings.map((setting) => (
                <MenuItem key={setting} onClick={handleCloseUserMenu}>
                  <Typography sx={{ textAlign: 'center', fontFamily: 'Montserrat' }}>
                    {setting}
                  </Typography>
                </MenuItem>
              ))}
            </Menu>
          </Box>

          {/* Settings drawer for small screens */}
          <Box sx={{ flexGrow: 1, display: { xs: 'flex', md: 'none' } }}>
            <Drawer anchor="right" open={settingsDrawerOpen} onClose={toggleSettingsDrawer(false)}>
              <Box
                sx={{ width: 250 }}
                role="presentation"
                onClick={toggleSettingsDrawer(false)}
                onKeyDown={toggleSettingsDrawer(false)}
              >
                <List>
                  {settings.map((text) => (
                    <ListItem key={text}>
                      <ListItemText primary={text} />
                    </ListItem>
                  ))}
                </List>
              </Box>
            </Drawer>
          </Box>
        </Toolbar>
      </Container>
    </AppBar>
  )
}

export default NavBar
