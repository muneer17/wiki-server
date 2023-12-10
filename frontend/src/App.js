import logo from './logo.svg';
import './App.css';
import React from 'react';
import { BrowserRouter, Routes, Route } from 'react-router-dom';
import IndexPage from './IndexPage';
import ArticlePage from './ArticlePage';
import EditPage from './EditPage';

function App() {
   return ( 
       <BrowserRouter>
           <Routes>
               <Route path="/" element={<IndexPage />} />
             
               <Route path="/:name" element={<ArticlePage />} />
               <Route path="/edit/:name" element={<EditPage />} />
           </Routes>
       </BrowserRouter>
   );
}

export default App;