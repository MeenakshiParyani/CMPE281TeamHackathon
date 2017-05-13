Vue.component('rb-order-item-list', {
  template: `
    <ul class="list-group">
      <li v-for="(item, i) in items" class="list-group-item list-group-item-action d-flex-inline justify-content-between">
        <span>{{ item.qty }} {{ item.size }} {{ item.name }} with {{ item.milk }}</span>
        <button v-if="showRemove" type="button" class="close" @click="remove(i)">
          <span>&times;</span>
        </button>
      </li>
    </ul>
  `,

  props: {
    items: Array,
    showRemove: {
      type: Boolean,
      default: true
    }
  },

  methods: {
    remove: function(index) {
      this.$emit('remove', index);
    }
  }
});