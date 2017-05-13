Vue.component('rb-orders-list', {
  template: `
    <div>
      <div v-for="order in orders" class="card my-3">
        <div class="card-block pb-0">
          <label>Order Status: {{ order.status }}</label>
          <button type="button" class="close" @click="removeOrder(order._id)">
            <span aria-hidden="true">&times;</span>
          </button>
        </div>

        <rb-order-item-list class="my-3 list-group-flush" :items="order.items" :showRemove="false"></rb-order-item-list>
      </div>
    </div>
  `,

  props: {
    orders: Array
  },

  methods: {
    removeOrder: function(orderId) {
      this.$emit('remove', orderId);
    }
  }
});