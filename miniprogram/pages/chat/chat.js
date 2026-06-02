Page({
  data: {
    chats: []
  },

  onLoad() {
    this.loadChats()
  },

  onShow() {
    this.loadChats()
  },

  async loadChats() {
    // TODO: 调用 API 获取聊天列表
  },

  goChat(e) {
    const id = e.currentTarget.dataset.id
    // TODO: 进入聊天详情
  }
})