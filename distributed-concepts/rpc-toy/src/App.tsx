import reactLogo from './assets/react.svg'
import reactLogo2 from './assets/delete-1.png'
import viteLogo from '/vite.svg'
import './App.css'

function App() {

  return (
    <>
      <div>
        <a href="https://vite.dev" target="_blank">
          <img src={ viteLogo } className="logo" alt="Vite logo" />
        </a>
        <a href="https://react.dev" target="_blank">
          <img src={ reactLogo } className="logo react" alt="React logo 2" />
        </a>
        <a href="https://react.dev" target="_blank">
          <img src={ reactLogo2 } className="logo react" alt="React logo 2" />
        </a>
      </div>
      <h1>K8s Learning...</h1>
      <div className="card">
        <h2><i><u> Click to Begin </u></i></h2>
      </div>
      <p className="read-the-docs">
        Overview:
        <br />
        This project will explore the performance of varios RPC frameworks and serialization formats and will also contain <br />
        flame graphs and internals of kafka and other distributed systems. <br />

        Kafka contains of a storage layer and a compute layer. The storage layer is responsible for storing the data and the compute layer is responsible for processing the data. <br />

      </p>
    </>
  )
}

export default App
