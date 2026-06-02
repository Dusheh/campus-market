Page({
  data: {
    service: {}
  },

  onLoad(options) {
    const { id } = options
    // TODO: 调用 API 获取服务详情
    console.log('加载服务详情:', id)
  },

  buyService() {
    wx.navigateTo({ url: '/pages/order/order' })
  }
})