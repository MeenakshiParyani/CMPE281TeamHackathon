Vue.component('rb-orders', {
  template: `
    <div class="card mx-auto mt-5" id="orders">
      <div class="card-block">
        <div class="card-title text-center pb-3">
          <img src="https://www.starbucks.com/static/images/global/logo.svg">
          <h4 class="my-3">{{ location ? location.name : '' }}</h4>
        </div>

        <hr>

        <div class="my-5 mx-3">
          <div v-if="!location">
            <h4 class="text-center mb-3">Select a location</h4>
            <select class="form-control" v-model="location">
              <option v-for="location in validLocations" :value="location">{{ location.name }}</option>
            </select>
          </div>
          
          <div v-else-if="state === states.SHOW_ORDERS">
            <a href @click.prevent="state = states.ADD_ORDER">Place an order</a>

            <rb-orders-list class="mt-3" :orders="orders" @remove="removeOrder"></rb-orders-list>
          </div>

          <rb-order-form v-else-if="state === states.ADD_ORDER" @submit="placeOrder"></rb-order-form>
        </div>
      </div>
    </div>
  `,

  data: function() {
    return {
      orders: [],
      location: null,
      state: null
    };
  },

  created: function() {
    this.states = {
      ADD_ORDER: 0,
      SHOW_ORDERS: 1
    };
    this.state = this.states.SHOW_ORDERS;

    this.validLocations = [
      { code: 'fr', name: 'Fremont' },
      { code: 'sj', name: 'San Jose' },
      { code: 'sv', name: 'Sunnyvale' }
    ];
  },

  watch: {
    location: function(val) {
      var self = this;
      
      if (val) {
        this.getOrders().done(function() {
          if (self.orders.length === 0) self.state = self.states.ADD_ORDER;
        });
      }
    }
  },

  methods: {
    getOrders: function() {
      var self = this;

      return jQuery.get('/api/orders?location=' + this.location.code, function(res) {
        self.orders = res.orders || res;
      });
    },

    placeOrder: function(order) {
      var self = this;

      jQuery.ajax({
        type: 'POST',
        url: '/api/order?location=' + this.location.code,
        contentType: 'application/json; charset=utf-8',
        data: JSON.stringify(order),

        success: function() {
          self.getOrders();
          self.state = self.states.SHOW_ORDERS;
        }
      });
    },

    removeOrder: function(orderId) {
      var self = this;

      jQuery.ajax({
        type: 'DELETE',
        url: '/api/order/' + orderId + '?location=' + this.location.code,

        success: function() {
          self.getOrders();
        }
      });
    }
  }
})