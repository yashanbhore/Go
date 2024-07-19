import React from 'react';
import 'bootstrap/dist/css/bootstrap.css'
import Entries from './component/entries.component';
import { Navbar } from './component/Navbar';



function App() {
  return (
    <div>
      <Navbar/>
      <Entries/>
    </div>
  );
}

export default App;