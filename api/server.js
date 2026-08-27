const express = require("express");
const cors = require("cors");

const API_PORT = process.env.API_PORT || 8080;
const BACKEND_HOST = process.env.BACKEND_HOST || "http://locahost:8081";

const app = express();

app.use(cors());
app.use(express.json());

app.get("/users", async (req, res) => {
  // This is api gateway we can have multiple middlewares here
  // Suppose for authentication, rate limiting, tls termination etc.
  const resp = await fetch(BACKEND_HOST + "/users");
  const result = await resp.json();
  return res.json(result);
});

app.post("/users", (req, res) => {
  let name = req.body.name;

  const resp = fetch(BACKEND_HOST+"/users", {
    method: "POST",
    body: JSON.stringify({name}),
    headers: {
        "Content-Type": "application/json"
    }
  })

  const result = resp.json()

  return res.json(result)
});

app.listen(API_PORT, "0.0.0.0", () => {
  console.log("listening on API_port ", API_PORT);
});
