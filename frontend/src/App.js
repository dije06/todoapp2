import React, { useState, useEffect } from 'react';

function App() {
  const [todos, setTodos] = useState([]); // Start with null to detect loading state
  const [title, setTitle] = useState("");

  const fetchTodos = async () => {
    try {
      const res = await fetch("http://localhost:8080/api/todos");
      const data = await res.json();
      setTodos(data);
    } catch (err) {
      console.error("Failed to fetch todos:", err);
      setTodos([]); // fallback to empty array on error
    }
  };

  const addTodo = async () => {
    if (!title.trim()) return;
    await fetch("http://localhost:8080/api/todos", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ title }),
    });
    setTitle("");
    fetchTodos();
  };

  const deleteTodo = async (id) => {
    await fetch(`http://localhost:8080/api/todos/${id}`, {
      method: "DELETE",
    });
    fetchTodos();
  };

  useEffect(() => {
    fetchTodos();
  }, []);

  return (
    <div style={{ padding: 20 }}>
      <h2>To-Do List</h2>
      <input
        value={title}
        onChange={(e) => setTitle(e.target.value)}
        placeholder="Enter task"
      />
      <button onClick={addTodo}>Add</button>

      <ul>
        {Array.isArray(todos) ? (
          todos.length > 0 ? (
            todos.map((todo, index) => (
            <li key={todo.ID ?? index}>
              {todo.Title ?? "Untitled"} <button onClick={() => deleteTodo(todo.ID)}>X</button>
            </li>
          ))
          ) : (
            <li>No tasks available.</li>
          )
        ) : (
          <li>Loading...</li>
        )}
      </ul>
    </div>
  );
}

export default App;
