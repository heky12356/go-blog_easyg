import React from 'react';
import Container from 'react-bootstrap/Container';
import {Routes, Route, Link} from 'react-router-dom';
import Test from './pages/test';
import Home from './pages/home';
export default function App () {
  return (
    <div>
      <Container>
        <h1 className='pt-4'>Blog</h1>
        <hr />
        <Routes>
          <Route path="/" element={<Home />} />
          <Route path="/test" element={<Test />} />
        </Routes>
      </Container>
    </div>
  )
}