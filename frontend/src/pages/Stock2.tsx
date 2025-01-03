import React, { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import styled from 'styled-components';
import Header from '../components/header';
import axios from 'axios';
import '../components/card2.css';

const Stock2: React.FC = () => {
  const [products, setProducts] = useState<any[]>([]); // ใช้ any[] หรือปรับเป็น interface สำหรับ Product
  const navigate = useNavigate();

  // ฟังก์ชันดึงข้อมูลจาก API
  useEffect(() => {
    axios
      .get("http://localhost:8000/products")
      .then((response) => {
        console.log("API Response:", response.data); // ตรวจสอบว่าโครงสร้างเป็นอย่างไร
        setProducts(response.data.products || []); // ดึงเฉพาะ `products` ถ้ามี
      })
      .catch((error) => {
        console.error("Error fetching products:", error);
      });
  }, []);
  

  const handleNext3Click = (productId: number) => {
    navigate(`/Stock3`, { state: { productId } }); // ใช้ state เพื่อส่ง productId ไปยังหน้า Stock3
  };
  

  return (
    <div>
      <Header />
      <a href="/" style={{ position: "absolute", top: "100px", right: "1400px" }}>
        <img style={{ width: "50px", height: "auto" }} src="/assets/back.png" alt="Back" />
      </a>

      <h1 style={{ marginTop: "100px", marginLeft: "400px" }}>กรุณาเลือกสินค้า</h1>
      <div className="product-grid2">
        {products.map((product) => (
          <button
            key={product.id}
            onClick={() => handleNext3Click(product.id)}
            className="product-card2"
          >
            <img className="product-card-img2" src={product.image} alt={product.name} />
            <div className="product-content2">
              <h2 className="product-title2">{product.name}</h2>
              <p className="product-description2">{product.description}</p>
            </div>
          </button>
        ))}
      </div>
    </div>
  );
};

export default Stock2;
