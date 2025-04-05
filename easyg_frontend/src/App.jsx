import React from 'react';
import Container from 'react-bootstrap/Container';
import {Routes, Route, Link} from 'react-router-dom';
import Test from './pages/test';
import Home from './pages/home';
import Post from './pages/post';
import About from './pages/about';
import Admin from './pages/admin';
export default function App () {
  return (
    <div>
      <Container>
        <h1 className='pt-4'>
          <Link to ={'/'} className='text-decoration-none text-black'>Blog</Link>
        </h1>
        <hr />
        <Routes>
          <Route path="/" element={<Home />} />
          <Route path="/test" element={<Test />} />
          <Route path="/about" element={<About />} />
          <Route path="/admin" element={<Admin />} />
          <Route path='/post' element={<Post />} >
            <Route path=':uid' element={<Post />} />
          </Route>
        </Routes>
      </Container>
    </div>
  )
}