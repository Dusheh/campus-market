Page({
  data: {
    isLogin: false,
    userInfo: {}
  },

  onLoad() {
    const token = wx.getStorageSync('token')
    if (token) {
      this.setData({ isLogin: true })
      this.loadUserInfo()
    }
  },

  async loadUserInfo() {
    // TODO: 调用 API 获取用户信息
  },

  onLogin(e) {
    // TODO: 微信登录逻辑
    wx.login({
      success: (res) => {
        if (res.code) {
          // 将 code 发送到后端换取 token
          console.log('登录 code:', res.code)
          this.setData({ isLogin: true })
        }
      }
    })
  },

  goPage(e) {
    const page = e.currentTarget.dataset.page
    wx.navigateTo({ url: `/pages/${page}/${page}` })
  }
})