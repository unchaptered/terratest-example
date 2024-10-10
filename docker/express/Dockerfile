# 베이스 이미지로 Node.js 선택
FROM node:20

# 앱 디렉토리 생성
WORKDIR /app

# package.json과 package-lock.json 복사
COPY package*.json ./

# 의존성 설치
RUN npm install

# 서버 코드 복사
COPY . .

# 앱이 7000번 포트에서 실행될 예정
EXPOSE 7000

# 서버 실행
CMD ["node", "server.js"]
