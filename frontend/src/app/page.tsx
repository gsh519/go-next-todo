"use client";
import { useEffect, useState } from "react";
import TodoComponent from "./components/Todo";
import { Todo, TodoType } from "./entities/todo";

export default function Home() {
  const [todos, setTodos] = useState<Todo[]>([]);

  useEffect(() => {
    fetchTodos();
  }, []);

  const fetchTodos = async () => {
    try {
      const response = await fetch('http://localhost:8080/todos');

      if (!response.ok) {
        throw new Error(`TODOの取得に失敗しました。`);
      }

      const json = await response.json();
      setTodos(() => json.map((row: TodoType) => new Todo(row)));
    } catch (error) {
      console.log(error);
    }
  };

  return (
    <body className="antialiased bg-slate-200 text-slate-700 mx-2">
      <div className="max-w-lg mx-auto my-10 bg-white p-8 rounded-xl shadow shadow-slate-300">
        <div className="flex flex-row justify-between items-center">
          <div>
            <h1 className="text-3xl font-medium text-slate-500">TODO List</h1>
          </div>
        </div>
        <p className="text-slate-500">Hello, here are your latest tasks</p>

        <div id="tasks" className="mt-5 mb-3">
          {todos.map((todo) => (
            <TodoComponent key={todo.todoId} todo={todo}></TodoComponent>
          ))}
        </div>
        <div className="flex justify-end">
          <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" strokeWidth={1.5} stroke="currentColor" className="size-9 fill-gray-500">
            <path strokeLinecap="round" strokeLinejoin="round" d="M12 9v6m3-3H9m12 0a9 9 0 1 1-18 0 9 9 0 0 1 18 0Z" />
          </svg>
        </div>

        <p className="mt-2 text-xs text-slate-500 text-center">Last updated 12 minutes ago</p>
      </div>

    </body>
  );
}
