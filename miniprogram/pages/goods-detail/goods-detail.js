Page({
  data: {
    goods: {}
  },

  onLoad(options) {
    const { id } = options
    // TODO: 调用 API 获取物品详情
    console.log('加载物品详情:', id)
  },

  buyGoods() {
    wx.navigateTo({ url: '/pages/order/order' })
  }
})