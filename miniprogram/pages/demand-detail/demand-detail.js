Page({
  data: {
    demand: {}
  },

  onLoad(options) {
    const { id } = options
    // TODO: 调用 API 获取需求详情
    console.log('加载需求详情:', id)
  },

  offerService() {
    // TODO: 接单逻辑
    wx.showToast({ title: '接单成功', icon: 'success' })
  }
})