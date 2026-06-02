Page({
  data: {
    currentTab: 'all',
    orders: []
  },

  onLoad() {
    this.loadOrders()
  },

  async loadOrders() {
    // TODO: 调用 API 获取订㕵列表
  },

  switchTab(e) {
    this.setData({ currentTab: e.currentTarget.dataset.tab })
    this.loadOrders()
  }
})