# 部署文档

## 服务器要求

- Ubuntu 20.04 / 22.04
- 2核4G 以上配置
- 已备案域名（用于 HTTPS 和微信小程序 API 白名单）

## 一键部署

### 1. 连接服务器
```bash
ssh username@your-server-ip
```

### 2. 上传初始化脚本
```bash
# 在本地
scp scripts/server_init.sh username@your-server-ip:~

# 在服务器上
sudo bash server_init.sh
```

### 3. 克隆项目
```bash
git clone https://github.com/Dusheh/campus-market.git
cd campus-market
```

### 4. 配置后端
```bash
cd backend
# 修改数据库密码
vim config/config.yaml
# 安装依赖
go mod tidy
# 编译
go build -o server ./cmd/server/
```

### 5. 配置 systemd 服务
```bash
sudo vim /etc/systemd/system/campus-market.service
```

```ini
[Unit]
Description=Campus Market Service
After=network.target mysql.service redis.service

[Service]
Type=simple
User=ubuntu
WorkingDirectory=/home/ubuntu/campus-market/backend
ExecStart=/home/ubuntu/campus-market/backend/server
Restart=always
RestartSec=5

[Install]
WantedBy=multi-user.target
```

```bash
sudo systemctl daemon-reload
sudo systemctl enable campus-market
sudo systemctl start campus-market
```

### 6. 配置 HTTPS（可选）
```bash
sudo apt-get install -y certbot python3-certbot-nginx
sudo certbot --nginx -d api.your-domain.com
```

### 7. 验证
```bash
curl http://localhost:8080/api/services
```

## 微信小程序配置

1. 登录 [微信公众平台](https://mp.weixin.qq.com/)
2. 开发管理 → 开发设置 → 服务器域名
3. 添加 `request合法域名`: `https://api.your-domain.com`
4. 添加 `socket合法域名`: `wss://api.your-domain.com`

## 数据库备份

```bash
# 创建备份脚本
cat > ~/backup.sh <<'EOF'
#!/bin/bash
DATE=$(date +%Y%m%d_%H%M%S)
mysqldump -u campus -p'campus_market_2024' campus_market > ~/backups/campus_market_$DATE.sql
# 保留最近 7 天的备份
find ~/backups -name "*.sql" -mtime +7 -delete
EOF

mkdir -p ~/backups
chmod +x ~/backup.sh

# 添加 crontab 每日凌晨 2 点备份
(crontab -l 2>/dev/null; echo "0 2 * * * ~/backup.sh") | crontab -
```

## 常用命令

```bash
# 查看服务状态
sudo systemctl status campus-market

# 查看日志
sudo journalctl -u campus-market -f

# 重启服务
sudo systemctl restart campus-market

# 查看端口占用
sudo netstat -tlnp | grep 8080
```