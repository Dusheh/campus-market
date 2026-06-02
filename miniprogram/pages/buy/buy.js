Page({
  data: {
    keyword: '',
    currentTab: 'all',
    demands: []
  },

  onLoad() {
    this.loadDemands()
  },

  onShow() {
    this.loadDemands()
  },

  onPullDownRefresh() {
    this.loadDemands().then(() => wx.stopPullDownRefresh())
  },

  async loadDemands() {
    try {
      // TODO: 调用 API 获取需求列表
    } catch (err) {
      console.error(err)
    }
  },

  onSearch(e) {
    this.setData({ keyword: e.detail.value })
    this.loadDemands()
  },

  switchTab(e) {
    const tab = e.currentTarget.dataset.tab
    this.setData({ currentTab: tab })
    this.loadDemands()
  },

  goPublish() {
    wx.navigateTo({ url: '/pages/publish-buy/publish-buy' })
  }
})