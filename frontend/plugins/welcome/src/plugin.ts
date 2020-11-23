import { createPlugin } from '@backstage/core';
import WelcomePage from './components/WelcomePage';
import createBookings from './components/Booking';
import Login from './components/login';
 
export const plugin = createPlugin({
  id: 'welcome',
  register({ router }) {
    router.registerRoute('/', WelcomePage);
    router.registerRoute('/booking', createBookings);
    router.registerRoute('/login', Login);
  },
});
