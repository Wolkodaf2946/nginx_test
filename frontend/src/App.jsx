import { useEffect, useState } from "react";

// База API: в dev-режиме — относительный /api (Vite проксирует на бэкенд),
// в docker-сборке подставляется полный URL бэкенда через VITE_API_URL.
const API_BASE = import.meta.env.VITE_API_URL || "/api";
const API = `${API_BASE}/todos`;

export default function App() {
  const [todos, setTodos] = useState([]);
  const [title, setTitle] = useState("");
  const [error, setError] = useState(null);

  async function load() {
    try {
      const res = await fetch(API);
      const json = await res.json();
      setTodos(json.data || []);
      setError(null);
    } catch (e) {
      setError("Не удалось загрузить задачи");
    }
  }

  useEffect(() => {
    load();
  }, []);

  async function addTodo(e) {
    e.preventDefault();
    if (!title.trim()) return;
    await fetch(API, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ title: title.trim() }),
    });
    setTitle("");
    load();
  }

  async function toggle(todo) {
    await fetch(`${API}/${todo.id}`, {
      method: "PATCH",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ done: !todo.done }),
    });
    load();
  }

  async function remove(id) {
    await fetch(`${API}/${id}`, { method: "DELETE" });
    load();
  }

  return (
    <div className="wrap">
      <h1>Todo</h1>
      <p className="hint">React → Go backend</p>

      <form onSubmit={addTodo} className="row">
        <input
          value={title}
          onChange={(e) => setTitle(e.target.value)}
          placeholder="Новая задача…"
        />
        <button type="submit">Добавить</button>
      </form>

      {error && <p className="error">{error}</p>}

      <ul>
        {todos.map((t) => (
          <li key={t.id} className={t.done ? "done" : ""}>
            <label>
              <input
                type="checkbox"
                checked={t.done}
                onChange={() => toggle(t)}
              />
              <span>{t.title}</span>
            </label>
            <button className="del" onClick={() => remove(t.id)}>
              ✕
            </button>
          </li>
        ))}
      </ul>

      {todos.length === 0 && !error && <p className="hint">Пока пусто.</p>}
    </div>
  );
}
