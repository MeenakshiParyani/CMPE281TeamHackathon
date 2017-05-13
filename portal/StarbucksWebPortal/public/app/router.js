router = new VueRouter({
  mode: 'history',
  
  routes: [
    {
      path: '/',
      redirect: '/orders'
    },
    {
      name: 'orders',
      path: '/orders',
      component: Vue.component('rb-orders')
    }
  ]
});