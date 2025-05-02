import React from 'react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import { Layout } from 'antd';
import './App.css';

const { Header, Content, Footer } = Layout;

function App() {
  return (
    <Router>
      <Layout className="layout">
        <Header>
          <div className="logo">E-Commerce Platform</div>
        </Header>
        <Content style={{ padding: '0 50px' }}>
          <Routes>
            <Route path="/" element={<Home />} />
            <Route path="/products" element={<Products />} />
            <Route path="/seckill" element={<Seckill />} />
            <Route path="/user" element={<UserCenter />} />
          </Routes>
        </Content>
        <Footer style={{ textAlign: 'center' }}>
          E-Commerce ©2025
        </Footer>
      </Layout>
    </Router>
  );
}

function Home() {
  return <div>Home Page</div>;
}

function Products() {
  return <div>Products Page</div>;
}

function Seckill() {
  return <div>Seckill Page</div>;
}

function UserCenter() {
  return <div>User Center</div>;
}

export default App;
