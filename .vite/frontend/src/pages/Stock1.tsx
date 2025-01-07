import React, { useState, useEffect } from "react";
import axios from "axios";
import styled from 'styled-components';
import Header from '../components/header';
import p1 from '../../public/assets/sofa.jpg';  // ใช้เป็น placeholder ถ้าภาพจาก API ไม่มี
import p2 from '../../public/assets/table.jpg';
import p3 from '../../public/assets/chair.jpg';
import promotion from '../../public/assets/promotion.png';
import Card from '../components/card1'

const ProductGrid = styled.div`
  display: flex;
  gap: 20px;
  flex-wrap: wrap;
  justify-content: center;
  padding: 20px;
`;

const Stock1: React.FC = () => {
  const [products, setProducts] = useState<any[]>([]);  // กำหนดประเภทข้อมูลของ products ให้เป็น array

  useEffect(() => {
    // เรียก API เพื่อดึงข้อมูลสินค้า
    axios.get("http://localhost:8000/stock")  // ใช้ endpoint ที่ถูกต้อง
      .then((response) => {
        setProducts(response.data);  // ตั้งค่าผลลัพธ์ที่ได้รับจาก API ไปยัง state
      })
      .catch((error) => {
        console.error("Error fetching products:", error);
      });
  }, []);

  return (
    <div>
      <Header />

      <h1 style={{ marginTop: "100px" }}></h1>
      <a href="/" style={{ display: "flex", justifyContent: "center", alignItems: "center" }}>
        <img style={{ width: "700px", height: "auto" }} src={promotion} alt="promotion" />
      </a>

      {/* ใช้การ map ข้อมูลจาก products เพื่อแสดงสินค้า */}
      <ProductGrid>
        {products.map((product, index) => (
          <Card 
            key={index}
            title={product.title}  // สมมติว่าใน API มีชื่อสินค้า
            price={product.price}  // สมมติว่าใน API มีราคาสินค้า
            image={product.image || p1}  // ใช้ภาพจาก API หรือใช้ภาพ placeholder หากไม่มี
          />
        ))}
      </ProductGrid>
    </div>
  );
};

export default Stock1;
