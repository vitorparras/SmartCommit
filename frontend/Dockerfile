FROM node:18-alpine

# Define o diretório de trabalho dentro do container
WORKDIR /usr/src/app

# Copia apenas os arquivos de configuração para aproveitar o cache
COPY package*.json ./

# Instala as dependências
RUN npm install

# Copia o restante do código-fonte (em produção, isso seria necessário)
# Em desenvolvimento, o volume sobrescreve esta pasta
COPY . .

# Porta usada pelo Vite para hot-reload
EXPOSE 3000

# Comando de inicialização
CMD ["npm", "run", "dev"]
