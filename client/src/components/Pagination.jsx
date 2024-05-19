import { useEffect, useState } from "react"

const Pagination = ({currentPage, total, limit, setPage}) => {
  const [pagesArray, setPagesArray] = useState([])

  useEffect(() => {
    pages()
  }, [currentPage, total, limit])

  const pages = () => {
    let pageCount = Math.floor(total/limit)
    const rest = total % limit

    if (rest !== 0) {
      pageCount += 1
    }

    setPagesArray(Array.from({ length: pageCount }, (value, index) => index + 1))
  }

  const handleClick = (page) => {
    if (page < 1 || page === currentPage) return
    setPage(page)
  }

  if (!total) {
    return null
  }

  return <div className="flex flex-row mt-3 mb-3">
    <button onClick={() => {handleClick(currentPage - 1)}} className="p-1 rounded-xl bg-sky-300 mr-1 cursor-pointer">Prev</button>
    {
      pagesArray.map((page) => {
        return <button key={page} onClick={() => {handleClick(page)}} className="w-6 p-1 rounded-xl bg-sky-300 mr-1 cursor-pointer">
          {currentPage === page && <span className="underline ">{page}</span>}
          {currentPage !== page && <span>{page}</span>}
        </button>
      })
    }
    <button onClick={() => {handleClick(currentPage + 1)}} className="p-1 rounded-xl bg-sky-300 cursor-pointer">Next</button>
  </div>
}

export default Pagination
