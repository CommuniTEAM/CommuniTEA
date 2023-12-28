import { useTheme } from '@mui/material/styles';
import useMediaQuery from '@mui/material/useMediaQuery';

const useResponsiveHeight = (): string => {
  const theme = useTheme();

  // Define media queries at the top level
  const extraLargeScreen = useMediaQuery(theme.breakpoints.up('xl'));
  const largeScreen = useMediaQuery(theme.breakpoints.up('lg'));
  const mediumScreen = useMediaQuery(theme.breakpoints.up('md'));
  const smallScreen = useMediaQuery(theme.breakpoints.up('sm'));

  if (extraLargeScreen) return '60vh';
  if (largeScreen) return '50vh';
  if (mediumScreen) return '40vh';
  if (smallScreen) return '35vh';

  return '30vh'; // Fallback for extra small screens
};

export default useResponsiveHeight;
