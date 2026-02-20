# 🚀 GoTask API | Task Management System

![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
![Gin Gonic](https://img.shields.io/badge/gin-white?style=for-the-badge&logo=gin)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-316192?style=for-the-badge&logo=postgresql&logoColor=white)
![JWT](https://img.shields.io/badge/JWT-black?style=for-the-badge&logo=JSON%20web%20tokens)
![Render](https://img.shields.io/badge/Render-%46E3B7?style=for-the-badge&logo=render&logoColor=white)

API REST de alto desempenho para gerenciamento de tarefas, construída com foco em segurança, escalabilidade e simplicidade.

---

## 🏗️ Arquitetura do Sistema



O projeto segue o padrão de responsabilidade única, onde cada pacote tem um papel definido:
- **Handlers:** Processamento de requisições e respostas JSON.
- **Models:** Estrutura de dados e comunicação com o Banco via GORM.
- **Middleware:** Filtro de segurança para validação de tokens JWT.

---

## ⚡ Principais Funcionalidades

- [x] **Autenticação Segura:** Login com geração de Token JWT.
- [x] **CRUD Completo:** Listar, criar, atualizar e deletar tarefas.
- [x] **Persistência Cloud:** Conectado ao Neon (PostgreSQL) com SSL.
- [x] **Auto Migration:** O banco de dados se ajusta automaticamente ao iniciar o app.
- [x] **Segurança de Dados:** Variáveis sensíveis protegidas via `.env`.

---

## 📡 API Endpoints

### 🔐 Autenticação
| Método | Rota | Descrição |
| :--- | :--- | :--- |
| `POST` | `/login` | Recebe credenciais e retorna o Token de acesso. |

### 📝 Tasks (Requer Header: `Authorization: Bearer <token>`)
| Método | Rota | Descrição |
| :--- | :--- | :--- |
| `GET` | `/tasks` | Retorna todas as tarefas cadastradas. |
| `POST` | `/tasks` | Cria uma nova tarefa. |
| `PUT` | `/tasks/:id` | Altera o status ou título de uma tarefa existente. |
| `DELETE` | `/tasks/:id` | Remove uma tarefa do banco de dados. |

---

## 🛠️ Como Rodar o Projeto

1. **Clone o repositório:**
   ```bash
   git clone [https://github.com/Lucasdip/gotask-app.git](https://github.com/Lucasdip/gotask-app.git)

👤 Autor
Desenvolvido por [Lucas Lima] – Sinta-se à vontade para entrar em contato!
