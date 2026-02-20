# ⚔️ Skyrim Quest Log | GoTask API

![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
![Gin Gonic](https://img.shields.io/badge/gin-white?style=for-the-badge&logo=gin)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-316192?style=for-the-badge&logo=postgresql&logoColor=white)
![JWT](https://img.shields.io/badge/JWT-black?style=for-the-badge&logo=JSON%20web%20tokens)

"I used to be an adventurer like you, then I took a Task in the knee." 🏹

Este é um sistema de gerenciamento de missões (Quest Log) inspirado no universo de The Elder Scrolls V: Skyrim. O projeto utiliza uma arquitetura moderna com um **Backend em Go** e um **Frontend SPA imersivo**.



---

## 📜 Funcionalidades do Pergaminho

- [x] **Autenticação Dragonborn:** Login seguro com geração de Token JWT.
- [x] **Quest Log (CRUD):** Adicione novas missões, visualize seu progresso e abandone (delete) quests.
- [x] **Status de Conclusão:** Marque missões como concluídas com feedback visual (estilo Skyrim).
- [x] **Interface Imersiva:** Design Single Page Application (SPA) com fontes Cinzel e MedievalSharp.
- [x] **Persistência em Oblivion:** Banco de dados PostgreSQL hospedado no Neon.tech.

---

## 🛡️ Tecnologias Utilizadas

- **Backend:** Golang com Framework Gin Gonic.
- **ORM:** GORM para interações fluidas com o banco de dados.
- **Segurança:** Middleware de autenticação JWT e proteção de CORS.
- **Frontend:** Vanilla JavaScript (Fetch API) e CSS temático (Google Fonts).

---

## 🏗️ Estrutura da API

### 🔐 Autenticação
- `POST /login`: Valida as credenciais e entrega o Token de acesso.

### 📝 Quests (Necessita de Token no Header)
- `GET /tasks`: Lista todas as missões do seu diário.
- `POST /tasks`: Adiciona uma nova missão à sua jornada.
- `PUT /tasks/:id`: Atualiza o status de conclusão da missão.
- `DELETE /tasks/:id`: Remove uma missão do pergaminho.

---

## ⚙️ Configuração da sua Jornada


1. **Clone o repositório:**
   ```bash
   git clone [https://github.com/Lucasdip/gotask-app.git](https://github.com/Lucasdip/gotask-app.git)

👤 Autor
Desenvolvido por [Lucas Lima] – Sinta-se à vontade para entrar em contato!
