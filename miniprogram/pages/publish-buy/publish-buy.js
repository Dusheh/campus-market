Page({
  data: {
    form: {
      type: 'goods',
      title: '',
      budget: '',
      description: '',
      isUrgent: false
    }
  },

  switchType(e) {
    this.setData({ 'form.type': e.currentTarget.dataset.type })
  },

  onInput(e) {
    const field = e.currentTarget.dataset.field
    this.setData({ [`form.${field}`]: e.detail.value })
  },

  onUrgentChange(e) {
    this.setData({ 'form.isUrgent': e.detail.value })
  },

  submit() {
    const { form } = this.data
    if (!form.title) {
      wx.showToast({ title: '请输入标题', icon: 'none' })
      return
    }
    // TODO: 调用 API 发布需求
    console.log('提交需求:', form)
    wx.showToast({ title: '发布成功', icon: 'success' })
    setTimeout(() => wx.navigateBack(), 1500)
  }
})