Vue.use(VueMaterial.default)
Vue.prototype.$http = axios

var vm = new Vue({
  el: '#app',
  components: {
    'app-devices': appDevices,
    'app-device-passwords': appDevicePasswords,
    'auto-search': autosearch,
    'manual-search': manualsearch,
    'surv': surv,
  }
})