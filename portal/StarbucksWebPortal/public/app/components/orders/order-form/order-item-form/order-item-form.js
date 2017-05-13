Vue.component('rb-order-item-form', {
  template: `
    <form @submit.prevent="onSubmit">
      <div class="d-flex flex-wrap justify-content-around drinks-list my-3">
        <div v-for="drink in drinks" class="drink text-center px-3 py-1" :class="{ active: item.name === drink.name }"
          @click="onClickDrink(drink)">
          <img :src="drink.image" class="img-thumbnail img-fluid">
          <span class="d-block">{{ drink.name }}</span>
        </div>
      </div>

      <div>
        <div class="form-group">
          <label>Quantity</label>  
          <div>
            <input type="number" class="form-control" v-model.number="item.qty">
          </div>
        </div>

        <div class="form-group">
          <label>Milk</label>
          <div class="btn-group">
            <button v-for="option in milkOptions" type="button" class="btn"
              :class="{ active: item.milk === option, 'btn-secondary': item.milk !== option, 'btn-success': item.milk === option }"
              @click="item.milk = option">
              {{ option }}
            </button>
          </div>
        </div>

        <div class="form-group">
          <label>Size</label>
          <div class="d-flex justify-content-around align-items-end">
            <a href @click.prevent="item.size = 'Large'"><span class="fa fa-coffee fa-4x" :class="{ active: item.size === 'Large' }"></span></a>
            <a href @click.prevent="item.size = 'Medium'"><span class="fa fa-coffee fa-3x" :class="{ active: item.size === 'Medium' }"></span></a>
            <a href @click.prevent="item.size = 'Small'"><span class="fa fa-coffee fa-2x" :class="{ active: item.size === 'Small' }"></span></a>
          </div>
        </div>
      </div>

      <button type="submit" class="btn btn-block btn-success">Add Item</button>
    </form>
  `,

  data: function() {
    return {
      item: {
        name: '',
        qty: 1,
        milk: '',
        size: ''
      }
    };
  },

  created: function() {
    this.drinks = [
      { name: 'Chai Latte', image: 'https://www.starbucks.com/assets/00e28c23d66a4e61897fb50e927d5309.jpg' },
      { name: 'Green Tea Latte', image: 'https://www.starbucks.com/assets/0cc56b6221984cd8b578566c2f0473eb.jpg' },
      { name: 'Iced Chai Latte', image: 'https://www.starbucks.com/assets/5a8b490c1f3b4acca697521ff523023e.jpg' },
      { name: 'Iced Green Tea Latte', image: 'https://www.starbucks.com/assets/3ff72860b3464466a6d2795b76d9c4bb.jpg' },
      { name: 'Shaken Sweet Tea', image: 'https://www.starbucks.com/assets/8bffaf8cd08743b6add5b20cb1f3eb2b.jpg' },
      { name: 'Iced Coffee', image: 'https://www.starbucks.com/assets/a70ad13b5afa4f15b3ceaf8c05c6d880.jpg' },
      { name: 'Iced Coffee with Milk', image: 'https://www.starbucks.com/assets/d7afa1b6cdae4d7dbbb14a29efb90999.jpg' },
      { name: 'Caffe Americano', image: 'https://www.starbucks.com/assets/bd5d1810c0d44ad6b2646e5314975adf.jpg' }
    ];

    this.milkOptions = ['Whole Milk', 'Half and Half', 'Organic', 'No Milk'];
  },

  methods: {
    onClickDrink: function(drink) {
      this.item.name = drink.name;
    },

    onSubmit: function() {
      this.$emit('submit', this.item);
      this.item = {
        name: '',
        qty: 0,
        milk: '',
        size: ''
      };
    }
  }
})