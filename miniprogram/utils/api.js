const API_BASE = 'https://api.campus-market.example.com'

const request = (url, method = 'GET', data = {}) => {
  const token = wx.getStorageSync('token')
  return new Promise((resolve, reject) => {
    wx.request({
      url: `${API_BASE}${url}`,
      method,
      data,
      header: {
        'Content-Type': 'application/json',
        'Authorization': token ? `Bearer ${token}` : ''
      },
      success(res) {
        if (res.statusCode === 200) {
          resolve(res.data)
        } else if (res.statusCode === 401) {
          wx.removeStorageSync('token')
          wx.navigateTo({ url: '/pages/profile/profile' })
          reject(res)
        } else {
          wx.showToast({ title: res.data.message || '请求失败', icon: 'none' })
          reject(res)
        }
      },
      fail(err) {
        wx.showToast({ title: '网络错误', icon: 'none' })
        reject(err)
      }
    })
  })
}

// 登录
const login = (code) => request('/api/auth/login', 'POST', { code })

// 服务相关
const getServices = (params) => request('/api/services', 'GET', params)
const getServiceDetail = (id) => request(`/api/services/${id}`, 'GET')
const createService = (data) => request('/api/services', 'POST', data)

// 物品相关
const getGoods = (params) => request('/api/goods', 'GET', params)
const getGoodsDetail = (id) => request(`/api/goods/${id}`, 'GET')
const createGoods = (data) => request('/api/goods', 'POST', data)

// 需求相关
const getDemands = (params) => request('/api/demands', 'GET', params)
const getDemandDetail = (id) => request(`/api/demands/${id}`, 'GET')
const createDemand = (data) => request('/api/demands', 'POST', data)

// 订单相关
const getOrders = (params) => request('/api/orders', 'GET', params)
const createOrder = (data) => request('/api/orders', 'POST', data)
const updateOrderStatus = (id, status) => request(`/api/orders/${id}/status`, 'PUT', { status })

// 用户相关
const getUserProfile = () => request('/api/user/profile', 'GET')
const updateUserProfile = (data) => request('/api/user/profile', 'PUT', data)

module.exports = {
  request,
  login,
  getServices,
  getServiceDetail,
  createService,
  getGoods,
  getGoodsDetail,
  createGoods,
  getDemands,
  getDemandDetail,
  createDemand,
  getOrders,
  createOrder,
  updateOrderStatus,
  getUserProfile,
  updateUserProfile
}