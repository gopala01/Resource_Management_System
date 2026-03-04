import { useState } from "react";
import reactLogo from "./assets/react.svg";
import viteLogo from "/vite.svg";
import "./App.css";

function App() {
  const [count, setCount] = useState(0);

  return (
    <div className="App">
      <h1>Job Manager</h1>
      <div className="card-container">
        <button onClick={() => console.log("Create")}>Create Job</button>
        <button onClick={() => console.log("Read")}>Read Job</button>
        <button onClick={() => console.log("Update")}>Update Job</button>
        <button onClick={() => console.log("Delete")}>Delete Job</button>
      </div>
    </div>
  );
}

export default App;
