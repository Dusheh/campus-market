Page({
  data: {
    keyword: '',
    currentTab: 'all',
    items: []
  },

  onLoad() {
    this.loadItems()
  },

  onShow() {
    this.loadItems()
  },

  onPullDownRefresh() {
    this.loadItems().then(() => wx.stopPullDownRefresh())
  },

  async loadItems() {
    try {
      // TODO: 调用 API 获取服务+物品混合列表
    } catch (err) {
      console.error(err)
    }
  },

  onSearch(e) {
    this.setData({ keyword: e.detail.value })
    this.loadItems()
  },

  switchTab(e) {
    const tab = e.currentTarget.dataset.tab
    this.setData({ currentTab: tab })
    this.loadItems()
  },

  goPublish() {
    wx.navigateTo({ url: '/pages/publish-sell/publish-sell' })
  }
})