import { useMediaQuery } from '@mui/material';

const useResponsiveHeight = (): string => {
  const extraLargeScreen = useMediaQuery('(min-width:2300px)');
  const largeScreen = useMediaQuery('(min-width:1500px)');
  const mediumScreen = useMediaQuery('(min-width:1000px)');
  const smallScreen = useMediaQuery('(min-width:400px)');

  if (extraLargeScreen) return '60vh';
  if (largeScreen) return '50vh';
  if (mediumScreen) return '40vh';
  if (smallScreen) return '30vw';

  return 'auto';
};

export default useResponsiveHeight;
