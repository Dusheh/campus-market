Page({
  data: {
    publishType: 'service', // 'service' | 'goods'
    serviceCategories: ['代取快递', '代课代签', '家教辅导', '维修维护', '摄影设计', '跑腿代办', '其他服务'],
    goodsCategories: ['教材书籍', '电子产品', '生活用品', '运动器材', '服饰鞋包', '其他物品'],
    form: {
      title: '',
      category: '',
      price: '',
      description: '',
      images: []
    }
  },

  onLoad(options) {
    if (options.type) {
      this.setData({ publishType: options.type })
    }
  },

  switchType(e) {
    this.setData({
      publishType: e.currentTarget.dataset.type,
      form: { title: '', category: '', price: '', description: '', images: [] }
    })
  },

  onInput(e) {
    const field = e.currentTarget.dataset.field
    this.setData({ [`form.${field}`]: e.detail.value })
  },

  onCategoryChange(e) {
    this.setData({ 'form.category': e.detail.value })
  },

  chooseImage() {
    wx.chooseImage({
      count: 9 - this.data.form.images.length,
      success: (res) => {
        const images = this.data.form.images.concat(res.tempFilePaths)
        this.setData({ 'form.images': images })
      }
    })
  },

  submit() {
    const { form, publishType } = this.data
    if (!form.title) {
      wx.showToast({ title: '请输入标题', icon: 'none' })
      return
    }
    if (!form.price) {
      wx.showToast({ title: '请输入价格', icon: 'none' })
      return
    }
    // TODO: 调用 API 发布
    console.log('提交数据:', { ...form, type: publishType })
    wx.showToast({ title: '发布成功', icon: 'success' })
    setTimeout(() => wx.navigateBack(), 1500)
  }
})