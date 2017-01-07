defmodule Senem.PageController do
  use Senem.Web, :controller

  def index(conn, _params) do
    render conn, "index.html", page_title: "Início"
  end
end
