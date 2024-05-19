import { useEffect, useRef, useState } from "react"
import TodoItem from "./TodoItem"
import Pagination from "./Pagination"

const TodoList = () => {
  const textRef = useRef(null)
  const [todos, setTodos] = useState([])
  const [somethingWrong, setSomethingWrong] = useState(false)
  const [showCreateError, setShowCreateError] = useState(false)
  const [createError, setCreateError] = useState("")
  const [page, setPage] = useState(1)
  const [limit, setLimit] = useState(5)
  const [pagination, setPagination] = useState({})

  useEffect(() => {
    fetchTodos()
  }, [page])

  const refetch = () => {
    fetchTodos()
  }

  const fetchTodos = async () => {
    const response = await fetch(`http://localhost:8080/api/todos?page=${page}&limit=${limit}`)
    if (!response.ok) {
      setSomethingWrong(true)
    }

    const responseJson = await response.json()
    setTodos(
      responseJson.data
    )
    setPagination(responseJson.pagination)
  }

  const handleCreate = async (event) => {
    event.preventDefault()
    setCreateError("")
    setShowCreateError(false)
    const response = await fetch("http://localhost:8080/api/todos", {
      method: "POST", 
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        text: textRef.current.value
      })
    });
    if (!response.ok) {
      const errorBody = await response.json()
      if (errorBody.errors) {
        setCreateError(errorBody.errors[0].text)
        setShowCreateError(true)
      }
    } else {
      refetch()
    }
  }

  if (somethingWrong) {
    return <>
      <h1 className="text-red-400">Could not fetch todos</h1>
    </>
  }

  return <>
    <div className="flex flex-col">
      {todos.map((todo) => {
        return <TodoItem key={todo.id} todoItem={todo} refetch={refetch}></TodoItem>
      })}
    </div>
    <Pagination total={pagination.total} limit={limit} currentPage={page} setPage={setPage}/>
    <form onSubmit={handleCreate}>
      <div className="flex flex-row justify-between">
        <div className="w-3/4">
          <input type="text" ref={textRef} className="rounded-xl p-2 w-full"/>
          { showCreateError && <p className="text-red-400">{createError}</p>}

        </div>
        <button type="submit" className="p-3 w-28 bg-sky-500 rounded-xl">Create</button>
      </div>  
    </form>
    
    
  </>
}

export default TodoList
