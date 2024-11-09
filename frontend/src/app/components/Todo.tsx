import { FocusEvent, ChangeEvent, useState, useRef } from "react";
import { Todo as TodoEntity } from "../entities/todo";

export default function Todo(props: { todo: TodoEntity }) {
  const [todo, setTodo] = useState<TodoEntity>(props.todo);
  const [isError, setIsError] = useState<boolean>(false);
  const [errorMessage, setErrorMessage] = useState<string>('');

  const handleChangeContent = (event: ChangeEvent<HTMLInputElement>) => {
    const newContent = event.target.value;

    if (newContent) {
      setIsError(false);
      setErrorMessage('');
    }

    setTodo((prevTodo) => ({...prevTodo, content: newContent}));
  };

  const inputRef = useRef<HTMLInputElement>(null);

  const handleBlur = async (event: FocusEvent<HTMLInputElement>) => {
    // 空白の場合フォーカスアウトさせない
    if (!event.target.value) {
      event.preventDefault();

      if (!inputRef.current) {
        return;
      }

      setIsError(true);
      setErrorMessage('必ず入力してください');
      inputRef.current.focus();
      return;
    }

    // 空白でなければリセット
    setIsError(false);
    setErrorMessage('');

    const response = await fetch(`http://localhost:8080/todo/${todo.todoId}`, {
      method: 'PUT',
      body: JSON.stringify({
        content: event.target.value,
      })
    });

    if (!response.ok) {
      setIsError(true);
      setErrorMessage('※更新に失敗しました');
    }
  };

  return (
    <div>
      <div id="task" className="flex justify-between items-center border-b border-slate-200 py-3 px-2 border-l-4  border-l-transparent">
        <div className="inline-flex items-center">
          <label className="flex items-center cursor-pointer relative" htmlFor="check-2">
            <input
              type="checkbox"
              className="peer h-5 w-5 cursor-pointer transition-all appearance-none rounded shadow hover:shadow-md border border-slate-300 checked:bg-slate-800 checked:border-slate-800"
              id="check-2"
              value={todo.todoId}
            ></input>
            <span className="absolute text-white opacity-0 peer-checked:opacity-100 top-1/2 left-1/2 transform -translate-x-1/2 -translate-y-1/2">
              <svg xmlns="http://www.w3.org/2000/svg" className="h-3.5 w-3.5" viewBox="0 0 20 20" fill="currentColor"
                stroke="currentColor" stroke-width="1">
                <path fill-rule="evenodd"
                  d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z"
                  clip-rule="evenodd"></path>
              </svg>
            </span>
          </label>
          <div className="w-full max-w-sm min-w-96">
            <input
              className="w-full placeholder:text-slate-400 text-slate-700 text-sm px-3 py-2 transition ease focus:outline-none focus:border-slate-400 hover:border-slate-300"
              placeholder="todo here"
              value={todo.content}
              ref={inputRef}
              onChange={handleChangeContent}
              onBlur={handleBlur}
            ></input>
          </div>
        </div>
        <div>
          <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" className="w-4 h-4 text-slate-500 hover:text-slate-700 hover:cursor-pointer">
            <path stroke-linecap="round" stroke-linejoin="round" d="M14.74 9l-.346 9m-4.788 0L9.26 9m9.968-3.21c.342.052.682.107 1.022.166m-1.022-.165L18.16 19.673a2.25 2.25 0 01-2.244 2.077H8.084a2.25 2.25 0 01-2.244-2.077L4.772 5.79m14.456 0a48.108 48.108 0 00-3.478-.397m-12 .562c.34-.059.68-.114 1.022-.165m0 0a48.11 48.11 0 013.478-.397m7.5 0v-.916c0-1.18-.91-2.164-2.09-2.201a51.964 51.964 0 00-3.32 0c-1.18.037-2.09 1.022-2.09 2.201v.916m7.5 0a48.667 48.667 0 00-7.5 0" />
          </svg>
        </div>
      </div>
      {isError && <p className="text-red-700 text-xs">{errorMessage}</p>}
    </div>
  );
}
