import axios from "axios";
import React from "react";
import Moment from "react-moment";
import { Todo } from "types";
import { instance } from "utils/http";

const TodoCard: React.FC<Todo> = ({ id, text, status, created_at }) => {
  function deleteTodo() {
    instance.delete(`delete-todo/${id}`).then((res) => console.log(res));
  }

  return (
    <div className="flex flex-col bg-primary w-full p-3 rounded-md">
      <div className="flex flex-row justify-between items-center mb-4">
        <p className="font-semibold text-accent text-lg">{text}</p>
        <div className="flex flex-row">
          <button className="bg-primary-600 rounded-md px-2 py-1 mr-2">
            Update
          </button>
          <button
            className="bg-primary-600 rounded-md px-2 py-1"
            onClick={deleteTodo}
          >
            Delete
          </button>
        </div>
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
