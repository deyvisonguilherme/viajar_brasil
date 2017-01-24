# Senem

Simulado para prova do exame nacional do ensino médio no Brasil, desenvolvido
utilizando as linguagens de programação:
  * GO
  * Elixir

Para iniciar o aplicativo:

  * Entre no diretório do aplicativo
  * Instale as dependências `mix deps.get`
  * Configure a conexão com o banco no arquivo `config/config.exs`
  * Crie a base de dados e as migrações `mix ecto.create && mix ecto.migrate`
  * Inicie o aplicativo `mix phoenix.server`

O aplicativo estará acessível pela url [`localhost:4000`](http://localhost:4000) em seu browser.
