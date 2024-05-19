import { useRef, useState } from "react";

const TodoItem = ({todoItem, refetch}) => {
  const textRef = useRef(null)
  const [showEdit, setShowEdit] = useState(false)
  const [showUpdateError, setShowUpdateError] = useState(false)
  const [updateError, setUpdateError] = useState("")
  const [text, setText] = useState(todoItem.text)

  const handleDelete = async () => {
    const response = await fetch(`http://localhost:8080/api/todos/${todoItem.id}`, {
      method: "DELETE", 
      headers: {
        "Content-Type": "application/json",
      }
    });
    if (response.ok) {
      refetch()
    }
  }

  const handleUpdate = async (event) => {
    event.preventDefault()
    setUpdateError("")
    setShowUpdateError(false)
    const response = await fetch(`http://localhost:8080/api/todos/${todoItem.id}`, {
      method: "PUT", 
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        text: text
      })
    });
    if (!response.ok) {
      const errorBody = await response.json()
      if (errorBody.errors) {
        setUpdateError(errorBody.errors[0].text)
        setShowUpdateError(true)
      }
    } else {
      toggleEdit()
      refetch()
    }
  }

  const toggleEdit = () => {
    setShowEdit((prev) => !prev)
    textRef.current = todoItem.text
  }

  return <div className="w-full m-5 p-2">
    {
      !showEdit &&
      <div className="flex flex-row justify-between">
        <p className="w-3/4">{todoItem.text}</p>
        <div>
          <button onClick={toggleEdit} className="p-3 w-28 bg-lime-200 rounded-xl mr-2">Edit</button>
          <button onClick={handleDelete} className="p-3 w-28 bg-rose-400 rounded-xl">Delete</button>
        </div>
      </div>
    }
    { showEdit &&
      <form className="flex flex-row justify-between" onSubmit={handleUpdate}>
        <div className="w-3/4">
          <input type="text" value={text} onChange={(event) => {setText(event.target.value)}} className="rounded-xl p-2 w-full"/>
          { showUpdateError &&
            <p className="text-red-400">{updateError}</p>
          }
        </div>
        <div>
          <button type='submit' className="p-3 w-28 bg-lime-200 rounded-xl mr-2">Save</button>
          <button type='button' className="p-3 w-28 bg-rose-400 rounded-xl" onClick={toggleEdit}>Cancel</button>
        </div>
      </form>
    }
  </div>
}

export default TodoItem
