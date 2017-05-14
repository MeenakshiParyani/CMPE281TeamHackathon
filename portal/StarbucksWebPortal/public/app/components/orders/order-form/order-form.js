Vue.component('rb-order-form', {
  template: `
    <form @submit.prevent="placeOrder">
      <div class="my-5 mx-3">
        <div>
          <h4 class="text-center mb-3">Place an order</h4>
          
          <div v-if="state === states.ADD_ITEM">
            <div>
              <h5><span class="badge badge-pill badge-default">1</span> Add an item</h5>
              <rb-order-item-form @submit="addItem"></rb-order-item-form>
            </div>
          </div>

          <div v-else-if="state === states.REVIEW_ORDER">
            <div>
              <h5><span class="badge badge-pill badge-default">2</span> Review order</h5>
              <rb-order-item-list class="my-3" :items="items" @remove="removeItem"></rb-order-item-list>
              <button @click="state = states.ADD_ITEM" class="btn btn-success">
                <span class="fa fa-plus"></span> Add more items
              </button>

              <hr>

              <div>
                <label class="radio-inline">
                  <input type="radio" value="in" checked="checked" v-model="orderMode">
                  In Store
                </label> 

                <label class="radio-inline">
                  <input type="radio" value="out" v-model="orderMode">
                  Take Out
                </label>
              </div>

              <div class="form-group">
                <label>Name</label>
                <input type="text" class="form-control" placeholder="Please enter your name" v-model="customerName">
              </div>

              <button type="submit" class="btn btn-block btn-success">Place Order</button>
            </div>
          </div>
        </div>
      </div>
    </form>
  `,

  data: function() {
    return {
      items: [],
      orderMode: '',
      customerName: '',

      state: null
    };
  },

  created: function() {
    this.states = {
      ADD_ITEM: 0,
      REVIEW_ORDER: 1
    };
    this.state = this.states.ADD_ITEM;
  },

  methods: {
    addItem: function(item) {
      this.items.push(item);
      this.state = this.states.REVIEW_ORDER;
    },

    removeItem: function(index) {
      this.items.splice(index, 1);
      if (this.items.length === 0) {
        this.state = this.states.ADD_ITEM;
      }
    },

    placeOrder: function() {
      var order = {
        customerName: this.customerName,
        location: this.orderMode,
        items: this.items
      };

      this.$emit("submit", order);
    }
  }
});