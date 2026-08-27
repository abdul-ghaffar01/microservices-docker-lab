"use client";
import { useEffect, useState } from "react";

export default function Home() {
  const [name, setName] = useState("");
  const [loading, setLoading] = useState(false);
  const [users, setUsers] = useState([{ id: "", name: "" }]);

  const fetchUsers = async () => {
    setLoading(true);

    // fetching users from api
    const resp = await fetch(`http://localhost:8080/users`);
    const res = await resp.json();
    console.log(res);

    setUsers(res.users);
    setLoading(false);
  };

  const addUser = async () => {
    // TODO: make api call

    const resp = await fetch("http://localhost:8080/users", {
      method: "POST",
      body: JSON.stringify({ name }),
      headers: {
        "Content-Type": "application/json",
      },
    });

    const res = await resp.json();
    console.log(res);

    if (res.error) {
      console.log(res.message);
      return;
    }

    setUsers((prev) => [...prev, res.user]); // Hope so this works I am not sure; Wohhh worked
  };
  useEffect(() => {
    fetchUsers();
  }, []);

  return (
    <div className="flex justify-center gap-10 py-10 min-h-screen">
      {/* To show all the users from database */}
      <div className="flex-1 text-center">
        <div>
          <h1 className="text-3xl">Fetch all users</h1>
          <hr className="w-[80%] mx-auto my-3" />
          <button
            onClick={fetchUsers}
            className="px-5 py-2 border border-white rounded-sm mt-2 cursor-pointer bg-white text-black"
          >
            Fetch
          </button>
        </div>

        <div className="w-full flex flex-col items-center mt-6 overflow-y-scroll">
          {loading && <p>Loading users...</p>}
          {!loading &&
            users.map((elem, index) => {
              return (
                <div
                  key={index}
                  className="w-[600px] bg-gray-800 text-black rounded-xl m-2 overflow-hidden"
                >
                  {/* User ID */}
                  <div className="text-xl bg-red-200 p-2">
                    ID: <span>{elem.id}</span>
                  </div>
                  {/* User name */}
                  <div className="text-2xl bg-yellow-300 p-2">
                    Name: <span>{elem.name}</span>
                  </div>
                </div>
              );
            })}
        </div>
      </div>

      {/* To add a new user to the database */}
      <div className="flex-1 text-center border-l">
        <h1 className="text-3xl">Add a new user</h1>
        <div className="">
          <input
            className="border border-white rounded-sm outline-none p-2 mx-2 min-w-[300px]"
            value={name}
            onChange={(e) => setName(e.target.value)}
            type="text"
            placeholder="Name"
          />
          <button
            onClick={addUser}
            className="px-5 py-2 border border-white rounded-sm mt-2 cursor-pointer"
          >
            Add
          </button>
        </div>
      </div>
    </div>
  );
}
