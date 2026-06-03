#!/bin/bash
# ============================================================
# 校园服务买卖平台 - 服务器初始化脚本
# 适用于 Ubuntu 20.04 / 22.04
# 用法: sudo bash server_init.sh
# ============================================================

set -e

GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m'

echo -e "${GREEN}========================================${NC}"
echo -e "${GREEN}  校园服务买卖平台 - 服务器初始化${NC}"
echo -e "${GREEN}========================================${NC}"

# 1. 更新系统
echo -e "${YELLOW}[1/8] 更新系统包...${NC}"
apt-get update -y && apt-get upgrade -y

# 2. 安装基础工具
echo -e "${YELLOW}[2/8] 安装基础工具...${NC}"
apt-get install -y curl wget git vim build-essential nginx

# 3. 安装 Go 1.22
echo -e "${YELLOW}[3/8] 安装 Go 1.22...${NC}"
if ! command -v go &> /dev/null; then
    wget -q https://go.dev/dl/go1.22.2.linux-amd64.tar.gz -O /tmp/go.tar.gz
    tar -C /usr/local -xzf /tmp/go.tar.gz
    echo 'export PATH=$PATH:/usr/local/go/bin' >> /etc/profile
    echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
    export PATH=$PATH:/usr/local/go/bin
    rm /tmp/go.tar.gz
fi
echo "Go version: $(go version)"

# 4. 安装 Docker
echo -e "${YELLOW}[4/8] 安装 Docker...${NC}"
if ! command -v docker &> /dev/null; then
    curl -fsSL https://get.docker.com | bash
    systemctl enable docker
    systemctl start docker
fi
echo "Docker version: $(docker --version)"

# 5. 安装 Docker Compose
echo -e "${YELLOW}[5/8] 安装 Docker Compose...${NC}"
if ! command -v docker-compose &> /dev/null; then
    curl -L "https://github.com/docker/compose/releases/latest/download/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
    chmod +x /usr/local/bin/docker-compose
fi
echo "Docker Compose version: $(docker-compose --version)"

# 6. 安装 MySQL
echo -e "${YELLOW}[6/8] 安装 MySQL 8.0...${NC}"
if ! command -v mysql &> /dev/null; then
    apt-get install -y mysql-server
    systemctl enable mysql
    systemctl start mysql
fi
echo "MySQL installed"

# 7. 创建数据库和用户
echo -e "${YELLOW}[7/8] 创建数据库...${NC}"
mysql -u root <<EOF
CREATE DATABASE IF NOT EXISTS campus_market CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
CREATE USER IF NOT EXISTS 'campus'@'localhost' IDENTIFIED BY 'campus_market_2024';
GRANT ALL PRIVILEGES ON campus_market.* TO 'campus'@'localhost';
FLUSH PRIVILEGES;
EOF
echo "Database 'campus_market' created"

# 8. 配置 Nginx
echo -e "${YELLOW}[8/8] 配置 Nginx...${NC}"
cat > /etc/nginx/sites-available/campus-market <<'NGINX_CONF'
server {
    listen 80;
    server_name api.campus-market.example.com;

    location / {
        proxy_pass http://127.0.0.1:8080;
        proxy_http_version 1.1;
        proxy_set_header Upgrade $http_upgrade;
        proxy_set_header Connection 'upgrade';
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
        proxy_cache_bypass $http_upgrade;
    }

    # 静态文件
    location /static/ {
        alias /var/www/campus-market/static/;
        expires 30d;
    }
}
NGINX_CONF

ln -sf /etc/nginx/sites-available/campus-market /etc/nginx/sites-enabled/
nginx -t && systemctl reload nginx || echo "Nginx config test failed, please check"

echo -e "${GREEN}========================================${NC}"
echo -e "${GREEN}  初始化完成！${NC}"
echo -e "${GREEN}========================================${NC}"
echo ""
echo "下一步："
echo "  1. 修改 /etc/nginx/sites-available/campus-market 中的域名"
echo "  2. 安装 Redis: apt-get install -y redis-server"
echo "  3. 修改 backend/config/config.yaml 中的数据库密码"
echo "  4. cd backend && go mod tidy && go build -o server ./cmd/server/"
echo "  5. SSL 证书: certbot --nginx -d api.campus-market.example.com"
echo ""