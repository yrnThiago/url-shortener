import { Button } from "@/components/ui/button"
import { Input } from "@/components/ui/input"
import { useState } from "react";

function App() {
  const [url, setUrl] = useState("");

  const submit = () => {
    setUrl("asd");
    console.log("OPA");
  }
  return (
    <div className="bg-neutral-950 flex flex-col items-center justify-center min-h-svh">
    
      <div className="text-white">
        <h1>Shortener URLs</h1>
      </div>

      <div className="flex w-full max-w-sm items-center space-x-2">
        <Input placeholder="https://www.urlexample.com" />
        <Button type="submit" onClick={submit}>Confirm</Button>
      </div>

    </div>
  )
}

export default App

