import React, { useState } from "react";
import Moment from "react-moment";
import { Todo } from "types";
import { instance } from "utils/http";
import { BsFillTrashFill } from "react-icons/bs";

const TodoCard: React.FC<Todo> = ({ id, text, status, created_at }) => {
  const [done, setDone] = useState<boolean>(status == true ? true : false);

  function handleUpdate() {
    const updatedTodo: Todo = {
      id: id,
      text: text,
      status: !done,
      created_at: created_at,
    };
    instance.patch(`todo/${id}`, JSON.stringify(updatedTodo));
  }
  function deleteTodo() {
    instance.delete(`delete-todo/${id}`).then((res) => console.log(res));
  }

  return (
    <div className="flex flex-col bg-primary w-full p-3 rounded-md shadow-lg duration-500 hover:-translate-y-1 ease-in-out">
      <div className="flex flex-row justify-between items-center mb-4">
        <div className="flex flex-row justify-center items-center space-x-3">
          <input
            type="checkbox"
            checked={done}
            onChange={() => {
              setDone(!done);
              handleUpdate();
            }}
            className="checked:bg-secondary"
          />
          <p className="font-semibold text-accent text-lg">{text}</p>
        </div>
        <button
          className=" rounded-md px-2 py-1 hover:text-red-400 duration-500 ease-in-out"
          onClick={deleteTodo}
        >
          <BsFillTrashFill />
        </button>
      </div>
      <div className="flex flex-row justify-between">
        <p className="text-sm font-semibold text-accent">
          <Moment fromNow>{created_at}</Moment>
        </p>
      </div>
    </div>
  );
};

export default TodoCard;
