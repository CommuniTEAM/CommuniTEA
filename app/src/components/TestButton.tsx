import * as React from 'react'
import { Button, Modal, Box, Typography } from '@mui/material'

function MyButton (): React.ReactElement {
  const [open, setOpen] = React.useState(false)

  const handleOpen = (): void => {
    setOpen(true)
  }

  const handleClose = (): void => {
    setOpen(false)
  }

  return (
    <>
      <Button variant="contained" onClick={handleOpen}>
        CLick me!
      </Button>
      <Modal
        open={open}
        onClose={handleClose}
        aria-labelledby="modal-modal-title"
        aria-describedby="modal-modal-description"
      >
        <Box
          sx={{
            position: 'absolute' as 'absolute',
            top: '50%',
            left: '50%',
            transform: 'translate(-50%, -50%)',
            width: 400,
            bgcolor: 'background.paper',
            border: '2px solid #000',
            boxShadow: 24,
            p: 4
          }}
        >
          <Typography id="modal-modal-title" variant="h6" component="h2" sx={{ color: 'green' }}>
            Notification
          </Typography>
          <Typography id="modal-modal-description" sx={{ mt: 2, color: 'black' }}>
            Material UI has successfully been installed!
          </Typography>
        </Box>
      </Modal>
    </>
  )
}

export default MyButton
