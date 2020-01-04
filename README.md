# Testapi

Uma api escrita em Golang.

## Requesitos
- Go 1.13
- PostgresSQL

## Descrição 

| Método | Endpoints          | Descrição                            |
|--------|--------------------|--------------------------------------|
| Get    | /api/products      | Mostra todos os produtos             |
| Post   | /api/products      | Cria um novo produto                 |
| Put    | /api/products/:id  | Atualiza um produto                  |
| Delete | /api/products/:id  | Remove um produto pelo id            |
| Get    | /api/products/:id  | Mostra os detalhes do produto        |
| Post   | /api/login         | Faz a autenticação e retorna o token |
| Post   | /api/register      | Registra um novo usuário             |
| Get    | /api/customers     | Mostra todos os clientes             |
| Post   | /api/customers     | Cria um novo cliente                 |
| Put    | /api/customers/:id | Atualiza o dados do cliente          |
| Delete | /api/customers/:id | Remove o cliente pelo id             |
| Post   | /api/orders        | Cria um novo pedido                  |
| Get    | /api/orders        | Mostra todos os pedidos do usuário   |
| Put    | /api/orders/:id    | Atualiza o pedido pelo id            |
| Get    | /api/orders/:id    | Mostra os detalhes do pedido         |

