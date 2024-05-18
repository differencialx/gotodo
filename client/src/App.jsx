import { useEffect, useState } from 'react'
import './App.css'

function App() {
  const [todos, setTodos] = useState({})
  useEffect(() => {
    const fetchTodos = async () => {
      const response = await fetch("/api/todos")
      if (!response.ok) {
        return
      }
      
      setTodos(await response.json())
    }

    fetchTodos()
  })

  return (
    <>
      {todos}
    </>
  )
}

export default App
