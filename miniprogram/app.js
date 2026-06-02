App({
  globalData: {
    userInfo: null,
    isLogin: false,
    baseUrl: 'https://api.campus-market.example.com'
  },

  onLaunch() {
    // 检查登录状态
    const token = wx.getStorageSync('token')
    if (token) {
      this.globalData.isLogin = true
      // 获取用户信息
      this.getUserInfo()
    }
  },

  getUserInfo() {
    // 从服务器获取用户信息
    wx.request({
      url: `${this.globalData.baseUrl}/api/user/profile`,
      method: 'GET',
      header: {
        'Authorization': `Bearer ${wx.getStorageSync('token')}`
      },
      success: (res) => {
        if (res.statusCode === 200) {
          this.globalData.userInfo = res.data.data
        }
      }
    })
  },

  checkLogin() {
    return new Promise((resolve, reject) => {
      if (this.globalData.isLogin) {
        resolve(true)
      } else {
        wx.navigateTo({
          url: '/pages/profile/profile'
        })
        reject(false)
      }
    })
  }
})