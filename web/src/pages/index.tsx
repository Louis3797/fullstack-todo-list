import TodoCard from "@/components/TodoCard";
import moment from "moment";
import type { GetServerSideProps, NextPage } from "next";
import { useState } from "react";
import { Todo } from "types";
import { instance } from "utils/http";
import { v4 } from "uuid";

interface HomeProps {
  data: Todo[];
}

const Home: NextPage<HomeProps> = ({ data }) => {
  const [input, setInput] = useState<string>("");

  async function handleSubmit() {
    if (input.length > 0) {
      const newTodo: Todo = {
        id: v4(),
        text: input,
        status: false,
        created_at: Date.now(),
      };

      instance
        .post("create-todo", JSON.stringify(newTodo))
        .then((res) => {
          console.log(res);
        })
        .catch(function (error) {
          console.log(error);
        });
    }
  }

  return (
    <div className="flex flex-col w-screen h-full justify-start items-center">
      <div className="flex flex-col h-96 py-5 justify-around items-start w-2/4">
        <h1 className="text-5xl font-bold text-secondary mt-10 mb-10">
          Todo List
        </h1>

        <form
          className="flex flex-row justify-between items-center px-1 w-full"
          onSubmit={handleSubmit}
        >
          <input
            type="text"
            onChange={(e) => setInput(e.target.value)}
            value={input}
            minLength={1}
            className="rounded-md w-full mr-20 bg-primary p-2 focus:outline-none text-accent font-semibold"
            placeholder="What else need you todo ??"
          />
          <button
            type="submit"
            className="px-4 py-1 bg-primary rounded-md text-accent font-semibold hover:bg-accent hover:text-primary duration-500 ease-int-out"
          >
            Submit
          </button>
        </form>
      </div>
      <div className="flex flex-col w-2/4 space-y-4 pb-10">
        {data.map((d: Todo) => {
          return (
            <TodoCard
              key={d.id}
              id={d.id}
              text={d.text}
              status={d.status}
              created_at={d.created_at}
            />
          );
        })}
      </div>
    </div>
  );
};

export const getServerSideProps: GetServerSideProps = async () => {
  const res = await instance.get(`http://localhost:4000`);

  const data: Todo[] = await res.data;
  data.sort((a: Todo, b: Todo) => {
    return moment(moment.unix(b.created_at)).diff(moment.unix(a.created_at));
  });

  return {
    props: { data },
  };
};

export default Home;
